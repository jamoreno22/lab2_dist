// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package DataNode

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeClient interface {
	DistributeChunks(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error)
	UploadBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error)
	DownloadBook(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Book, error)
	//Distribuido
	SendProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*Message, error)
}

type dataNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeClient(cc grpc.ClientConnInterface) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) DistributeChunks(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/DataNode.DataNode/DistributeChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) UploadBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/DataNode.DataNode/UploadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) DownloadBook(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/DataNode.DataNode/DownloadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) SendProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/DataNode.DataNode/SendProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
// All implementations must embed UnimplementedDataNodeServer
// for forward compatibility
type DataNodeServer interface {
	DistributeChunks(context.Context, *Book) (*Message, error)
	UploadBook(context.Context, *Book) (*Message, error)
	DownloadBook(context.Context, *Message) (*Book, error)
	//Distribuido
	SendProposal(context.Context, *Proposal) (*Message, error)
	mustEmbedUnimplementedDataNodeServer()
}

// UnimplementedDataNodeServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (UnimplementedDataNodeServer) DistributeChunks(context.Context, *Book) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeChunks not implemented")
}
func (UnimplementedDataNodeServer) UploadBook(context.Context, *Book) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (UnimplementedDataNodeServer) DownloadBook(context.Context, *Message) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadBook not implemented")
}
func (UnimplementedDataNodeServer) SendProposal(context.Context, *Proposal) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}
func (UnimplementedDataNodeServer) mustEmbedUnimplementedDataNodeServer() {}

// UnsafeDataNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServer will
// result in compilation errors.
type UnsafeDataNodeServer interface {
	mustEmbedUnimplementedDataNodeServer()
}

func RegisterDataNodeServer(s grpc.ServiceRegistrar, srv DataNodeServer) {
	s.RegisterService(&_DataNode_serviceDesc, srv)
}

func _DataNode_DistributeChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).DistributeChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataNode.DataNode/DistributeChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).DistributeChunks(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_UploadBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).UploadBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataNode.DataNode/UploadBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).UploadBook(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_DownloadBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).DownloadBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataNode.DataNode/DownloadBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).DownloadBook(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNode_SendProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Proposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServer).SendProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataNode.DataNode/SendProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).SendProposal(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DataNode.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeChunks",
			Handler:    _DataNode_DistributeChunks_Handler,
		},
		{
			MethodName: "UploadBook",
			Handler:    _DataNode_UploadBook_Handler,
		},
		{
			MethodName: "DownloadBook",
			Handler:    _DataNode_DownloadBook_Handler,
		},
		{
			MethodName: "SendProposal",
			Handler:    _DataNode_SendProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DataNode.proto",
}
