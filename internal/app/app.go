package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pentops/grpc.go/protovalidatemw"
	"github.com/pentops/grpc.go/versionmw"
	"github.com/pentops/log.go/grpc_log"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/internal/github"
	"github.com/pentops/o5-builds/internal/state"
	"github.com/pentops/o5-messaging/outbox"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
)

type App struct {
	QueryService   *GithubQueryService
	CommandService *GithubCommandService
	WebhookWorker  *github.WebhookWorker
	ReplyWorker    *ReplyWorker
}

func NewApp(db sqrlx.Transactor, githubClient IClient) (*App, error) {

	outboxPub, err := outbox.NewDirectPublisher(db, outbox.DefaultSender)
	if err != nil {
		return nil, fmt.Errorf("failed to create outbox publisher: %w", err)
	}

	states, err := state.NewStateMachines()
	if err != nil {
		return nil, fmt.Errorf("failed to create state machines: %w", err)
	}

	refs, err := NewRefStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create ref store: %w", err)
	}

	githubHandler, err := NewGithubHandler(refs, githubClient, outboxPub)
	if err != nil {
		return nil, fmt.Errorf("failed to create github handler: %w", err)
	}

	webhookWorker, err := github.NewWebhookWorker(githubHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create webhook worker: %w", err)
	}

	replyWorker, err := NewReplyWorker(githubClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create reply worker: %w", err)
	}

	commandService, err := NewGithubCommandService(db, states, githubHandler)
	if err != nil {
		return nil, fmt.Errorf("failed to create github command service: %w", err)
	}

	queryService, err := NewGithubQueryService(db, states)
	if err != nil {
		return nil, fmt.Errorf("failed to create github query service: %w", err)
	}

	return &App{
		WebhookWorker:  webhookWorker,
		ReplyWorker:    replyWorker,
		CommandService: commandService,
		QueryService:   queryService,
	}, nil
}

func (aa *App) RegisterGRPC(srv *grpc.Server) {
	aa.WebhookWorker.RegisterGRPC(srv)
	aa.ReplyWorker.RegisterGRPC(srv)
	aa.CommandService.RegisterGRPC(srv)
	aa.QueryService.RegisterGRPC(srv)
}

func GRPCMiddleware(version string) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpc_log.UnaryServerInterceptor(log.DefaultContext, log.DefaultTrace, log.DefaultLogger),
		j5auth.GRPCMiddleware,
		protovalidatemw.UnaryServerInterceptor(),
		versionmw.UnaryServerInterceptor(version),
	}
}

type DBConfig struct {
	URL string `env:"POSTGRES_URL"`
}

func (cfg *DBConfig) OpenDatabase(ctx context.Context) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	// Default is unlimited connections, use a cap to prevent hammering the database if it's the bottleneck.
	// 10 was selected as a conservative number and will likely be revised later.
	db.SetMaxOpenConns(10)

	for {
		if err := db.Ping(); err != nil {
			log.WithError(ctx, err).Error("pinging PG")
			time.Sleep(time.Second)
			continue
		}
		break
	}

	return db, nil
}
