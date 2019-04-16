// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generic-signalling.proto

package v1_generics

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

type StatusCode int32

const (
	StatusCode_SUCC                StatusCode = 0
	StatusCode_ERR_GENERIC         StatusCode = 1
	StatusCode_ERR_GENERIC_SERVER  StatusCode = 2
	StatusCode_ERR_GENERIC_REQUEST StatusCode = 3
	StatusCode_ERR_VALIDATION      StatusCode = 4
	StatusCode_SUCC_PENDING        StatusCode = 100
	StatusCode_SUCC_NOACTION       StatusCode = 101
)

var StatusCode_name = map[int32]string{
	0:   "SUCC",
	1:   "ERR_GENERIC",
	2:   "ERR_GENERIC_SERVER",
	3:   "ERR_GENERIC_REQUEST",
	4:   "ERR_VALIDATION",
	100: "SUCC_PENDING",
	101: "SUCC_NOACTION",
}

var StatusCode_value = map[string]int32{
	"SUCC":                0,
	"ERR_GENERIC":         1,
	"ERR_GENERIC_SERVER":  2,
	"ERR_GENERIC_REQUEST": 3,
	"ERR_VALIDATION":      4,
	"SUCC_PENDING":        100,
	"SUCC_NOACTION":       101,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d6cf46d0cec1f40f, []int{0}
}

type StatusResponse struct {
	Status               StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=v1.generics.StatusCode" json:"status,omitempty"`
	Httpstatus           int32      `protobuf:"varint,2,opt,name=httpstatus,proto3" json:"httpstatus,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6cf46d0cec1f40f, []int{0}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() StatusCode {
	if m != nil {
		return m.Status
	}
	return StatusCode_SUCC
}

func (m *StatusResponse) GetHttpstatus() int32 {
	if m != nil {
		return m.Httpstatus
	}
	return 0
}

func (m *StatusResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("v1.generics.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*StatusResponse)(nil), "v1.generics.StatusResponse")
}

func init() { proto.RegisterFile("generic-signalling.proto", fileDescriptor_d6cf46d0cec1f40f) }

var fileDescriptor_d6cf46d0cec1f40f = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xb6, 0x56, 0x9d, 0x6a, 0x5c, 0x47, 0xb0, 0x39, 0x49, 0xf0, 0x14, 0x04, 0x23,
	0xea, 0x13, 0x84, 0xed, 0x52, 0x02, 0xb2, 0xd5, 0x49, 0xdb, 0x6b, 0x88, 0x76, 0x89, 0x81, 0x9a,
	0x84, 0xee, 0xea, 0xc5, 0xa7, 0xf0, 0x8d, 0x25, 0x31, 0xe2, 0x1e, 0xe7, 0xfb, 0xbf, 0xf9, 0x0f,
	0x3f, 0xf8, 0x85, 0xae, 0xf4, 0xae, 0x7c, 0xbd, 0x31, 0x65, 0x51, 0xe5, 0xdb, 0x6d, 0x59, 0x15,
	0x51, 0xb3, 0xab, 0x6d, 0x8d, 0x93, 0xcf, 0xbb, 0xa8, 0x0f, 0xcd, 0xd5, 0x17, 0x78, 0xa9, 0xcd,
	0xed, 0x87, 0x21, 0x6d, 0x9a, 0xba, 0x32, 0x1a, 0x6f, 0x61, 0x6c, 0x3a, 0xe2, 0xb3, 0x80, 0x85,
	0xde, 0xfd, 0x34, 0x72, 0xfc, 0xe8, 0x57, 0x16, 0xf5, 0x46, 0x53, 0xaf, 0xe1, 0x25, 0xc0, 0x9b,
	0xb5, 0x4d, 0xff, 0x34, 0x08, 0x58, 0xb8, 0x4f, 0x0e, 0x41, 0x1f, 0x0e, 0xde, 0xb5, 0x31, 0x79,
	0xa1, 0xfd, 0x61, 0xc0, 0xc2, 0x23, 0xfa, 0x3b, 0xaf, 0xbf, 0x19, 0xc0, 0x7f, 0x21, 0x1e, 0xc2,
	0x28, 0x5d, 0x09, 0xc1, 0xf7, 0xf0, 0x14, 0x26, 0x92, 0x28, 0x9b, 0x4b, 0x25, 0x29, 0x11, 0x9c,
	0xe1, 0x05, 0xa0, 0x03, 0xb2, 0x54, 0xd2, 0x5a, 0x12, 0x1f, 0xe0, 0x14, 0xce, 0x5d, 0x4e, 0xf2,
	0x79, 0x25, 0xd3, 0x25, 0x1f, 0x22, 0x82, 0xd7, 0x06, 0xeb, 0xf8, 0x31, 0x99, 0xc5, 0xcb, 0x64,
	0xa1, 0xf8, 0x08, 0x39, 0x1c, 0xb7, 0xfd, 0xd9, 0x93, 0x54, 0xb3, 0x44, 0xcd, 0xf9, 0x06, 0xcf,
	0xe0, 0xa4, 0x23, 0x6a, 0x11, 0x8b, 0x4e, 0xd2, 0x2f, 0xe3, 0x6e, 0xa4, 0x87, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc1, 0x81, 0xae, 0x66, 0x40, 0x01, 0x00, 0x00,
}
