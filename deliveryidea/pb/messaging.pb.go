// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messaging.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Message_Type int32

const (
	Message_NONE  Message_Type = 0
	Message_EMAIL Message_Type = 1
	Message_TEXT  Message_Type = 2
)

var Message_Type_name = map[int32]string{
	0: "NONE",
	1: "EMAIL",
	2: "TEXT",
}

var Message_Type_value = map[string]int32{
	"NONE":  0,
	"EMAIL": 1,
	"TEXT":  2,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{0, 0}
}

type Message struct {
	Message              string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Recipients           []string     `protobuf:"bytes,2,rep,name=recipients,proto3" json:"recipients,omitempty"`
	Type                 Message_Type `protobuf:"varint,3,opt,name=type,proto3,enum=pb.Message_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_42a1718997f046ec, []int{0}
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

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_NONE
}

func init() {
	proto.RegisterEnum("pb.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("messaging.proto", fileDescriptor_42a1718997f046ec) }

var fileDescriptor_42a1718997f046ec = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0xcf, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92,
	0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0xa7, 0xe6,
	0x16, 0x94, 0x54, 0x42, 0x14, 0x28, 0x4d, 0x60, 0xe4, 0x62, 0xf7, 0x05, 0x6b, 0x4a, 0x15, 0x92,
	0xe0, 0x62, 0x87, 0xe8, 0x4f, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xe4,
	0xb8, 0xb8, 0x8a, 0x52, 0x93, 0x33, 0x0b, 0x32, 0x53, 0xf3, 0x4a, 0x8a, 0x25, 0x98, 0x14, 0x98,
	0x35, 0x38, 0x83, 0x90, 0x44, 0x84, 0x54, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98, 0x15,
	0x18, 0x35, 0xf8, 0x8c, 0x04, 0xf4, 0x0a, 0x92, 0xf4, 0xa0, 0x86, 0xea, 0x85, 0x54, 0x16, 0xa4,
	0x06, 0x81, 0x65, 0x95, 0x54, 0xb9, 0x58, 0x40, 0x3c, 0x21, 0x0e, 0x2e, 0x16, 0x3f, 0x7f, 0x3f,
	0x57, 0x01, 0x06, 0x21, 0x4e, 0x2e, 0x56, 0x57, 0x5f, 0x47, 0x4f, 0x1f, 0x01, 0x46, 0x90, 0x60,
	0x88, 0x6b, 0x44, 0x88, 0x00, 0x93, 0x91, 0x3d, 0x17, 0x1f, 0x54, 0x73, 0x70, 0x6a, 0x51, 0x59,
	0x66, 0x72, 0xaa, 0x90, 0x2e, 0x17, 0x4b, 0x70, 0x6a, 0x5e, 0x8a, 0x10, 0x37, 0x92, 0xc1, 0x52,
	0x62, 0x7a, 0x10, 0x7f, 0xe9, 0xc1, 0xfc, 0xa5, 0xe7, 0x0a, 0xf2, 0x97, 0x12, 0x43, 0x12, 0x1b,
	0x58, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x71, 0x68, 0xe4, 0xa6, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	//3. send message to delivery guys
	//7. send message to customer that order is in transit
	//10. send message to customer that message is delivered
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.MessageService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	//3. send message to delivery guys
	//7. send message to customer that order is in transit
	//10. send message to customer that message is delivered
	Send(context.Context, *Message) (*empty.Empty, error)
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) Send(ctx context.Context, req *Message) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MessageService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messaging.proto",
}
