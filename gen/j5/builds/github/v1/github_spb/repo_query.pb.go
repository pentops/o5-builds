// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/builds/github/v1/service/repo_query.proto

package github_spb

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/pentops/j5/gen/j5/ext/v1/ext_j5pb"
	list_j5pb "github.com/pentops/j5/gen/j5/list/v1/list_j5pb"
	github_pb "github.com/pentops/o5-builds/gen/j5/builds/github/v1/github_pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRepoRequest) Reset() {
	*x = GetRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepoRequest) ProtoMessage() {}

func (x *GetRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepoRequest.ProtoReflect.Descriptor instead.
func (*GetRepoRequest) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{0}
}

func (x *GetRepoRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo *github_pb.RepoState `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *GetRepoResponse) Reset() {
	*x = GetRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepoResponse) ProtoMessage() {}

func (x *GetRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepoResponse.ProtoReflect.Descriptor instead.
func (*GetRepoResponse) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetRepoResponse) GetRepo() *github_pb.RepoState {
	if x != nil {
		return x.Repo
	}
	return nil
}

type ListReposRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *list_j5pb.PageRequest  `protobuf:"bytes,100,opt,name=page,proto3" json:"page,omitempty"`
	Query *list_j5pb.QueryRequest `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListReposRequest) Reset() {
	*x = ListReposRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposRequest) ProtoMessage() {}

func (x *ListReposRequest) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposRequest.ProtoReflect.Descriptor instead.
func (*ListReposRequest) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{2}
}

func (x *ListReposRequest) GetPage() *list_j5pb.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListReposRequest) GetQuery() *list_j5pb.QueryRequest {
	if x != nil {
		return x.Query
	}
	return nil
}

type ListReposResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repos []*github_pb.RepoState  `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
	Page  *list_j5pb.PageResponse `protobuf:"bytes,100,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListReposResponse) Reset() {
	*x = ListReposResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReposResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReposResponse) ProtoMessage() {}

func (x *ListReposResponse) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReposResponse.ProtoReflect.Descriptor instead.
func (*ListReposResponse) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{3}
}

func (x *ListReposResponse) GetRepos() []*github_pb.RepoState {
	if x != nil {
		return x.Repos
	}
	return nil
}

