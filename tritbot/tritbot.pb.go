// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tritbot.proto

package tritbot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// InternalMessage has the details of a message that is going to
// be sent under the corporate account. It contains details that
// wouldn't be sent to the Tritter service, such as username of
// the user that made the request.
type InternalMessage struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalMessage) Reset()         { *m = InternalMessage{} }
func (m *InternalMessage) String() string { return proto.CompactTextString(m) }
func (*InternalMessage) ProtoMessage()    {}
func (*InternalMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a7d20fb8dea166, []int{0}
}

func (m *InternalMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalMessage.Unmarshal(m, b)
}
func (m *InternalMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalMessage.Marshal(b, m, deterministic)
}
func (m *InternalMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalMessage.Merge(m, src)
}
func (m *InternalMessage) XXX_Size() int {
	return xxx_messageInfo_InternalMessage.Size(m)
}
func (m *InternalMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InternalMessage proto.InternalMessageInfo

func (m *InternalMessage) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InternalMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InternalMessage)(nil), "tritbot.InternalMessage")
}

func init() { proto.RegisterFile("tritbot.proto", fileDescriptor_d0a7d20fb8dea166) }

var fileDescriptor_d0a7d20fb8dea166 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x29, 0xca, 0x2c,
	0x49, 0xca, 0x2f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xec, 0xb9,
	0xf8, 0x3d, 0xf3, 0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53,
	0x85, 0x84, 0xb8, 0x58, 0x4a, 0x8b, 0x53, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0,
	0x6c, 0x21, 0x09, 0x2e, 0xf6, 0x5c, 0x88, 0xb4, 0x04, 0x13, 0x58, 0x18, 0xc6, 0x75, 0xd2, 0x88,
	0x52, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x28, 0x2d,
	0x49, 0xce, 0xc8, 0xcc, 0x2b, 0xce, 0xcf, 0xd3, 0x07, 0x59, 0x51, 0x92, 0x5a, 0xa4, 0x0f, 0xb5,
	0x2a, 0x89, 0x0d, 0x6c, 0xb5, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xd6, 0x3f, 0x6e, 0x8b,
	0x00, 0x00, 0x00,
}
