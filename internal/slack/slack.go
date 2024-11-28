package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pentops/golib/gl"
	"github.com/pentops/log.go/log"
	"github.com/pentops/o5-builds/gen/j5/builds/builder/v1/builder_pb"
)

type SlackMessage struct {
	Text   string                   `json:"text"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"`
}

type Publisher struct {
	URL string
}

func NewPublisher(url string) *Publisher {
	return &Publisher{
		URL: url,
	}
}

func (ss *Publisher) Send(ctx context.Context, msg *SlackMessage) error {
	json, err := json.Marshal(msg)
	if err != nil {
		log.WithError(ctx, err).Error("Couldn't convert dead letter to slack message")
		msg.Text = fmt.Sprintf("<Error json.Marshal SlackMessage: %s>", err)
	}
	res, err := http.Post(ss.URL, "application/json", bytes.NewReader([]byte(json)))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		log.WithFields(ctx, map[string]interface{}{
			"status": res.StatusCode,
			"req":    string(json),
		}).Error("Failed to send slack message")
		return fmt.Errorf("got status code %d: %s", res.StatusCode, string(body))
	}
	return nil
}

func (ss *Publisher) PublishBuildReport(ctx context.Context, msg *builder_pb.BuildReport) error {
	out := buildReport(msg)
	return ss.Send(ctx, out)
}

var statusText = map[builder_pb.BuildStatus]string{
	builder_pb.BuildStatus_PENDING:  "Pending",
	builder_pb.BuildStatus_PROGRESS: "Progress",
	builder_pb.BuildStatus_FAILURE:  "Failed",
	builder_pb.BuildStatus_SUCCESS:  "Success",
}

func buildReport(msg *builder_pb.BuildReport) *SlackMessage {

	headerText, ok := statusText[msg.Status]
	if !ok {
		headerText = msg.Status.ShortString()
	}

	headerText = fmt.Sprintf("%s %s: %s", msg.Build.Commit.Repo, msg.Build.Name, headerText)

	outMsg := &SlackMessage{
		Blocks: []map[string]interface{}{{
			"type": "header",
			"text": map[string]interface{}{
				"type": "plain_text",
				"text": headerText,
			},
		}, {
			"type": "section",
			"fields": []map[string]interface{}{{
				"type": "mrkdwn",
				"text": fmt.Sprintf("*Repo:*\n%s/%s", msg.Build.Commit.Owner, msg.Build.Commit.Repo),
			}, {
				"type": "mrkdwn",
				"text": fmt.Sprintf("*Branch:*\n%s", gl.Coalesce("", msg.Build.Commit.Ref)),
			}},
		}, {
			"type": "section",
			"fields": []map[string]interface{}{{
				"type": "mrkdwn",
				"text": fmt.Sprintf("*Commit:*\n%s", msg.Build.Commit.Sha),
			}, {
				"type": "mrkdwn",
				"text": fmt.Sprintf("*Build:*\n%s", msg.Build.Name),
			}},
		}},
	}

	if msg.Output != nil {
		lines := []string{}
		if msg.Output.Title != "" {
			lines = append(lines, fmt.Sprintf("*%s*", msg.Output.Title))
		}
		if msg.Output.Summary != "" {
			lines = append(lines, msg.Output.Summary)
		} else if msg.Output.Text != nil {
			lines = append(lines, *msg.Output.Text)
		}

		if len(lines) > 0 {
			txt := strings.Join(lines, "\n")

			outMsg.Blocks = append(outMsg.Blocks, map[string]interface{}{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": txt,
				},
			})
		}
	}

	return outMsg
}
