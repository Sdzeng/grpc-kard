// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: crawler.proto

package pbs

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q string `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"` //搜索的文字
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type NewMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *NewMessage) Reset() {
	*x = NewMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMessage) ProtoMessage() {}

func (x *NewMessage) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMessage.ProtoReflect.Descriptor instead.
func (*NewMessage) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{2}
}

func (x *NewMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewMessage) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nm *NewMessage `protobuf:"bytes,1,opt,name=nm,proto3" json:"nm,omitempty"` //搜索的文字
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{3}
}

func (x *NewRequest) GetNm() *NewMessage {
	if x != nil {
		return x.Nm
	}
	return nil
}

type NewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *NewResponse) Reset() {
	*x = NewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewResponse) ProtoMessage() {}

func (x *NewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewResponse.ProtoReflect.Descriptor instead.
func (*NewResponse) Descriptor() ([]byte, []int) {
	return file_crawler_proto_rawDescGZIP(), []int{4}
}

func (x *NewResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_crawler_proto protoreflect.FileDescriptor

var file_crawler_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x62, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x71, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4c, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x6e, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x02, 0x6e, 0x6d, 0x22, 0x21, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0xa3, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x72,
	0x77, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x2f, 0x7b, 0x71, 0x7d, 0x12, 0x45, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x0f,
	0x2e, 0x70, 0x62, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x77, 0x3a, 0x02, 0x6e, 0x6d, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawler_proto_rawDescOnce sync.Once
	file_crawler_proto_rawDescData = file_crawler_proto_rawDesc
)

func file_crawler_proto_rawDescGZIP() []byte {
	file_crawler_proto_rawDescOnce.Do(func() {
		file_crawler_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawler_proto_rawDescData)
	})
	return file_crawler_proto_rawDescData
}

var file_crawler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_crawler_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),         // 0: pbs.SearchRequest
	(*SearchResponse)(nil),        // 1: pbs.SearchResponse
	(*NewMessage)(nil),            // 2: pbs.NewMessage
	(*NewRequest)(nil),            // 3: pbs.NewRequest
	(*NewResponse)(nil),           // 4: pbs.NewResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_crawler_proto_depIdxs = []int32{
	5, // 0: pbs.NewMessage.time:type_name -> google.protobuf.Timestamp
	2, // 1: pbs.NewRequest.nm:type_name -> pbs.NewMessage
	0, // 2: pbs.CarwlerService.Search:input_type -> pbs.SearchRequest
	3, // 3: pbs.CarwlerService.New:input_type -> pbs.NewRequest
	1, // 4: pbs.CarwlerService.Search:output_type -> pbs.SearchResponse
	4, // 5: pbs.CarwlerService.New:output_type -> pbs.NewResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_crawler_proto_init() }
func file_crawler_proto_init() {
	if File_crawler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_crawler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_crawler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMessage); i {
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
		file_crawler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_crawler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewResponse); i {
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
			RawDescriptor: file_crawler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crawler_proto_goTypes,
		DependencyIndexes: file_crawler_proto_depIdxs,
		MessageInfos:      file_crawler_proto_msgTypes,
	}.Build()
	File_crawler_proto = out.File
	file_crawler_proto_rawDesc = nil
	file_crawler_proto_goTypes = nil
	file_crawler_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarwlerServiceClient is the client API for CarwlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarwlerServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error)
}

type carwlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarwlerServiceClient(cc grpc.ClientConnInterface) CarwlerServiceClient {
	return &carwlerServiceClient{cc}
}

func (c *carwlerServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/pbs.CarwlerService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carwlerServiceClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*NewResponse, error) {
	out := new(NewResponse)
	err := c.cc.Invoke(ctx, "/pbs.CarwlerService/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarwlerServiceServer is the server API for CarwlerService service.
type CarwlerServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	New(context.Context, *NewRequest) (*NewResponse, error)
}

// UnimplementedCarwlerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarwlerServiceServer struct {
}

func (*UnimplementedCarwlerServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedCarwlerServiceServer) New(context.Context, *NewRequest) (*NewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}

func RegisterCarwlerServiceServer(s *grpc.Server, srv CarwlerServiceServer) {
	s.RegisterService(&_CarwlerService_serviceDesc, srv)
}

func _CarwlerService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarwlerServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CarwlerService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarwlerServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarwlerService_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarwlerServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CarwlerService/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarwlerServiceServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarwlerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.CarwlerService",
	HandlerType: (*CarwlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _CarwlerService_Search_Handler,
		},
		{
			MethodName: "New",
			Handler:    _CarwlerService_New_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crawler.proto",
}
