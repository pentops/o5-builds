package app

import (
	"context"
	"fmt"

	"github.com/pentops/j5/lib/j5codec"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
	"github.com/pentops/o5-deploy-aws/gen/o5/aws/deployer/v1/awsdeployer_tpb"
	"github.com/pentops/registry/gen/j5/registry/v1/registry_tpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReplyWorker struct {
	publishers []IBuildPublisher

	registry_tpb.UnimplementedBuildReplyTopicServer
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
	registry_tpb.RegisterBuildReplyTopicServer(srv, rw)
	awsdeployer_tpb.RegisterDeploymentReplyTopicServer(srv, rw)
}

var j5StatusMap = map[registry_tpb.BuildStatus]builder_pb.BuildStatus{
	registry_tpb.BuildStatus_BUILD_STATUS_FAILURE:     builder_pb.BuildStatus_BUILD_STATUS_FAILURE,
	registry_tpb.BuildStatus_BUILD_STATUS_IN_PROGRESS: builder_pb.BuildStatus_BUILD_STATUS_PROGRESS,
	registry_tpb.BuildStatus_BUILD_STATUS_SUCCESS:     builder_pb.BuildStatus_BUILD_STATUS_SUCCESS,
}

func (ww *ReplyWorker) J5BuildStatus(ctx context.Context, message *registry_tpb.J5BuildStatusMessage) (*emptypb.Empty, error) {

	log.WithFields(ctx, map[string]any{
		"gh-status":  message.Status,
		"gh-outcome": message.Output,
	}).Debug("BuildStatus")

	if message.Request == nil || len(message.Request.Context) == 0 {
		log.Debug(ctx, "no build context, NOP")
		return &emptypb.Empty{}, nil
	}

	buildContext := &builder_pb.BuildContext{}
	err := j5codec.Global.JSONToProto(message.Request.Context, buildContext.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("unmarshal check context: %w", err)
	}

	status, ok := j5StatusMap[message.Status]
	if !ok {
		return nil, fmt.Errorf("unknown status: %v", message.Status)
	}

	rep := &builder_pb.BuildReport{
		Build:  buildContext,
		Status: status,
	}
	if message.Output != nil {
		rep.Output = &builder_pb.Output{
			Title:   message.Output.Title,
			Summary: message.Output.Summary,
			Text:    message.Output.Text,
		}
	}

	for _, publisher := range ww.publishers {
		if err := publisher.PublishBuildReport(ctx, rep); err != nil {
			return nil, fmt.Errorf("publish build report: %w", err)
		}
	}

	return &emptypb.Empty{}, nil
}

var o5StatusMap = map[awsdeployer_tpb.DeploymentStatus]builder_pb.BuildStatus{
	awsdeployer_tpb.DeploymentStatus_FAILED:      builder_pb.BuildStatus_BUILD_STATUS_FAILURE,
	awsdeployer_tpb.DeploymentStatus_IN_PROGRESS: builder_pb.BuildStatus_BUILD_STATUS_PROGRESS,
	awsdeployer_tpb.DeploymentStatus_PENDING:     builder_pb.BuildStatus_BUILD_STATUS_PENDING,
	awsdeployer_tpb.DeploymentStatus_SUCCESS:     builder_pb.BuildStatus_BUILD_STATUS_SUCCESS,
}

func (ww *ReplyWorker) DeploymentStatus(ctx context.Context, message *awsdeployer_tpb.DeploymentStatusMessage) (*emptypb.Empty, error) {

	log.WithFields(ctx, map[string]any{
		"gh-status": message.Status,
	}).Debug("BuildStatus")

	if message.Request == nil || len(message.Request.Context) == 0 {
		log.Debug(ctx, "no build context, NOP")
		return &emptypb.Empty{}, nil
	}

	buildContext := &builder_pb.BuildContext{}
	err := j5codec.Global.JSONToProto(message.Request.Context, buildContext.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("unmarshal check context: %w", err)
	}

	mappedStatus, ok := o5StatusMap[message.Status]
	if !ok {
		return nil, fmt.Errorf("unknown status: %v", message.Status)
	}

	rep := &builder_pb.BuildReport{
		Build:  buildContext,
		Status: mappedStatus,
	}
	if message.Message != "" {
		rep.Output = &builder_pb.Output{
			Title:   "Detail",
			Summary: message.Message,
		}
	} else {
		rep.Output = &builder_pb.Output{
			Title: fmt.Sprintf("Status: %s", message.Status.ShortString()),
		}
	}

	for _, publisher := range ww.publishers {
		if err := publisher.PublishBuildReport(ctx, rep); err != nil {
			return nil, fmt.Errorf("publish build report: %w", err)
		}
	}

	return &emptypb.Empty{}, nil
}
