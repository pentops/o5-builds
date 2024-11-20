package state

import (
	"fmt"

	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/protostate/psm"
)

type StateMachines struct {
	Repo *github_pb.RepoPSM
}

func NewStateMachines() (*StateMachines, error) {
	repo, err := NewRepoPSM()
	if err != nil {
		return nil, fmt.Errorf("NewDeploymentEventer: %w", err)
	}

	return &StateMachines{
		Repo: repo,
	}, nil
}

func (sm *StateMachines) TableSpecs() []psm.QueryTableSpec {
	return []psm.QueryTableSpec{
		sm.Repo.StateTableSpec(),
	}
}
