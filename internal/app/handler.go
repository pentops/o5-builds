package app

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/google/uuid"
	"github.com/pentops/j5/gen/j5/config/v1/config_j5pb"
	"github.com/pentops/j5/gen/j5/messaging/v1/messaging_j5pb"
	"github.com/pentops/j5/gen/j5/source/v1/source_j5pb"
	"github.com/pentops/j5/lib/j5codec"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"github.com/pentops/o5-builds/internal/github"
	"github.com/pentops/o5-deploy-aws/gen/o5/application/v1/application_pb"
	"github.com/pentops/o5-deploy-aws/gen/o5/aws/deployer/v1/awsdeployer_tpb"
	"github.com/pentops/o5-messaging/gen/o5/messaging/v1/messaging_tpb"
	"github.com/pentops/o5-messaging/o5msg"
	"github.com/pentops/registry/gen/j5/registry/v1/registry_tpb"
	"google.golang.org/protobuf/proto"
)

type IClient interface {
	PullConfig(ctx context.Context, ref *github_pb.Commit, into proto.Message, tryPaths []string) error
	GetCommit(ctx context.Context, ref *github_pb.Commit) (*source_j5pb.CommitInfo, error)
	CreateCheckRun(ctx context.Context, ref *github_pb.Commit, name string, status *builder_pb.BuildReport) (*github_pb.CheckRun, error)
	PublishBuildReport(ctx context.Context, status *builder_pb.BuildReport) error
	BranchHead(ctx context.Context, ref *github_pb.Commit) (string, error)
}

type RefMatcher interface {
	GetRepo(ctx context.Context, owner string, name string) (*github_pb.RepoState, error)
}

type GithubHandler struct {
	github    IClient
	refs      RefMatcher
	publisher Publisher

	messaging_tpb.UnimplementedRawMessageTopicServer
}

type Publisher interface {
	Publish(ctx context.Context, msg o5msg.Message) error
}

func NewGithubHandler(refs RefMatcher, githubClient IClient, publisher Publisher) (*GithubHandler, error) {
	return &GithubHandler{
		github:    githubClient,
		publisher: publisher,
		refs:      refs,
	}, nil
}

func (ww *GithubHandler) Push(ctx context.Context, event *github.PushEvent) error {
	ctx = log.WithFields(ctx, map[string]any{
		"owner":     event.Commit.Owner,
		"repo":      event.Commit.Repo,
		"commitSha": event.Commit.Sha,
		"ref":       event.Commit.Ref,
	})
	log.Debug(ctx, "Push")

	if event.Commit.Ref == nil || !strings.HasPrefix(*event.Commit.Ref, "refs/heads/") {
		log.Info(ctx, "Not a branch push, nothing to do")
	}

	branchName := strings.TrimPrefix(*event.Commit.Ref, "refs/heads/")
	return ww.kickOffChecks(ctx, event.Commit, branchName)
}

func (ww *GithubHandler) CheckSuite(ctx context.Context, event *github.CheckSuiteEvent) error {
	ctx = log.WithFields(ctx, map[string]any{
		"owner":   event.CheckSuite.Commit.Owner,
		"repo":    event.CheckSuite.Commit.Repo,
		"branch":  event.CheckSuite.Branch,
		"commit":  event.CheckSuite.Commit.Sha,
		"suiteId": event.CheckSuite.CheckSuiteId,
	})
	log.Debug(ctx, "CheckSuite")

	switch event.Action {
	case "requested", "rerequested":
		return ww.kickOffChecks(ctx, event.CheckSuite.Commit, event.CheckSuite.Branch)
	}
	return nil
}

func (ww *GithubHandler) kickOffChecks(ctx context.Context, commit *github_pb.Commit, branchName string) error {
	buildTargets, repo, err := ww.buildTasksForBranch(ctx, commit, branchName)
	if err != nil {
		if repo == nil || !repo.Data.ChecksEnabled {
			return err
		}
		checkRunError := &CheckRunError{}
		if !errors.As(err, checkRunError) {
			return err
		}
		_, err = ww.github.CreateCheckRun(ctx, commit, checkRunError.RunName, &builder_pb.BuildReport{
			Output: &builder_pb.Output{
				Title:   checkRunError.Title,
				Summary: checkRunError.Summary,
			},
			Status: builder_pb.BuildStatus_BUILD_STATUS_FAILURE,
		})
		if err != nil {
			return fmt.Errorf("create check run: %w", err)
		}
		return nil
	}
	if repo == nil {
		log.Info(ctx, "No repo config, nothing to do")
		return nil
	}

	if len(buildTargets) < 1 {
		log.Info(ctx, "No build targets, nothing to do")
		return nil
	}

	if err := ww.addBuildContext(ctx, commit, buildTargets, repo.Data.ChecksEnabled); err != nil {
		return err
	}

	if err := ww.publishTasks(ctx, buildTargets); err != nil {
		return err
	}

	return nil
}

