package main

import (
	"context"
	"fmt"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pentops/grpc.go/grpcbind"
	"github.com/pentops/j5/lib/psm/psmigrate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pentops/j5/lib/j5grpc"
	"github.com/pentops/o5-builds/internal/app"
	"github.com/pentops/o5-builds/internal/github"
	"github.com/pentops/o5-builds/internal/slack"
	"github.com/pentops/o5-builds/internal/state"
	"github.com/pentops/runner/commander"
	"github.com/pentops/sqrlx.go/pgenv"
	"github.com/pressly/goose"
)

var Version string

func main() {

	mainGroup := commander.NewCommandSet()

	mainGroup.Add("serve", commander.NewCommand(runServe))
	mainGroup.Add("migrate", commander.NewCommand(runMigrate))
	mainGroup.Add("psm-tables", commander.NewCommand(runPSMTables))
	mainGroup.Add("info", commander.NewCommand(runServiceInfo))

	mainGroup.RunMain("registration", Version)
}

func runPSMTables(ctx context.Context, cfg struct {
}) error {

	stateMachines, err := state.NewStateMachines()
	if err != nil {
		return err
	}

	migrationFile, err := psmigrate.BuildStateMachineMigrations(stateMachines.TableSpecs()...)
	if err != nil {
		return fmt.Errorf("build migration file: %w", err)
	}

	fmt.Println(string(migrationFile))
	return nil
}

func runMigrate(ctx context.Context, cfg struct {
	MigrationsDir string `env:"MIGRATIONS_DIR" default:"./ext/db"`
	pgenv.DatabaseConfig
}) error {

	db, err := cfg.OpenPostgres(ctx)
	if err != nil {
		return err
	}

	return goose.Up(db, cfg.MigrationsDir)
}

func runServiceInfo(_ context.Context, _ struct{}) error {
	grpcServer := grpc.NewServer()
	ordersSet, err := app.NewApp(nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to build OMS: %w", err)
	}
	ordersSet.RegisterGRPC(grpcServer)
	return j5grpc.PrintServerInfo(os.Stdout, grpcServer)
}

func runServe(ctx context.Context, config struct {
	GRPC     grpcbind.EnvConfig
	DB       pgenv.DatabaseConfig
	SlackURL string `env:"SLACK_URL" default:""`
}) error {

	db, err := config.DB.OpenPostgresTransactor(ctx)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		app.GRPCMiddleware(Version)...,
	)))

	githubClient, err := github.NewEnvClient(ctx)
	if err != nil {
		return err
	}

	extraPublishers := []app.IBuildPublisher{}
	if config.SlackURL != "" {
		slackClient := slack.NewPublisher(config.SlackURL)
		extraPublishers = append(extraPublishers, slackClient)

	}

	ordersSet, err := app.NewApp(db, githubClient, extraPublishers)
	if err != nil {
		return fmt.Errorf("failed to build OMS: %w", err)
	}
	ordersSet.RegisterGRPC(grpcServer)
	reflection.Register(grpcServer)

	return config.GRPC.ListenAndServe(ctx, grpcServer)
}
