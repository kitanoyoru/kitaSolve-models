// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/kitasolve_models.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Solve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Solve) Reset() {
	*x = Solve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kitasolve_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solve) ProtoMessage() {}

func (x *Solve) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kitasolve_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solve.ProtoReflect.Descriptor instead.
func (*Solve) Descriptor() ([]byte, []int) {
	return file_proto_kitasolve_models_proto_rawDescGZIP(), []int{0}
}

func (x *Solve) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Solve) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Solve) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_proto_kitasolve_models_proto protoreflect.FileDescriptor

var file_proto_kitasolve_models_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x53, 0x0a, 0x05, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xf0, 0x02, 0x0a, 0x0f, 0x6b, 0x69, 0x74, 0x61, 0x53, 0x6f,
	0x6c, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x61, 0x64, 0x64,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x1a, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x41, 0x0a, 0x08,
	0x67, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x17, 0x2e, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12,
	0x47, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x17, 0x2e, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x30, 0x01,
	0x12, 0x46, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12,
	0x17, 0x2e, 0x6b, 0x69, 0x74, 0x61, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kitasolve_models_proto_rawDescOnce sync.Once
	file_proto_kitasolve_models_proto_rawDescData = file_proto_kitasolve_models_proto_rawDesc
)

func file_proto_kitasolve_models_proto_rawDescGZIP() []byte {
	file_proto_kitasolve_models_proto_rawDescOnce.Do(func() {
		file_proto_kitasolve_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kitasolve_models_proto_rawDescData)
	})
	return file_proto_kitasolve_models_proto_rawDescData
}

var file_proto_kitasolve_models_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_kitasolve_models_proto_goTypes = []interface{}{
	(*Solve)(nil),                  // 0: kitasolve_models.Solve
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.BoolValue)(nil),   // 2: google.protobuf.BoolValue
}
var file_proto_kitasolve_models_proto_depIdxs = []int32{
	0, // 0: kitasolve_models.kitaSolveModels.addSolve:input_type -> kitasolve_models.Solve
	1, // 1: kitasolve_models.kitaSolveModels.getSolve:input_type -> google.protobuf.StringValue
	1, // 2: kitasolve_models.kitaSolveModels.deleteSolve:input_type -> google.protobuf.StringValue
	1, // 3: kitasolve_models.kitaSolveModels.searchSolve:input_type -> google.protobuf.StringValue
	0, // 4: kitasolve_models.kitaSolveModels.updateSolve:input_type -> kitasolve_models.Solve
	1, // 5: kitasolve_models.kitaSolveModels.addSolve:output_type -> google.protobuf.StringValue
	0, // 6: kitasolve_models.kitaSolveModels.getSolve:output_type -> kitasolve_models.Solve
	2, // 7: kitasolve_models.kitaSolveModels.deleteSolve:output_type -> google.protobuf.BoolValue
	0, // 8: kitasolve_models.kitaSolveModels.searchSolve:output_type -> kitasolve_models.Solve
	1, // 9: kitasolve_models.kitaSolveModels.updateSolve:output_type -> google.protobuf.StringValue
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kitasolve_models_proto_init() }
func file_proto_kitasolve_models_proto_init() {
	if File_proto_kitasolve_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kitasolve_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solve); i {
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
			RawDescriptor: file_proto_kitasolve_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kitasolve_models_proto_goTypes,
		DependencyIndexes: file_proto_kitasolve_models_proto_depIdxs,
		MessageInfos:      file_proto_kitasolve_models_proto_msgTypes,
	}.Build()
	File_proto_kitasolve_models_proto = out.File
	file_proto_kitasolve_models_proto_rawDesc = nil
	file_proto_kitasolve_models_proto_goTypes = nil
	file_proto_kitasolve_models_proto_depIdxs = nil
}
