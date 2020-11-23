// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcserver.proto

package private

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetRsp struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRsp) Reset()         { *m = GetRsp{} }
func (m *GetRsp) String() string { return proto.CompactTextString(m) }
func (*GetRsp) ProtoMessage()    {}
func (*GetRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa6debe97205904, []int{0}
}

func (m *GetRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRsp.Unmarshal(m, b)
}
func (m *GetRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRsp.Marshal(b, m, deterministic)
}
func (m *GetRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRsp.Merge(m, src)
}
func (m *GetRsp) XXX_Size() int {
	return xxx_messageInfo_GetRsp.Size(m)
}
func (m *GetRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRsp proto.InternalMessageInfo

func (m *GetRsp) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type SetReq struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReq) Reset()         { *m = SetReq{} }
func (m *SetReq) String() string { return proto.CompactTextString(m) }
func (*SetReq) ProtoMessage()    {}
func (*SetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa6debe97205904, []int{1}
}

func (m *SetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReq.Unmarshal(m, b)
}
func (m *SetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReq.Marshal(b, m, deterministic)
}
func (m *SetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReq.Merge(m, src)
}
func (m *SetReq) XXX_Size() int {
	return xxx_messageInfo_SetReq.Size(m)
}
func (m *SetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetReq proto.InternalMessageInfo

func (m *SetReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type PurgeReq struct {
	OlderThenDays        int64    `protobuf:"varint,1,opt,name=OlderThenDays,proto3" json:"OlderThenDays,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeReq) Reset()         { *m = PurgeReq{} }
func (m *PurgeReq) String() string { return proto.CompactTextString(m) }
func (*PurgeReq) ProtoMessage()    {}
func (*PurgeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa6debe97205904, []int{2}
}

func (m *PurgeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeReq.Unmarshal(m, b)
}
func (m *PurgeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeReq.Marshal(b, m, deterministic)
}
func (m *PurgeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeReq.Merge(m, src)
}
func (m *PurgeReq) XXX_Size() int {
	return xxx_messageInfo_PurgeReq.Size(m)
}
func (m *PurgeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeReq.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeReq proto.InternalMessageInfo

func (m *PurgeReq) GetOlderThenDays() int64 {
	if m != nil {
		return m.OlderThenDays
	}
	return 0
}

type PurgeResp struct {
	Qty                  int64    `protobuf:"varint,1,opt,name=Qty,proto3" json:"Qty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurgeResp) Reset()         { *m = PurgeResp{} }
func (m *PurgeResp) String() string { return proto.CompactTextString(m) }
func (*PurgeResp) ProtoMessage()    {}
func (*PurgeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa6debe97205904, []int{3}
}

func (m *PurgeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurgeResp.Unmarshal(m, b)
}
func (m *PurgeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurgeResp.Marshal(b, m, deterministic)
}
func (m *PurgeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurgeResp.Merge(m, src)
}
func (m *PurgeResp) XXX_Size() int {
	return xxx_messageInfo_PurgeResp.Size(m)
}
func (m *PurgeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PurgeResp.DiscardUnknown(m)
}

var xxx_messageInfo_PurgeResp proto.InternalMessageInfo

func (m *PurgeResp) GetQty() int64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

type Event struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Date,proto3" json:"Date,omitempty"`
	UserID               int64                `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa6debe97205904, []int{4}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Event) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func init() {
	proto.RegisterType((*GetRsp)(nil), "private.GetRsp")
	proto.RegisterType((*SetReq)(nil), "private.SetReq")
	proto.RegisterType((*PurgeReq)(nil), "private.PurgeReq")
	proto.RegisterType((*PurgeResp)(nil), "private.PurgeResp")
	proto.RegisterType((*Event)(nil), "private.Event")
}

func init() { proto.RegisterFile("grpcserver.proto", fileDescriptor_afa6debe97205904) }

var fileDescriptor_afa6debe97205904 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6b, 0xea, 0x50,
	0x10, 0x85, 0x8d, 0xd1, 0xbc, 0xe7, 0xc8, 0xf3, 0xf9, 0x86, 0x87, 0x84, 0x94, 0xd2, 0x10, 0x4a,
	0xc9, 0x2a, 0x8a, 0xdd, 0xb4, 0x8b, 0x6e, 0x4a, 0x44, 0xdc, 0xd4, 0x36, 0xda, 0x4d, 0x77, 0x51,
	0xc7, 0x34, 0x10, 0x4d, 0x7a, 0xef, 0x28, 0xf8, 0xd7, 0xfa, 0xeb, 0x4a, 0x72, 0x13, 0x4b, 0x2d,
	0xdd, 0xe5, 0x9e, 0x39, 0x73, 0x0e, 0xf3, 0x11, 0xe8, 0x46, 0x22, 0x5b, 0x4a, 0x12, 0x7b, 0x12,
	0x5e, 0x26, 0x52, 0x4e, 0xf1, 0x57, 0x26, 0xe2, 0x7d, 0xc8, 0x64, 0x5d, 0x44, 0x69, 0x1a, 0x25,
	0xd4, 0x2f, 0xe4, 0xc5, 0x6e, 0xdd, 0xe7, 0x78, 0x43, 0x92, 0xc3, 0x4d, 0xa6, 0x9c, 0xd6, 0xd9,
	0xa9, 0x81, 0x36, 0x19, 0x1f, 0xd4, 0xd0, 0x19, 0x80, 0x31, 0x26, 0x0e, 0x64, 0x86, 0x57, 0x60,
	0x8c, 0xf6, 0xb4, 0x65, 0x69, 0x6a, 0xb6, 0xee, 0xb6, 0x87, 0x1d, 0xaf, 0x6c, 0xf0, 0x0a, 0x39,
	0x28, 0xa7, 0x8e, 0x09, 0xc6, 0x8c, 0x38, 0xa0, 0x37, 0xec, 0x40, 0x7d, 0xe2, 0x9b, 0x9a, 0xad,
	0xb9, 0x7a, 0x50, 0x9f, 0xf8, 0xce, 0x00, 0x7e, 0x3f, 0xee, 0x44, 0x44, 0xf9, 0xec, 0x12, 0xfe,
	0x4c, 0x93, 0x15, 0x89, 0xf9, 0x2b, 0x6d, 0xfd, 0xf0, 0x20, 0x4b, 0xdb, 0x57, 0xd1, 0x39, 0x87,
	0x56, 0xb9, 0x21, 0x33, 0xec, 0x82, 0xfe, 0xc4, 0x87, 0xd2, 0x98, 0x7f, 0x3a, 0x3b, 0x68, 0x16,
	0xa5, 0xa7, 0x4d, 0xf8, 0x1f, 0x9a, 0xf3, 0x98, 0x13, 0x32, 0xeb, 0xb6, 0xe6, 0xb6, 0x02, 0xf5,
	0x40, 0x0f, 0x1a, 0x7e, 0xc8, 0x64, 0xea, 0xb6, 0xe6, 0xb6, 0x87, 0x96, 0xa7, 0xee, 0xf6, 0xaa,
	0xbb, 0xbd, 0x79, 0x05, 0x26, 0x28, 0x7c, 0xd8, 0x03, 0xe3, 0x59, 0x92, 0x98, 0xf8, 0x66, 0xa3,
	0x48, 0x2e, 0x5f, 0xc3, 0x77, 0x0d, 0x1a, 0x39, 0x6f, 0xbc, 0x83, 0xee, 0x98, 0xf8, 0x21, 0xe5,
	0x78, 0x1d, 0x2f, 0x43, 0x8e, 0xd3, 0xad, 0xc4, 0xde, 0xb7, 0xd8, 0x51, 0x8e, 0xd3, 0xfa, 0x7b,
	0xc4, 0xa5, 0x78, 0x3a, 0x35, 0xbc, 0x81, 0xf6, 0xac, 0x5a, 0xa7, 0x15, 0x7e, 0x3a, 0x14, 0x3f,
	0xeb, 0x87, 0x28, 0xa7, 0x86, 0xb7, 0xd0, 0x29, 0xb8, 0x4c, 0x93, 0x95, 0xa2, 0x8e, 0xff, 0x8e,
	0xcb, 0x15, 0x62, 0x0b, 0x4f, 0xa5, 0xbc, 0xf4, 0xbe, 0xf5, 0x52, 0xfd, 0x19, 0x0b, 0xa3, 0xc8,
	0xbd, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xac, 0x01, 0x89, 0x3d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GrpcClient is the client API for Grpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcClient interface {
	GetNotifications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetRsp, error)
	SetNotified(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*empty.Empty, error)
	PurgeOldEvents(ctx context.Context, in *PurgeReq, opts ...grpc.CallOption) (*PurgeResp, error)
}

type grpcClient struct {
	cc *grpc.ClientConn
}

func NewGrpcClient(cc *grpc.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) GetNotifications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetRsp, error) {
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, "/private.grpc/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) SetNotified(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/private.grpc/SetNotified", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) PurgeOldEvents(ctx context.Context, in *PurgeReq, opts ...grpc.CallOption) (*PurgeResp, error) {
	out := new(PurgeResp)
	err := c.cc.Invoke(ctx, "/private.grpc/PurgeOldEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
type GrpcServer interface {
	GetNotifications(context.Context, *empty.Empty) (*GetRsp, error)
	SetNotified(context.Context, *SetReq) (*empty.Empty, error)
	PurgeOldEvents(context.Context, *PurgeReq) (*PurgeResp, error)
}

// UnimplementedGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcServer struct {
}

func (*UnimplementedGrpcServer) GetNotifications(ctx context.Context, req *empty.Empty) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (*UnimplementedGrpcServer) SetNotified(ctx context.Context, req *SetReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotified not implemented")
}
func (*UnimplementedGrpcServer) PurgeOldEvents(ctx context.Context, req *PurgeReq) (*PurgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeOldEvents not implemented")
}

func RegisterGrpcServer(s *grpc.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.grpc/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetNotifications(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_SetNotified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).SetNotified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.grpc/SetNotified",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).SetNotified(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_PurgeOldEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).PurgeOldEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/private.grpc/PurgeOldEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).PurgeOldEvents(ctx, req.(*PurgeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "private.grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNotifications",
			Handler:    _Grpc_GetNotifications_Handler,
		},
		{
			MethodName: "SetNotified",
			Handler:    _Grpc_SetNotified_Handler,
		},
		{
			MethodName: "PurgeOldEvents",
			Handler:    _Grpc_PurgeOldEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcserver.proto",
}
