package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_spb"
	"github.com/pentops/o5-builds/internal/state"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GithubCommandService struct {
	db sqrlx.Transactor

	stateMachines *state.StateMachines
	*github_spb.UnimplementedRepoCommandServiceServer

	builder targetBuilder
	refs    RefMatcher

	githubLookup GithubLookup
}

type GithubLookup interface {
	BranchHead(context.Context, *github_pb.Commit) (string, error)
}

type targetBuilder interface {
	buildTarget(ctx context.Context, ref *github_pb.Commit, target *github_pb.DeployTargetType) error
}

func NewGithubCommandService(db sqrlx.Transactor, sm *state.StateMachines, builder targetBuilder, lookup GithubLookup) (*GithubCommandService, error) {

	refs, err := NewRefStore(db)
	if err != nil {
		return nil, err
	}

	return &GithubCommandService{
		db:            db,
		stateMachines: sm,
		builder:       builder,
		refs:          refs,
		githubLookup:  lookup,
	}, nil
}

func (ss *GithubCommandService) RegisterGRPC(srv *grpc.Server) {
	github_spb.RegisterRepoCommandServiceServer(srv, ss)
}

func (ss *GithubCommandService) ConfigureRepo(ctx context.Context, req *github_spb.ConfigureRepoRequest) (*github_spb.ConfigureRepoResponse, error) {

	action, err := j5auth.GetAuthenticatedAction(ctx)
	if err != nil {
		return nil, err
	}

	evt := &github_pb.RepoPSMEventSpec{
		Keys: &github_pb.RepoKeys{
			Owner: req.Owner,
			Name:  req.Name,
		},
		EventID:   uuid.NewString(),
		Timestamp: time.Now(),
		Action:    action,
		Event:     req.Config,
	}

	newState, err := ss.stateMachines.Repo.Transition(ctx, ss.db, evt)
	if err != nil {
		return nil, err
	}

	return &github_spb.ConfigureRepoResponse{
		Repo: newState,
	}, nil
}

func (ss *GithubCommandService) Trigger(ctx context.Context, req *github_spb.TriggerRequest) (*github_spb.TriggerResponse, error) {

	_, err := j5auth.GetAuthenticatedAction(ctx)
	if err != nil {
		return nil, err
	}

	repo, err := ss.refs.GetRepo(ctx, req.Owner, req.Repo)
	if err != nil {
		return nil, fmt.Errorf("get repo: %w", err)
	}
	if repo == nil {
		return nil, status.Error(codes.NotFound, "repo not found")
	}

	ref := &github_pb.Commit{
		Owner: repo.Keys.Owner,
		Repo:  repo.Keys.Name,
		Sha:   req.Commit,
	}

	if strings.HasPrefix(req.Commit, "refs/") {
		sha, err := ss.githubLookup.BranchHead(ctx, ref)
		if err != nil {
			return nil, fmt.Errorf("get branch head: %w", err)
		}

		ref.Sha = sha
	}

	err = ss.builder.buildTarget(ctx, ref, req.Target)
	if err != nil {
		return nil, fmt.Errorf("build targets: %w", err)
	}

	return &github_spb.TriggerResponse{}, nil
}