func (ww *GithubHandler) publishTasks(ctx context.Context, tasks []*buildTask) error {
	for _, task := range tasks {
		err := ww.publisher.Publish(ctx, task.message)
		if err != nil {
			return fmt.Errorf("publish: %w", err)
		}
	}
	return nil
}

func (ww *GithubHandler) addBuildContext(ctx context.Context, commit *github_pb.Commit, tasks []*buildTask, runGithubChecks bool) error {
	for _, task := range tasks {
		cc := &builder_pb.BuildContext{
			Commit: commit,
			Name:   task.label,
		}

		if runGithubChecks {
			checkRun, err := ww.github.CreateCheckRun(ctx, commit, task.uniqueName, nil)
			if err != nil {
				return fmt.Errorf("create check run: %w", err)
			}
			cc.GithubCheckRun = checkRun
		}

		contextData, err := j5codec.Global.ProtoToJSON(cc.ProtoReflect())
		if err != nil {
			return fmt.Errorf("marshal check run: %w", err)
		}
		task.message.SetJ5RequestMetadata(&messaging_j5pb.RequestMetadata{
			Context: contextData,
		})
	}
	return nil
}

func (ww *GithubHandler) buildTasksForBranch(ctx context.Context, commit *github_pb.Commit, branchName string) ([]*buildTask, *github_pb.RepoState, error) {
	repo, err := ww.refs.GetRepo(ctx, commit.Owner, commit.Repo)
	if err != nil {
		return nil, nil, fmt.Errorf("get repo: %w", err)
	}
	if repo == nil {
		log.Info(ctx, "No repo config, nothing to do")
		return nil, nil, nil
	}

	targets := make([]*github_pb.DeployTargetType, 0, len(repo.Data.Branches))
	for _, target := range repo.Data.Branches {
		if target.BranchName == branchName || target.BranchName == "*" {
			targets = append(targets, target.DeployTargets...)
		}
	}

	if len(targets) < 1 {
		log.Info(ctx, "No deploy targets, nothing to do")
		return nil, repo, nil
	}

	t2, err := ww.buildTargets(ctx, commit, targets)
	if err != nil {
		return nil, repo, fmt.Errorf("build targets: %w", err)
	}
	return t2, repo, nil
}

type taskMessage interface {
	o5msg.Message
	SetJ5RequestMetadata(*messaging_j5pb.RequestMetadata)
}

type buildTask struct {
	uniqueName string
	label      string
	message    taskMessage
}

func (ww *GithubHandler) buildTarget(ctx context.Context, commit *github_pb.Commit, target *github_pb.DeployTargetType) error {

	buildMessages, err := ww.buildTargets(ctx, commit, []*github_pb.DeployTargetType{target})
	if err != nil {
		return err
	}

	for _, msg := range buildMessages {
		err := ww.publisher.Publish(ctx, msg.message)
		if err != nil {
			return fmt.Errorf("publish: %w", err)
		}
	}

	return nil
}

func (ww *GithubHandler) buildTargets(ctx context.Context, commit *github_pb.Commit, targets []*github_pb.DeployTargetType) ([]*buildTask, error) {

	o5Envs := []string{}
	j5Build := false

	buildMessages := []*buildTask{}

	for _, target := range targets {
		switch target := target.Type.(type) {
		case *github_pb.DeployTargetType_O5Build_:
			o5Envs = append(o5Envs, target.O5Build.Environment)

		case *github_pb.DeployTargetType_J5Build_:
			j5Build = true
		default:
			return nil, fmt.Errorf("unknown target type: %T", target)
		}
	}

	if j5Build {
		builds, err := ww.j5Build(ctx, commit)
		if err != nil {
			return nil, err

		}
		for _, apiBuild := range builds.APIBuilds {
			buildMessages = append(buildMessages, &buildTask{
				uniqueName: "j5-image",
				label:      "J5 Image",
				message:    apiBuild,
			})
		}

		for _, protoBuild := range builds.ProtoBuilds {
			buildMessages = append(buildMessages, &buildTask{
				uniqueName: fmt.Sprintf("j5-proto-%s", protoBuild.Name),
				message:    protoBuild,
			})
		}
	}

	if len(o5Envs) > 0 {
		builds, err := ww.o5Build(ctx, commit, o5Envs)
		if err != nil {
			return nil, fmt.Errorf("o5 build: %w", err)
		}

		for _, build := range builds {
			buildMessages = append(buildMessages, &buildTask{
				uniqueName: fmt.Sprintf("o5-deploy-%s", build.EnvironmentId),
				label:      "O5 Deploy",
				message:    build,
			})
		}
	}

	return buildMessages, nil
}

