// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: api/proto/cutter.proto

package cutter

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

type AddURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url            string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	CustomEndpoint string `protobuf:"bytes,2,opt,name=customEndpoint,proto3" json:"customEndpoint,omitempty"`
}

func (x *AddURLRequest) Reset() {
	*x = AddURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cutter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddURLRequest) ProtoMessage() {}

func (x *AddURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cutter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddURLRequest.ProtoReflect.Descriptor instead.
func (*AddURLRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_cutter_proto_rawDescGZIP(), []int{0}
}

func (x *AddURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddURLRequest) GetCustomEndpoint() string {
	if x != nil {
		return x.CustomEndpoint
	}
	return ""
}

type AddURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Url     *ShortenedURL `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *AddURLResponse) Reset() {
	*x = AddURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cutter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddURLResponse) ProtoMessage() {}

func (x *AddURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cutter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddURLResponse.ProtoReflect.Descriptor instead.
func (*AddURLResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_cutter_proto_rawDescGZIP(), []int{1}
}

func (x *AddURLResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AddURLResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddURLResponse) GetUrl() *ShortenedURL {
	if x != nil {
		return x.Url
	}
	return nil
}

type ShortenedURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalURL  string `protobuf:"bytes,1,opt,name=originalURL,proto3" json:"originalURL,omitempty"`
	ShortenedURL string `protobuf:"bytes,2,opt,name=shortenedURL,proto3" json:"shortenedURL,omitempty"`
}

func (x *ShortenedURL) Reset() {
	*x = ShortenedURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cutter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenedURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenedURL) ProtoMessage() {}

func (x *ShortenedURL) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cutter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenedURL.ProtoReflect.Descriptor instead.
func (*ShortenedURL) Descriptor() ([]byte, []int) {
	return file_api_proto_cutter_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenedURL) GetOriginalURL() string {
	if x != nil {
		return x.OriginalURL
	}
	return ""
}

func (x *ShortenedURL) GetShortenedURL() string {
	if x != nil {
		return x.ShortenedURL
	}
	return ""
}

type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cutter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cutter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_cutter_proto_rawDescGZIP(), []int{3}
}

func (x *GetURLRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url []*ShortenedURL `protobuf:"bytes,1,rep,name=url,proto3" json:"url,omitempty"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_cutter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_cutter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_cutter_proto_rawDescGZIP(), []int{4}
}

func (x *GetURLResponse) GetUrl() []*ShortenedURL {
	if x != nil {
		return x.Url
	}
	return nil
}

var File_api_proto_cutter_proto protoreflect.FileDescriptor

var file_api_proto_cutter_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x49, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x67, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x54, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x52,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x52, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64,
	0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x22, 0x21, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x35, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x32, 0x78, 0x0a, 0x0c, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x12, 0x33, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c,
	0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x63, 0x75, 0x74, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_cutter_proto_rawDescOnce sync.Once
	file_api_proto_cutter_proto_rawDescData = file_api_proto_cutter_proto_rawDesc
)

func file_api_proto_cutter_proto_rawDescGZIP() []byte {
	file_api_proto_cutter_proto_rawDescOnce.Do(func() {
		file_api_proto_cutter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_cutter_proto_rawDescData)
	})
	return file_api_proto_cutter_proto_rawDescData
}

var file_api_proto_cutter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_cutter_proto_goTypes = []interface{}{
	(*AddURLRequest)(nil),  // 0: api.AddURLRequest
	(*AddURLResponse)(nil), // 1: api.AddURLResponse
	(*ShortenedURL)(nil),   // 2: api.ShortenedURL
	(*GetURLRequest)(nil),  // 3: api.GetURLRequest
	(*GetURLResponse)(nil), // 4: api.GetURLResponse
}
var file_api_proto_cutter_proto_depIdxs = []int32{
	2, // 0: api.AddURLResponse.url:type_name -> api.ShortenedURL
	2, // 1: api.GetURLResponse.url:type_name -> api.ShortenedURL
	0, // 2: api.URLShortener.AddURL:input_type -> api.AddURLRequest
	3, // 3: api.URLShortener.GetURL:input_type -> api.GetURLRequest
	1, // 4: api.URLShortener.AddURL:output_type -> api.AddURLResponse
	4, // 5: api.URLShortener.GetURL:output_type -> api.GetURLResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_cutter_proto_init() }
func file_api_proto_cutter_proto_init() {
	if File_api_proto_cutter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_cutter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddURLRequest); i {
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
		file_api_proto_cutter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddURLResponse); i {
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
		file_api_proto_cutter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenedURL); i {
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
		file_api_proto_cutter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLRequest); i {
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
		file_api_proto_cutter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLResponse); i {
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
			RawDescriptor: file_api_proto_cutter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_cutter_proto_goTypes,
		DependencyIndexes: file_api_proto_cutter_proto_depIdxs,
		MessageInfos:      file_api_proto_cutter_proto_msgTypes,
	}.Build()
	File_api_proto_cutter_proto = out.File
	file_api_proto_cutter_proto_rawDesc = nil
	file_api_proto_cutter_proto_goTypes = nil
	file_api_proto_cutter_proto_depIdxs = nil
}
