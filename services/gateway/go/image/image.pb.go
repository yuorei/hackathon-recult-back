// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: image.proto

package image

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

// ファイルメタ情報
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Meta) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 最初のリクエストはmeta 2回目以降のリクエストはdata
	//
	// Types that are assignable to Value:
	//
	//	*UploadRequest_Meta
	//	*UploadRequest_Data
	Value isUploadRequest_Value `protobuf_oneof:"value"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (m *UploadRequest) GetValue() isUploadRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *UploadRequest) GetMeta() *Meta {
	if x, ok := x.GetValue().(*UploadRequest_Meta); ok {
		return x.Meta
	}
	return nil
}

func (x *UploadRequest) GetData() []byte {
	if x, ok := x.GetValue().(*UploadRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isUploadRequest_Value interface {
	isUploadRequest_Value()
}

type UploadRequest_Meta struct {
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3,oneof"`
}

type UploadRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*UploadRequest_Meta) isUploadRequest_Value() {}

func (*UploadRequest_Data) isUploadRequest_Value() {}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x22, 0x51, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x32, 0x46, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x14, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_image_proto_goTypes = []interface{}{
	(*Meta)(nil),           // 0: image.Meta
	(*UploadRequest)(nil),  // 1: image.UploadRequest
	(*UploadResponse)(nil), // 2: image.UploadResponse
}
var file_image_proto_depIdxs = []int32{
	0, // 0: image.UploadRequest.meta:type_name -> image.Meta
	1, // 1: image.UploadImage.Upload:input_type -> image.UploadRequest
	2, // 2: image.UploadImage.Upload:output_type -> image.UploadResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
	file_image_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UploadRequest_Meta)(nil),
		(*UploadRequest_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}
