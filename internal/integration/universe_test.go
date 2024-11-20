package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pentops/flowtest"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_tpb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_spb"
	"github.com/pentops/o5-builds/internal/app"
	"github.com/pentops/o5-builds/internal/integration/mocks"
	"github.com/pentops/o5-messaging/gen/o5/messaging/v1/messaging_tpb"
	"github.com/pentops/o5-messaging/outbox/outboxtest"
	"github.com/pentops/pgtest.go/pgtest"
	"github.com/pentops/sqrlx.go/sqrlx"
)

type Universe struct {
	Outbox *outboxtest.OutboxAsserter

	RepoCommand  github_spb.RepoCommandServiceClient
	RepoQuery    github_spb.RepoQueryServiceClient
	RawTopic     messaging_tpb.RawMessageTopicClient
	BuilderReply builder_tpb.BuilderReplyTopicClient

	Github *mocks.GithubMock
}

func NewUniverse(t *testing.T) (*flowtest.Stepper[*testing.T], *Universe) {
	name := t.Name()
	stepper := flowtest.NewStepper[*testing.T](name)
	uu := &Universe{}

	stepper.Setup(func(ctx context.Context, t flowtest.Asserter) error {
		log.DefaultLogger = log.NewCallbackLogger(stepper.LevelLog)
		setupUniverse(ctx, t, uu)
		return nil
	})

	stepper.PostStepHook(func(ctx context.Context, t flowtest.Asserter) error {
		uu.Outbox.AssertEmpty(t)
		return nil
	})

	return stepper, uu
}

const TestVersion = "test-version"

// setupUniverse should only be called from the Setup callback, it is effectively
// a method but shouldn't show up there.
func setupUniverse(ctx context.Context, t flowtest.Asserter, uu *Universe) {
	t.Helper()

	conn := pgtest.GetTestDB(t, pgtest.WithDir("../../ext/db"))
	db := sqrlx.NewPostgres(conn)

	uu.Outbox = outboxtest.NewOutboxAsserter(t, conn)
	uu.Github = mocks.NewGithubMock()
	appSet, err := app.NewApp(db, uu.Github)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	grpcPair := flowtest.NewGRPCPair(t, app.GRPCMiddleware()...)

	appSet.RegisterGRPC(grpcPair.Server)

	uu.RawTopic = messaging_tpb.NewRawMessageTopicClient(grpcPair.Client)
	uu.RepoCommand = github_spb.NewRepoCommandServiceClient(grpcPair.Client)
	uu.RepoQuery = github_spb.NewRepoQueryServiceClient(grpcPair.Client)
	uu.BuilderReply = builder_tpb.NewBuilderReplyTopicClient(grpcPair.Client)

	grpcPair.ServeUntilDone(t, ctx)
}

func (uu *Universe) GithubEvent(t flowtest.TB, eventType string, event interface{}) {

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	msg := &messaging_tpb.RawMessage{
		Topic:   fmt.Sprintf("github:%s", eventType),
		Payload: payload,
	}

	_, err = uu.RawTopic.Raw(context.Background(), msg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}