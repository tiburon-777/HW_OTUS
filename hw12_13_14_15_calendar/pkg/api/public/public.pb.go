// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public.proto

package public

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryRange int32

const (
	QueryRange_DAY   QueryRange = 0
	QueryRange_WEEK  QueryRange = 1
	QueryRange_MONTH QueryRange = 2
)

var QueryRange_name = map[int32]string{
	0: "DAY",
	1: "WEEK",
	2: "MONTH",
}

var QueryRange_value = map[string]int32{
	"DAY":   0,
	"WEEK":  1,
	"MONTH": 2,
}

func (x QueryRange) String() string {
	return proto.EnumName(QueryRange_name, int32(x))
}

func (QueryRange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{0}
}

type Event struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Date,proto3" json:"Date,omitempty"`
	Latency              *duration.Duration   `protobuf:"bytes,4,opt,name=Latency,proto3" json:"Latency,omitempty"`
	Note                 string               `protobuf:"bytes,5,opt,name=Note,proto3" json:"Note,omitempty"`
	UserID               int64                `protobuf:"varint,6,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NotifyTime           *duration.Duration   `protobuf:"bytes,7,opt,name=NotifyTime,proto3" json:"NotifyTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{0}
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

func (m *Event) GetLatency() *duration.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *Event) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *Event) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Event) GetNotifyTime() *duration.Duration {
	if m != nil {
		return m.NotifyTime
	}
	return nil
}

type CreateReq struct {
	Title                string               `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Date,proto3" json:"Date,omitempty"`
	Latency              *duration.Duration   `protobuf:"bytes,4,opt,name=Latency,proto3" json:"Latency,omitempty"`
	Note                 string               `protobuf:"bytes,5,opt,name=Note,proto3" json:"Note,omitempty"`
	UserID               int64                `protobuf:"varint,6,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NotifyTime           *duration.Duration   `protobuf:"bytes,7,opt,name=NotifyTime,proto3" json:"NotifyTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateReq) Reset()         { *m = CreateReq{} }
func (m *CreateReq) String() string { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()    {}
func (*CreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{1}
}

func (m *CreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReq.Unmarshal(m, b)
}
func (m *CreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReq.Marshal(b, m, deterministic)
}
func (m *CreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReq.Merge(m, src)
}
func (m *CreateReq) XXX_Size() int {
	return xxx_messageInfo_CreateReq.Size(m)
}
func (m *CreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReq proto.InternalMessageInfo

func (m *CreateReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateReq) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *CreateReq) GetLatency() *duration.Duration {
	if m != nil {
		return m.Latency
	}
	return nil
}

func (m *CreateReq) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *CreateReq) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CreateReq) GetNotifyTime() *duration.Duration {
	if m != nil {
		return m.NotifyTime
	}
	return nil
}

type CreateRsp struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRsp) Reset()         { *m = CreateRsp{} }
func (m *CreateRsp) String() string { return proto.CompactTextString(m) }
func (*CreateRsp) ProtoMessage()    {}
func (*CreateRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{2}
}

func (m *CreateRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRsp.Unmarshal(m, b)
}
func (m *CreateRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRsp.Marshal(b, m, deterministic)
}
func (m *CreateRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRsp.Merge(m, src)
}
func (m *CreateRsp) XXX_Size() int {
	return xxx_messageInfo_CreateRsp.Size(m)
}
func (m *CreateRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRsp proto.InternalMessageInfo

func (m *CreateRsp) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type UpdateReq struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=Event,proto3" json:"Event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{3}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UpdateReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type DeleteReq struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{4}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type ListResp struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResp) Reset()         { *m = ListResp{} }
func (m *ListResp) String() string { return proto.CompactTextString(m) }
func (*ListResp) ProtoMessage()    {}
func (*ListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{5}
}

func (m *ListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResp.Unmarshal(m, b)
}
func (m *ListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResp.Marshal(b, m, deterministic)
}
func (m *ListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResp.Merge(m, src)
}
func (m *ListResp) XXX_Size() int {
	return xxx_messageInfo_ListResp.Size(m)
}
func (m *ListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListResp proto.InternalMessageInfo

func (m *ListResp) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetByIDReq struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDReq) Reset()         { *m = GetByIDReq{} }
func (m *GetByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetByIDReq) ProtoMessage()    {}
func (*GetByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{6}
}

func (m *GetByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDReq.Unmarshal(m, b)
}
func (m *GetByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDReq.Marshal(b, m, deterministic)
}
func (m *GetByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDReq.Merge(m, src)
}
func (m *GetByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetByIDReq.Size(m)
}
func (m *GetByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDReq proto.InternalMessageInfo

func (m *GetByIDReq) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetByIDResp struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDResp) Reset()         { *m = GetByIDResp{} }
func (m *GetByIDResp) String() string { return proto.CompactTextString(m) }
func (*GetByIDResp) ProtoMessage()    {}
func (*GetByIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{7}
}

