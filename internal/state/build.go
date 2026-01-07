package state

import (
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
)

func NewBuildPSM() (*builder_pb.BuildPSM, error) {
	sm, err := builder_pb.BuildPSMBuilder().
		BuildStateMachine()
	if err != nil {
		return nil, err
	}

	return sm, nil
}
