// Code generated by protoc-gen-go.
// source: pkg/proto/user.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:

	pkg/proto/user.proto

It has these top-level messages:

	UserInfo
	UserId
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type UserInfo struct {
	ID        string `protobuf:"bytes,1,opt,name=ID"          json:"id,omitempty"          bson:"_id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=UserName"    json:"user_name"             bson:"user_name"             validate:"required"`
	FirstName string `protobuf:"bytes,3,opt,name=FirstName"   json:"first_name,omitempty"  bson:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=LastName"    json:"last_name,omitempty"   bson:"last_name,omitempty"`
	Age       int32  `protobuf:"varint,5,opt,name=Age"        json:"age,omitempty"         bson:"age,omitempty"         validate:"gte=0,lte=130"`
	Email     string `protobuf:"bytes,6,opt,name=Email"       json:"email"                 bson:"email"                 validate:"required,email"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserInfo) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserInfo) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserId struct {
	Value string `protobuf:"bytes,1,opt,name=Value"       json:"value,omitempty"       bson:"value,omitempty"`
}

func (m *UserId) Reset()                    { *m = UserId{} }
func (m *UserId) String() string            { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()               {}
func (*UserId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserId) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*UserId)(nil), "user.UserId")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserInfo, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.UserService/GetUserById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	GetUserById(context.Context, *UserId) (*UserInfo, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/user.proto",
}

func init() { proto.RegisterFile("pkg/proto/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x33, 0x85, 0x58, 0x40,
	0x6c, 0xa5, 0x19, 0x8c, 0x5c, 0x1c, 0xa1, 0xc5, 0xa9, 0x45, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x7c,
	0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x52,
	0x10, 0x39, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x28, 0x9c, 0x2f, 0x24, 0xc3, 0xc5, 0xe9,
	0x96, 0x59, 0x54, 0x5c, 0x02, 0x96, 0x64, 0x06, 0x4b, 0x22, 0x04, 0x40, 0x3a, 0x7d, 0x12, 0xa1,
	0x92, 0x2c, 0x10, 0x9d, 0x30, 0xbe, 0x90, 0x00, 0x17, 0xb3, 0x63, 0x7a, 0xaa, 0x04, 0xab, 0x02,
	0xa3, 0x06, 0x6b, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0xea, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1,
	0x06, 0x56, 0x0a, 0xe1, 0x28, 0xc9, 0x71, 0xb1, 0x81, 0x5d, 0x96, 0x02, 0x92, 0x0f, 0x4b, 0xcc,
	0x29, 0x4d, 0x85, 0x3a, 0x0d, 0xc2, 0x31, 0xb2, 0xe1, 0xe2, 0x06, 0xc9, 0x07, 0xa7, 0x16, 0x95,
	0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x72, 0x71, 0xbb, 0xa7, 0x96, 0x80, 0x44, 0x9c, 0x2a, 0x3d, 0x53,
	0x84, 0x78, 0xf4, 0xc0, 0x7e, 0x85, 0x98, 0x20, 0xc5, 0x87, 0xc4, 0xcb, 0x4b, 0xcb, 0x57, 0x62,
	0x70, 0xe2, 0x88, 0x62, 0x03, 0x07, 0x4b, 0x52, 0x12, 0x1b, 0x38, 0x3c, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xc5, 0xe7, 0x65, 0x27, 0x01, 0x00, 0x00,
}