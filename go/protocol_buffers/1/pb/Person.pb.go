// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

package person

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

type PhoneType int32

const (
	PhoneType_NONE       PhoneType = 0
	PhoneType_STATIONARY PhoneType = 1
	PhoneType_CELLULAR   PhoneType = 2
	PhoneType_OTHER      PhoneType = 3
)

var PhoneType_name = map[int32]string{
	0: "NONE",
	1: "STATIONARY",
	2: "CELLULAR",
	3: "OTHER",
}

var PhoneType_value = map[string]int32{
	"NONE":       0,
	"STATIONARY": 1,
	"CELLULAR":   2,
	"OTHER":      3,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

type Phone struct {
	Number               string    `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=person.PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_NONE
}

type Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phones               []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

func init() {
	proto.RegisterEnum("person.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "person.Phone")
	proto.RegisterType((*Person)(nil), "person.Person")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x00, 0xf3, 0x94, 0xdc, 0xb8,
	0x58, 0x03, 0x32, 0xf2, 0xf3, 0x52, 0x85, 0xc4, 0xb8, 0xd8, 0xf2, 0x4a, 0x73, 0x93, 0x52, 0x8b,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x55, 0x2e, 0x96, 0x92, 0xca, 0x82,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x41, 0x3d, 0x88, 0x3e, 0x3d, 0xb0, 0xa6, 0x90,
	0xca, 0x82, 0xd4, 0x20, 0xb0, 0xb4, 0x52, 0x30, 0x17, 0x1b, 0xc4, 0x7c, 0x21, 0x3e, 0x2e, 0xa6,
	0xcc, 0x14, 0xb0, 0x21, 0xac, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9,
	0x10, 0x03, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x55, 0x2e, 0xb6, 0x02, 0x90, 0x01, 0xc5, 0x12, 0xcc,
	0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xbc, 0x28, 0xc6, 0x06, 0x41, 0x25, 0xb5, 0xec, 0xb8, 0x38, 0xe1,
	0xf6, 0x08, 0x71, 0x70, 0xb1, 0xf8, 0xf9, 0xfb, 0xb9, 0x0a, 0x30, 0x08, 0xf1, 0x71, 0x71, 0x05,
	0x87, 0x38, 0x86, 0x78, 0xfa, 0xfb, 0x39, 0x06, 0x45, 0x0a, 0x30, 0x0a, 0xf1, 0x70, 0x71, 0x38,
	0xbb, 0xfa, 0xf8, 0x84, 0xfa, 0x38, 0x06, 0x09, 0x30, 0x09, 0x71, 0x72, 0xb1, 0xfa, 0x87, 0x78,
	0xb8, 0x06, 0x09, 0x30, 0x27, 0xb1, 0x81, 0xfd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0x17, 0xfa, 0x44, 0xfb, 0x00, 0x00, 0x00,
}
