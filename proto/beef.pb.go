// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: beef.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetAllBeefResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllBeefResponse) Reset() {
	*x = GetAllBeefResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beef_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBeefResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBeefResponse) ProtoMessage() {}

func (x *GetAllBeefResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beef_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBeefResponse.ProtoReflect.Descriptor instead.
func (*GetAllBeefResponse) Descriptor() ([]byte, []int) {
	return file_beef_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllBeefResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_beef_proto protoreflect.FileDescriptor

var file_beef_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x65,
	0x65, 0x66, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x65, 0x65, 0x66, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4f, 0x0a, 0x0b, 0x42, 0x65, 0x65,
	0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x42, 0x65, 0x65, 0x66, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18,
	0x2e, 0x62, 0x65, 0x65, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x65, 0x65, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x61, 0x6b, 0x6f, 0x72,
	0x6e, 0x73, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x65, 0x65, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beef_proto_rawDescOnce sync.Once
	file_beef_proto_rawDescData = file_beef_proto_rawDesc
)

func file_beef_proto_rawDescGZIP() []byte {
	file_beef_proto_rawDescOnce.Do(func() {
		file_beef_proto_rawDescData = protoimpl.X.CompressGZIP(file_beef_proto_rawDescData)
	})
	return file_beef_proto_rawDescData
}

var file_beef_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beef_proto_goTypes = []any{
	(*GetAllBeefResponse)(nil), // 0: beef.GetAllBeefResponse
	(*empty.Empty)(nil),        // 1: google.protobuf.Empty
}
var file_beef_proto_depIdxs = []int32{
	1, // 0: beef.BeefService.GetAllBeef:input_type -> google.protobuf.Empty
	0, // 1: beef.BeefService.GetAllBeef:output_type -> beef.GetAllBeefResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_beef_proto_init() }
func file_beef_proto_init() {
	if File_beef_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beef_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllBeefResponse); i {
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
			RawDescriptor: file_beef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beef_proto_goTypes,
		DependencyIndexes: file_beef_proto_depIdxs,
		MessageInfos:      file_beef_proto_msgTypes,
	}.Build()
	File_beef_proto = out.File
	file_beef_proto_rawDesc = nil
	file_beef_proto_goTypes = nil
	file_beef_proto_depIdxs = nil
}
