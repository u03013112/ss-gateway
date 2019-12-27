// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ios/ios.proto

/*
Package ios is a generated protocol buffer package.

It is generated from these files:
	ios/ios.proto

It has these top-level messages:
	LoginRequest
	LoginReply
	PurchaseRequest
	PurchaseReply
	GetConfigRequest
	GetConfigReply
*/
package ios

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type LoginRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type LoginReply struct {
	Token       string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	ExpiresDate int64  `protobuf:"varint,2,opt,name=expiresDate" json:"expiresDate,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

type PurchaseRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *PurchaseRequest) Reset()                    { *m = PurchaseRequest{} }
func (m *PurchaseRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseRequest) ProtoMessage()               {}
func (*PurchaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PurchaseRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PurchaseRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PurchaseReply struct {
	ExpiresDate int64 `protobuf:"varint,1,opt,name=expiresDate" json:"expiresDate,omitempty"`
}

func (m *PurchaseReply) Reset()                    { *m = PurchaseReply{} }
func (m *PurchaseReply) String() string            { return proto.CompactTextString(m) }
func (*PurchaseReply) ProtoMessage()               {}
func (*PurchaseReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PurchaseReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

type GetConfigRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetConfigRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetConfigReply struct {
	IP          string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Port        string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	Method      string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Passwd      string `protobuf:"bytes,4,opt,name=passwd" json:"passwd,omitempty"`
	ExpiresDate int64  `protobuf:"varint,5,opt,name=expiresDate" json:"expiresDate,omitempty"`
}

func (m *GetConfigReply) Reset()                    { *m = GetConfigReply{} }
func (m *GetConfigReply) String() string            { return proto.CompactTextString(m) }
func (*GetConfigReply) ProtoMessage()               {}
func (*GetConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetConfigReply) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *GetConfigReply) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *GetConfigReply) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *GetConfigReply) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *GetConfigReply) GetExpiresDate() int64 {
	if m != nil {
		return m.ExpiresDate
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "ios.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "ios.LoginReply")
	proto.RegisterType((*PurchaseRequest)(nil), "ios.PurchaseRequest")
	proto.RegisterType((*PurchaseReply)(nil), "ios.PurchaseReply")
	proto.RegisterType((*GetConfigRequest)(nil), "ios.GetConfigRequest")
	proto.RegisterType((*GetConfigReply)(nil), "ios.GetConfigReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IOS service

type IOSClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseReply, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error)
}

type iOSClient struct {
	cc *grpc.ClientConn
}

func NewIOSClient(cc *grpc.ClientConn) IOSClient {
	return &iOSClient{cc}
}

func (c *iOSClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/ios.IOS/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iOSClient) Purchase(ctx context.Context, in *PurchaseRequest, opts ...grpc.CallOption) (*PurchaseReply, error) {
	out := new(PurchaseReply)
	err := grpc.Invoke(ctx, "/ios.IOS/Purchase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iOSClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigReply, error) {
	out := new(GetConfigReply)
	err := grpc.Invoke(ctx, "/ios.IOS/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IOS service

type IOSServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Purchase(context.Context, *PurchaseRequest) (*PurchaseReply, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigReply, error)
}

func RegisterIOSServer(s *grpc.Server, srv IOSServer) {
	s.RegisterService(&_IOS_serviceDesc, srv)
}

func _IOS_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOSServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ios.IOS/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOSServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IOS_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOSServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ios.IOS/Purchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOSServer).Purchase(ctx, req.(*PurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IOS_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOSServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ios.IOS/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOSServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IOS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ios.IOS",
	HandlerType: (*IOSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IOS_Login_Handler,
		},
		{
			MethodName: "Purchase",
			Handler:    _IOS_Purchase_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _IOS_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ios/ios.proto",
}

func init() { proto.RegisterFile("ios/ios.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0xae, 0xd3, 0x30,
	0x10, 0x86, 0xe5, 0xa4, 0x7d, 0xa2, 0x03, 0xed, 0x7b, 0xb8, 0x2d, 0x0a, 0x01, 0xa1, 0xc8, 0xab,
	0xaa, 0x12, 0x4d, 0xd3, 0xee, 0xca, 0x0e, 0x8a, 0x50, 0x25, 0x24, 0x4a, 0xd8, 0xb1, 0x73, 0x1b,
	0x93, 0x5a, 0xa4, 0xb1, 0x89, 0x1d, 0xa0, 0x5b, 0x16, 0x5c, 0x80, 0xa3, 0x71, 0x05, 0x6e, 0xc0,
	0x05, 0x90, 0xdd, 0x04, 0xa2, 0x14, 0xbd, 0x9d, 0xe7, 0x9f, 0x7f, 0xbe, 0x19, 0xcd, 0x18, 0xfa,
	0x5c, 0xa8, 0x90, 0x0b, 0x35, 0x93, 0x85, 0xd0, 0x02, 0xbb, 0x5c, 0x28, 0xff, 0x71, 0x2a, 0x44,
	0x9a, 0xb1, 0x90, 0x4a, 0x1e, 0xd2, 0x3c, 0x17, 0x9a, 0x6a, 0x2e, 0xf2, 0xca, 0x42, 0x08, 0xdc,
	0x7b, 0x2d, 0x52, 0x9e, 0xc7, 0xec, 0x53, 0xc9, 0x94, 0xc6, 0x18, 0x3a, 0x65, 0xc9, 0x13, 0x0f,
	0x05, 0x68, 0xd2, 0x8b, 0xed, 0x9b, 0xac, 0x01, 0x2a, 0x8f, 0xcc, 0x4e, 0x78, 0x04, 0x5d, 0x2d,
	0x3e, 0xb2, 0xbc, 0xb2, 0x9c, 0x03, 0x1c, 0xc0, 0x5d, 0xf6, 0x55, 0xf2, 0x82, 0xa9, 0x35, 0xd5,
	0xcc, 0x73, 0x02, 0x34, 0x71, 0xe3, 0xa6, 0x44, 0x9e, 0xc1, 0xf5, 0xb6, 0x2c, 0xf6, 0x07, 0xaa,
	0x58, 0xdd, 0xec, 0xff, 0x28, 0x0c, 0x9d, 0x84, 0x6a, 0x6a, 0x19, 0xbd, 0xd8, 0xbe, 0x49, 0x04,
	0xfd, 0x7f, 0xc5, 0x66, 0x8a, 0x56, 0x3f, 0x74, 0xd9, 0x6f, 0x02, 0x37, 0xaf, 0x98, 0x7e, 0x21,
	0xf2, 0x0f, 0x3c, 0xbd, 0xb5, 0x21, 0xf9, 0x8e, 0x60, 0xd0, 0xb0, 0x1a, 0xfc, 0x00, 0x9c, 0xcd,
	0xb6, 0x72, 0x39, 0x9b, 0xad, 0x99, 0x49, 0x8a, 0x42, 0xd7, 0x33, 0x99, 0x37, 0x7e, 0x00, 0x57,
	0x47, 0xa6, 0x0f, 0x22, 0xf1, 0x5c, 0xab, 0x56, 0x91, 0xd1, 0x25, 0x55, 0xea, 0x4b, 0xe2, 0x75,
	0xce, 0xfa, 0x39, 0x6a, 0x8f, 0xdc, 0xbd, 0x18, 0x79, 0xf1, 0x1b, 0x81, 0xbb, 0x79, 0xf3, 0x0e,
	0xbf, 0x84, 0xae, 0x5d, 0x38, 0xbe, 0x3f, 0x33, 0xc7, 0x6c, 0x1e, 0xc8, 0xbf, 0x6e, 0x4a, 0x32,
	0x3b, 0x11, 0xef, 0xdb, 0xcf, 0x5f, 0x3f, 0x1c, 0x4c, 0xfa, 0xe1, 0xe7, 0xc8, 0xdc, 0x3e, 0xcc,
	0x4c, 0x6e, 0x85, 0xa6, 0xf8, 0x2d, 0xdc, 0xa9, 0x97, 0x86, 0x47, 0xb6, 0xac, 0x75, 0x00, 0x1f,
	0xb7, 0x54, 0xc3, 0x7b, 0x64, 0x79, 0x63, 0x72, 0x53, 0xf3, 0x64, 0x95, 0x36, 0xc8, 0x18, 0x7a,
	0x7f, 0x37, 0x85, 0xc7, 0xb6, 0xba, 0xbd, 0x64, 0x7f, 0xd8, 0x96, 0x0d, 0xf5, 0xa1, 0xa5, 0x0e,
	0xc9, 0xa0, 0xa6, 0xee, 0x6d, 0x72, 0x85, 0xa6, 0xcf, 0x83, 0xf7, 0x4f, 0x52, 0xae, 0x0f, 0xe5,
	0x6e, 0xb6, 0x17, 0xc7, 0xb0, 0x9c, 0x2f, 0xe7, 0xd1, 0x32, 0x8a, 0x16, 0xa1, 0x52, 0x4f, 0xe5,
	0xce, 0x78, 0x77, 0x57, 0xf6, 0xaf, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x33, 0x7a,
	0x65, 0xdf, 0x02, 0x00, 0x00,
}
