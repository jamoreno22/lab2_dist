// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DataNode.proto

package lab2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Book struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Chunks               [][]byte `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4c9d8f36887fb, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Book) GetChunks() [][]byte {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4c9d8f36887fb, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Proposal struct {
	Ids                  []int32  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4c9d8f36887fb, []int{2}
}

func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "lab2.Book")
	proto.RegisterType((*Message)(nil), "lab2.Message")
	proto.RegisterType((*Proposal)(nil), "lab2.Proposal")
}

func init() {
	proto.RegisterFile("DataNode.proto", fileDescriptor_2cc4c9d8f36887fb)
}

var fileDescriptor_2cc4c9d8f36887fb = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0x49, 0x2c, 0x49,
	0xf4, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x49, 0x4c, 0x32,
	0x52, 0x32, 0xe2, 0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x92, 0x33, 0x4a,
	0xf3, 0xb2, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35, 0x78, 0x82, 0xa0, 0x3c, 0x25, 0x59, 0x2e, 0x76,
	0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x90, 0xb6, 0x92, 0xd4, 0x8a, 0x12, 0x98, 0x36, 0x10,
	0x5b, 0x49, 0x86, 0x8b, 0x23, 0xa0, 0x28, 0xbf, 0x20, 0xbf, 0x38, 0x31, 0x47, 0x48, 0x80, 0x8b,
	0x39, 0x33, 0xa5, 0x18, 0x28, 0xcd, 0xac, 0xc1, 0x1a, 0x04, 0x62, 0x1a, 0x1d, 0x61, 0xe4, 0xe2,
	0x80, 0xb9, 0x44, 0x48, 0x9f, 0x4b, 0xc0, 0x25, 0xb3, 0xb8, 0xa4, 0x28, 0x33, 0xa9, 0xb4, 0x24,
	0xd5, 0x19, 0x6c, 0xba, 0x10, 0x97, 0x1e, 0xc8, 0x61, 0x7a, 0x20, 0x57, 0x49, 0xf1, 0x42, 0xd8,
	0x50, 0xdb, 0x94, 0x18, 0x84, 0x34, 0xb9, 0xb8, 0x42, 0x0b, 0x72, 0xf2, 0x13, 0x53, 0xc0, 0x8e,
	0xc6, 0xab, 0x54, 0x9b, 0x8b, 0xc7, 0x25, 0xbf, 0x3c, 0x0f, 0xae, 0x18, 0x55, 0x81, 0x14, 0x92,
	0x5e, 0xa0, 0x62, 0x7d, 0x2e, 0x9e, 0xe0, 0xd4, 0xbc, 0x14, 0xb8, 0xbb, 0xf9, 0x20, 0xb2, 0x30,
	0x3e, 0x86, 0xe9, 0x49, 0x6c, 0xe0, 0x40, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xcc,
	0x22, 0x5c, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataNodeClient is the client API for DataNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	err := c.cc.Invoke(ctx, "/lab2.DataNode/DistributeChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) UploadBook(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab2.DataNode/UploadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) DownloadBook(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/lab2.DataNode/DownloadBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeClient) SendProposal(ctx context.Context, in *Proposal, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab2.DataNode/SendProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServer is the server API for DataNode service.
type DataNodeServer interface {
	DistributeChunks(context.Context, *Book) (*Message, error)
	UploadBook(context.Context, *Book) (*Message, error)
	DownloadBook(context.Context, *Message) (*Book, error)
	//Distribuido
	SendProposal(context.Context, *Proposal) (*Message, error)
}

// UnimplementedDataNodeServer can be embedded to have forward compatible implementations.
type UnimplementedDataNodeServer struct {
}

func (*UnimplementedDataNodeServer) DistributeChunks(ctx context.Context, req *Book) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeChunks not implemented")
}
func (*UnimplementedDataNodeServer) UploadBook(ctx context.Context, req *Book) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadBook not implemented")
}
func (*UnimplementedDataNodeServer) DownloadBook(ctx context.Context, req *Message) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadBook not implemented")
}
func (*UnimplementedDataNodeServer) SendProposal(ctx context.Context, req *Proposal) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProposal not implemented")
}

func RegisterDataNodeServer(s *grpc.Server, srv DataNodeServer) {
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
		FullMethod: "/lab2.DataNode/DistributeChunks",
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
		FullMethod: "/lab2.DataNode/UploadBook",
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
		FullMethod: "/lab2.DataNode/DownloadBook",
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
		FullMethod: "/lab2.DataNode/SendProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServer).SendProposal(ctx, req.(*Proposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab2.DataNode",
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
