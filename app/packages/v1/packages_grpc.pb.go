// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// PackageServiceClient is the client API for PackageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PackageServiceClient interface {
	// CreatePackage uploads a package to the cloud
	CreatePackage(ctx context.Context, opts ...grpc.CallOption) (PackageService_CreatePackageClient, error)
	// DeletePackage removes the given package versions
	DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error)
	// GetPackage returns the metadata for a requested package version. It also returns a URL
	// for downloading the package if one is requested.
	GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error)
	// ListPackages gets the metadata for the requested packages. Include package name, version, and/or
	// type to filter beyond the required organization_id. ListPackages also returns URLs for
	// downloading each package if they are requested.
	ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error)
}

type packageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPackageServiceClient(cc grpc.ClientConnInterface) PackageServiceClient {
	return &packageServiceClient{cc}
}

func (c *packageServiceClient) CreatePackage(ctx context.Context, opts ...grpc.CallOption) (PackageService_CreatePackageClient, error) {
	stream, err := c.cc.NewStream(ctx, &PackageService_ServiceDesc.Streams[0], "/viam.app.packages.v1.PackageService/CreatePackage", opts...)
	if err != nil {
		return nil, err
	}
	x := &packageServiceCreatePackageClient{stream}
	return x, nil
}

type PackageService_CreatePackageClient interface {
	Send(*CreatePackageRequest) error
	CloseAndRecv() (*CreatePackageResponse, error)
	grpc.ClientStream
}

type packageServiceCreatePackageClient struct {
	grpc.ClientStream
}

func (x *packageServiceCreatePackageClient) Send(m *CreatePackageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *packageServiceCreatePackageClient) CloseAndRecv() (*CreatePackageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreatePackageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *packageServiceClient) DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error) {
	out := new(DeletePackageResponse)
	err := c.cc.Invoke(ctx, "/viam.app.packages.v1.PackageService/DeletePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error) {
	out := new(GetPackageResponse)
	err := c.cc.Invoke(ctx, "/viam.app.packages.v1.PackageService/GetPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packageServiceClient) ListPackages(ctx context.Context, in *ListPackagesRequest, opts ...grpc.CallOption) (*ListPackagesResponse, error) {
	out := new(ListPackagesResponse)
	err := c.cc.Invoke(ctx, "/viam.app.packages.v1.PackageService/ListPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PackageServiceServer is the server API for PackageService service.
// All implementations must embed UnimplementedPackageServiceServer
// for forward compatibility
type PackageServiceServer interface {
	// CreatePackage uploads a package to the cloud
	CreatePackage(PackageService_CreatePackageServer) error
	// DeletePackage removes the given package versions
	DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error)
	// GetPackage returns the metadata for a requested package version. It also returns a URL
	// for downloading the package if one is requested.
	GetPackage(context.Context, *GetPackageRequest) (*GetPackageResponse, error)
	// ListPackages gets the metadata for the requested packages. Include package name, version, and/or
	// type to filter beyond the required organization_id. ListPackages also returns URLs for
	// downloading each package if they are requested.
	ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error)
	mustEmbedUnimplementedPackageServiceServer()
}

// UnimplementedPackageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPackageServiceServer struct {
}

func (UnimplementedPackageServiceServer) CreatePackage(PackageService_CreatePackageServer) error {
	return status.Errorf(codes.Unimplemented, "method CreatePackage not implemented")
}
func (UnimplementedPackageServiceServer) DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackage not implemented")
}
func (UnimplementedPackageServiceServer) GetPackage(context.Context, *GetPackageRequest) (*GetPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackage not implemented")
}
func (UnimplementedPackageServiceServer) ListPackages(context.Context, *ListPackagesRequest) (*ListPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPackages not implemented")
}
func (UnimplementedPackageServiceServer) mustEmbedUnimplementedPackageServiceServer() {}

// UnsafePackageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PackageServiceServer will
// result in compilation errors.
type UnsafePackageServiceServer interface {
	mustEmbedUnimplementedPackageServiceServer()
}

func RegisterPackageServiceServer(s grpc.ServiceRegistrar, srv PackageServiceServer) {
	s.RegisterService(&PackageService_ServiceDesc, srv)
}

func _PackageService_CreatePackage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PackageServiceServer).CreatePackage(&packageServiceCreatePackageServer{stream})
}

type PackageService_CreatePackageServer interface {
	SendAndClose(*CreatePackageResponse) error
	Recv() (*CreatePackageRequest, error)
	grpc.ServerStream
}

type packageServiceCreatePackageServer struct {
	grpc.ServerStream
}

func (x *packageServiceCreatePackageServer) SendAndClose(m *CreatePackageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *packageServiceCreatePackageServer) Recv() (*CreatePackageRequest, error) {
	m := new(CreatePackageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PackageService_DeletePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).DeletePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.packages.v1.PackageService/DeletePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).DeletePackage(ctx, req.(*DeletePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_GetPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).GetPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.packages.v1.PackageService/GetPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).GetPackage(ctx, req.(*GetPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PackageService_ListPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PackageServiceServer).ListPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.app.packages.v1.PackageService/ListPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PackageServiceServer).ListPackages(ctx, req.(*ListPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PackageService_ServiceDesc is the grpc.ServiceDesc for PackageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PackageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.app.packages.v1.PackageService",
	HandlerType: (*PackageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeletePackage",
			Handler:    _PackageService_DeletePackage_Handler,
		},
		{
			MethodName: "GetPackage",
			Handler:    _PackageService_GetPackage_Handler,
		},
		{
			MethodName: "ListPackages",
			Handler:    _PackageService_ListPackages_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreatePackage",
			Handler:       _PackageService_CreatePackage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "app/packages/v1/packages.proto",
}
