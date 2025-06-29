// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: calculator/proto/sqrt.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SqrtRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int32                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SqrtRequest) Reset() {
	*x = SqrtRequest{}
	mi := &file_calculator_proto_sqrt_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SqrtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtRequest) ProtoMessage() {}

func (x *SqrtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_sqrt_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtRequest.ProtoReflect.Descriptor instead.
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_sqrt_proto_rawDescGZIP(), []int{0}
}

func (x *SqrtRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SqrtResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SqrtResponse) Reset() {
	*x = SqrtResponse{}
	mi := &file_calculator_proto_sqrt_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SqrtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SqrtResponse) ProtoMessage() {}

func (x *SqrtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_sqrt_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SqrtResponse.ProtoReflect.Descriptor instead.
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_sqrt_proto_rawDescGZIP(), []int{1}
}

func (x *SqrtResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto_sqrt_proto protoreflect.FileDescriptor

const file_calculator_proto_sqrt_proto_rawDesc = "" +
	"\n" +
	"\x1bcalculator/proto/sqrt.proto\x12\n" +
	"calculator\"%\n" +
	"\vSqrtRequest\x12\x16\n" +
	"\x06number\x18\x01 \x01(\x05R\x06number\"&\n" +
	"\fSqrtResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06resultB6Z4github.com/Akshat120/grpc-go-course/calculator/protob\x06proto3"

var (
	file_calculator_proto_sqrt_proto_rawDescOnce sync.Once
	file_calculator_proto_sqrt_proto_rawDescData []byte
)

func file_calculator_proto_sqrt_proto_rawDescGZIP() []byte {
	file_calculator_proto_sqrt_proto_rawDescOnce.Do(func() {
		file_calculator_proto_sqrt_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_proto_sqrt_proto_rawDesc), len(file_calculator_proto_sqrt_proto_rawDesc)))
	})
	return file_calculator_proto_sqrt_proto_rawDescData
}

var file_calculator_proto_sqrt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_sqrt_proto_goTypes = []any{
	(*SqrtRequest)(nil),  // 0: calculator.SqrtRequest
	(*SqrtResponse)(nil), // 1: calculator.SqrtResponse
}
var file_calculator_proto_sqrt_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_sqrt_proto_init() }
func file_calculator_proto_sqrt_proto_init() {
	if File_calculator_proto_sqrt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_proto_sqrt_proto_rawDesc), len(file_calculator_proto_sqrt_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_calculator_proto_sqrt_proto_goTypes,
		DependencyIndexes: file_calculator_proto_sqrt_proto_depIdxs,
		MessageInfos:      file_calculator_proto_sqrt_proto_msgTypes,
	}.Build()
	File_calculator_proto_sqrt_proto = out.File
	file_calculator_proto_sqrt_proto_goTypes = nil
	file_calculator_proto_sqrt_proto_depIdxs = nil
}
