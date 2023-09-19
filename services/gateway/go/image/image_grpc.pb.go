// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: image.proto

package image

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UploadImage_Upload_FullMethodName = "/image.UploadImage/Upload"
)

// UploadImageClient is the client API for UploadImage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadImageClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (UploadImage_UploadClient, error)
}

type uploadImageClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadImageClient(cc grpc.ClientConnInterface) UploadImageClient {
	return &uploadImageClient{cc}
}

func (c *uploadImageClient) Upload(ctx context.Context, opts ...grpc.CallOption) (UploadImage_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadImage_ServiceDesc.Streams[0], UploadImage_Upload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadImageUploadClient{stream}
	return x, nil
}

type UploadImage_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploadImageUploadClient struct {
	grpc.ClientStream
}

func (x *uploadImageUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadImageUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadImageServer is the server API for UploadImage service.
// All implementations must embed UnimplementedUploadImageServer
// for forward compatibility
type UploadImageServer interface {
	Upload(UploadImage_UploadServer) error
	mustEmbedUnimplementedUploadImageServer()
}

// UnimplementedUploadImageServer must be embedded to have forward compatible implementations.
type UnimplementedUploadImageServer struct {
}

func (UnimplementedUploadImageServer) Upload(UploadImage_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadImageServer) mustEmbedUnimplementedUploadImageServer() {}

// UnsafeUploadImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadImageServer will
// result in compilation errors.
type UnsafeUploadImageServer interface {
	mustEmbedUnimplementedUploadImageServer()
}

func RegisterUploadImageServer(s grpc.ServiceRegistrar, srv UploadImageServer) {
	s.RegisterService(&UploadImage_ServiceDesc, srv)
}

func _UploadImage_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadImageServer).Upload(&uploadImageUploadServer{stream})
}

type UploadImage_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type uploadImageUploadServer struct {
	grpc.ServerStream
}

func (x *uploadImageUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadImageUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadImage_ServiceDesc is the grpc.ServiceDesc for UploadImage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadImage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image.UploadImage",
	HandlerType: (*UploadImageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _UploadImage_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "image.proto",
}