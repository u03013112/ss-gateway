// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tester/tester.proto

/*
Package tester is a generated protocol buffer package.

It is generated from these files:
	tester/tester.proto

It has these top-level messages:
	GetSSLineListRequest
	SSLine
	GetSSLineListReply
	GetSSLineConfigRequest
	GetSSLineConfigReply
*/
package tester

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

type GetSSLineListRequest struct {
}

func (m *GetSSLineListRequest) Reset()                    { *m = GetSSLineListRequest{} }
func (m *GetSSLineListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSSLineListRequest) ProtoMessage()               {}
func (*GetSSLineListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SSLine struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *SSLine) Reset()                    { *m = SSLine{} }
func (m *SSLine) String() string            { return proto.CompactTextString(m) }
func (*SSLine) ProtoMessage()               {}
func (*SSLine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SSLine) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SSLine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SSLine) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SSLine) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SSLine) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetSSLineListReply struct {
	List []*SSLine `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *GetSSLineListReply) Reset()                    { *m = GetSSLineListReply{} }
func (m *GetSSLineListReply) String() string            { return proto.CompactTextString(m) }
func (*GetSSLineListReply) ProtoMessage()               {}
func (*GetSSLineListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetSSLineListReply) GetList() []*SSLine {
	if m != nil {
		return m.List
	}
	return nil
}

type GetSSLineConfigRequest struct {
	LineID int64 `protobuf:"varint,1,opt,name=lineID" json:"lineID,omitempty"`
}

func (m *GetSSLineConfigRequest) Reset()                    { *m = GetSSLineConfigRequest{} }
func (m *GetSSLineConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSSLineConfigRequest) ProtoMessage()               {}
func (*GetSSLineConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetSSLineConfigRequest) GetLineID() int64 {
	if m != nil {
		return m.LineID
	}
	return 0
}

type GetSSLineConfigReply struct {
	IP     string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Port   string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	Method string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Passwd string `protobuf:"bytes,4,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSSLineConfigReply) Reset()                    { *m = GetSSLineConfigReply{} }
func (m *GetSSLineConfigReply) String() string            { return proto.CompactTextString(m) }
func (*GetSSLineConfigReply) ProtoMessage()               {}
func (*GetSSLineConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetSSLineConfigReply) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *GetSSLineConfigReply) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *GetSSLineConfigReply) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *GetSSLineConfigReply) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSSLineListRequest)(nil), "tester.GetSSLineListRequest")
	proto.RegisterType((*SSLine)(nil), "tester.SSLine")
	proto.RegisterType((*GetSSLineListReply)(nil), "tester.GetSSLineListReply")
	proto.RegisterType((*GetSSLineConfigRequest)(nil), "tester.GetSSLineConfigRequest")
	proto.RegisterType((*GetSSLineConfigReply)(nil), "tester.GetSSLineConfigReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SSTester service

type SSTesterClient interface {
	GetSSLineList(ctx context.Context, in *GetSSLineListRequest, opts ...grpc.CallOption) (*GetSSLineListReply, error)
	GetSSLineConfig(ctx context.Context, in *GetSSLineConfigRequest, opts ...grpc.CallOption) (*GetSSLineConfigReply, error)
}

type sSTesterClient struct {
	cc *grpc.ClientConn
}

func NewSSTesterClient(cc *grpc.ClientConn) SSTesterClient {
	return &sSTesterClient{cc}
}

func (c *sSTesterClient) GetSSLineList(ctx context.Context, in *GetSSLineListRequest, opts ...grpc.CallOption) (*GetSSLineListReply, error) {
	out := new(GetSSLineListReply)
	err := grpc.Invoke(ctx, "/tester.SSTester/GetSSLineList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSTesterClient) GetSSLineConfig(ctx context.Context, in *GetSSLineConfigRequest, opts ...grpc.CallOption) (*GetSSLineConfigReply, error) {
	out := new(GetSSLineConfigReply)
	err := grpc.Invoke(ctx, "/tester.SSTester/GetSSLineConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SSTester service

type SSTesterServer interface {
	GetSSLineList(context.Context, *GetSSLineListRequest) (*GetSSLineListReply, error)
	GetSSLineConfig(context.Context, *GetSSLineConfigRequest) (*GetSSLineConfigReply, error)
}

func RegisterSSTesterServer(s *grpc.Server, srv SSTesterServer) {
	s.RegisterService(&_SSTester_serviceDesc, srv)
}

func _SSTester_GetSSLineList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSSLineListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSTesterServer).GetSSLineList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tester.SSTester/GetSSLineList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSTesterServer).GetSSLineList(ctx, req.(*GetSSLineListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSTester_GetSSLineConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSSLineConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSTesterServer).GetSSLineConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tester.SSTester/GetSSLineConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSTesterServer).GetSSLineConfig(ctx, req.(*GetSSLineConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSTester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tester.SSTester",
	HandlerType: (*SSTesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSSLineList",
			Handler:    _SSTester_GetSSLineList_Handler,
		},
		{
			MethodName: "GetSSLineConfig",
			Handler:    _SSTester_GetSSLineConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tester/tester.proto",
}

func init() { proto.RegisterFile("tester/tester.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xe5, 0xb4, 0x04, 0xe6, 0x89, 0x22, 0x19, 0x54, 0x59, 0x51, 0x81, 0xe0, 0x09, 0xa9,
	0x9a, 0x44, 0xd3, 0xb4, 0x37, 0x68, 0x97, 0x80, 0x84, 0x2a, 0xed, 0x62, 0x4a, 0xb8, 0xe2, 0x2e,
	0x6d, 0x4c, 0x66, 0x94, 0xda, 0x26, 0xc7, 0x19, 0xda, 0x1d, 0xe2, 0x15, 0x78, 0x34, 0x5e, 0x81,
	0x27, 0xe0, 0x09, 0x90, 0x1d, 0x07, 0xad, 0x5d, 0x7b, 0x15, 0xff, 0xe7, 0x3f, 0xf6, 0xef, 0xef,
	0xc4, 0xf8, 0xa9, 0xe1, 0x60, 0x78, 0x93, 0x74, 0x9f, 0x99, 0x6e, 0x94, 0x51, 0x24, 0xec, 0x54,
	0x34, 0xa9, 0x94, 0xaa, 0x6a, 0x9e, 0x14, 0x5a, 0x24, 0x85, 0x94, 0xca, 0x14, 0x46, 0x28, 0x09,
	0x5d, 0x17, 0x1b, 0xe3, 0x67, 0x1f, 0xb9, 0xc9, 0xf3, 0x4b, 0x21, 0xf9, 0xa5, 0x00, 0x93, 0xf1,
	0x6f, 0x2d, 0x07, 0xc3, 0x7e, 0x20, 0x1c, 0x76, 0x55, 0x32, 0xc2, 0x81, 0x28, 0x29, 0x8a, 0xd1,
	0x74, 0x90, 0x05, 0xa2, 0x24, 0x04, 0x0f, 0x65, 0xb1, 0xe5, 0x34, 0x88, 0xd1, 0xf4, 0x24, 0x73,
	0x6b, 0x42, 0xf1, 0xc3, 0xa2, 0x2c, 0x1b, 0x0e, 0x40, 0x07, 0xae, 0xdc, 0x4b, 0x12, 0xe3, 0xd3,
	0x92, 0xc3, 0xa6, 0x11, 0xda, 0xc6, 0xd2, 0xa1, 0x73, 0xef, 0x96, 0xc8, 0x18, 0x87, 0x60, 0x0a,
	0xd3, 0x02, 0x7d, 0xe0, 0x4c, 0xaf, 0xd8, 0x5b, 0x4c, 0xf6, 0xae, 0xa6, 0xeb, 0x5b, 0xc2, 0xf0,
	0xb0, 0x16, 0x60, 0x28, 0x8a, 0x07, 0xd3, 0xd3, 0xc5, 0x68, 0xe6, 0x99, 0xbb, 0xb6, 0xcc, 0x79,
	0x6c, 0x8e, 0xc7, 0xff, 0x77, 0xbe, 0x57, 0xf2, 0x8b, 0xa8, 0x3c, 0x96, 0xcd, 0xaa, 0x85, 0xe4,
	0xab, 0x0f, 0x9e, 0xc7, 0x2b, 0xf6, 0xf5, 0xce, 0x18, 0xfa, 0x1d, 0x36, 0x6d, 0x84, 0x83, 0xd5,
	0x95, 0xeb, 0x3d, 0xc9, 0x82, 0xd5, 0x95, 0x65, 0xd7, 0xaa, 0x31, 0x3d, 0xbb, 0x5d, 0xdb, 0x33,
	0xb7, 0xdc, 0x5c, 0xab, 0xd2, 0xa3, 0x7b, 0x65, 0xeb, 0xba, 0x00, 0xf8, 0x5e, 0x7a, 0x68, 0xaf,
	0x16, 0x7f, 0x11, 0x7e, 0x94, 0xe7, 0x9f, 0xdc, 0xbd, 0x89, 0xc4, 0x8f, 0x77, 0x20, 0xc9, 0xa4,
	0x27, 0x3a, 0xf4, 0x5b, 0xa2, 0xe8, 0x88, 0xab, 0xeb, 0x5b, 0x76, 0xf6, 0xf3, 0xf7, 0x9f, 0x5f,
	0xc1, 0x73, 0x46, 0x93, 0x9b, 0xd4, 0x3f, 0x85, 0xa4, 0xe2, 0x06, 0xc0, 0x32, 0xda, 0xb9, 0x5c,
	0xa0, 0x73, 0x72, 0x83, 0x9f, 0xec, 0x81, 0x92, 0x17, 0xf7, 0xce, 0xdc, 0x99, 0x59, 0x34, 0x39,
	0xea, 0xdb, 0xd4, 0xd7, 0x2e, 0xf5, 0x25, 0x8b, 0x0e, 0xa5, 0x6e, 0x5c, 0xe3, 0x05, 0x3a, 0x7f,
	0x77, 0xf6, 0xf9, 0x55, 0x25, 0xcc, 0x75, 0xbb, 0x9e, 0x6d, 0xd4, 0x36, 0x69, 0xe7, 0xcb, 0x79,
	0xba, 0x4c, 0xd3, 0x45, 0x02, 0xf0, 0x46, 0xaf, 0xfd, 0xbe, 0x75, 0xe8, 0xde, 0xe4, 0xf2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x51, 0xa5, 0x89, 0xd0, 0x02, 0x00, 0x00,
}
