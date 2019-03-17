// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volume.proto

package volume_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Info struct {
	Action               string   `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	ErrorCode            string   `protobuf:"bytes,2,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{1}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Info) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Info) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *Info) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type File struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
	ErrorCode            string   `protobuf:"bytes,3,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,4,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	Meta                 []byte   `protobuf:"bytes,5,opt,name=Meta,proto3" json:"Meta,omitempty"`
	Data                 []byte   `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{2}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *File) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *File) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *File) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *File) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "volume_pb.Empty")
	proto.RegisterType((*Info)(nil), "volume_pb.Info")
	proto.RegisterType((*File)(nil), "volume_pb.File")
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor_498b213ad3bcd5ad) }

var fileDescriptor_498b213ad3bcd5ad = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xbb, 0xed, 0xb6, 0xdf, 0xd7, 0xa1, 0xd6, 0x30, 0x07, 0x09, 0xc1, 0x43, 0xd9, 0x53,
	0xf0, 0x10, 0x44, 0x41, 0x7a, 0x15, 0x8d, 0x20, 0x52, 0x90, 0x44, 0xbc, 0xca, 0xb6, 0x19, 0x4b,
	0xa0, 0xd9, 0x0d, 0x9b, 0xb5, 0xd0, 0x5f, 0x22, 0xf8, 0x6b, 0x25, 0xdb, 0x52, 0xd3, 0xaa, 0xd8,
	0xdb, 0x3b, 0xcf, 0xee, 0x90, 0x27, 0xb3, 0x03, 0x83, 0xa5, 0x5e, 0xbc, 0x15, 0x14, 0x95, 0x46,
	0x5b, 0x8d, 0xfd, 0x75, 0xf5, 0x52, 0x4e, 0xc5, 0x3f, 0xe8, 0xc6, 0x45, 0x69, 0x57, 0xc2, 0x02,
	0xbf, 0x57, 0xaf, 0x1a, 0x4f, 0xa0, 0x77, 0x3d, 0xb3, 0xb9, 0x56, 0x3e, 0x1b, 0xb1, 0xb0, 0x9f,
	0x6c, 0x2a, 0x3c, 0x85, 0x7e, 0x6c, 0x8c, 0x36, 0x37, 0x3a, 0x23, 0xbf, 0xed, 0x8e, 0xbe, 0x00,
	0x0a, 0x18, 0xb8, 0x62, 0x42, 0x55, 0x25, 0xe7, 0xe4, 0x77, 0xdc, 0x85, 0x1d, 0x86, 0x08, 0xfc,
	0x56, 0x5a, 0xe9, 0x73, 0x77, 0xe6, 0xb2, 0xf8, 0x60, 0xc0, 0xef, 0xf2, 0x05, 0xa1, 0x07, 0x9d,
	0x07, 0x5a, 0x6d, 0xbe, 0x59, 0xc7, 0x86, 0x48, 0xfb, 0x77, 0x91, 0xce, 0x5f, 0x22, 0xfc, 0x67,
	0x91, 0x09, 0x59, 0xe9, 0x77, 0x47, 0x2c, 0x1c, 0x24, 0x2e, 0x6f, 0xe5, 0x7a, 0x6b, 0x56, 0xe7,
	0x8b, 0xf7, 0x36, 0x1c, 0x3d, 0xbb, 0x49, 0xa5, 0x64, 0x96, 0xf9, 0x8c, 0x70, 0x0c, 0xc3, 0xd4,
	0x1a, 0x92, 0x45, 0x4a, 0x2a, 0x73, 0xde, 0xc7, 0xd1, 0x76, 0x96, 0x51, 0x0d, 0x82, 0x7d, 0x20,
	0x5a, 0x21, 0x3b, 0x67, 0x18, 0xc1, 0xff, 0x84, 0xe4, 0xc1, 0x3d, 0x38, 0x06, 0xef, 0xc9, 0x48,
	0x55, 0x49, 0xf7, 0xd3, 0xa9, 0x95, 0xc6, 0xa2, 0xd7, 0xb8, 0xe6, 0x1e, 0x2d, 0xf8, 0x46, 0x44,
	0x0b, 0xaf, 0x60, 0xd8, 0xe8, 0x8c, 0x55, 0x76, 0x60, 0xdf, 0x19, 0xf0, 0xc7, 0x5c, 0xcd, 0x77,
	0xec, 0xea, 0x8d, 0x08, 0xf6, 0x81, 0x68, 0x4d, 0x7b, 0x6e, 0x8f, 0x2e, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x35, 0x49, 0xd9, 0x57, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VolumeServiceClient is the client API for VolumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VolumeServiceClient interface {
	StreamSendFile(ctx context.Context, opts ...grpc.CallOption) (VolumeService_StreamSendFileClient, error)
	ReadFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	TransactionStart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	TransactionEnd(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Ping(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error)
}

type volumeServiceClient struct {
	cc *grpc.ClientConn
}

func NewVolumeServiceClient(cc *grpc.ClientConn) VolumeServiceClient {
	return &volumeServiceClient{cc}
}

func (c *volumeServiceClient) StreamSendFile(ctx context.Context, opts ...grpc.CallOption) (VolumeService_StreamSendFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VolumeService_serviceDesc.Streams[0], "/volume_pb.VolumeService/StreamSendFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &volumeServiceStreamSendFileClient{stream}
	return x, nil
}

type VolumeService_StreamSendFileClient interface {
	Send(*File) error
	Recv() (*File, error)
	grpc.ClientStream
}

type volumeServiceStreamSendFileClient struct {
	grpc.ClientStream
}

func (x *volumeServiceStreamSendFileClient) Send(m *File) error {
	return x.ClientStream.SendMsg(m)
}

func (x *volumeServiceStreamSendFileClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *volumeServiceClient) ReadFile(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, "/volume_pb.VolumeService/ReadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) TransactionStart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/volume_pb.VolumeService/TransactionStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) TransactionEnd(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/volume_pb.VolumeService/TransactionEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) Ping(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/volume_pb.VolumeService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeServiceServer is the server API for VolumeService service.
type VolumeServiceServer interface {
	StreamSendFile(VolumeService_StreamSendFileServer) error
	ReadFile(context.Context, *File) (*File, error)
	TransactionStart(context.Context, *Empty) (*Empty, error)
	TransactionEnd(context.Context, *Empty) (*Empty, error)
	Ping(context.Context, *Info) (*Info, error)
}

func RegisterVolumeServiceServer(s *grpc.Server, srv VolumeServiceServer) {
	s.RegisterService(&_VolumeService_serviceDesc, srv)
}

func _VolumeService_StreamSendFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VolumeServiceServer).StreamSendFile(&volumeServiceStreamSendFileServer{stream})
}

type VolumeService_StreamSendFileServer interface {
	Send(*File) error
	Recv() (*File, error)
	grpc.ServerStream
}

type volumeServiceStreamSendFileServer struct {
	grpc.ServerStream
}

func (x *volumeServiceStreamSendFileServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func (x *volumeServiceStreamSendFileServer) Recv() (*File, error) {
	m := new(File)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VolumeService_ReadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).ReadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume_pb.VolumeService/ReadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).ReadFile(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_TransactionStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).TransactionStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume_pb.VolumeService/TransactionStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).TransactionStart(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_TransactionEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).TransactionEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume_pb.VolumeService/TransactionEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).TransactionEnd(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/volume_pb.VolumeService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).Ping(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

var _VolumeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "volume_pb.VolumeService",
	HandlerType: (*VolumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadFile",
			Handler:    _VolumeService_ReadFile_Handler,
		},
		{
			MethodName: "TransactionStart",
			Handler:    _VolumeService_TransactionStart_Handler,
		},
		{
			MethodName: "TransactionEnd",
			Handler:    _VolumeService_TransactionEnd_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _VolumeService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSendFile",
			Handler:       _VolumeService_StreamSendFile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "volume.proto",
}
