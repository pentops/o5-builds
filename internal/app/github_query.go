package app

import (
	"github.com/pentops/j5/lib/psm"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_spb"
	"github.com/pentops/o5-builds/internal/state"
	"github.com/pentops/sqrlx.go/sqrlx"
	"google.golang.org/grpc"
)

type GithubQueryService struct {
	db sqrlx.Transactor

	impl *github_spb.RepoQueryServiceImpl
}

func NewGithubQueryService(db sqrlx.Transactor, states *state.StateMachines) (*GithubQueryService, error) {

	querySpec := github_spb.DefaultRepoPSMQuerySpec(states.Repo.StateTableSpec())
	querySet, err := github_spb.NewRepoPSMQuerySet(querySpec, psm.StateQueryOptions{})
	if err != nil {
		return nil, err
	}
	queryImpl := github_spb.NewRepoQueryServiceImpl(db, querySet)

	return &GithubQueryService{
		impl: queryImpl,
	}, nil
}

func (ds *GithubQueryService) RegisterGRPC(srv *grpc.Server) {
	github_spb.RegisterRepoQueryServiceServer(srv, ds.impl)
}
