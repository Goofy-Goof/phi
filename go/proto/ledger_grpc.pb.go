// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: ledger.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BeanAccountServiceClient is the client API for BeanAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeanAccountServiceClient interface {
	SyncLedger(ctx context.Context, in *SyncMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DownLoadBeanAccountFile(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (BeanAccountService_DownLoadBeanAccountFileClient, error)
}

type beanAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBeanAccountServiceClient(cc grpc.ClientConnInterface) BeanAccountServiceClient {
	return &beanAccountServiceClient{cc}
}

func (c *beanAccountServiceClient) SyncLedger(ctx context.Context, in *SyncMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/phi.BeanAccountService/SyncLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beanAccountServiceClient) DownLoadBeanAccountFile(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (BeanAccountService_DownLoadBeanAccountFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &BeanAccountService_ServiceDesc.Streams[0], "/phi.BeanAccountService/DownLoadBeanAccountFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &beanAccountServiceDownLoadBeanAccountFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BeanAccountService_DownLoadBeanAccountFileClient interface {
	Recv() (*FileChunkMessage, error)
	grpc.ClientStream
}

type beanAccountServiceDownLoadBeanAccountFileClient struct {
	grpc.ClientStream
}

func (x *beanAccountServiceDownLoadBeanAccountFileClient) Recv() (*FileChunkMessage, error) {
	m := new(FileChunkMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BeanAccountServiceServer is the server API for BeanAccountService service.
// All implementations must embed UnimplementedBeanAccountServiceServer
// for forward compatibility
type BeanAccountServiceServer interface {
	SyncLedger(context.Context, *SyncMessage) (*emptypb.Empty, error)
	DownLoadBeanAccountFile(*StringMessage, BeanAccountService_DownLoadBeanAccountFileServer) error
	mustEmbedUnimplementedBeanAccountServiceServer()
}

// UnimplementedBeanAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBeanAccountServiceServer struct {
}

func (UnimplementedBeanAccountServiceServer) SyncLedger(context.Context, *SyncMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncLedger not implemented")
}
func (UnimplementedBeanAccountServiceServer) DownLoadBeanAccountFile(*StringMessage, BeanAccountService_DownLoadBeanAccountFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownLoadBeanAccountFile not implemented")
}
func (UnimplementedBeanAccountServiceServer) mustEmbedUnimplementedBeanAccountServiceServer() {}

// UnsafeBeanAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeanAccountServiceServer will
// result in compilation errors.
type UnsafeBeanAccountServiceServer interface {
	mustEmbedUnimplementedBeanAccountServiceServer()
}

func RegisterBeanAccountServiceServer(s grpc.ServiceRegistrar, srv BeanAccountServiceServer) {
	s.RegisterService(&BeanAccountService_ServiceDesc, srv)
}

func _BeanAccountService_SyncLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeanAccountServiceServer).SyncLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phi.BeanAccountService/SyncLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeanAccountServiceServer).SyncLedger(ctx, req.(*SyncMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeanAccountService_DownLoadBeanAccountFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BeanAccountServiceServer).DownLoadBeanAccountFile(m, &beanAccountServiceDownLoadBeanAccountFileServer{stream})
}

type BeanAccountService_DownLoadBeanAccountFileServer interface {
	Send(*FileChunkMessage) error
	grpc.ServerStream
}

type beanAccountServiceDownLoadBeanAccountFileServer struct {
	grpc.ServerStream
}

func (x *beanAccountServiceDownLoadBeanAccountFileServer) Send(m *FileChunkMessage) error {
	return x.ServerStream.SendMsg(m)
}

// BeanAccountService_ServiceDesc is the grpc.ServiceDesc for BeanAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BeanAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phi.BeanAccountService",
	HandlerType: (*BeanAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncLedger",
			Handler:    _BeanAccountService_SyncLedger_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownLoadBeanAccountFile",
			Handler:       _BeanAccountService_DownLoadBeanAccountFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ledger.proto",
}
