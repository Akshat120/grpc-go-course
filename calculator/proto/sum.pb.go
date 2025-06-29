// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: calculator/proto/sum.proto

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

type SumRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             int32                  `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B             int32                  `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	mi := &file_calculator_proto_sum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_sum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_sum_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SumRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        int32                  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	mi := &file_calculator_proto_sum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_sum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_sum_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto_sum_proto protoreflect.FileDescriptor

const file_calculator_proto_sum_proto_rawDesc = "" +
	"\n" +
	"\x1acalculator/proto/sum.proto\x12\n" +
	"calculator\"(\n" +
	"\n" +
	"SumRequest\x12\f\n" +
	"\x01a\x18\x01 \x01(\x05R\x01a\x12\f\n" +
	"\x01b\x18\x02 \x01(\x05R\x01b\"%\n" +
	"\vSumResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x05R\x06resultB6Z4github.com/Akshat120/grpc-go-course/calculator/protob\x06proto3"

var (
	file_calculator_proto_sum_proto_rawDescOnce sync.Once
	file_calculator_proto_sum_proto_rawDescData []byte
)

func file_calculator_proto_sum_proto_rawDescGZIP() []byte {
	file_calculator_proto_sum_proto_rawDescOnce.Do(func() {
		file_calculator_proto_sum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_proto_sum_proto_rawDesc), len(file_calculator_proto_sum_proto_rawDesc)))
	})
	return file_calculator_proto_sum_proto_rawDescData
}

var file_calculator_proto_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_sum_proto_goTypes = []any{
	(*SumRequest)(nil),  // 0: calculator.SumRequest
	(*SumResponse)(nil), // 1: calculator.SumResponse
}
var file_calculator_proto_sum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_sum_proto_init() }
func file_calculator_proto_sum_proto_init() {
	if File_calculator_proto_sum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_proto_sum_proto_rawDesc), len(file_calculator_proto_sum_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_calculator_proto_sum_proto_goTypes,
		DependencyIndexes: file_calculator_proto_sum_proto_depIdxs,
		MessageInfos:      file_calculator_proto_sum_proto_msgTypes,
	}.Build()
	File_calculator_proto_sum_proto = out.File
	file_calculator_proto_sum_proto_goTypes = nil
	file_calculator_proto_sum_proto_depIdxs = nil
}
