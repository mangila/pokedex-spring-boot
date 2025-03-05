// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: image_converter.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloWorld struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hello         string                 `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloWorld) Reset() {
	*x = HelloWorld{}
	mi := &file_image_converter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorld) ProtoMessage() {}

func (x *HelloWorld) ProtoReflect() protoreflect.Message {
	mi := &file_image_converter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorld.ProtoReflect.Descriptor instead.
func (*HelloWorld) Descriptor() ([]byte, []int) {
	return file_image_converter_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorld) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

var File_image_converter_proto protoreflect.FileDescriptor

var file_image_converter_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x44, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x12, 0x0b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x50,
	0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x61, 0x6e,
	0x67, 0x69, 0x6c, 0x61, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x64, 0x65, 0x78, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x50, 0x01, 0x5a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_image_converter_proto_rawDescOnce sync.Once
	file_image_converter_proto_rawDescData []byte
)

func file_image_converter_proto_rawDescGZIP() []byte {
	file_image_converter_proto_rawDescOnce.Do(func() {
		file_image_converter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_image_converter_proto_rawDesc), len(file_image_converter_proto_rawDesc)))
	})
	return file_image_converter_proto_rawDescData
}

var file_image_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_image_converter_proto_goTypes = []any{
	(*HelloWorld)(nil),             // 0: HelloWorld
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_image_converter_proto_depIdxs = []int32{
	0, // 0: ConvertImage.Convert:input_type -> HelloWorld
	1, // 1: ConvertImage.Convert:output_type -> google.protobuf.StringValue
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_image_converter_proto_init() }
func file_image_converter_proto_init() {
	if File_image_converter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_image_converter_proto_rawDesc), len(file_image_converter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_converter_proto_goTypes,
		DependencyIndexes: file_image_converter_proto_depIdxs,
		MessageInfos:      file_image_converter_proto_msgTypes,
	}.Build()
	File_image_converter_proto = out.File
	file_image_converter_proto_goTypes = nil
	file_image_converter_proto_depIdxs = nil
}
