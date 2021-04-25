// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoReq) Reset()         { *m = CreateTodoReq{} }
func (m *CreateTodoReq) String() string { return proto.CompactTextString(m) }
func (*CreateTodoReq) ProtoMessage()    {}
func (*CreateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoReq.Unmarshal(m, b)
}
func (m *CreateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoReq.Marshal(b, m, deterministic)
}
func (m *CreateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoReq.Merge(m, src)
}
func (m *CreateTodoReq) XXX_Size() int {
	return xxx_messageInfo_CreateTodoReq.Size(m)
}
func (m *CreateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoReq proto.InternalMessageInfo

func (m *CreateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoReq) Reset()         { *m = UpdateTodoReq{} }
func (m *UpdateTodoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoReq) ProtoMessage()    {}
func (*UpdateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *UpdateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoReq.Unmarshal(m, b)
}
func (m *UpdateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoReq.Marshal(b, m, deterministic)
}
func (m *UpdateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoReq.Merge(m, src)
}
func (m *UpdateTodoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoReq.Size(m)
}
func (m *UpdateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoReq proto.InternalMessageInfo

func (m *UpdateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteTodoReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoReq) Reset()         { *m = DeleteTodoReq{} }
func (m *DeleteTodoReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoReq) ProtoMessage()    {}
func (*DeleteTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *DeleteTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoReq.Unmarshal(m, b)
}
func (m *DeleteTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoReq.Marshal(b, m, deterministic)
}
func (m *DeleteTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoReq.Merge(m, src)
}
func (m *DeleteTodoReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoReq.Size(m)
}
func (m *DeleteTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoReq proto.InternalMessageInfo

func (m *DeleteTodoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadTodoReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadTodoReq) Reset()         { *m = ReadTodoReq{} }
func (m *ReadTodoReq) String() string { return proto.CompactTextString(m) }
func (*ReadTodoReq) ProtoMessage()    {}
func (*ReadTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ReadTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadTodoReq.Unmarshal(m, b)
}
func (m *ReadTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadTodoReq.Marshal(b, m, deterministic)
}
func (m *ReadTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadTodoReq.Merge(m, src)
}
func (m *ReadTodoReq) XXX_Size() int {
	return xxx_messageInfo_ReadTodoReq.Size(m)
}
func (m *ReadTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadTodoReq proto.InternalMessageInfo

func (m *ReadTodoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTodoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoReq) Reset()         { *m = ListTodoReq{} }
func (m *ListTodoReq) String() string { return proto.CompactTextString(m) }
func (*ListTodoReq) ProtoMessage()    {}
func (*ListTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ListTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoReq.Unmarshal(m, b)
}
func (m *ListTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoReq.Marshal(b, m, deterministic)
}
func (m *ListTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoReq.Merge(m, src)
}
func (m *ListTodoReq) XXX_Size() int {
	return xxx_messageInfo_ListTodoReq.Size(m)
}
func (m *ListTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoReq proto.InternalMessageInfo

type CreateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRes) Reset()         { *m = CreateTodoRes{} }
func (m *CreateTodoRes) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRes) ProtoMessage()    {}
func (*CreateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *CreateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRes.Unmarshal(m, b)
}
func (m *CreateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRes.Marshal(b, m, deterministic)
}
func (m *CreateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRes.Merge(m, src)
}
func (m *CreateTodoRes) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRes.Size(m)
}
func (m *CreateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRes proto.InternalMessageInfo

func (m *CreateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRes) Reset()         { *m = UpdateTodoRes{} }
func (m *UpdateTodoRes) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRes) ProtoMessage()    {}
func (*UpdateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *UpdateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRes.Unmarshal(m, b)
}
func (m *UpdateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRes.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRes.Merge(m, src)
}
func (m *UpdateTodoRes) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRes.Size(m)
}
func (m *UpdateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRes proto.InternalMessageInfo

func (m *UpdateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteTodoRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRes) Reset()         { *m = DeleteTodoRes{} }
func (m *DeleteTodoRes) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRes) ProtoMessage()    {}
func (*DeleteTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *DeleteTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRes.Unmarshal(m, b)
}
func (m *DeleteTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRes.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRes.Merge(m, src)
}
func (m *DeleteTodoRes) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRes.Size(m)
}
func (m *DeleteTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRes proto.InternalMessageInfo

func (m *DeleteTodoRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ReadTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadTodoRes) Reset()         { *m = ReadTodoRes{} }
func (m *ReadTodoRes) String() string { return proto.CompactTextString(m) }
func (*ReadTodoRes) ProtoMessage()    {}
func (*ReadTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *ReadTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadTodoRes.Unmarshal(m, b)
}
func (m *ReadTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadTodoRes.Marshal(b, m, deterministic)
}
func (m *ReadTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadTodoRes.Merge(m, src)
}
func (m *ReadTodoRes) XXX_Size() int {
	return xxx_messageInfo_ReadTodoRes.Size(m)
}
func (m *ReadTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadTodoRes proto.InternalMessageInfo

func (m *ReadTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type ListTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRes) Reset()         { *m = ListTodoRes{} }
func (m *ListTodoRes) String() string { return proto.CompactTextString(m) }
func (*ListTodoRes) ProtoMessage()    {}
func (*ListTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *ListTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRes.Unmarshal(m, b)
}
func (m *ListTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRes.Marshal(b, m, deterministic)
}
func (m *ListTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRes.Merge(m, src)
}
func (m *ListTodoRes) XXX_Size() int {
	return xxx_messageInfo_ListTodoRes.Size(m)
}
func (m *ListTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRes proto.InternalMessageInfo

func (m *ListTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "proto.Todo")
	proto.RegisterType((*CreateTodoReq)(nil), "proto.CreateTodoReq")
	proto.RegisterType((*UpdateTodoReq)(nil), "proto.UpdateTodoReq")
	proto.RegisterType((*DeleteTodoReq)(nil), "proto.DeleteTodoReq")
	proto.RegisterType((*ReadTodoReq)(nil), "proto.ReadTodoReq")
	proto.RegisterType((*ListTodoReq)(nil), "proto.ListTodoReq")
	proto.RegisterType((*CreateTodoRes)(nil), "proto.CreateTodoRes")
	proto.RegisterType((*UpdateTodoRes)(nil), "proto.UpdateTodoRes")
	proto.RegisterType((*DeleteTodoRes)(nil), "proto.DeleteTodoRes")
	proto.RegisterType((*ReadTodoRes)(nil), "proto.ReadTodoRes")
	proto.RegisterType((*ListTodoRes)(nil), "proto.ListTodoRes")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x69, 0xfe, 0x7d, 0xfb, 0xcf, 0x12, 0x0f, 0x43, 0xc1, 0x10, 0x90, 0x96, 0x9c, 0xf4,
	0x12, 0xa4, 0x5e, 0xbc, 0xeb, 0x45, 0xf0, 0x14, 0xf5, 0x2c, 0x35, 0x3b, 0xc8, 0x42, 0x71, 0x63,
	0x66, 0xeb, 0x17, 0xf2, 0x8b, 0x4a, 0x26, 0xef, 0x6d, 0xc4, 0xf6, 0x54, 0xe6, 0xf7, 0xec, 0xd3,
	0x79, 0x66, 0x32, 0xe0, 0x33, 0xe5, 0x5f, 0x26, 0xa5, 0x38, 0xcb, 0xad, 0xb3, 0x38, 0x91, 0x9f,
	0xe8, 0x1d, 0xc6, 0xcf, 0x56, 0x5b, 0x3c, 0x03, 0xcf, 0xe8, 0x60, 0xb4, 0x1a, 0x5d, 0xfe, 0x4f,
	0x3c, 0xa3, 0xf1, 0x1c, 0x66, 0x3b, 0xa6, 0xfc, 0xd5, 0xe8, 0xc0, 0x13, 0x38, 0x2d, 0xca, 0x07,
	0x8d, 0x0b, 0x98, 0x38, 0xe3, 0xb6, 0x14, 0xfc, 0x13, 0x5c, 0x16, 0xb8, 0x02, 0xa5, 0x89, 0xd3,
	0xdc, 0x64, 0xce, 0xd8, 0x8f, 0x60, 0x2c, 0x5a, 0x17, 0x45, 0xd7, 0xe0, 0xdf, 0xe5, 0xb4, 0x71,
	0x54, 0xb4, 0x4b, 0xe8, 0x13, 0x97, 0x30, 0x76, 0x56, 0x5b, 0xe9, 0xa9, 0xd6, 0xaa, 0x8c, 0x15,
	0x8b, 0x2a, 0x42, 0xe1, 0x78, 0xc9, 0xf4, 0x29, 0x8e, 0x25, 0xf8, 0xf7, 0xb4, 0xa5, 0xd6, 0xb1,
	0x37, 0x55, 0x74, 0x01, 0x2a, 0xa1, 0x8d, 0xfe, 0x4d, 0xf6, 0x41, 0x3d, 0x1a, 0x76, 0x95, 0xbc,
	0x1f, 0x99, 0x4f, 0x8e, 0x7c, 0x84, 0xe3, 0xaa, 0x1f, 0x99, 0x31, 0x80, 0x19, 0xef, 0xd2, 0x94,
	0x98, 0xc5, 0x34, 0x4f, 0xea, 0x32, 0x8a, 0xbb, 0xe1, 0x8f, 0xf8, 0xeb, 0xb8, 0x3b, 0xcd, 0xdf,
	0xef, 0xd7, 0xdf, 0x1e, 0xa8, 0xa2, 0x7c, 0x2a, 0xef, 0x04, 0x6f, 0x01, 0xda, 0xf1, 0x71, 0x51,
	0x19, 0x7a, 0x1f, 0x31, 0x1c, 0xa2, 0x8c, 0x6b, 0x98, 0xd7, 0x49, 0x11, 0xab, 0x17, 0x9d, 0xbd,
	0x87, 0x87, 0x8c, 0x8b, 0x6e, 0xed, 0xea, 0x9a, 0x6e, 0xbd, 0x03, 0x08, 0x87, 0xa8, 0x38, 0xdb,
	0x15, 0x36, 0xce, 0xde, 0x21, 0x84, 0x43, 0x54, 0x72, 0xd6, 0x1b, 0x6a, 0x72, 0x76, 0x0e, 0x20,
	0x3c, 0x64, 0xfc, 0x36, 0x15, 0x74, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x08, 0xf6, 0x85,
	0x4f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error)
	ReadTodo(ctx context.Context, in *ReadTodoReq, opts ...grpc.CallOption) (*ReadTodoRes, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error)
	ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (*ListTodoRes, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error) {
	out := new(CreateTodoRes)
	err := c.cc.Invoke(ctx, "/proto.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ReadTodo(ctx context.Context, in *ReadTodoReq, opts ...grpc.CallOption) (*ReadTodoRes, error) {
	out := new(ReadTodoRes)
	err := c.cc.Invoke(ctx, "/proto.TodoService/ReadTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error) {
	out := new(UpdateTodoRes)
	err := c.cc.Invoke(ctx, "/proto.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error) {
	out := new(DeleteTodoRes)
	err := c.cc.Invoke(ctx, "/proto.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (*ListTodoRes, error) {
	out := new(ListTodoRes)
	err := c.cc.Invoke(ctx, "/proto.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoRes, error)
	ReadTodo(context.Context, *ReadTodoReq) (*ReadTodoRes, error)
	UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoRes, error)
	DeleteTodo(context.Context, *DeleteTodoReq) (*DeleteTodoRes, error)
	ListTodo(context.Context, *ListTodoReq) (*ListTodoRes, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodo(ctx context.Context, req *CreateTodoReq) (*CreateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ReadTodo(ctx context.Context, req *ReadTodoReq) (*ReadTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTodo not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoReq) (*UpdateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoReq) (*DeleteTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ListTodo(ctx context.Context, req *ListTodoReq) (*ListTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ReadTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ReadTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/ReadTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ReadTodo(ctx, req.(*ReadTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*ListTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "ReadTodo",
			Handler:    _TodoService_ReadTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
