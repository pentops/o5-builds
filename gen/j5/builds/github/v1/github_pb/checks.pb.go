// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: j5/builds/github/v1/checks.proto

package github_pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckSuite *CheckSuite `protobuf:"bytes,1,opt,name=check_suite,json=checkSuite,proto3" json:"check_suite,omitempty"`
	CheckName  string      `protobuf:"bytes,3,opt,name=check_name,json=checkName,proto3" json:"check_name,omitempty"`
	CheckId    int64       `protobuf:"varint,4,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
}

func (x *CheckRun) Reset() {
	*x = CheckRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_checks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRun) ProtoMessage() {}

func (x *CheckRun) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_checks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRun.ProtoReflect.Descriptor instead.
func (*CheckRun) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_checks_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRun) GetCheckSuite() *CheckSuite {
	if x != nil {
		return x.CheckSuite
	}
	return nil
}

func (x *CheckRun) GetCheckName() string {
	if x != nil {
		return x.CheckName
	}
	return ""
}

func (x *CheckRun) GetCheckId() int64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type CheckSuite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckSuiteId int64   `protobuf:"varint,1,opt,name=check_suite_id,json=checkSuiteId,proto3" json:"check_suite_id,omitempty"`
	Branch       string  `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	Commit       *Commit `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *CheckSuite) Reset() {
	*x = CheckSuite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_checks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckSuite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckSuite) ProtoMessage() {}

func (x *CheckSuite) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_checks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckSuite.ProtoReflect.Descriptor instead.
func (*CheckSuite) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_checks_proto_rawDescGZIP(), []int{1}
}

func (x *CheckSuite) GetCheckSuiteId() int64 {
	if x != nil {
		return x.CheckSuiteId
	}
	return 0
}

func (x *CheckSuite) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *CheckSuite) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner string  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo  string  `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Sha   string  `protobuf:"bytes,3,opt,name=sha,proto3" json:"sha,omitempty"`
	Ref   *string `protobuf:"bytes,4,opt,name=ref,proto3,oneof" json:"ref,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_j5_builds_github_v1_checks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_j5_builds_github_v1_checks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_j5_builds_github_v1_checks_proto_rawDescGZIP(), []int{2}
}

func (x *Commit) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Commit) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Commit) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *Commit) GetRef() string {
	if x != nil && x.Ref != nil {
		return *x.Ref
	}
	return ""
}

var File_j5_builds_github_v1_checks_proto protoreflect.FileDescriptor

var file_j5_builds_github_v1_checks_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6a, 0x35, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x75, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x73, 0x75, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6a, 0x35,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x69, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x64, 0x22, 0x7f, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x75, 0x69, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x33,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6a, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x22, 0x63, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x15, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x72, 0x65, 0x66, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x72, 0x65, 0x66, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f,
	0x35, 0x2d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6a, 0x35, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_j5_builds_github_v1_checks_proto_rawDescOnce sync.Once
	file_j5_builds_github_v1_checks_proto_rawDescData = file_j5_builds_github_v1_checks_proto_rawDesc
)

func file_j5_builds_github_v1_checks_proto_rawDescGZIP() []byte {
	file_j5_builds_github_v1_checks_proto_rawDescOnce.Do(func() {
		file_j5_builds_github_v1_checks_proto_rawDescData = protoimpl.X.CompressGZIP(file_j5_builds_github_v1_checks_proto_rawDescData)
	})
	return file_j5_builds_github_v1_checks_proto_rawDescData
}

var file_j5_builds_github_v1_checks_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_j5_builds_github_v1_checks_proto_goTypes = []any{
	(*CheckRun)(nil),   // 0: j5.builds.github.v1.CheckRun
	(*CheckSuite)(nil), // 1: j5.builds.github.v1.CheckSuite
	(*Commit)(nil),     // 2: j5.builds.github.v1.Commit
}
var file_j5_builds_github_v1_checks_proto_depIdxs = []int32{
	1, // 0: j5.builds.github.v1.CheckRun.check_suite:type_name -> j5.builds.github.v1.CheckSuite
	2, // 1: j5.builds.github.v1.CheckSuite.commit:type_name -> j5.builds.github.v1.Commit
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_j5_builds_github_v1_checks_proto_init() }
func file_j5_builds_github_v1_checks_proto_init() {
	if File_j5_builds_github_v1_checks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_j5_builds_github_v1_checks_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CheckRun); i {
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
		file_j5_builds_github_v1_checks_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CheckSuite); i {
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
		file_j5_builds_github_v1_checks_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Commit); i {
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
	file_j5_builds_github_v1_checks_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_j5_builds_github_v1_checks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_j5_builds_github_v1_checks_proto_goTypes,
		DependencyIndexes: file_j5_builds_github_v1_checks_proto_depIdxs,
		MessageInfos:      file_j5_builds_github_v1_checks_proto_msgTypes,
	}.Build()
	File_j5_builds_github_v1_checks_proto = out.File
	file_j5_builds_github_v1_checks_proto_rawDesc = nil
	file_j5_builds_github_v1_checks_proto_goTypes = nil
	file_j5_builds_github_v1_checks_proto_depIdxs = nil
}
