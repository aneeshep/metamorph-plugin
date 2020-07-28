// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: proto/metamorph.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_metamorph_proto_rawDescGZIP(), []int{0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_metamorph_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type ResponseBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseBool) Reset() {
	*x = ResponseBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBool) ProtoMessage() {}

func (x *ResponseBool) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBool.ProtoReflect.Descriptor instead.
func (*ResponseBool) Descriptor() ([]byte, []int) {
	return file_proto_metamorph_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseBool) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type ResponseMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapInfo map[string]string `protobuf:"bytes,1,rep,name=mapInfo,proto3" json:"mapInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResponseMap) Reset() {
	*x = ResponseMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMap) ProtoMessage() {}

func (x *ResponseMap) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMap.ProtoReflect.Descriptor instead.
func (*ResponseMap) Descriptor() ([]byte, []int) {
	return file_proto_metamorph_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseMap) GetMapInfo() map[string]string {
	if x != nil {
		return x.MapInfo
	}
	return nil
}

var File_proto_metamorph_proto protoreflect.FileDescriptor

var file_proto_metamorph_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61,
	0x70, 0x12, 0x39, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x3a, 0x0a, 0x0c,
	0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xec, 0x02, 0x0a, 0x03, 0x62, 0x6d, 0x68,
	0x12, 0x27, 0x0a, 0x09, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x49, 0x53, 0x4f, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x52, 0x41, 0x49, 0x44, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x57, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x61, 0x70, 0x12, 0x26, 0x0a, 0x08, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4f, 0x66, 0x66, 0x12,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x4f, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x32, 0x31, 0x0a, 0x06, 0x69, 0x73, 0x6f, 0x67, 0x65,
	0x6e, 0x12, 0x27, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x53, 0x4f, 0x12, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_metamorph_proto_rawDescOnce sync.Once
	file_proto_metamorph_proto_rawDescData = file_proto_metamorph_proto_rawDesc
)

func file_proto_metamorph_proto_rawDescGZIP() []byte {
	file_proto_metamorph_proto_rawDescOnce.Do(func() {
		file_proto_metamorph_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metamorph_proto_rawDescData)
	})
	return file_proto_metamorph_proto_rawDescData
}

var file_proto_metamorph_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_metamorph_proto_goTypes = []interface{}{
	(*Empty)(nil),        // 0: proto.Empty
	(*Response)(nil),     // 1: proto.Response
	(*ResponseBool)(nil), // 2: proto.ResponseBool
	(*ResponseMap)(nil),  // 3: proto.ResponseMap
	nil,                  // 4: proto.ResponseMap.MapInfoEntry
}
var file_proto_metamorph_proto_depIdxs = []int32{
	4,  // 0: proto.ResponseMap.mapInfo:type_name -> proto.ResponseMap.MapInfoEntry
	0,  // 1: proto.bmh.DeployISO:input_type -> proto.Empty
	0,  // 2: proto.bmh.UpdateFirmware:input_type -> proto.Empty
	0,  // 3: proto.bmh.ConfigureRAID:input_type -> proto.Empty
	0,  // 4: proto.bmh.GetGUUID:input_type -> proto.Empty
	0,  // 5: proto.bmh.GetHWInventory:input_type -> proto.Empty
	0,  // 6: proto.bmh.PowerOff:input_type -> proto.Empty
	0,  // 7: proto.bmh.PowerOn:input_type -> proto.Empty
	0,  // 8: proto.bmh.GetPowerStatus:input_type -> proto.Empty
	0,  // 9: proto.isogen.CreateISO:input_type -> proto.Empty
	0,  // 10: proto.bmh.DeployISO:output_type -> proto.Empty
	0,  // 11: proto.bmh.UpdateFirmware:output_type -> proto.Empty
	0,  // 12: proto.bmh.ConfigureRAID:output_type -> proto.Empty
	1,  // 13: proto.bmh.GetGUUID:output_type -> proto.Response
	3,  // 14: proto.bmh.GetHWInventory:output_type -> proto.ResponseMap
	0,  // 15: proto.bmh.PowerOff:output_type -> proto.Empty
	0,  // 16: proto.bmh.PowerOn:output_type -> proto.Empty
	2,  // 17: proto.bmh.GetPowerStatus:output_type -> proto.ResponseBool
	0,  // 18: proto.isogen.CreateISO:output_type -> proto.Empty
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_metamorph_proto_init() }
func file_proto_metamorph_proto_init() {
	if File_proto_metamorph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metamorph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_metamorph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_metamorph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBool); i {
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
		file_proto_metamorph_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMap); i {
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
			RawDescriptor: file_proto_metamorph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_metamorph_proto_goTypes,
		DependencyIndexes: file_proto_metamorph_proto_depIdxs,
		MessageInfos:      file_proto_metamorph_proto_msgTypes,
	}.Build()
	File_proto_metamorph_proto = out.File
	file_proto_metamorph_proto_rawDesc = nil
	file_proto_metamorph_proto_goTypes = nil
	file_proto_metamorph_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BmhClient is the client API for Bmh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BmhClient interface {
	DeployISO(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	UpdateFirmware(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ConfigureRAID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetGUUID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	GetHWInventory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseMap, error)
	PowerOff(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	PowerOn(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetPowerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseBool, error)
}

type bmhClient struct {
	cc grpc.ClientConnInterface
}

func NewBmhClient(cc grpc.ClientConnInterface) BmhClient {
	return &bmhClient{cc}
}

func (c *bmhClient) DeployISO(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.bmh/DeployISO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) UpdateFirmware(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.bmh/UpdateFirmware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) ConfigureRAID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.bmh/ConfigureRAID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) GetGUUID(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.bmh/GetGUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) GetHWInventory(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseMap, error) {
	out := new(ResponseMap)
	err := c.cc.Invoke(ctx, "/proto.bmh/GetHWInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) PowerOff(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.bmh/PowerOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) PowerOn(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.bmh/PowerOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmhClient) GetPowerStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseBool, error) {
	out := new(ResponseBool)
	err := c.cc.Invoke(ctx, "/proto.bmh/GetPowerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BmhServer is the server API for Bmh service.
type BmhServer interface {
	DeployISO(context.Context, *Empty) (*Empty, error)
	UpdateFirmware(context.Context, *Empty) (*Empty, error)
	ConfigureRAID(context.Context, *Empty) (*Empty, error)
	GetGUUID(context.Context, *Empty) (*Response, error)
	GetHWInventory(context.Context, *Empty) (*ResponseMap, error)
	PowerOff(context.Context, *Empty) (*Empty, error)
	PowerOn(context.Context, *Empty) (*Empty, error)
	GetPowerStatus(context.Context, *Empty) (*ResponseBool, error)
}

// UnimplementedBmhServer can be embedded to have forward compatible implementations.
type UnimplementedBmhServer struct {
}

func (*UnimplementedBmhServer) DeployISO(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployISO not implemented")
}
func (*UnimplementedBmhServer) UpdateFirmware(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFirmware not implemented")
}
func (*UnimplementedBmhServer) ConfigureRAID(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureRAID not implemented")
}
func (*UnimplementedBmhServer) GetGUUID(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGUUID not implemented")
}
func (*UnimplementedBmhServer) GetHWInventory(context.Context, *Empty) (*ResponseMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHWInventory not implemented")
}
func (*UnimplementedBmhServer) PowerOff(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerOff not implemented")
}
func (*UnimplementedBmhServer) PowerOn(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowerOn not implemented")
}
func (*UnimplementedBmhServer) GetPowerStatus(context.Context, *Empty) (*ResponseBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPowerStatus not implemented")
}

func RegisterBmhServer(s *grpc.Server, srv BmhServer) {
	s.RegisterService(&_Bmh_serviceDesc, srv)
}

func _Bmh_DeployISO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).DeployISO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/DeployISO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).DeployISO(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_UpdateFirmware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).UpdateFirmware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/UpdateFirmware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).UpdateFirmware(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_ConfigureRAID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).ConfigureRAID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/ConfigureRAID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).ConfigureRAID(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_GetGUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).GetGUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/GetGUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).GetGUUID(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_GetHWInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).GetHWInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/GetHWInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).GetHWInventory(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_PowerOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).PowerOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/PowerOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).PowerOff(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_PowerOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).PowerOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/PowerOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).PowerOn(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bmh_GetPowerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmhServer).GetPowerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.bmh/GetPowerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmhServer).GetPowerStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bmh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.bmh",
	HandlerType: (*BmhServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployISO",
			Handler:    _Bmh_DeployISO_Handler,
		},
		{
			MethodName: "UpdateFirmware",
			Handler:    _Bmh_UpdateFirmware_Handler,
		},
		{
			MethodName: "ConfigureRAID",
			Handler:    _Bmh_ConfigureRAID_Handler,
		},
		{
			MethodName: "GetGUUID",
			Handler:    _Bmh_GetGUUID_Handler,
		},
		{
			MethodName: "GetHWInventory",
			Handler:    _Bmh_GetHWInventory_Handler,
		},
		{
			MethodName: "PowerOff",
			Handler:    _Bmh_PowerOff_Handler,
		},
		{
			MethodName: "PowerOn",
			Handler:    _Bmh_PowerOn_Handler,
		},
		{
			MethodName: "GetPowerStatus",
			Handler:    _Bmh_GetPowerStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metamorph.proto",
}

// IsogenClient is the client API for Isogen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IsogenClient interface {
	CreateISO(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type isogenClient struct {
	cc grpc.ClientConnInterface
}

func NewIsogenClient(cc grpc.ClientConnInterface) IsogenClient {
	return &isogenClient{cc}
}

func (c *isogenClient) CreateISO(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.isogen/CreateISO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsogenServer is the server API for Isogen service.
type IsogenServer interface {
	CreateISO(context.Context, *Empty) (*Empty, error)
}

// UnimplementedIsogenServer can be embedded to have forward compatible implementations.
type UnimplementedIsogenServer struct {
}

func (*UnimplementedIsogenServer) CreateISO(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateISO not implemented")
}

func RegisterIsogenServer(s *grpc.Server, srv IsogenServer) {
	s.RegisterService(&_Isogen_serviceDesc, srv)
}

func _Isogen_CreateISO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsogenServer).CreateISO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.isogen/CreateISO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsogenServer).CreateISO(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Isogen_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.isogen",
	HandlerType: (*IsogenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateISO",
			Handler:    _Isogen_CreateISO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metamorph.proto",
}
