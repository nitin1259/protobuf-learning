// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple/simple.proto

package simplepb

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

type Simple struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple             bool     `protobuf:"varint,2,opt,name=isSimple,proto3" json:"isSimple,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SampleList           []int32  `protobuf:"varint,4,rep,packed,name=sample_list,json=sampleList,proto3" json:"sample_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Simple) Reset()         { *m = Simple{} }
func (m *Simple) String() string { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()    {}
func (*Simple) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3c868e94d57426, []int{0}
}

func (m *Simple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Simple.Unmarshal(m, b)
}
func (m *Simple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Simple.Marshal(b, m, deterministic)
}
func (m *Simple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Simple.Merge(m, src)
}
func (m *Simple) XXX_Size() int {
	return xxx_messageInfo_Simple.Size(m)
}
func (m *Simple) XXX_DiscardUnknown() {
	xxx_messageInfo_Simple.DiscardUnknown(m)
}

var xxx_messageInfo_Simple proto.InternalMessageInfo

func (m *Simple) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Simple) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *Simple) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Simple) GetSampleList() []int32 {
	if m != nil {
		return m.SampleList
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple)(nil), "example.simple.Simple")
}

func init() { proto.RegisterFile("simple/simple.proto", fileDescriptor_9b3c868e94d57426) }

var fileDescriptor_9b3c868e94d57426 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0xa9, 0x15, 0x89,
	0x60, 0x2e, 0x44, 0x54, 0x29, 0x93, 0x8b, 0x2d, 0x18, 0xcc, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0x2c,
	0x86, 0xc8, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x3c, 0x17, 0x77,
	0x31, 0xd8, 0xe8, 0xf8, 0x9c, 0xcc, 0xe2, 0x12, 0x09, 0x16, 0x05, 0x66, 0x0d, 0xd6, 0x20, 0x2e,
	0x88, 0x90, 0x4f, 0x66, 0x71, 0x89, 0x13, 0x57, 0x14, 0x07, 0xc4, 0xd2, 0x82, 0xa4, 0x24, 0x36,
	0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x01, 0xa3, 0xea, 0xa4, 0x00, 0x00,
	0x00,
}