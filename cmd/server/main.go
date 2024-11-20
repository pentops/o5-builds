package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pentops/grpc.go/grpcbind"
	"github.com/pentops/log.go/log"
	"github.com/pentops/protostate/psmigrate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pentops/j5/lib/j5grpc"
	"github.com/pentops/o5-builds/internal/app"
	"github.com/pentops/o5-builds/internal/github"
	"github.com/pentops/o5-builds/internal/state"
	"github.com/pentops/runner/commander"
	"github.com/pentops/sqrlx.go/sqrlx"
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

func openDatabase(ctx context.Context, dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

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

func runMigrate(ctx context.Context, config struct {
	MigrationsDir string `env:"MIGRATIONS_DIR" default:"./ext/db"`
	app.DBConfig
}) error {

	db, err := config.OpenDatabase(ctx)
	if err != nil {
		return err
	}

	return goose.Up(db, config.MigrationsDir)
}

func runServiceInfo(_ context.Context, _ struct{}) error {
	grpcServer := grpc.NewServer()
	ordersSet, err := app.NewApp(nil, nil)
	if err != nil {
		return fmt.Errorf("failed to build OMS: %w", err)
	}
	ordersSet.RegisterGRPC(grpcServer)
	return j5grpc.PrintServerInfo(grpcServer)
}

func runServe(ctx context.Context, config struct {
	PublicAddr  string `env:"PUBLIC_ADDR" default:":8081"`
	PostgresURL string `env:"POSTGRES_URL"`
}) error {

	dbConn, err := openDatabase(ctx, config.PostgresURL)
	if err != nil {
		return err
	}

	db := sqrlx.NewPostgres(dbConn)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		app.GRPCMiddleware(Version)...,
	)))

	githubClient, err := github.NewEnvClient(ctx)
	if err != nil {
		return err
	}

	ordersSet, err := app.NewApp(db, githubClient)
	if err != nil {
		return fmt.Errorf("failed to build OMS: %w", err)
	}
	ordersSet.RegisterGRPC(grpcServer)
	reflection.Register(grpcServer)

	return grpcbind.ListenAndServe(ctx, grpcServer, config.PublicAddr)
}