func (m *GetByIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDResp.Unmarshal(m, b)
}
func (m *GetByIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDResp.Marshal(b, m, deterministic)
}
func (m *GetByIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDResp.Merge(m, src)
}
func (m *GetByIDResp) XXX_Size() int {
	return xxx_messageInfo_GetByIDResp.Size(m)
}
func (m *GetByIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDResp proto.InternalMessageInfo

func (m *GetByIDResp) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type GetByDateReq struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	Range                QueryRange           `protobuf:"varint,2,opt,name=Range,proto3,enum=public.QueryRange" json:"Range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetByDateReq) Reset()         { *m = GetByDateReq{} }
func (m *GetByDateReq) String() string { return proto.CompactTextString(m) }
func (*GetByDateReq) ProtoMessage()    {}
func (*GetByDateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{8}
}

func (m *GetByDateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByDateReq.Unmarshal(m, b)
}
func (m *GetByDateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByDateReq.Marshal(b, m, deterministic)
}
func (m *GetByDateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByDateReq.Merge(m, src)
}
func (m *GetByDateReq) XXX_Size() int {
	return xxx_messageInfo_GetByDateReq.Size(m)
}
func (m *GetByDateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByDateReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetByDateReq proto.InternalMessageInfo

func (m *GetByDateReq) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *GetByDateReq) GetRange() QueryRange {
	if m != nil {
		return m.Range
	}
	return QueryRange_DAY
}

type GetByDateResp struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByDateResp) Reset()         { *m = GetByDateResp{} }
func (m *GetByDateResp) String() string { return proto.CompactTextString(m) }
func (*GetByDateResp) ProtoMessage()    {}
func (*GetByDateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_413a91106d7bcce8, []int{9}
}

func (m *GetByDateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByDateResp.Unmarshal(m, b)
}
func (m *GetByDateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByDateResp.Marshal(b, m, deterministic)
}
func (m *GetByDateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByDateResp.Merge(m, src)
}
func (m *GetByDateResp) XXX_Size() int {
	return xxx_messageInfo_GetByDateResp.Size(m)
}
func (m *GetByDateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByDateResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetByDateResp proto.InternalMessageInfo

func (m *GetByDateResp) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterEnum("public.QueryRange", QueryRange_name, QueryRange_value)
	proto.RegisterType((*Event)(nil), "public.Event")
	proto.RegisterType((*CreateReq)(nil), "public.CreateReq")
	proto.RegisterType((*CreateRsp)(nil), "public.CreateRsp")
	proto.RegisterType((*UpdateReq)(nil), "public.UpdateReq")
	proto.RegisterType((*DeleteReq)(nil), "public.DeleteReq")
	proto.RegisterType((*ListResp)(nil), "public.ListResp")
	proto.RegisterType((*GetByIDReq)(nil), "public.GetByIDReq")
	proto.RegisterType((*GetByIDResp)(nil), "public.GetByIDResp")
	proto.RegisterType((*GetByDateReq)(nil), "public.GetByDateReq")
	proto.RegisterType((*GetByDateResp)(nil), "public.GetByDateResp")
}

func init() { proto.RegisterFile("public.proto", fileDescriptor_413a91106d7bcce8) }

var fileDescriptor_413a91106d7bcce8 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0xe2, 0x38, 0xf5, 0x34, 0x2d, 0xe9, 0x50, 0x8a, 0x71, 0xab, 0x12, 0x2d, 0x42,
	0x8a, 0x72, 0xb0, 0x45, 0x8a, 0x90, 0xe0, 0x44, 0x8b, 0x23, 0x6a, 0x51, 0x8a, 0xb0, 0x5a, 0xf1,
	0x71, 0x73, 0xd2, 0x6d, 0xb0, 0x94, 0xd8, 0x8b, 0xbd, 0x41, 0x8a, 0xaa, 0x5c, 0x78, 0x05, 0xde,
	0x8a, 0x2b, 0xaf, 0xc0, 0x95, 0x07, 0xe0, 0x86, 0xbc, 0x5e, 0x3b, 0xcd, 0x47, 0x45, 0xee, 0xdc,
	0xb2, 0xfb, 0x9f, 0xf9, 0xcd, 0xcc, 0xce, 0x5f, 0x31, 0xd4, 0xd8, 0xa8, 0x3b, 0x08, 0x7a, 0x16,
	0x8b, 0x23, 0x1e, 0xa1, 0x96, 0x9d, 0xcc, 0x07, 0xfd, 0x28, 0xea, 0x0f, 0xa8, 0x2d, 0x6e, 0xbb,
	0xa3, 0x4b, 0x9b, 0x07, 0x43, 0x9a, 0x70, 0x7f, 0xc8, 0xb2, 0x40, 0x73, 0x7f, 0x3e, 0xe0, 0x62,
	0x14, 0xfb, 0x3c, 0x88, 0x42, 0xa9, 0xef, 0xce, 0xeb, 0x74, 0xc8, 0xf8, 0x58, 0x8a, 0x7b, 0x52,
	0xf4, 0x59, 0x60, 0xfb, 0x61, 0x18, 0x71, 0x91, 0x99, 0x64, 0x2a, 0xf9, 0xa3, 0x40, 0xa5, 0xf3,
	0x95, 0x86, 0x1c, 0x37, 0xa1, 0xe4, 0x3a, 0x86, 0xd2, 0x50, 0x9a, 0x65, 0xaf, 0xe4, 0x3a, 0xb8,
	0x0d, 0x95, 0xb3, 0x80, 0x0f, 0xa8, 0x51, 0x6a, 0x28, 0x4d, 0xdd, 0xcb, 0x0e, 0x68, 0x81, 0xea,
	0xf8, 0x9c, 0x1a, 0xe5, 0x86, 0xd2, 0x5c, 0x6f, 0x9b, 0x56, 0x06, 0xb7, 0xf2, 0xca, 0xd6, 0x59,
	0xde, 0xba, 0x27, 0xe2, 0xf0, 0x00, 0xaa, 0x27, 0x3e, 0xa7, 0x61, 0x6f, 0x6c, 0xa8, 0x22, 0xe5,
	0xfe, 0x42, 0x8a, 0x23, 0x87, 0xf1, 0xf2, 0x48, 0x44, 0x50, 0x4f, 0x23, 0x4e, 0x8d, 0x8a, 0xa8,
	0x2c, 0x7e, 0xe3, 0x0e, 0x68, 0xe7, 0x09, 0x8d, 0x5d, 0xc7, 0xd0, 0x44, 0x8b, 0xf2, 0x84, 0xcf,
	0x00, 0x4e, 0x23, 0x1e, 0x5c, 0x8e, 0xd3, 0xca, 0x46, 0xf5, 0x5f, 0x35, 0xae, 0x05, 0x93, 0xdf,
	0x0a, 0xe8, 0x2f, 0x63, 0xea, 0x73, 0xea, 0xd1, 0x2f, 0xff, 0xc1, 0xbc, 0xbb, 0xc5, 0xb8, 0x09,
	0x9b, 0x5f, 0x37, 0x79, 0x01, 0xfa, 0x39, 0xbb, 0x90, 0x6f, 0x31, 0xef, 0x85, 0x87, 0xd2, 0x24,
	0xe2, 0x6d, 0xd6, 0xdb, 0x1b, 0x96, 0xf4, 0xb1, 0xb8, 0xf4, 0x32, 0x2d, 0xc5, 0x3b, 0x74, 0x40,
	0x97, 0x12, 0xc8, 0x63, 0x58, 0x3b, 0x09, 0x12, 0xee, 0xd1, 0x84, 0xe1, 0x23, 0xd0, 0x44, 0x46,
	0x62, 0x28, 0x8d, 0xf2, 0x22, 0x4e, 0x8a, 0x64, 0x0f, 0xe0, 0x15, 0xe5, 0x47, 0x63, 0xd7, 0x59,
	0x06, 0x7c, 0x02, 0xeb, 0x85, 0xba, 0x3a, 0xf3, 0x33, 0xd4, 0x44, 0x96, 0x23, 0x07, 0xcd, 0xd7,
	0xab, 0xac, 0xb8, 0xde, 0x26, 0x54, 0x3c, 0x3f, 0xec, 0x67, 0x26, 0xd9, 0x6c, 0x63, 0x5e, 0xe5,
	0xdd, 0x88, 0xc6, 0x63, 0xa1, 0x78, 0x59, 0x00, 0x79, 0x0a, 0x1b, 0xd7, 0x2a, 0xad, 0xdc, 0x61,
	0xab, 0x05, 0x30, 0x85, 0x61, 0x15, 0xca, 0xce, 0xe1, 0xc7, 0xfa, 0x2d, 0x5c, 0x03, 0xf5, 0x7d,
	0xa7, 0xf3, 0xba, 0xae, 0xa0, 0x0e, 0x95, 0x37, 0x6f, 0x4f, 0xcf, 0x8e, 0xeb, 0xa5, 0xf6, 0x8f,
	0x32, 0xa8, 0xfd, 0x98, 0xf5, 0xf0, 0x08, 0xb4, 0x6c, 0xb3, 0xb8, 0x95, 0x53, 0x0b, 0x63, 0x9b,
	0xf3, 0x57, 0x09, 0x23, 0xf8, 0xed, 0xe7, 0xaf, 0xef, 0xa5, 0x1a, 0xa9, 0xda, 0x54, 0x54, 0x7d,
	0xae, 0xb4, 0xf0, 0x04, 0xb4, 0xcc, 0x00, 0x53, 0x46, 0x61, 0x08, 0x73, 0x67, 0xe1, 0x65, 0x3a,
	0xe9, 0x5f, 0x0c, 0xb9, 0x27, 0x40, 0x5b, 0x66, 0x4d, 0x82, 0xec, 0x2b, 0xd7, 0x99, 0xa4, 0x34,
	0x17, 0xb4, 0xcc, 0x0c, 0x53, 0x5a, 0x61, 0x8e, 0x1b, 0x69, 0xdb, 0x82, 0xb6, 0xd9, 0x9a, 0xa1,
	0xe1, 0x21, 0xa8, 0xa9, 0x75, 0xf0, 0x86, 0x2c, 0xb3, 0x9e, 0x17, 0xc8, 0x0d, 0x46, 0x6e, 0x0b,
	0x8e, 0x8e, 0xf9, 0x78, 0x78, 0x0c, 0x55, 0x69, 0x16, 0x2c, 0x56, 0x36, 0xf5, 0x96, 0x79, 0x67,
	0xe1, 0x2e, 0x61, 0x79, 0x33, 0x38, 0xdb, 0xcc, 0x07, 0xd0, 0x8b, 0xb5, 0xe2, 0xf6, 0x4c, 0x9e,
	0xf4, 0x94, 0x79, 0x77, 0xc9, 0x6d, 0xc2, 0xc8, 0xbe, 0xe0, 0x19, 0x64, 0xa7, 0xe0, 0x89, 0x15,
	0x4f, 0xec, 0xab, 0x34, 0x64, 0x72, 0xb4, 0xf6, 0x49, 0x7e, 0x0f, 0xba, 0x9a, 0x18, 0xf0, 0xe0,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x37, 0x8b, 0x96, 0x2e, 0x06, 0x00, 0x00,
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
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRsp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResp, error)
	GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDResp, error)
	GetByDate(ctx context.Context, in *GetByDateReq, opts ...grpc.CallOption) (*GetByDateResp, error)
}

type grpcClient struct {
	cc *grpc.ClientConn
}

func NewGrpcClient(cc *grpc.ClientConn) GrpcClient {
	return &grpcClient{cc}
}

func (c *grpcClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRsp, error) {
	out := new(CreateRsp)
	err := c.cc.Invoke(ctx, "/public.grpc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/public.grpc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/public.grpc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/public.grpc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDResp, error) {
	out := new(GetByIDResp)
	err := c.cc.Invoke(ctx, "/public.grpc/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClient) GetByDate(ctx context.Context, in *GetByDateReq, opts ...grpc.CallOption) (*GetByDateResp, error) {
	out := new(GetByDateResp)
	err := c.cc.Invoke(ctx, "/public.grpc/GetByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServer is the server API for Grpc service.
type GrpcServer interface {
	Create(context.Context, *CreateReq) (*CreateRsp, error)
	Update(context.Context, *UpdateReq) (*empty.Empty, error)
	Delete(context.Context, *DeleteReq) (*empty.Empty, error)
	List(context.Context, *empty.Empty) (*ListResp, error)
	GetByID(context.Context, *GetByIDReq) (*GetByIDResp, error)
	GetByDate(context.Context, *GetByDateReq) (*GetByDateResp, error)
}

// UnimplementedGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcServer struct {
}

func (*UnimplementedGrpcServer) Create(ctx context.Context, req *CreateReq) (*CreateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGrpcServer) Update(ctx context.Context, req *UpdateReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGrpcServer) Delete(ctx context.Context, req *DeleteReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedGrpcServer) List(ctx context.Context, req *empty.Empty) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGrpcServer) GetByID(ctx context.Context, req *GetByIDReq) (*GetByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (*UnimplementedGrpcServer) GetByDate(ctx context.Context, req *GetByDateReq) (*GetByDateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByDate not implemented")
}

func RegisterGrpcServer(s *grpc.Server, srv GrpcServer) {
	s.RegisterService(&_Grpc_serviceDesc, srv)
}

func _Grpc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		sss := srv.(GrpcServer)
		rrrrr, err := sss.Create(ctx, in)
		return rrrrr, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).List(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetByID(ctx, req.(*GetByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpc_GetByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByDateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServer).GetByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/public.grpc/GetByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServer).GetByDate(ctx, req.(*GetByDateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "public.grpc",
	HandlerType: (*GrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Grpc_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Grpc_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Grpc_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Grpc_List_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Grpc_GetByID_Handler,
		},
		{
			MethodName: "GetByDate",
			Handler:    _Grpc_GetByDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public.proto",
}
