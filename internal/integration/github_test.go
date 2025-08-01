package integration

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"buf.build/go/protoyaml"
	github_api "github.com/google/go-github/v47/github"
	"github.com/google/uuid"
	"github.com/pentops/flowtest"
	"github.com/pentops/golib/gl"
	"github.com/pentops/j5/gen/j5/config/v1/config_j5pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_spb"
	"github.com/pentops/o5-builds/internal/integration/mocks"
	"github.com/pentops/o5-deploy-aws/gen/o5/aws/deployer/v1/awsdeployer_tpb"
	"github.com/pentops/o5-messaging/outbox/outboxtest"
	"github.com/pentops/realms/j5auth"
	"github.com/pentops/registry/gen/j5/registry/v1/registry_tpb"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func withTestActor(ctx context.Context) context.Context {
	jwt := map[string]any{
		"sub":                         "test/" + uuid.NewString(),
		"claims.pentops.com/tenant":   "test",
		"claims.pentops.com/tenantid": "test",
	}
	jwtJSON, err := json.Marshal(jwt)
	if err != nil {
		panic(err)
	}

	md := metadata.MD{j5auth.VerifiedJWTHeader: []string{
		string(jwtJSON),
	}}

	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

func TestO5Trigger(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	request := &awsdeployer_tpb.RequestDeploymentMessage{}
	environmentID := uuid.NewString()

	flow.Step("ConfigureRepo", func(ctx context.Context, t flowtest.Asserter) {
		ctx = withTestActor(ctx)
		res, err := uu.RepoCommand.ConfigureRepo(ctx, &github_spb.ConfigureRepoRequest{
			Owner: "owner",
			Name:  "repo",
			Config: &github_pb.RepoEventType_Configure{
				ChecksEnabled: false,
				Branches: []*github_pb.Branch{{
					BranchName: "ref1",
					DeployTargets: []*github_pb.DeployTargetType{{
						Type: &github_pb.DeployTargetType_O5Build_{
							O5Build: &github_pb.DeployTargetType_O5Build{
								Environment: environmentID,
							},
						},
					}},
				}},
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		t.Equal("owner", res.Repo.Keys.Owner)

	})

	flow.Step("O5 Build", func(ctx context.Context, t flowtest.Asserter) {
		uu.Github.TestPush("owner", "repo", mocks.GithubCommit{
			SHA: "after",
			Files: map[string]string{
				"ext/o5/app.yaml": strings.Join([]string{
					"name: appname",
				}, "\n")},
		})

		uu.GithubEvent(t, "push", &github_api.PushEvent{
			Ref: gl.Ptr("refs/heads/ref1"),
			Repo: &github_api.PushEventRepository{
				Owner: &github_api.User{
					Name: gl.Ptr("owner"),
				},
				Name: gl.Ptr("repo"),
			},
			After:  gl.Ptr("after"),
			Before: gl.Ptr("before"),
		})

		uu.Outbox.PopMessage(t, request)

		t.Equal(environmentID, request.EnvironmentId)
		t.Equal("appname", request.Application.Name)
		t.Equal("after", request.Version)

	})
}

func mustMarshal(t flowtest.TB, pb proto.Message) string {
	t.Helper()
	b, err := protoyaml.Marshal(pb)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return string(b)
}

func TestJ5Trigger(t *testing.T) {
	flow, uu := NewUniverse(t)
	defer flow.RunSteps(t)

	flow.Step("ConfigureRepo", func(ctx context.Context, t flowtest.Asserter) {
		ctx = withTestActor(ctx)
		res, err := uu.RepoCommand.ConfigureRepo(ctx, &github_spb.ConfigureRepoRequest{
			Owner: "owner",
			Name:  "repo",
			Config: &github_pb.RepoEventType_Configure{
				ChecksEnabled: true,
				Branches: []*github_pb.Branch{{
					BranchName: "ref1",
					DeployTargets: []*github_pb.DeployTargetType{{
						Type: &github_pb.DeployTargetType_J5Build_{
							J5Build: &github_pb.DeployTargetType_J5Build{},
						},
					}},
				}},
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		t.Equal("owner", res.Repo.Keys.Owner)

	})

	var buildRoot *registry_tpb.BuildAPIMessage
	flow.Step("J5 Build", func(ctx context.Context, t flowtest.Asserter) {

		uu.Github.TestPush("owner", "repo", mocks.GithubCommit{
			SHA: "after",
			Files: map[string]string{
				"j5.yaml": mustMarshal(t, &config_j5pb.RepoConfigFile{
					Bundles: []*config_j5pb.BundleReference{{
						Dir:  "proto/b1",
						Name: "bundle1",
					}, {
						Dir:  "proto/b2",
						Name: "bundle2",
					}},
					Registry: &config_j5pb.RegistryConfig{
						Owner: "owner",
						Name:  "repo",
					},
				}),
				"proto/b1/j5.yaml": mustMarshal(t, &config_j5pb.BundleConfigFile{
					Registry: nil,
				}),
				"proto/b2/j5.yaml": mustMarshal(t, &config_j5pb.BundleConfigFile{
					Registry: &config_j5pb.RegistryConfig{
						Owner: "owner",
						Name:  "repo1",
					},
				}),
			},
		})

		uu.GithubEvent(t, "push", &github_api.PushEvent{
			Ref: gl.Ptr("refs/heads/ref1"),
			Repo: &github_api.PushEventRepository{
				Owner: &github_api.User{
					Name: gl.Ptr("owner"),
				},
				Name: gl.Ptr("repo"),
			},
			After:  gl.Ptr("after"),
			Before: gl.Ptr("before"),
		})

		buildRoot = &registry_tpb.BuildAPIMessage{}
		uu.Outbox.PopMessage(t, buildRoot, outboxtest.MessageBodyMatches(func(b *registry_tpb.BuildAPIMessage) bool {
			return b.Bundle == ""
		}))
		build2 := &registry_tpb.BuildAPIMessage{}
		uu.Outbox.PopMessage(t, build2, outboxtest.MessageBodyMatches(func(b *registry_tpb.BuildAPIMessage) bool {
			t.Logf("bundle: %q", b.Bundle)
			return b.Bundle == "bundle2"
		}))

		t.NotEmpty(buildRoot.Request)
		t.NotEmpty(build2.Request)

		t.Equal("", buildRoot.Bundle)
		t.Equal("bundle2", build2.Bundle)

	})

	flow.Step("J5 Reply", func(ctx context.Context, t flowtest.Asserter) {
		t.Logf("buildAPI: %v", buildRoot.Request)
		_, err := uu.RegistryReply.J5BuildStatus(ctx, &registry_tpb.J5BuildStatusMessage{
			Request: buildRoot.Request,
			Status:  registry_tpb.BuildStatus_BUILD_STATUS_SUCCESS,
		})
		t.NoError(err)

		gotStatus := uu.Github.CheckRunUpdates
		if len(gotStatus) != 1 {
			t.Fatalf("unexpected number of check runs: %d", len(gotStatus))
		}
		got := gotStatus[0]

		t.Equal("owner", got.Build.GithubCheckRun.CheckSuite.Commit.Owner)
		t.Equal("repo", got.Build.GithubCheckRun.CheckSuite.Commit.Repo)

	})
}
