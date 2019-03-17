// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package main

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

type NumberOfHits struct {
	NumHits              int32    `protobuf:"varint,1,opt,name=numHits,proto3" json:"numHits,omitempty"`
	HitKey               string   `protobuf:"bytes,2,opt,name=hitKey,proto3" json:"hitKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberOfHits) Reset()         { *m = NumberOfHits{} }
func (m *NumberOfHits) String() string { return proto.CompactTextString(m) }
func (*NumberOfHits) ProtoMessage()    {}
func (*NumberOfHits) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *NumberOfHits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberOfHits.Unmarshal(m, b)
}
func (m *NumberOfHits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberOfHits.Marshal(b, m, deterministic)
}
func (m *NumberOfHits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberOfHits.Merge(m, src)
}
func (m *NumberOfHits) XXX_Size() int {
	return xxx_messageInfo_NumberOfHits.Size(m)
}
func (m *NumberOfHits) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberOfHits.DiscardUnknown(m)
}

var xxx_messageInfo_NumberOfHits proto.InternalMessageInfo

func (m *NumberOfHits) GetNumHits() int32 {
	if m != nil {
		return m.NumHits
	}
	return 0
}

func (m *NumberOfHits) GetHitKey() string {
	if m != nil {
		return m.HitKey
	}
	return ""
}

func init() {
	proto.RegisterType((*NumberOfHits)(nil), "main.NumberOfHits")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf1, 0xcd, 0x4f, 0x49,
	0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x72,
	0xe0, 0xe2, 0xf1, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xf2, 0x4f, 0xf3, 0xc8, 0x2c, 0x29, 0x16, 0x92,
	0xe0, 0x62, 0xcf, 0x2b, 0xcd, 0x05, 0x31, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x60, 0x5c,
	0x21, 0x31, 0x2e, 0xb6, 0x8c, 0xcc, 0x12, 0xef, 0xd4, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x28, 0x2f, 0x89, 0x0d, 0x6c, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x99, 0x01, 0x56,
	0xf4, 0x5e, 0x00, 0x00, 0x00,
}