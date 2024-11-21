package app

import (
	"context"
	"fmt"

	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_tpb"
	"github.com/pentops/o5-deploy-aws/gen/o5/aws/deployer/v1/awsdeployer_tpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReplyWorker struct {
	publishers []IBuildPublisher

	builder_tpb.UnimplementedBuilderReplyTopicServer
	awsdeployer_tpb.UnimplementedDeploymentReplyTopicServer
}

type IBuildPublisher interface {
	PublishBuildReport(ctx context.Context, message *builder_pb.BuildReport) error
}

func NewReplyWorker(publishers ...IBuildPublisher) (*ReplyWorker, error) {
	return &ReplyWorker{
		publishers: publishers,
	}, nil
}

func (rw *ReplyWorker) RegisterGRPC(srv *grpc.Server) {
	builder_tpb.RegisterBuilderReplyTopicServer(srv, rw)
	awsdeployer_tpb.RegisterDeploymentReplyTopicServer(srv, rw)
}

func (ww *ReplyWorker) BuildStatus(ctx context.Context, message *builder_tpb.BuildStatusMessage) (*emptypb.Empty, error) {

	log.WithFields(ctx, map[string]interface{}{
		"gh-status":  message.Status,
		"gh-outcome": message.Output,
	}).Debug("BuildStatus")

	buildContext := &builder_pb.BuildContext{}
	err := protojson.Unmarshal(message.Request.Context, buildContext)
	if err != nil {
		return nil, fmt.Errorf("unmarshal check context: %w", err)
	}

	rep := &builder_pb.BuildReport{
		Build:  buildContext,
		Status: message.Status,
		Output: message.Output,
	}

	for _, publisher := range ww.publishers {
		if err := publisher.PublishBuildReport(ctx, rep); err != nil {
			return nil, fmt.Errorf("publish build report: %w", err)
		}
	}

	return &emptypb.Empty{}, nil
}
