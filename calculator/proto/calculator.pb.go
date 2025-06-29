// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: calculator/proto/calculator.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_calculator_proto_calculator_proto protoreflect.FileDescriptor

const file_calculator_proto_calculator_proto_rawDesc = "" +
	"\n" +
	"!calculator/proto/calculator.proto\x12\n" +
	"calculator\x1a\x1acalculator/proto/sum.proto\x1a\x1ccalculator/proto/prime.proto\x1a\x1ecalculator/proto/average.proto\x1a\x1acalculator/proto/max.proto\x1a\x1bcalculator/proto/sqrt.proto2\xc8\x02\n" +
	"\x11CalculatorService\x126\n" +
	"\x03Sum\x12\x16.calculator.SumRequest\x1a\x17.calculator.SumResponse\x12>\n" +
	"\x05Prime\x12\x18.calculator.PrimeRequest\x1a\x19.calculator.PrimeResponse0\x01\x12D\n" +
	"\aAverage\x12\x1a.calculator.AverageRequest\x1a\x1b.calculator.AverageResponse(\x01\x12:\n" +
	"\x03Max\x12\x16.calculator.MaxRequest\x1a\x17.calculator.MaxResponse(\x010\x01\x129\n" +
	"\x04Sqrt\x12\x17.calculator.SqrtRequest\x1a\x18.calculator.SqrtResponseB6Z4github.com/Akshat120/grpc-go-course/calculator/protob\x06proto3"

var file_calculator_proto_calculator_proto_goTypes = []any{
	(*SumRequest)(nil),      // 0: calculator.SumRequest
	(*PrimeRequest)(nil),    // 1: calculator.PrimeRequest
	(*AverageRequest)(nil),  // 2: calculator.AverageRequest
	(*MaxRequest)(nil),      // 3: calculator.MaxRequest
	(*SqrtRequest)(nil),     // 4: calculator.SqrtRequest
	(*SumResponse)(nil),     // 5: calculator.SumResponse
	(*PrimeResponse)(nil),   // 6: calculator.PrimeResponse
	(*AverageResponse)(nil), // 7: calculator.AverageResponse
	(*MaxResponse)(nil),     // 8: calculator.MaxResponse
	(*SqrtResponse)(nil),    // 9: calculator.SqrtResponse
}
var file_calculator_proto_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Sum:input_type -> calculator.SumRequest
	1, // 1: calculator.CalculatorService.Prime:input_type -> calculator.PrimeRequest
	2, // 2: calculator.CalculatorService.Average:input_type -> calculator.AverageRequest
	3, // 3: calculator.CalculatorService.Max:input_type -> calculator.MaxRequest
	4, // 4: calculator.CalculatorService.Sqrt:input_type -> calculator.SqrtRequest
	5, // 5: calculator.CalculatorService.Sum:output_type -> calculator.SumResponse
	6, // 6: calculator.CalculatorService.Prime:output_type -> calculator.PrimeResponse
	7, // 7: calculator.CalculatorService.Average:output_type -> calculator.AverageResponse
	8, // 8: calculator.CalculatorService.Max:output_type -> calculator.MaxResponse
	9, // 9: calculator.CalculatorService.Sqrt:output_type -> calculator.SqrtResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_calculator_proto_init() }
func file_calculator_proto_calculator_proto_init() {
	if File_calculator_proto_calculator_proto != nil {
		return
	}
	file_calculator_proto_sum_proto_init()
	file_calculator_proto_prime_proto_init()
	file_calculator_proto_average_proto_init()
	file_calculator_proto_max_proto_init()
	file_calculator_proto_sqrt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_proto_calculator_proto_rawDesc), len(file_calculator_proto_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_calculator_proto_depIdxs,
	}.Build()
	File_calculator_proto_calculator_proto = out.File
	file_calculator_proto_calculator_proto_goTypes = nil
	file_calculator_proto_calculator_proto_depIdxs = nil
}
