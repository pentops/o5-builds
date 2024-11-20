package service

import (
	"context"
	"fmt"

	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_tpb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/o5-builds/internal/github"
	"github.com/pentops/o5-deploy-aws/gen/o5/aws/deployer/v1/awsdeployer_tpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ReplyWorker struct {
	github IClient
	builder_tpb.UnimplementedBuilderReplyTopicServer
	awsdeployer_tpb.UnimplementedDeploymentReplyTopicServer
}

func NewReplyWorker(githubClient IClient) *ReplyWorker {
	return &ReplyWorker{
		github: githubClient,
	}
}

func (rw *ReplyWorker) RegisterGRPC(srv *grpc.Server) {
	builder_tpb.RegisterBuilderReplyTopicServer(srv, rw)
	awsdeployer_tpb.RegisterDeploymentReplyTopicServer(srv, rw)
}

func (ww *WebhookWorker) BuildStatus(ctx context.Context, message *builder_tpb.BuildStatusMessage) (*emptypb.Empty, error) {

	checkContext := &github_pb.CheckRun{}
	err := protojson.Unmarshal(message.Request.Context, checkContext)
	if err != nil {
		return nil, fmt.Errorf("unmarshal check context: %w", err)
	}

	status := github.CheckRunUpdate{}

	switch message.Status {
	case builder_tpb.BuildStatus_IN_PROGRESS:
		status.Status = github.CheckRunStatusInProgress

	case builder_tpb.BuildStatus_FAILURE:
		status.Status = github.CheckRunStatusCompleted
		status.Conclusion = some(github.CheckRunConclusionFailure)

	case builder_tpb.BuildStatus_SUCCESS:
		status.Status = github.CheckRunStatusCompleted
		status.Conclusion = some(github.CheckRunConclusionSuccess)
	}

	log.WithFields(ctx, map[string]interface{}{
		"gh-status":  message.Status,
		"gh-outcome": message.Output,
	}).Debug("BuildStatus")

	if message.Output != nil {
		status.Output = &github.CheckRunOutput{
			Title:   message.Output.Title,
			Summary: message.Output.Summary,
			Text:    message.Output.Text,
		}
	}

	if err := ww.github.UpdateCheckRun(ctx, checkContext, status); err != nil {
		return nil, fmt.Errorf("update check run: %w", err)
	}

	return &emptypb.Empty{}, nil
}
