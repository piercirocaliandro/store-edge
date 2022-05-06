// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_storage

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

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	// Put operation. OpStatus is a message containing the result of the operation
	Put(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error)
	// Get operation. This is a server-side stream RPC, because the server may send multiple values
	Get(ctx context.Context, opts ...grpc.CallOption) (Storage_GetClient, error)
	// Del operation. OpStatus is a message containing the result of the operation
	Del(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error)
	// Append operation
	Append(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error)
	//Update operation, just used server-by-server side
	Update(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (Storage_UpdateClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Put(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error) {
	out := new(OpStatus)
	err := c.cc.Invoke(ctx, "/storage.Storage/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Get(ctx context.Context, opts ...grpc.CallOption) (Storage_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], "/storage.Storage/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageGetClient{stream}
	return x, nil
}

type Storage_GetClient interface {
	Send(*StorageRecordMessage) error
	Recv() (*StorageRecordMessage, error)
	grpc.ClientStream
}

type storageGetClient struct {
	grpc.ClientStream
}

func (x *storageGetClient) Send(m *StorageRecordMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storageGetClient) Recv() (*StorageRecordMessage, error) {
	m := new(StorageRecordMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Del(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error) {
	out := new(OpStatus)
	err := c.cc.Invoke(ctx, "/storage.Storage/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Append(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (*OpStatus, error) {
	out := new(OpStatus)
	err := c.cc.Invoke(ctx, "/storage.Storage/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Update(ctx context.Context, in *StorageRecordMessage, opts ...grpc.CallOption) (Storage_UpdateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[1], "/storage.Storage/Update", opts...)
	if err != nil {
		return nil, err
	}
	x := &storageUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_UpdateClient interface {
	Recv() (*StorageRecordMessage, error)
	grpc.ClientStream
}

type storageUpdateClient struct {
	grpc.ClientStream
}

func (x *storageUpdateClient) Recv() (*StorageRecordMessage, error) {
	m := new(StorageRecordMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	// Put operation. OpStatus is a message containing the result of the operation
	Put(context.Context, *StorageRecordMessage) (*OpStatus, error)
	// Get operation. This is a server-side stream RPC, because the server may send multiple values
	Get(Storage_GetServer) error
	// Del operation. OpStatus is a message containing the result of the operation
	Del(context.Context, *StorageRecordMessage) (*OpStatus, error)
	// Append operation
	Append(context.Context, *StorageRecordMessage) (*OpStatus, error)
	//Update operation, just used server-by-server side
	Update(*StorageRecordMessage, Storage_UpdateServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Put(context.Context, *StorageRecordMessage) (*OpStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedStorageServer) Get(Storage_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) Del(context.Context, *StorageRecordMessage) (*OpStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedStorageServer) Append(context.Context, *StorageRecordMessage) (*OpStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedStorageServer) Update(*StorageRecordMessage, Storage_UpdateServer) error {
	return status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRecordMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Put(ctx, req.(*StorageRecordMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Get(&storageGetServer{stream})
}

type Storage_GetServer interface {
	Send(*StorageRecordMessage) error
	Recv() (*StorageRecordMessage, error)
	grpc.ServerStream
}

type storageGetServer struct {
	grpc.ServerStream
}

func (x *storageGetServer) Send(m *StorageRecordMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storageGetServer) Recv() (*StorageRecordMessage, error) {
	m := new(StorageRecordMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Storage_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRecordMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Del(ctx, req.(*StorageRecordMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageRecordMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Append(ctx, req.(*StorageRecordMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_Update_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StorageRecordMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Update(m, &storageUpdateServer{stream})
}

type Storage_UpdateServer interface {
	Send(*StorageRecordMessage) error
	grpc.ServerStream
}

type storageUpdateServer struct {
	grpc.ServerStream
}

func (x *storageUpdateServer) Send(m *StorageRecordMessage) error {
	return x.ServerStream.SendMsg(m)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Storage_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Storage_Del_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _Storage_Append_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Storage_Get_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Update",
			Handler:       _Storage_Update_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb_storage/Storage.proto",
}