func (x *ListReposResponse) GetPage() *list_j5pb.PageResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListRepoEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string                  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name  string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Page  *list_j5pb.PageRequest  `protobuf:"bytes,100,opt,name=page,proto3" json:"page,omitempty"`
	Query *list_j5pb.QueryRequest `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ListRepoEventsRequest) Reset() {
	*x = ListRepoEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoEventsRequest) ProtoMessage() {}

func (x *ListRepoEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoEventsRequest.ProtoReflect.Descriptor instead.
func (*ListRepoEventsRequest) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{4}
}

func (x *ListRepoEventsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ListRepoEventsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListRepoEventsRequest) GetPage() *list_j5pb.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListRepoEventsRequest) GetQuery() *list_j5pb.QueryRequest {
	if x != nil {
		return x.Query
	}
	return nil
}

type ListRepoEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*github_pb.RepoEvent  `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	Page   *list_j5pb.PageResponse `protobuf:"bytes,100,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRepoEventsResponse) Reset() {
	*x = ListRepoEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoEventsResponse) ProtoMessage() {}

func (x *ListRepoEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_service_repo_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoEventsResponse.ProtoReflect.Descriptor instead.
func (*ListRepoEventsResponse) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP(), []int{5}
}

func (x *ListRepoEventsResponse) GetEvents() []*github_pb.RepoEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ListRepoEventsResponse) GetPage() *list_j5pb.PageResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_j5_builds_github_v1_service_repo_query_proto protoreflect.FileDescriptor

var file_j5_builds_github_v1_service_repo_query_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6a, 0x35, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6a, 0x35, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6a, 0x35, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6a, 0x35, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6a, 0x35, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x6f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6a, 0x35, 0x2e, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6a, 0x35, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x35,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x92, 0x01, 0x02, 0x10, 0x0a, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6a, 0x35, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6a,
	0x35, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6a, 0x35, 0x2e,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x88, 0x01, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x10, 0x0a,
	0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6a, 0x35, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x97, 0x04, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x2b, 0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x38, 0xc2, 0xff, 0x8e, 0x02, 0x04, 0x52, 0x02, 0x08, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x95, 0x01, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x2d, 0x2e, 0x6a, 0x35, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6a, 0x35, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0xc2, 0xff, 0x8e, 0x02, 0x04,
	0x52, 0x02, 0x10, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x12, 0xba, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6a, 0x35,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3f, 0xc2, 0xff, 0x8e, 0x02, 0x04, 0x52, 0x02, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x30, 0x12, 0x2e, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x7b, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x0d, 0xea, 0x85, 0x8f, 0x02, 0x08, 0x0a, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x35, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_j5_builds_github_v1_service_repo_query_proto_rawDescOnce sync.Once
	file_j5_builds_github_v1_service_repo_query_proto_rawDescData = file_j5_builds_github_v1_service_repo_query_proto_rawDesc
)

func file_j5_builds_github_v1_service_repo_query_proto_rawDescGZIP() []byte {
	file_j5_builds_github_v1_service_repo_query_proto_rawDescOnce.Do(func() {
		file_j5_builds_github_v1_service_repo_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_builds_github_v1_service_repo_query_proto_rawDescData)
	})
	return file_j5_builds_github_v1_service_repo_query_proto_rawDescData
}

var file_j5_builds_github_v1_service_repo_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_j5_builds_github_v1_service_repo_query_proto_goTypes = []any{
	(*GetRepoRequest)(nil),         // 0: j5.builds.github.v1.service.GetRepoRequest
	(*GetRepoResponse)(nil),        // 1: j5.builds.github.v1.service.GetRepoResponse
	(*ListReposRequest)(nil),       // 2: j5.builds.github.v1.service.ListReposRequest
	(*ListReposResponse)(nil),      // 3: j5.builds.github.v1.service.ListReposResponse
	(*ListRepoEventsRequest)(nil),  // 4: j5.builds.github.v1.service.ListRepoEventsRequest
	(*ListRepoEventsResponse)(nil), // 5: j5.builds.github.v1.service.ListRepoEventsResponse
	(*github_pb.RepoState)(nil),    // 6: j5.builds.github.v1.RepoState
	(*list_j5pb.PageRequest)(nil),  // 7: j5.list.v1.PageRequest
	(*list_j5pb.QueryRequest)(nil), // 8: j5.list.v1.QueryRequest
	(*list_j5pb.PageResponse)(nil), // 9: j5.list.v1.PageResponse
	(*github_pb.RepoEvent)(nil),    // 10: j5.builds.github.v1.RepoEvent
}
var file_j5_builds_github_v1_service_repo_query_proto_depIdxs = []int32{
	6,  // 0: j5.builds.github.v1.service.GetRepoResponse.repo:type_name -> j5.builds.github.v1.RepoState
	7,  // 1: j5.builds.github.v1.service.ListReposRequest.page:type_name -> j5.list.v1.PageRequest
	8,  // 2: j5.builds.github.v1.service.ListReposRequest.query:type_name -> j5.list.v1.QueryRequest
	6,  // 3: j5.builds.github.v1.service.ListReposResponse.repos:type_name -> j5.builds.github.v1.RepoState
	9,  // 4: j5.builds.github.v1.service.ListReposResponse.page:type_name -> j5.list.v1.PageResponse
	7,  // 5: j5.builds.github.v1.service.ListRepoEventsRequest.page:type_name -> j5.list.v1.PageRequest
	8,  // 6: j5.builds.github.v1.service.ListRepoEventsRequest.query:type_name -> j5.list.v1.QueryRequest
	10, // 7: j5.builds.github.v1.service.ListRepoEventsResponse.events:type_name -> j5.builds.github.v1.RepoEvent
	9,  // 8: j5.builds.github.v1.service.ListRepoEventsResponse.page:type_name -> j5.list.v1.PageResponse
	0,  // 9: j5.builds.github.v1.service.RepoQueryService.GetRepo:input_type -> j5.builds.github.v1.service.GetRepoRequest
	2,  // 10: j5.builds.github.v1.service.RepoQueryService.ListRepos:input_type -> j5.builds.github.v1.service.ListReposRequest
	4,  // 11: j5.builds.github.v1.service.RepoQueryService.ListRepoEvents:input_type -> j5.builds.github.v1.service.ListRepoEventsRequest
	1,  // 12: j5.builds.github.v1.service.RepoQueryService.GetRepo:output_type -> j5.builds.github.v1.service.GetRepoResponse
	3,  // 13: j5.builds.github.v1.service.RepoQueryService.ListRepos:output_type -> j5.builds.github.v1.service.ListReposResponse
	5,  // 14: j5.builds.github.v1.service.RepoQueryService.ListRepoEvents:output_type -> j5.builds.github.v1.service.ListRepoEventsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_j5_builds_github_v1_service_repo_query_proto_init() }
func file_j5_builds_github_v1_service_repo_query_proto_init() {
	if File_j5_builds_github_v1_service_repo_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRepoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetRepoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListReposRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListReposResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListRepoEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_j5_builds_github_v1_service_repo_query_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListRepoEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_j5_builds_github_v1_service_repo_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_j5_builds_github_v1_service_repo_query_proto_goTypes,
		DependencyIndexes: file_j5_builds_github_v1_service_repo_query_proto_depIdxs,
		MessageInfos:      file_j5_builds_github_v1_service_repo_query_proto_msgTypes,
	}.Build()
	File_j5_builds_github_v1_service_repo_query_proto = out.File
	file_j5_builds_github_v1_service_repo_query_proto_rawDesc = nil
	file_j5_builds_github_v1_service_repo_query_proto_goTypes = nil
	file_j5_builds_github_v1_service_repo_query_proto_depIdxs = nil
}
