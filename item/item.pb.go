// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package item

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Item struct {
	// Item ID
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// collection ID i.e. the parent
	CollectionId int64 `protobuf:"varint,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// ordering integer
	Order uint32 `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	// the content of an item
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// creation timestamp
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	// last updated timestamp
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetCollectionId() int64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *Item) GetOrder() uint32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *Item) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Item) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Item) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateItemRequest struct {
	// The parent resource name where the Item is to be created. (the collection)
	Parent int64 `protobuf:"varint,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The Item resource to create.
	// The field name should match the Noun in the method name.
	Item                 *Item    `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateItemRequest) Reset()         { *m = CreateItemRequest{} }
func (m *CreateItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateItemRequest) ProtoMessage()    {}
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{1}
}

func (m *CreateItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItemRequest.Unmarshal(m, b)
}
func (m *CreateItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItemRequest.Merge(m, src)
}
func (m *CreateItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateItemRequest.Size(m)
}
func (m *CreateItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItemRequest proto.InternalMessageInfo

func (m *CreateItemRequest) GetParent() int64 {
	if m != nil {
		return m.Parent
	}
	return 0
}

func (m *CreateItemRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateItemRequest struct {
	// The Item resource which replaces the resource on the server.
	Item                 *Item    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateItemRequest) Reset()         { *m = UpdateItemRequest{} }
func (m *UpdateItemRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateItemRequest) ProtoMessage()    {}
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{2}
}

func (m *UpdateItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateItemRequest.Unmarshal(m, b)
}
func (m *UpdateItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateItemRequest.Marshal(b, m, deterministic)
}
func (m *UpdateItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateItemRequest.Merge(m, src)
}
func (m *UpdateItemRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateItemRequest.Size(m)
}
func (m *UpdateItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateItemRequest proto.InternalMessageInfo

func (m *UpdateItemRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type DeleteItemRequest struct {
	// The resource name of the Item to be deleted.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteItemRequest) Reset()         { *m = DeleteItemRequest{} }
func (m *DeleteItemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteItemRequest) ProtoMessage()    {}
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6007f868cf6553df, []int{3}
}

func (m *DeleteItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteItemRequest.Unmarshal(m, b)
}
func (m *DeleteItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteItemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteItemRequest.Merge(m, src)
}
func (m *DeleteItemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteItemRequest.Size(m)
}
func (m *DeleteItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteItemRequest proto.InternalMessageInfo

func (m *DeleteItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "Item")
	proto.RegisterType((*CreateItemRequest)(nil), "CreateItemRequest")
	proto.RegisterType((*UpdateItemRequest)(nil), "UpdateItemRequest")
	proto.RegisterType((*DeleteItemRequest)(nil), "DeleteItemRequest")
}

func init() {
	proto.RegisterFile("item.proto", fileDescriptor_6007f868cf6553df)
}

var fileDescriptor_6007f868cf6553df = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x80, 0xbb, 0xfd, 0x7a, 0xe9, 0xf4, 0xad, 0xd0, 0x41, 0x4a, 0x8c, 0x07, 0x43, 0x7a, 0x89,
	0x97, 0x2d, 0xd4, 0x8b, 0x78, 0x13, 0x3f, 0xa0, 0xd7, 0xa8, 0x67, 0x69, 0xb3, 0x63, 0x59, 0x48,
	0xba, 0x31, 0x9d, 0x0a, 0xfe, 0x1a, 0x7f, 0x9a, 0x7f, 0x45, 0x76, 0xd3, 0x18, 0x6d, 0x0e, 0x1e,
	0x67, 0xe6, 0x99, 0xd9, 0x67, 0x86, 0x05, 0xd0, 0x4c, 0x99, 0xcc, 0x0b, 0xc3, 0xc6, 0x3f, 0x5b,
	0x1b, 0xb3, 0x4e, 0x69, 0xe6, 0xa2, 0xd5, 0xee, 0x65, 0xc6, 0x3a, 0xa3, 0x2d, 0x2f, 0xb3, 0x7c,
	0x0f, 0x9c, 0x1e, 0x02, 0x94, 0xe5, 0xfc, 0x5e, 0x16, 0xc3, 0x4f, 0x01, 0xdd, 0x05, 0x53, 0x86,
	0x47, 0xd0, 0xd6, 0xca, 0x13, 0x81, 0x88, 0x3a, 0x71, 0x5b, 0x2b, 0x9c, 0xc2, 0x28, 0x31, 0x69,
	0x4a, 0x09, 0x6b, 0xb3, 0x79, 0xd6, 0xca, 0x6b, 0xbb, 0xd2, 0xff, 0x3a, 0xb9, 0x50, 0x78, 0x0c,
	0x3d, 0x53, 0x28, 0x2a, 0xbc, 0x4e, 0x20, 0xa2, 0x51, 0x5c, 0x06, 0xe8, 0xc1, 0xbf, 0xc4, 0x6c,
	0x98, 0x36, 0xec, 0x75, 0x03, 0x11, 0x0d, 0xe2, 0x2a, 0xc4, 0x4b, 0x18, 0x24, 0x05, 0x2d, 0x99,
	0xd4, 0x35, 0x7b, 0xbd, 0x40, 0x44, 0xc3, 0xb9, 0x2f, 0x4b, 0x3d, 0x59, 0xe9, 0xc9, 0xc7, 0xca,
	0x3f, 0xae, 0x61, 0xdb, 0xb9, 0xcb, 0xd5, 0xbe, 0xb3, 0xff, 0x77, 0xe7, 0x37, 0x1c, 0xde, 0xc3,
	0xf8, 0xc6, 0x8d, 0xb1, 0x6b, 0xc6, 0xf4, 0xba, 0xa3, 0x2d, 0xe3, 0x04, 0xfa, 0xf9, 0xb2, 0xb0,
	0x86, 0xe5, 0xc6, 0xfb, 0x08, 0x4f, 0xa0, 0x6b, 0x4f, 0xeb, 0xf6, 0x19, 0xce, 0x7b, 0xd2, 0xf5,
	0xb8, 0x54, 0x28, 0x61, 0xfc, 0xe4, 0x86, 0xfe, 0x9c, 0x53, 0xf1, 0xa2, 0xc9, 0x4f, 0x61, 0x7c,
	0x4b, 0x29, 0xfd, 0xe6, 0x0f, 0xae, 0x3c, 0xff, 0x10, 0x30, 0xb4, 0xf5, 0x07, 0x2a, 0xde, 0x74,
	0x42, 0x78, 0x0e, 0x50, 0xcb, 0x22, 0xca, 0x86, 0xb9, 0x5f, 0xbe, 0x11, 0xb6, 0x2c, 0x5a, 0xfb,
	0x20, 0xca, 0x86, 0x5c, 0x8d, 0x5e, 0x01, 0xd4, 0x2a, 0x88, 0xb2, 0xe1, 0xe5, 0x4f, 0x1a, 0xb7,
	0xbc, 0xb3, 0x9f, 0x24, 0x6c, 0xad, 0xfa, 0x2e, 0x73, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0x37, 0x34, 0xd1, 0x73, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ItemServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Item, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Item, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ItemService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ItemService/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ItemService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
type ItemServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*Item, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*Item, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*empty.Empty, error)
}

// UnimplementedItemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (*UnimplementedItemServiceServer) CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (*UnimplementedItemServiceServer) UpdateItem(ctx context.Context, req *UpdateItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (*UnimplementedItemServiceServer) DeleteItem(ctx context.Context, req *DeleteItemRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}

func RegisterItemServiceServer(s *grpc.Server, srv ItemServiceServer) {
	s.RegisterService(&_ItemService_serviceDesc, srv)
}

func _ItemService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ItemService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _ItemService_CreateItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ItemService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ItemService_DeleteItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "item.proto",
}
