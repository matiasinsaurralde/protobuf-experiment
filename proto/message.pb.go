// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package experiment is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
*/
package experiment

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

type Message struct {
	Body             *string `protobuf:"bytes,1,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "experiment.Message")
}

var fileDescriptor0 = []byte{
	// 72 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xad, 0x28, 0x48, 0x2d,
	0xca, 0xcc, 0x4d, 0xcd, 0x2b, 0x51, 0x12, 0xe7, 0x62, 0xf7, 0x85, 0x48, 0x0a, 0xf1, 0x70, 0xb1,
	0x24, 0xe5, 0xa7, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0xdf, 0x69, 0x52, 0x34, 0x00, 0x00, 0x00,
}