var o5ConfigPaths = []string{
	"ext/o5/app.yaml",
	"ext/o5/app.yml",
	"o5.yaml",
	"o5.yml",
}

func (ww *GithubHandler) o5Build(ctx context.Context, commit *github_pb.Commit, targetEnvs []string) ([]*awsdeployer_tpb.RequestDeploymentMessage, *CheckRunError) {
	cfg := &application_pb.Application{}
	err := ww.github.PullConfig(ctx, commit, cfg, o5ConfigPaths)
	if err != nil {
		return nil, &CheckRunError{
			RunName: "o5-config",
			Title:   "o5 config error",
			Summary: err.Error(),
		}
	}

	triggers := make([]*awsdeployer_tpb.RequestDeploymentMessage, 0, len(targetEnvs))

	for _, envID := range targetEnvs {
		triggers = append(triggers, &awsdeployer_tpb.RequestDeploymentMessage{
			DeploymentId:  uuid.NewString(),
			Application:   cfg,
			Version:       commit.Sha,
			EnvironmentId: envID,
		})
	}

	return triggers, nil
}

type CheckRunError struct {
	RunName string
	Title   string
	Summary string
}

func (e CheckRunError) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Summary)
}

type j5Buildset struct {
	APIBuilds   []*registry_tpb.BuildAPIMessage
	ProtoBuilds []*registry_tpb.PublishMessage
}

var configPaths = []string{
	"j5.yaml",
	"j5.repo.yaml",
	"ext/j5/j5.yaml",
}

var bundleConfigPaths = []string{
	"j5.yaml",
	"j5.bundle.yaml",
}

func (ww *GithubHandler) j5Build(ctx context.Context, commit *github_pb.Commit) (*j5Buildset, error) {

	commitInfo, err := ww.github.GetCommit(ctx, commit)
	if err != nil {
		return nil, fmt.Errorf("get commit: %w", err)
	}

	commit.Sha = commitInfo.Hash

	cfg := &config_j5pb.RepoConfigFile{}
	err = ww.github.PullConfig(ctx, commit, cfg, configPaths)
	if err != nil {
		log.WithError(ctx, err).Error("Config Error")
		return nil, &CheckRunError{
			RunName: "j5-config",
			Title:   "j5 config error",
			Summary: err.Error(),
		}
	}

	if cfg.Git != nil {
		github.ExpandGitAliases(cfg.Git.Main, commitInfo)
	}

	type namedBundle struct {
		name     string
		registry *config_j5pb.RegistryConfig
		publish  []*config_j5pb.PublishConfig
	}

	bundles := make([]namedBundle, 0, len(cfg.Bundles)+1)
	for _, bundle := range cfg.Bundles {
		bundleConfig := &config_j5pb.BundleConfigFile{}
		paths := make([]string, 0, len(bundleConfigPaths))
		for _, configPath := range bundleConfigPaths {
			paths = append(paths, path.Join(bundle.Dir, configPath))
		}
		if err := ww.github.PullConfig(ctx, commit, bundleConfig, paths); err != nil {
			return nil, &CheckRunError{
				RunName: "j5-config",
				Title:   "j5 bundle config error",
				Summary: fmt.Sprintf("Pulling %s/j5.yaml: %s", bundle.Dir, err.Error()),
			}
		}
		bundles = append(bundles, namedBundle{
			registry: bundleConfig.Registry,
			name:     bundle.Name,
			publish:  bundleConfig.Publish,
		})
	}

	if cfg.Registry != nil {
		// root is also a bundle.
		bundles = append(bundles, namedBundle{
			name:     "",
			registry: cfg.Registry,
			publish:  cfg.Publish,
		})

	}

	output := &j5Buildset{}

	for _, bundle := range bundles {
		if bundle.registry != nil {
			req := &registry_tpb.BuildAPIMessage{
				Commit: commitInfo,
				Bundle: bundle.name,
			}
			output.APIBuilds = append(output.APIBuilds, req)
		}

		for _, publish := range bundle.publish {
			req := &registry_tpb.PublishMessage{
				Commit: commitInfo,
				Name:   publish.Name,
				Bundle: bundle.name,
			}
			output.ProtoBuilds = append(output.ProtoBuilds, req)
		}
	}

	return output, nil
}
