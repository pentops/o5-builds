package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/go-github/v58/github"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/o5-messaging/gen/o5/messaging/v1/messaging_tpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PushEvent struct {
	Commit *github_pb.Commit
	Ref    string
}

type CheckSuiteEvent struct {
	Commit     *github_pb.Commit
	CheckSuite *github_pb.CheckSuite
	Action     string
}

type CheckRunEvent struct {
	CheckRun *github_pb.CheckRun
	Action   string
}

type WebhookHandler interface {
	Push(ctx context.Context, event *PushEvent) error
	CheckSuite(ctx context.Context, event *CheckSuiteEvent) error
	CheckRun(ctx context.Context, event *CheckRunEvent) error
}

const emptyCommit = "0000000000000000000000000000000000000000"

type WebhookWorker struct {
	handler WebhookHandler
	messaging_tpb.UnimplementedRawMessageTopicServer
}

func NewWebhookWorker(handler WebhookHandler) (*WebhookWorker, error) {
	return &WebhookWorker{
		handler: handler,
	}, nil
}

func (ww *WebhookWorker) RegisterGRPC(srv *grpc.Server) {
	messaging_tpb.RegisterRawMessageTopicServer(srv, ww)
}

func (ww *WebhookWorker) RawMessage(ctx context.Context, message *messaging_tpb.RawMessage) (*emptypb.Empty, error) {

	parts := strings.SplitN(message.Topic, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid topic")
	}
	if parts[0] != "github" {
		return nil, errors.New("not a github message")
	}

	msgType := parts[1]

	payload, err := github.ParseWebHook(msgType, message.Payload)
	if err != nil {
		return nil, fmt.Errorf("parse webhook: %w", err)
	}

	switch event := payload.(type) {
	case *github.PushEvent:
		return &empty.Empty{}, ww.handlePushEvent(ctx, event)

	case *github.CheckSuiteEvent:
		return &empty.Empty{}, ww.handleCheckSuiteEvent(ctx, event)

	case *github.CheckRunEvent:
		return &empty.Empty{}, ww.handleCheckRunEvent(ctx, event)

	default:
		return nil, nil
	}
}

func validateRepo(repo *github.Repository) error {
	if repo == nil {
		return fmt.Errorf("nil 'repo' on check_run event")
	}
	if repo.Owner == nil {
		return fmt.Errorf("nil 'repo.owner' on check_run event")
	}
	if repo.Owner.Login == nil {
		return fmt.Errorf("nil 'repo.owner.login' on check_run event")
	}
	if repo.Name == nil {
		return fmt.Errorf("nil 'repo.name' on check_run event")
	}
	return nil
}

func (ww *WebhookWorker) handleCheckSuiteEvent(ctx context.Context, event *github.CheckSuiteEvent) error {
	if event.Action == nil {
		return fmt.Errorf("nil 'action' on check_suite event")
	}
	if err := validateRepo(event.Repo); err != nil {
		return err
	}
	if event.CheckSuite == nil {
		return fmt.Errorf("nil 'check_suite' on check_suite event")
	}
	if event.CheckSuite.HeadBranch == nil {
		return fmt.Errorf("nil 'check_suite.head_branch' on check_suite event")
	}
	if event.CheckSuite.BeforeSHA == nil {
		return fmt.Errorf("nil 'check_suite.before_sha' on check_suite event")
	}
	if event.CheckSuite.AfterSHA == nil {
		return fmt.Errorf("nil 'check_suite.after_sha' on check_suite event")
	}

	msg := &CheckSuiteEvent{
		Action: *event.Action,
		CheckSuite: &github_pb.CheckSuite{
			Commit: &github_pb.Commit{
				Owner: *event.Repo.Owner.Login,
				Repo:  *event.Repo.Name,
				Sha:   *event.CheckSuite.AfterSHA,
			},
			Branch:       *event.CheckSuite.HeadBranch,
			CheckSuiteId: *event.CheckSuite.ID,
		},
	}

	return ww.handler.CheckSuite(ctx, msg)
}

func (ww *WebhookWorker) handleCheckRunEvent(ctx context.Context, event *github.CheckRunEvent) error {

	if event.Action == nil {
		return fmt.Errorf("nil 'action' on check_run event")
	}
	if err := validateRepo(event.Repo); err != nil {
		return err
	}
	if event.CheckRun == nil {
		return fmt.Errorf("nil 'check_run' on check_run event")
	}
	if event.CheckRun.Name == nil {
		return fmt.Errorf("nil 'check_run.name' on check_run event")
	}
	if event.CheckRun.ID == nil {
		return fmt.Errorf("nil 'check_run.id' on check_run event")
	}

	if event.CheckRun.CheckSuite == nil {
		return fmt.Errorf("nil 'check_run.check_suite' on check_run event")
	}
	if event.CheckRun.CheckSuite.HeadBranch == nil {
		return fmt.Errorf("nil 'check_run.check_suite.head_branch' on check_run event")
	}
	if event.CheckRun.CheckSuite.BeforeSHA == nil {
		return fmt.Errorf("nil 'check_run.check_suite.before_sha' on check_run event")
	}
	if event.CheckRun.CheckSuite.AfterSHA == nil {
		return fmt.Errorf("nil 'check_run.check_suite.after_sha' on check_run event")
	}

	msg := &CheckRunEvent{
		Action: *event.Action,
		CheckRun: &github_pb.CheckRun{
			CheckName: *event.CheckRun.Name,
			CheckId:   *event.CheckRun.ID,
			CheckSuite: &github_pb.CheckSuite{
				Commit: &github_pb.Commit{
					Owner: *event.Repo.Owner.Login,
					Repo:  *event.Repo.Name,
					Sha:   *event.CheckRun.CheckSuite.AfterSHA,
				},
				Branch:       *event.CheckRun.CheckSuite.HeadBranch,
				CheckSuiteId: *event.CheckRun.CheckSuite.ID,
			},
		},
	}

	return ww.handler.CheckRun(ctx, msg)
}

func (ww *WebhookWorker) handlePushEvent(ctx context.Context, event *github.PushEvent) error {

	if event.Ref == nil {
		return fmt.Errorf("nil 'ref' on push event")
	}

	if event.Repo == nil {
		return fmt.Errorf("nil 'repo' on check_run event")
	}

	if event.Repo.Owner == nil {
		return fmt.Errorf("nil 'repo.owner' on check_run event")
	}

	if event.Repo.Owner.Name == nil {
		return fmt.Errorf("nil 'repo.owner.name' on check_run event")
	}

	if event.Repo.Name == nil {
		return fmt.Errorf("nil 'repo.name' on check_run event")
	}

	if event.After == nil {
		return fmt.Errorf("nil 'after' on push event")
	}

	if event.Before == nil {
		return fmt.Errorf("nil 'before' on push event")
	}

	if *event.After == emptyCommit {
		return nil
	}

	// Send Message to SNS
	msg := &PushEvent{
		Commit: &github_pb.Commit{
			Sha:   *event.After,
			Repo:  *event.Repo.Name,
			Owner: *event.Repo.Owner.Name,
		},
		Ref: *event.Ref,
	}

	return ww.handler.Push(ctx, msg)
}
