package mocks

import (
	"context"
	"fmt"
	"math/rand"

	"buf.build/go/protoyaml"
	"github.com/pentops/j5/gen/j5/source/v1/source_j5pb"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	"google.golang.org/protobuf/proto"
)

type GithubMock struct {
	Repos map[string]GithubRepo

	CheckRunUpdates []*builder_pb.BuildReport
}

type GithubRepo struct {
	Commits map[string]GithubCommit
}

type GithubCommit struct {
	SHA   string
	info  *source_j5pb.CommitInfo
	Files map[string]string
}

func NewGithubMock() *GithubMock {
	return &GithubMock{
		Repos: make(map[string]GithubRepo),
	}
}

func (gh *GithubMock) TestPush(owner, name string, commit GithubCommit, refs ...string) {
	r, ok := gh.Repos[owner+"/"+name]
	if !ok {
		r = GithubRepo{
			Commits: make(map[string]GithubCommit),
		}
		gh.Repos[owner+"/"+name] = r
	}

	commit.info = &source_j5pb.CommitInfo{
		Owner: owner,
		Repo:  name,
		Hash:  commit.SHA,
	}

	commit.info.Aliases = append(commit.info.Aliases, refs...)

	r.Commits[commit.SHA] = commit
}

func (gh *GithubMock) PullConfig(ctx context.Context, ref *github_pb.Commit, into proto.Message, tryPaths []string) error {
	repo, ok := gh.Repos[ref.Owner+"/"+ref.Repo]
	if !ok {
		return fmt.Errorf("repo not found")
	}

	commit, ok := repo.Commits[ref.Sha]
	if !ok {
		return fmt.Errorf("commit not found")
	}

	for _, path := range tryPaths {
		data, ok := commit.Files[path]
		if !ok {
			continue
		}

		if err := protoyaml.Unmarshal([]byte(data), into); err != nil {
			return fmt.Errorf("unmarshalling yaml: %s", err)
		}

		return nil
	}

	return fmt.Errorf("no config found")
}

func (gh *GithubMock) GetCommit(ctx context.Context, ref *github_pb.Commit) (*source_j5pb.CommitInfo, error) {
	repoName := ref.Owner + "/" + ref.Repo
	repo, ok := gh.Repos[repoName]
	if !ok {
		return nil, fmt.Errorf("repo '%s' not found", repoName)
	}

	commit, ok := repo.Commits[ref.Sha]
	if !ok {
		return nil, fmt.Errorf("ref '%s' not found", ref.Sha)
	}
	return commit.info, nil
}

func (gh *GithubMock) CreateCheckRun(ctx context.Context, ref *github_pb.Commit, name string, report *builder_pb.BuildReport) (*github_pb.CheckRun, error) {
	return &github_pb.CheckRun{
		CheckSuite: &github_pb.CheckSuite{
			Commit:       ref,
			CheckSuiteId: rand.Int63(),
			Branch:       "main",
		},
		CheckName: name,
		CheckId:   rand.Int63(),
	}, nil
}

func (gh *GithubMock) PublishBuildReport(ctx context.Context, status *builder_pb.BuildReport) error {
	gh.CheckRunUpdates = append(gh.CheckRunUpdates, status)
	return nil
}
