// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery.proto

package pb

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

type Delivery_Status int32

const (
	Delivery_NONE       Delivery_Status = 0
	Delivery_AVAILABLE  Delivery_Status = 1
	Delivery_DELIVERING Delivery_Status = 2
)

var Delivery_Status_name = map[int32]string{
	0: "NONE",
	1: "AVAILABLE",
	2: "DELIVERING",
}

var Delivery_Status_value = map[string]int32{
	"NONE":       0,
	"AVAILABLE":  1,
	"DELIVERING": 2,
}

func (x Delivery_Status) String() string {
	return proto.EnumName(Delivery_Status_name, int32(x))
}

func (Delivery_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{0, 0}
}

type Delivery struct {
	DeliveryId           string          `protobuf:"bytes,1,opt,name=deliveryId,proto3" json:"deliveryId,omitempty"`
	Location             string          `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string          `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Status               Delivery_Status `protobuf:"varint,5,opt,name=status,proto3,enum=pb.Delivery_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Delivery) Reset()         { *m = Delivery{} }
func (m *Delivery) String() string { return proto.CompactTextString(m) }
func (*Delivery) ProtoMessage()    {}
func (*Delivery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{0}
}

func (m *Delivery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delivery.Unmarshal(m, b)
}
func (m *Delivery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delivery.Marshal(b, m, deterministic)
}
func (m *Delivery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delivery.Merge(m, src)
}
func (m *Delivery) XXX_Size() int {
	return xxx_messageInfo_Delivery.Size(m)
}
func (m *Delivery) XXX_DiscardUnknown() {
	xxx_messageInfo_Delivery.DiscardUnknown(m)
}

var xxx_messageInfo_Delivery proto.InternalMessageInfo

func (m *Delivery) GetDeliveryId() string {
	if m != nil {
		return m.DeliveryId
	}
	return ""
}

func (m *Delivery) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Delivery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Delivery) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Delivery) GetStatus() Delivery_Status {
	if m != nil {
		return m.Status
	}
	return Delivery_NONE
}

type DeliveryOrder struct {
	DeliveryId           string   `protobuf:"bytes,1,opt,name=deliveryId,proto3" json:"deliveryId,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliveryOrder) Reset()         { *m = DeliveryOrder{} }
func (m *DeliveryOrder) String() string { return proto.CompactTextString(m) }
func (*DeliveryOrder) ProtoMessage()    {}
func (*DeliveryOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf387bcb4e23d880, []int{1}
}

func (m *DeliveryOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryOrder.Unmarshal(m, b)
}
func (m *DeliveryOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryOrder.Marshal(b, m, deterministic)
}
func (m *DeliveryOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryOrder.Merge(m, src)
}
func (m *DeliveryOrder) XXX_Size() int {
	return xxx_messageInfo_DeliveryOrder.Size(m)
}
func (m *DeliveryOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryOrder.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryOrder proto.InternalMessageInfo

func (m *DeliveryOrder) GetDeliveryId() string {
	if m != nil {
		return m.DeliveryId
	}
	return ""
}

func (m *DeliveryOrder) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.Delivery_Status", Delivery_Status_name, Delivery_Status_value)
	proto.RegisterType((*Delivery)(nil), "pb.Delivery")
	proto.RegisterType((*DeliveryOrder)(nil), "pb.DeliveryOrder")
}

func init() { proto.RegisterFile("delivery.proto", fileDescriptor_bf387bcb4e23d880) }

var fileDescriptor_bf387bcb4e23d880 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x51, 0x6b, 0x82, 0x60,
	0x14, 0x4d, 0x57, 0xce, 0xee, 0xca, 0xda, 0xdd, 0x18, 0xe2, 0x60, 0x84, 0x4f, 0xc1, 0xc0, 0xb6,
	0x7a, 0xde, 0xc0, 0x2d, 0x19, 0x42, 0xd4, 0x30, 0xe8, 0x3d, 0xf5, 0x56, 0x82, 0xfa, 0x89, 0x5a,
	0xb0, 0x9f, 0xb3, 0x7f, 0xb4, 0x9f, 0x34, 0xfa, 0xd4, 0x68, 0x7b, 0x58, 0xb0, 0xb7, 0x7b, 0xce,
	0xbd, 0xe7, 0x7c, 0xf7, 0x7e, 0x07, 0x14, 0x9f, 0xc2, 0x60, 0x47, 0xe9, 0x87, 0x91, 0xa4, 0x2c,
	0x67, 0x28, 0x26, 0xae, 0x76, 0xbb, 0x66, 0x6c, 0x1d, 0xd2, 0x80, 0x33, 0xee, 0x76, 0x35, 0xa0,
	0x28, 0xc9, 0xcb, 0x01, 0xed, 0x82, 0xa5, 0x3e, 0xa5, 0x05, 0xd0, 0xbf, 0x04, 0x90, 0xc7, 0xa5,
	0x01, 0xde, 0x01, 0x54, 0x66, 0xb6, 0xaf, 0x0a, 0x3d, 0xa1, 0xdf, 0x74, 0x8e, 0x18, 0xd4, 0x40,
	0x0e, 0x99, 0xb7, 0xcc, 0x03, 0x16, 0xab, 0x22, 0xef, 0x1e, 0x30, 0x22, 0xd4, 0xe3, 0x65, 0x44,
	0xea, 0x19, 0xe7, 0x79, 0x8d, 0xd7, 0xd0, 0x48, 0x36, 0x2c, 0x26, 0xb5, 0xce, 0xc9, 0x02, 0xe0,
	0x3d, 0x48, 0x59, 0xbe, 0xcc, 0xb7, 0x99, 0xda, 0xe8, 0x09, 0x7d, 0x65, 0x78, 0x65, 0x24, 0xae,
	0x51, 0xed, 0x60, 0xcc, 0x79, 0xcb, 0x29, 0x47, 0xf4, 0x47, 0x90, 0x0a, 0x06, 0x65, 0xa8, 0x4f,
	0x67, 0x53, 0xab, 0x5b, 0xc3, 0x36, 0x34, 0xcd, 0x85, 0x69, 0x4f, 0xcc, 0x97, 0x89, 0xd5, 0x15,
	0x50, 0x01, 0x18, 0x5b, 0x13, 0x7b, 0x61, 0x39, 0xf6, 0xf4, 0xad, 0x2b, 0xea, 0x36, 0xb4, 0x2b,
	0xb7, 0xd9, 0xfe, 0xd2, 0x93, 0x67, 0xa9, 0x70, 0xce, 0xbf, 0xc4, 0xf6, 0xcb, 0xab, 0x2a, 0x38,
	0xfc, 0x14, 0xa1, 0x53, 0x79, 0xcd, 0x29, 0xdd, 0x05, 0x1e, 0x61, 0x1f, 0x64, 0x87, 0xd6, 0x41,
	0x96, 0x53, 0x8a, 0xad, 0xe3, 0xd5, 0xb5, 0x1f, 0x48, 0xaf, 0xe1, 0x03, 0x48, 0x0e, 0x45, 0x6c,
	0x47, 0xbf, 0xe6, 0x6e, 0x8c, 0x22, 0x1e, 0xa3, 0x8a, 0xc7, 0xb0, 0xf6, 0xf1, 0xe8, 0x35, 0x1c,
	0x41, 0xeb, 0x7d, 0xeb, 0x86, 0x41, 0xb6, 0x29, 0x36, 0x6f, 0xee, 0x75, 0xbc, 0xfc, 0x43, 0xf4,
	0x04, 0x8a, 0xe9, 0x79, 0x94, 0xe4, 0x87, 0x1c, 0x2f, 0x8f, 0x9f, 0x3b, 0x25, 0x7f, 0x86, 0xce,
	0x2b, 0x8b, 0x57, 0x41, 0x1a, 0xfd, 0x4b, 0xef, 0x4a, 0x9c, 0x19, 0x7d, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x47, 0xfa, 0xd7, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	Register(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*Delivery, error)
	Remove(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*empty.Empty, error)
	//2. filters delivery guys by status, calls Messaging/SendMessage
	PublishOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*empty.Empty, error)
	//5. delivery person accepts delivery, change Delivery status
	//5. calls Order/DeliveringOrder, calls Messaging/SendMessage
	AcceptDelivery(ctx context.Context, in *DeliveryOrder, opts ...grpc.CallOption) (*empty.Empty, error)
	//8. change status of Delivery to available
	//8. calls Order/DeliveredOrder
	//8. calls Messaging/SendMessage
	ConfirmDelivery(ctx context.Context, in *DeliveryOrder, opts ...grpc.CallOption) (*empty.Empty, error)
}

type deliveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeliveryServiceClient(cc *grpc.ClientConn) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) Register(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*Delivery, error) {
	out := new(Delivery)
	err := c.cc.Invoke(ctx, "/pb.DeliveryService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) Remove(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DeliveryService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) PublishOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DeliveryService/PublishOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) AcceptDelivery(ctx context.Context, in *DeliveryOrder, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DeliveryService/AcceptDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) ConfirmDelivery(ctx context.Context, in *DeliveryOrder, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DeliveryService/ConfirmDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
type DeliveryServiceServer interface {
	Register(context.Context, *Delivery) (*Delivery, error)
	Remove(context.Context, *Delivery) (*empty.Empty, error)
	//2. filters delivery guys by status, calls Messaging/SendMessage
	PublishOrder(context.Context, *Order) (*empty.Empty, error)
	//5. delivery person accepts delivery, change Delivery status
	//5. calls Order/DeliveringOrder, calls Messaging/SendMessage
	AcceptDelivery(context.Context, *DeliveryOrder) (*empty.Empty, error)
	//8. change status of Delivery to available
	//8. calls Order/DeliveredOrder
	//8. calls Messaging/SendMessage
	ConfirmDelivery(context.Context, *DeliveryOrder) (*empty.Empty, error)
}

// UnimplementedDeliveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeliveryServiceServer struct {
}

func (*UnimplementedDeliveryServiceServer) Register(ctx context.Context, req *Delivery) (*Delivery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedDeliveryServiceServer) Remove(ctx context.Context, req *Delivery) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedDeliveryServiceServer) PublishOrder(ctx context.Context, req *Order) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishOrder not implemented")
}
func (*UnimplementedDeliveryServiceServer) AcceptDelivery(ctx context.Context, req *DeliveryOrder) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDelivery not implemented")
}
func (*UnimplementedDeliveryServiceServer) ConfirmDelivery(ctx context.Context, req *DeliveryOrder) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDelivery not implemented")
}

func RegisterDeliveryServiceServer(s *grpc.Server, srv DeliveryServiceServer) {
	s.RegisterService(&_DeliveryService_serviceDesc, srv)
}

func _DeliveryService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeliveryService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).Register(ctx, req.(*Delivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeliveryService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).Remove(ctx, req.(*Delivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_PublishOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).PublishOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeliveryService/PublishOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).PublishOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_AcceptDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).AcceptDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeliveryService/AcceptDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).AcceptDelivery(ctx, req.(*DeliveryOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_ConfirmDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).ConfirmDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DeliveryService/ConfirmDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).ConfirmDelivery(ctx, req.(*DeliveryOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeliveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DeliveryService_Register_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _DeliveryService_Remove_Handler,
		},
		{
			MethodName: "PublishOrder",
			Handler:    _DeliveryService_PublishOrder_Handler,
		},
		{
			MethodName: "AcceptDelivery",
			Handler:    _DeliveryService_AcceptDelivery_Handler,
		},
		{
			MethodName: "ConfirmDelivery",
			Handler:    _DeliveryService_ConfirmDelivery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery.proto",
}
