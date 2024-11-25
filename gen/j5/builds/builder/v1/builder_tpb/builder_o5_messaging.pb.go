// Code generated by protoc-gen-go-o5-messaging. DO NOT EDIT.
// versions:
// - protoc-gen-go-o5-messaging 0.0.0
// source: j5/builds/builder/v1/topic/builder.proto

package builder_tpb

import (
	context "context"
	messaging_j5pb "github.com/pentops/j5/gen/j5/messaging/v1/messaging_j5pb"
	messaging_pb "github.com/pentops/o5-messaging/gen/o5/messaging/v1/messaging_pb"
	o5msg "github.com/pentops/o5-messaging/o5msg"
)

// Service: BuilderReplyTopic
// Expose Request Metadata
func (msg *BuildStatusMessage) SetJ5RequestMetadata(md *messaging_j5pb.RequestMetadata) {
	msg.Request = md
}
func (msg *BuildStatusMessage) GetJ5RequestMetadata() *messaging_j5pb.RequestMetadata {
	return msg.Request
}

type BuilderReplyTopicTxSender[C any] struct {
	sender o5msg.TxSender[C]
}

func NewBuilderReplyTopicTxSender[C any](sender o5msg.TxSender[C]) *BuilderReplyTopicTxSender[C] {
	sender.Register(o5msg.TopicDescriptor{
		Service: "j5.builds.builder.v1.topic.BuilderReplyTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "BuildStatus",
				Message: (*BuildStatusMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &BuilderReplyTopicTxSender[C]{sender: sender}
}

type BuilderReplyTopicCollector[C any] struct {
	collector o5msg.Collector[C]
}

func NewBuilderReplyTopicCollector[C any](collector o5msg.Collector[C]) *BuilderReplyTopicCollector[C] {
	collector.Register(o5msg.TopicDescriptor{
		Service: "j5.builds.builder.v1.topic.BuilderReplyTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "BuildStatus",
				Message: (*BuildStatusMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &BuilderReplyTopicCollector[C]{collector: collector}
}

type BuilderReplyTopicPublisher struct {
	publisher o5msg.Publisher
}

func NewBuilderReplyTopicPublisher(publisher o5msg.Publisher) *BuilderReplyTopicPublisher {
	publisher.Register(o5msg.TopicDescriptor{
		Service: "j5.builds.builder.v1.topic.BuilderReplyTopic",
		Methods: []o5msg.MethodDescriptor{
			{
				Name:    "BuildStatus",
				Message: (*BuildStatusMessage).ProtoReflect(nil).Descriptor(),
			},
		},
	})
	return &BuilderReplyTopicPublisher{publisher: publisher}
}

// Method: BuildStatus

func (msg *BuildStatusMessage) O5MessageHeader() o5msg.Header {
	header := o5msg.Header{
		GrpcService:      "j5.builds.builder.v1.topic.BuilderReplyTopic",
		GrpcMethod:       "BuildStatus",
		Headers:          map[string]string{},
		DestinationTopic: "build_reply",
	}
	if msg.Request != nil {
		header.Extension = &messaging_pb.Message_Reply_{
			Reply: &messaging_pb.Message_Reply{
				ReplyTo: msg.Request.ReplyTo,
			},
		}
	}
	return header
}

func (send BuilderReplyTopicTxSender[C]) BuildStatus(ctx context.Context, sendContext C, msg *BuildStatusMessage) error {
	return send.sender.Send(ctx, sendContext, msg)
}

func (collect BuilderReplyTopicCollector[C]) BuildStatus(sendContext C, msg *BuildStatusMessage) {
	collect.collector.Collect(sendContext, msg)
}

func (publish BuilderReplyTopicPublisher) BuildStatus(ctx context.Context, msg *BuildStatusMessage) error {
	return publish.publisher.Publish(ctx, msg)
}
