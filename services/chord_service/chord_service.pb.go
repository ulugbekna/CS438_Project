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
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x41, 0x68, 0xa0, 0x57, 0xeb, 0xe2, 0x2e, 0x5c, 0x8c, 0x1b, 0x77, 0x82, 0xe2, 0x04,
	0x2a, 0xbe, 0x80, 0x7f, 0x45, 0x02, 0xa5, 0xa8, 0x74, 0x2b, 0x69, 0x72, 0x12, 0x03, 0x69, 0xa7,
	0xce, 0x9d, 0x08, 0xf3, 0x56, 0x3e, 0xa2, 0x24, 0x99, 0x62, 0xbb, 0x88, 0xbb, 0xfb, 0x9d, 0x39,
	0x3f, 0x43, 0x57, 0x02, 0xfb, 0x5d, 0x65, 0x90, 0x38, 0xfb, 0x34, 0x36, 0xff, 0x08, 0x78, 0x48,
	0x7a, 0x6b, 0x8d, 0x33, 0x3c, 0xea, 0x44, 0x75, 0xfd, 0x6f, 0x64, 0x0d, 0x91, 0xb4, 0x84, 0xf4,
	0x19, 0x75, 0x5e, 0x1a, 0x53, 0xd6, 0x88, 0x3b, 0x5a, 0x35, 0x45, 0x8c, 0xf5, 0xd6, 0xf9, 0xfe,
	0x71, 0xfa, 0x73, 0x44, 0xa3, 0x87, 0x36, 0xc5, 0x97, 0x34, 0x79, 0xae, 0x36, 0xf9, 0x5b, 0x93,
	0x65, 0x10, 0x31, 0x96, 0xc7, 0xba, 0xab, 0xd3, 0x2f, 0x8f, 0xea, 0x38, 0x9c, 0x73, 0x93, 0x83,
	0x6f, 0x28, 0x9a, 0x1b, 0x57, 0x15, 0x9e, 0xf7, 0x65, 0x75, 0xa6, 0xfb, 0x1d, 0xbd, 0xdb, 0xd1,
	0x4f, 0xed, 0x0e, 0xdf, 0xd1, 0xe9, 0x0c, 0x6e, 0x61, 0x91, 0x23, 0x14, 0x0f, 0x38, 0x0f, 0x57,
	0x2e, 0x28, 0x9a, 0xc1, 0x25, 0xf0, 0x4c, 0x41, 0x4e, 0xe0, 0xd5, 0xee, 0x5e, 0xa6, 0x35, 0x4f,
	0x69, 0xbc, 0x68, 0x5a, 0x47, 0x0b, 0x93, 0x3f, 0xd3, 0x32, 0xad, 0x07, 0x3f, 0x73, 0x4f, 0x27,
	0xef, 0x36, 0xdd, 0x48, 0x01, 0x9b, 0xc0, 0x0b, 0xab, 0x10, 0xdb, 0x17, 0x5f, 0xf1, 0xd5, 0x40,
	0xdc, 0x50, 0xc7, 0x2a, 0xea, 0xf8, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xab, 0x95, 0xcf, 0x8a,
	0xb8, 0x01, 0x00, 0x00,
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
	GetKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error)
	PutKeyVal(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*empty.Empty, error)
	TransferKeys(ctx context.Context, in *TransferKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *chordClient) GetKey(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Val, error) {
	out := new(Val)
	err := c.cc.Invoke(ctx, "/chord.Chord/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) PutKeyVal(ctx context.Context, in *KeyVal, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chord.Chord/PutKeyVal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) TransferKeys(ctx context.Context, in *TransferKeysRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chord.Chord/TransferKeys", in, out, opts...)
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
	GetKey(context.Context, *Key) (*Val, error)
	PutKeyVal(context.Context, *KeyVal) (*empty.Empty, error)
	TransferKeys(context.Context, *TransferKeysRequest) (*empty.Empty, error)
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
func (*UnimplementedChordServer) GetKey(ctx context.Context, req *Key) (*Val, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (*UnimplementedChordServer) PutKeyVal(ctx context.Context, req *KeyVal) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutKeyVal not implemented")
}
func (*UnimplementedChordServer) TransferKeys(ctx context.Context, req *TransferKeysRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferKeys not implemented")
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

func _Chord_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).GetKey(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_PutKeyVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyVal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).PutKeyVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/PutKeyVal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).PutKeyVal(ctx, req.(*KeyVal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_TransferKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).TransferKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.Chord/TransferKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).TransferKeys(ctx, req.(*TransferKeysRequest))
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
		{
			MethodName: "GetKey",
			Handler:    _Chord_GetKey_Handler,
		},
		{
			MethodName: "PutKeyVal",
			Handler:    _Chord_PutKeyVal_Handler,
		},
		{
			MethodName: "TransferKeys",
			Handler:    _Chord_TransferKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/chord_service/chord_service.proto",
}
