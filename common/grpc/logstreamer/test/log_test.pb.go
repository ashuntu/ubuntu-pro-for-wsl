// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: log_test.proto

package test

import (
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

// this is the base object to decode into for client tests, including log messages
// that are later turned as Log object.
type EmptyLogTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyLogTest) Reset() {
	*x = EmptyLogTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyLogTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyLogTest) ProtoMessage() {}

func (x *EmptyLogTest) ProtoReflect() protoreflect.Message {
	mi := &file_log_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyLogTest.ProtoReflect.Descriptor instead.
func (*EmptyLogTest) Descriptor() ([]byte, []int) {
	return file_log_test_proto_rawDescGZIP(), []int{0}
}

var File_log_test_proto protoreflect.FileDescriptor

var file_log_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x73, 0x74,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x75, 0x62, 0x75, 0x6e, 0x74, 0x75, 0x2d,
	0x70, 0x72, 0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x77, 0x73, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_test_proto_rawDescOnce sync.Once
	file_log_test_proto_rawDescData = file_log_test_proto_rawDesc
)

func file_log_test_proto_rawDescGZIP() []byte {
	file_log_test_proto_rawDescOnce.Do(func() {
		file_log_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_test_proto_rawDescData)
	})
	return file_log_test_proto_rawDescData
}

var file_log_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_log_test_proto_goTypes = []any{
	(*EmptyLogTest)(nil), // 0: EmptyLogTest
}
var file_log_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_test_proto_init() }
func file_log_test_proto_init() {
	if File_log_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_test_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyLogTest); i {
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
			RawDescriptor: file_log_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_test_proto_goTypes,
		DependencyIndexes: file_log_test_proto_depIdxs,
		MessageInfos:      file_log_test_proto_msgTypes,
	}.Build()
	File_log_test_proto = out.File
	file_log_test_proto_rawDesc = nil
	file_log_test_proto_goTypes = nil
	file_log_test_proto_depIdxs = nil
}
