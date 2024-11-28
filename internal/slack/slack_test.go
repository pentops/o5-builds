package slack

import (
	"context"
	"os"
	"testing"

	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
	"github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
)

func TestSlack(t *testing.T) {
	testURL := os.Getenv("SLACK_TEST_URL")
	if testURL == "" {
		t.Skip("SLACK_TEST_URL not set")
	}

	publisher := NewPublisher(testURL)
	err := publisher.PublishBuildReport(context.Background(), &builder_pb.BuildReport{
		Build: &builder_pb.BuildContext{
			Commit: &github_pb.Commit{
				Repo:  "test-repo",
				Owner: "test-owner",
				Sha:   "test-sha",
			},
			Name: "test-name",
		},
		Status: builder_pb.BuildStatus_SUCCESS,
		Output: &builder_pb.Output{
			Title:   "test-title",
			Summary: "test-summary",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

}
