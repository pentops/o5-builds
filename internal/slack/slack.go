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
	Text   string       `json:"text"`
	Blocks []SlackBlock `json:"blocks,omitempty"`
}

type SlackBlock struct {
	Type   string      `json:"type"`
	Text   *SlackText  `json:"text,omitempty"`
	Fields []SlackText `json:"fields,omitempty"`
}

type SlackText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

const (
	PlainText = "plain_text"
	Markdown  = "mrkdwn"

	TypeHeader  = "header"
	TypeSection = "section"
	TypeDivider = "divider"
)

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
		log.WithFields(ctx, map[string]any{
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
		Blocks: []SlackBlock{{
			Type: TypeHeader,
			Text: &SlackText{
				Type: PlainText,
				Text: headerText,
			},
		}, {
			Type: TypeSection,
			Fields: []SlackText{{
				Type: Markdown,
				Text: fmt.Sprintf("*Repo:*\n%s/%s", msg.Build.Commit.Owner, msg.Build.Commit.Repo),
			}, {
				Type: Markdown,
				Text: fmt.Sprintf("*Branch:*\n%s", gl.Coalesce("", msg.Build.Commit.Ref)),
			}},
		}, {
			Type: TypeSection,
			Fields: []SlackText{{
				Type: Markdown,
				Text: fmt.Sprintf("*Commit:*\n%s", msg.Build.Commit.Sha),
			}, {
				Type: Markdown,
				Text: fmt.Sprintf("*Build:*\n%s", msg.Build.Name),
			}},
		},
		}}

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

			outMsg.Blocks = append(outMsg.Blocks, SlackBlock{
				Type: TypeSection,
				Text: &SlackText{
					Type: Markdown,
					Text: txt,
				},
			})
		}
	}

	outMsg.Blocks = append(outMsg.Blocks, SlackBlock{
		Type: TypeDivider,
	})

	return outMsg
}
