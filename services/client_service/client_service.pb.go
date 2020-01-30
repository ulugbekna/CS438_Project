// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/client_service/client_service.proto

package client_service

import (
	context "context"
	fmt "fmt"
	chord_service "github.com/dosarudaniel/CS438_Project/services/chord_service"
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
	proto.RegisterFile("services/client_service/client_service.proto", fileDescriptor_73cdb0b12f50e91e)
}

var fileDescriptor_73cdb0b12f50e91e = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x41, 0x50, 0x31, 0x2a, 0xc3, 0x0c, 0x3c, 0xd4, 0x8b, 0x77, 0x35, 0x05, 0xa7, 0xa0,
	0x07, 0x11, 0xdc, 0x1c, 0x8c, 0x21, 0xe8, 0xaa, 0x17, 0x2f, 0x23, 0x4b, 0xde, 0xba, 0x48, 0xdb,
	0x57, 0xf3, 0x12, 0xa1, 0xff, 0xc0, 0x9f, 0x2d, 0x4d, 0x3b, 0x36, 0x26, 0xe2, 0xf1, 0xeb, 0x7b,
	0xdf, 0xeb, 0xd7, 0xb2, 0x73, 0x02, 0xfb, 0x65, 0x14, 0x50, 0xac, 0x32, 0x03, 0x85, 0x9b, 0xb6,
	0xbc, 0x81, 0xa2, 0xb4, 0xe8, 0x90, 0x6f, 0xab, 0x05, 0x5a, 0x1d, 0x9d, 0xad, 0xa4, 0x9a, 0x57,
	0x4e, 0xa0, 0x1c, 0x88, 0x64, 0x0a, 0xd4, 0x38, 0xd1, 0xc5, 0x3f, 0x6f, 0xd8, 0x58, 0x3f, 0x49,
	0x11, 0xd3, 0x0c, 0xe2, 0x40, 0x33, 0x3f, 0x8f, 0x21, 0x2f, 0x5d, 0xd5, 0x0c, 0x2f, 0xbf, 0xb7,
	0xd8, 0x61, 0x3f, 0x68, 0x49, 0x73, 0x84, 0xf7, 0xd8, 0xfe, 0x04, 0x3e, 0x3d, 0x90, 0x1b, 0x9a,
	0x0c, 0x78, 0x57, 0x84, 0x06, 0x51, 0xc3, 0x13, 0x38, 0xa9, 0xa5, 0x93, 0x51, 0xa7, 0x7d, 0x38,
	0x01, 0x2a, 0xb1, 0x20, 0xe0, 0xd7, 0x8c, 0xbd, 0x95, 0x19, 0x4a, 0x1d, 0x9c, 0xce, 0x9a, 0x53,
	0xc8, 0x1c, 0xa2, 0x63, 0xd1, 0x34, 0x88, 0x65, 0x83, 0x78, 0xac, 0x1b, 0xb8, 0x60, 0x2c, 0x01,
	0x69, 0xd5, 0x22, 0x68, 0x07, 0xad, 0xf6, 0xe2, 0xc1, 0x56, 0x11, 0x5f, 0x3b, 0x32, 0x01, 0x85,
	0x56, 0x13, 0xbf, 0x65, 0xdd, 0xa1, 0x29, 0x74, 0xe2, 0x95, 0x02, 0x22, 0xb4, 0x4d, 0x39, 0x3f,
	0x6a, 0x57, 0x47, 0x1a, 0x0a, 0x67, 0xe6, 0x06, 0xec, 0xef, 0xc2, 0x53, 0xb6, 0x3b, 0x86, 0xea,
	0x15, 0x47, 0x03, 0xce, 0xda, 0xd9, 0x18, 0xaa, 0x68, 0x6f, 0xa9, 0x0e, 0x1e, 0xee, 0xdf, 0xef,
	0x52, 0xe3, 0x16, 0x7e, 0x26, 0x14, 0xe6, 0xb1, 0x46, 0x92, 0xd6, 0x6b, 0x59, 0x18, 0xc8, 0xe2,
	0x7e, 0x72, 0xd5, 0xbb, 0x99, 0x3e, 0x5b, 0xfc, 0x00, 0xe5, 0xe2, 0x3f, 0x7e, 0xff, 0x6c, 0x27,
	0x7c, 0x5d, 0xef, 0x27, 0x00, 0x00, 0xff, 0xff, 0x66, 0xa1, 0xf8, 0x60, 0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	RequestFile(ctx context.Context, in *FileMetadata, opts ...grpc.CallOption) (*Response, error)
	UploadFile(ctx context.Context, in *Filename, opts ...grpc.CallOption) (*empty.Empty, error)
	SearchFile(ctx context.Context, in *Query, opts ...grpc.CallOption) (*chord_service.FileRecords, error)
	FindSuccessorClient(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Response, error)
	KeyToID(ctx context.Context, in *chord_service.Key, opts ...grpc.CallOption) (*chord_service.ID, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) RequestFile(ctx context.Context, in *FileMetadata, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chord.ClientService/RequestFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UploadFile(ctx context.Context, in *Filename, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chord.ClientService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) SearchFile(ctx context.Context, in *Query, opts ...grpc.CallOption) (*chord_service.FileRecords, error) {
	out := new(chord_service.FileRecords)
	err := c.cc.Invoke(ctx, "/chord.ClientService/SearchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) FindSuccessorClient(ctx context.Context, in *Identifier, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chord.ClientService/FindSuccessorClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) KeyToID(ctx context.Context, in *chord_service.Key, opts ...grpc.CallOption) (*chord_service.ID, error) {
	out := new(chord_service.ID)
	err := c.cc.Invoke(ctx, "/chord.ClientService/KeyToID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	RequestFile(context.Context, *FileMetadata) (*Response, error)
	UploadFile(context.Context, *Filename) (*empty.Empty, error)
	SearchFile(context.Context, *Query) (*chord_service.FileRecords, error)
	FindSuccessorClient(context.Context, *Identifier) (*Response, error)
	KeyToID(context.Context, *chord_service.Key) (*chord_service.ID, error)
}

// UnimplementedClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (*UnimplementedClientServiceServer) RequestFile(ctx context.Context, req *FileMetadata) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestFile not implemented")
}
func (*UnimplementedClientServiceServer) UploadFile(ctx context.Context, req *Filename) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedClientServiceServer) SearchFile(ctx context.Context, req *Query) (*chord_service.FileRecords, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFile not implemented")
}
func (*UnimplementedClientServiceServer) FindSuccessorClient(ctx context.Context, req *Identifier) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorClient not implemented")
}
func (*UnimplementedClientServiceServer) KeyToID(ctx context.Context, req *chord_service.Key) (*chord_service.ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyToID not implemented")
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_RequestFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RequestFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ClientService/RequestFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RequestFile(ctx, req.(*FileMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filename)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ClientService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UploadFile(ctx, req.(*Filename))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_SearchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).SearchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ClientService/SearchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).SearchFile(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_FindSuccessorClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FindSuccessorClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ClientService/FindSuccessorClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FindSuccessorClient(ctx, req.(*Identifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_KeyToID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(chord_service.Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).KeyToID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chord.ClientService/KeyToID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).KeyToID(ctx, req.(*chord_service.Key))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chord.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestFile",
			Handler:    _ClientService_RequestFile_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _ClientService_UploadFile_Handler,
		},
		{
			MethodName: "SearchFile",
			Handler:    _ClientService_SearchFile_Handler,
		},
		{
			MethodName: "FindSuccessorClient",
			Handler:    _ClientService_FindSuccessorClient_Handler,
		},
		{
			MethodName: "KeyToID",
			Handler:    _ClientService_KeyToID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/client_service/client_service.proto",
}
