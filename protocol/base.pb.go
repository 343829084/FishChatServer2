// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	base.proto
	error.proto
	gateway.proto
	msg_server.proto

It has these top-level messages:
	Base
	Error
	ReqMsgServer
	ResSelectMsgServerForClient
	ReqLogin
	ResLogin
	ReqLogout
	ResLogout
	ReqSendP2PMsg
	ResSendP2PMsg
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// cmdNum = 1001
type Base struct {
	Cmd uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *Base) Reset()                    { *m = Base{} }
func (m *Base) String() string            { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()               {}
func (*Base) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Base)(nil), "protocol.Base")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 71 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x12, 0x5c,
	0x2c, 0x4e, 0x89, 0xc5, 0xa9, 0x42, 0x02, 0x5c, 0xcc, 0xc9, 0xb9, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xbc, 0x41, 0x20, 0x66, 0x12, 0x1b, 0x58, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xde,
	0x93, 0x41, 0xdb, 0x38, 0x00, 0x00, 0x00,
}
