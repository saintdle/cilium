// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.7
// source: envoy/type/hash_policy.proto

package _type

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Specifies the hash policy
type HashPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PolicySpecifier:
	//
	//	*HashPolicy_SourceIp_
	PolicySpecifier isHashPolicy_PolicySpecifier `protobuf_oneof:"policy_specifier"`
}

func (x *HashPolicy) Reset() {
	*x = HashPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_hash_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashPolicy) ProtoMessage() {}

func (x *HashPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_hash_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashPolicy.ProtoReflect.Descriptor instead.
func (*HashPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_type_hash_policy_proto_rawDescGZIP(), []int{0}
}

func (m *HashPolicy) GetPolicySpecifier() isHashPolicy_PolicySpecifier {
	if m != nil {
		return m.PolicySpecifier
	}
	return nil
}

func (x *HashPolicy) GetSourceIp() *HashPolicy_SourceIp {
	if x, ok := x.GetPolicySpecifier().(*HashPolicy_SourceIp_); ok {
		return x.SourceIp
	}
	return nil
}

type isHashPolicy_PolicySpecifier interface {
	isHashPolicy_PolicySpecifier()
}

type HashPolicy_SourceIp_ struct {
	SourceIp *HashPolicy_SourceIp `protobuf:"bytes,1,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

func (*HashPolicy_SourceIp_) isHashPolicy_PolicySpecifier() {}

// The source IP will be used to compute the hash used by hash-based load balancing
// algorithms.
type HashPolicy_SourceIp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HashPolicy_SourceIp) Reset() {
	*x = HashPolicy_SourceIp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_hash_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashPolicy_SourceIp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashPolicy_SourceIp) ProtoMessage() {}

func (x *HashPolicy_SourceIp) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_hash_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashPolicy_SourceIp.ProtoReflect.Descriptor instead.
func (*HashPolicy_SourceIp) Descriptor() ([]byte, []int) {
	return file_envoy_type_hash_policy_proto_rawDescGZIP(), []int{0, 0}
}

var File_envoy_type_hash_policy_proto protoreflect.FileDescriptor

var file_envoy_type_hash_policy_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0a, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x3e, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x70, 0x48, 0x00, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70,
	0x1a, 0x0a, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x42, 0x17, 0x0a, 0x10,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x68, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0f, 0x48, 0x61, 0x73, 0x68, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_hash_policy_proto_rawDescOnce sync.Once
	file_envoy_type_hash_policy_proto_rawDescData = file_envoy_type_hash_policy_proto_rawDesc
)

func file_envoy_type_hash_policy_proto_rawDescGZIP() []byte {
	file_envoy_type_hash_policy_proto_rawDescOnce.Do(func() {
		file_envoy_type_hash_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_hash_policy_proto_rawDescData)
	})
	return file_envoy_type_hash_policy_proto_rawDescData
}

var file_envoy_type_hash_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_type_hash_policy_proto_goTypes = []interface{}{
	(*HashPolicy)(nil),          // 0: envoy.type.HashPolicy
	(*HashPolicy_SourceIp)(nil), // 1: envoy.type.HashPolicy.SourceIp
}
var file_envoy_type_hash_policy_proto_depIdxs = []int32{
	1, // 0: envoy.type.HashPolicy.source_ip:type_name -> envoy.type.HashPolicy.SourceIp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_type_hash_policy_proto_init() }
func file_envoy_type_hash_policy_proto_init() {
	if File_envoy_type_hash_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_hash_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashPolicy); i {
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
		file_envoy_type_hash_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashPolicy_SourceIp); i {
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
	file_envoy_type_hash_policy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HashPolicy_SourceIp_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_hash_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_hash_policy_proto_goTypes,
		DependencyIndexes: file_envoy_type_hash_policy_proto_depIdxs,
		MessageInfos:      file_envoy_type_hash_policy_proto_msgTypes,
	}.Build()
	File_envoy_type_hash_policy_proto = out.File
	file_envoy_type_hash_policy_proto_rawDesc = nil
	file_envoy_type_hash_policy_proto_goTypes = nil
	file_envoy_type_hash_policy_proto_depIdxs = nil
}
