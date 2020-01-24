// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/chord_service/chord_service.proto

package chord

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

func init() {
	proto.RegisterFile("services/chord_service/chord_service.proto", fileDescriptor_89c86567def1976f)
}

var fileDescriptor_89c86567def1976f = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0xce, 0xc8, 0x2f, 0x4a, 0x89, 0x87, 0x72, 0x51, 0x79, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x41, 0x29, 0x6d, 0xbc, 0x5a, 0x72, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0x8b, 0x21, 0x7a, 0xa4, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5,
	0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0xa4, 0xd1, 0x34, 0x46,
	0x2e, 0x56, 0x67, 0x90, 0x2e, 0x21, 0x75, 0x2e, 0x5e, 0xb7, 0xcc, 0xbc, 0x94, 0xe0, 0xd2, 0xe4,
	0xe4, 0xd4, 0xe2, 0xe2, 0xfc, 0x22, 0x21, 0x4e, 0x3d, 0xb0, 0x71, 0x7a, 0x9e, 0x2e, 0x52, 0xdc,
	0x50, 0xa6, 0x5f, 0x7e, 0x4a, 0xaa, 0x90, 0x2e, 0x17, 0x9b, 0x5f, 0x7e, 0x49, 0x66, 0x5a, 0xa5,
	0x10, 0xb2, 0xb0, 0x94, 0x98, 0x1e, 0xc4, 0x1e, 0x3d, 0x98, 0x3d, 0x7a, 0xae, 0x20, 0x7b, 0x84,
	0x4c, 0xb9, 0xf8, 0xdc, 0x53, 0x4b, 0x02, 0x8a, 0x52, 0x53, 0x52, 0xa1, 0x06, 0xe3, 0x50, 0x89,
	0x62, 0x4b, 0x12, 0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x56, 0x44, 0xeb, 0x9a,
	0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordClient interface {
	FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPredecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error)
}

type chordClient struct {
	cc *grpc.ClientConn
}

func NewChordClient(cc *grpc.ClientConn) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/chord.Chord/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chord.Chord/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) GetPredecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/chord.Chord/GetPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
type ChordServer interface {
	FindSuccessor(context.Context, *ID) (*Node, error)
	Notify(context.Context, *Node) (*empty.Empty, error)
	GetPredecessor(context.Context, *empty.Empty) (*Node, error)
}

// UnimplementedChordServer can be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (*UnimplementedChordServer) FindSuccessor(ctx context.Context, req *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (*UnimplementedChordServer) Notify(ctx context.Context, req *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedChordServer) GetPredecessor(ctx context.Context, req *empty.Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}

func RegisterChordServer(s *grpc.Server, srv ChordServer) {
	s.RegisterService(&_Chord_serviceDesc, srv)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/GetPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetPredecessor(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chord.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Chord_Notify_Handler,
		},
		{
			MethodName: "GetPredecessor",
			Handler:    _Chord_GetPredecessor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chord_service/chord_service.proto",
}
