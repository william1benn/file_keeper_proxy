// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: proto/s3utils.proto

package proto

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

type S3FileStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	FileData []byte `protobuf:"bytes,2,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
}

func (x *S3FileStoreRequest) Reset() {
	*x = S3FileStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_s3utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3FileStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3FileStoreRequest) ProtoMessage() {}

func (x *S3FileStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_s3utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3FileStoreRequest.ProtoReflect.Descriptor instead.
func (*S3FileStoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_s3utils_proto_rawDescGZIP(), []int{0}
}

func (x *S3FileStoreRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *S3FileStoreRequest) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

type S3FileStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *S3FileStoreResponse) Reset() {
	*x = S3FileStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_s3utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3FileStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3FileStoreResponse) ProtoMessage() {}

func (x *S3FileStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_s3utils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3FileStoreResponse.ProtoReflect.Descriptor instead.
func (*S3FileStoreResponse) Descriptor() ([]byte, []int) {
	return file_proto_s3utils_proto_rawDescGZIP(), []int{1}
}

func (x *S3FileStoreResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type S3GetFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *S3GetFileData) Reset() {
	*x = S3GetFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_s3utils_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3GetFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3GetFileData) ProtoMessage() {}

func (x *S3GetFileData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_s3utils_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3GetFileData.ProtoReflect.Descriptor instead.
func (*S3GetFileData) Descriptor() ([]byte, []int) {
	return file_proto_s3utils_proto_rawDescGZIP(), []int{2}
}

func (x *S3GetFileData) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type S3ResponseFileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData []byte `protobuf:"bytes,1,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
}

func (x *S3ResponseFileData) Reset() {
	*x = S3ResponseFileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_s3utils_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3ResponseFileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3ResponseFileData) ProtoMessage() {}

func (x *S3ResponseFileData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_s3utils_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3ResponseFileData.ProtoReflect.Descriptor instead.
func (*S3ResponseFileData) Descriptor() ([]byte, []int) {
	return file_proto_s3utils_proto_rawDescGZIP(), []int{3}
}

func (x *S3ResponseFileData) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

var File_proto_s3utils_proto protoreflect.FileDescriptor

var file_proto_s3utils_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x33, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x62, 0x22, 0x43, 0x0a, 0x12, 0x53, 0x33, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x33, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x33, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x12, 0x53, 0x33, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0xaf, 0x01, 0x0a, 0x09, 0x53,
	0x33, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x33, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x46, 0x69, 0x6c, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x33, 0x12, 0x1b, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x33, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_s3utils_proto_rawDescOnce sync.Once
	file_proto_s3utils_proto_rawDescData = file_proto_s3utils_proto_rawDesc
)

func file_proto_s3utils_proto_rawDescGZIP() []byte {
	file_proto_s3utils_proto_rawDescOnce.Do(func() {
		file_proto_s3utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_s3utils_proto_rawDescData)
	})
	return file_proto_s3utils_proto_rawDescData
}

var file_proto_s3utils_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_s3utils_proto_goTypes = []interface{}{
	(*S3FileStoreRequest)(nil),  // 0: messaging_pb.S3FileStoreRequest
	(*S3FileStoreResponse)(nil), // 1: messaging_pb.S3FileStoreResponse
	(*S3GetFileData)(nil),       // 2: messaging_pb.S3GetFileData
	(*S3ResponseFileData)(nil),  // 3: messaging_pb.S3ResponseFileData
}
var file_proto_s3utils_proto_depIdxs = []int32{
	0, // 0: messaging_pb.S3Service.StoreFileS3:input_type -> messaging_pb.S3FileStoreRequest
	2, // 1: messaging_pb.S3Service.GetFileS3:input_type -> messaging_pb.S3GetFileData
	1, // 2: messaging_pb.S3Service.StoreFileS3:output_type -> messaging_pb.S3FileStoreResponse
	3, // 3: messaging_pb.S3Service.GetFileS3:output_type -> messaging_pb.S3ResponseFileData
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_s3utils_proto_init() }
func file_proto_s3utils_proto_init() {
	if File_proto_s3utils_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_s3utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3FileStoreRequest); i {
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
		file_proto_s3utils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3FileStoreResponse); i {
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
		file_proto_s3utils_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3GetFileData); i {
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
		file_proto_s3utils_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3ResponseFileData); i {
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
			RawDescriptor: file_proto_s3utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_s3utils_proto_goTypes,
		DependencyIndexes: file_proto_s3utils_proto_depIdxs,
		MessageInfos:      file_proto_s3utils_proto_msgTypes,
	}.Build()
	File_proto_s3utils_proto = out.File
	file_proto_s3utils_proto_rawDesc = nil
	file_proto_s3utils_proto_goTypes = nil
	file_proto_s3utils_proto_depIdxs = nil
}
