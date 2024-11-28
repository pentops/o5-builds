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
)

type GithubCommandService struct {
	db sqrlx.Transactor

	stateMachines *state.StateMachines
	*github_spb.UnimplementedRepoCommandServiceServer

	builder *GithubHandler
	refs    RefMatcher

	githubLookup GithubLookup
}

type GithubLookup interface {
	BranchHead(context.Context, *github_pb.Commit) (string, error)
}

func NewGithubCommandService(db sqrlx.Transactor, sm *state.StateMachines, builder *GithubHandler, lookup GithubLookup) (*GithubCommandService, error) {

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

	ref := &github_pb.Commit{
		Owner: req.Owner,
		Repo:  req.Repo,
		Sha:   req.Commit,
	}

	if strings.HasPrefix(req.Commit, "refs/") {
		sha, err := ss.githubLookup.BranchHead(ctx, ref)
		if err != nil {
			return nil, fmt.Errorf("get branch head: %w", err)
		}

		ref.Sha = sha
	}

	buildMessages, err := ss.builder.buildTargets(ctx, ref, []*github_pb.DeployTargetType{req.Target})
	if err != nil {
		return nil, err
	}

	err = ss.builder.addBuildContext(ctx, ref, buildMessages, false)
	if err != nil {
		return nil, fmt.Errorf("add build context: %w", err)
	}

	for _, msg := range buildMessages {
		err := ss.builder.publisher.Publish(ctx, msg.message)
		if err != nil {
			return nil, fmt.Errorf("publish: %w", err)
		}
	}

	err = ss.builder.buildTarget(ctx, ref, req.Target)
	if err != nil {
		return nil, fmt.Errorf("build targets: %w", err)
	}

	targets := make([]string, len(buildMessages))
	for i, msg := range buildMessages {
		targets[i] = msg.label
	}

	return &github_spb.TriggerResponse{
		Targets: targets,
	}, nil
}
