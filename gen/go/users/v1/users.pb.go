// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users/v1/users.proto

package v1pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string               `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_52ac4d72540ebedf, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_52ac4d72540ebedf, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_52ac4d72540ebedf, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "gomesh.users.v1.User")
	proto.RegisterType((*GetRequest)(nil), "gomesh.users.v1.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "gomesh.users.v1.CreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gomesh.users.v1.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/gomesh.users.v1.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Get(context.Context, *GetRequest) (*User, error)
	Create(context.Context, *CreateRequest) (*User, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gomesh.users.v1.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gomesh.users.v1.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gomesh.users.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users/v1/users.proto",
}

func init() { proto.RegisterFile("users/v1/users.proto", fileDescriptor_users_52ac4d72540ebedf) }

var fileDescriptor_users_52ac4d72540ebedf = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x99, 0x36, 0xed, 0xcb, 0x7b, 0xe3, 0x07, 0x8c, 0x5f, 0xa1, 0x82, 0xd6, 0xae, 0xda,
	0xcd, 0x84, 0xd6, 0x95, 0xb8, 0x10, 0xea, 0xa2, 0xb8, 0x91, 0x12, 0x74, 0xe3, 0x26, 0x24, 0xcd,
	0x35, 0x06, 0x9b, 0x4e, 0xcc, 0x4c, 0x02, 0xae, 0xfd, 0x27, 0xfe, 0x52, 0x99, 0x9b, 0x86, 0xd0,
	0x56, 0x77, 0xf7, 0x9e, 0x73, 0x66, 0x78, 0xee, 0x81, 0xe3, 0x42, 0x61, 0xae, 0xdc, 0x72, 0xec,
	0xd2, 0x20, 0xb2, 0x5c, 0x6a, 0xc9, 0x0f, 0x63, 0x99, 0xa2, 0x7a, 0x13, 0x95, 0x56, 0x8e, 0x7b,
	0x97, 0xb1, 0x94, 0xf1, 0x12, 0x5d, 0xb2, 0xc3, 0xe2, 0xd5, 0xd5, 0x49, 0x8a, 0x4a, 0x07, 0x69,
	0x56, 0xbd, 0x18, 0x7c, 0x33, 0xb0, 0x9e, 0x15, 0xe6, 0xfc, 0x00, 0x5a, 0x49, 0xe4, 0xb0, 0x3e,
	0x1b, 0xfe, 0xf7, 0x5a, 0x49, 0xc4, 0x4f, 0xa1, 0x9b, 0x05, 0x39, 0xae, 0xb4, 0xd3, 0x22, 0x6d,
	0xbd, 0x71, 0x0e, 0xd6, 0x2a, 0x48, 0xd1, 0x69, 0x93, 0x4a, 0x33, 0xbf, 0x82, 0xbd, 0x28, 0x51,
	0xd9, 0x32, 0xf8, 0xf4, 0xc9, 0xb3, 0xc8, 0xb3, 0xd7, 0xda, 0xa3, 0x89, 0xdc, 0x82, 0xbd, 0xc8,
	0x31, 0xd0, 0xe8, 0x1b, 0x02, 0xa7, 0xd3, 0x67, 0x43, 0x7b, 0xd2, 0x13, 0x15, 0x9e, 0xa8, 0xf1,
	0xc4, 0x53, 0x8d, 0xe7, 0x41, 0x15, 0x37, 0xc2, 0xa0, 0x0f, 0x30, 0x43, 0xed, 0xe1, 0x47, 0x81,
	0xaa, 0x21, 0x60, 0x0d, 0xc1, 0xe0, 0x1d, 0xf6, 0xef, 0x29, 0x5f, 0x87, 0x1a, 0x7c, 0xb6, 0x81,
	0x7f, 0x06, 0xff, 0x4c, 0x39, 0x7e, 0x12, 0xd5, 0x77, 0x99, 0xf5, 0x21, 0xe2, 0x23, 0xb0, 0xcc,
	0x44, 0x77, 0xd9, 0x93, 0x13, 0xb1, 0xd5, 0xa4, 0x30, 0x25, 0x79, 0x14, 0x99, 0x7c, 0x31, 0xe8,
	0x98, 0x55, 0xf1, 0x1b, 0x68, 0xcf, 0x50, 0xf3, 0xf3, 0x9d, 0x74, 0x83, 0xdb, 0xfb, 0xfd, 0x2b,
	0x7e, 0x07, 0xdd, 0x8a, 0x98, 0x5f, 0xec, 0x04, 0x36, 0x4e, 0xf9, 0xe3, 0x83, 0xe9, 0x08, 0x8e,
	0x16, 0x32, 0xdd, 0xf6, 0xa6, 0x40, 0x64, 0x73, 0x53, 0xe8, 0x9c, 0xbd, 0x58, 0xe5, 0x38, 0x0b,
	0xc3, 0x2e, 0xf5, 0x7b, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x10, 0xc8, 0xd8, 0xd8, 0x35, 0x02,
	0x00, 0x00,
}
