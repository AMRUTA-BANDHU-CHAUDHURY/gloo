// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/subset_spec.proto

package options

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SubsetSpec struct {
	Selectors            []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubsetSpec) Reset()         { *m = SubsetSpec{} }
func (m *SubsetSpec) String() string { return proto.CompactTextString(m) }
func (*SubsetSpec) ProtoMessage()    {}
func (*SubsetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbea30dcd5183d4d, []int{0}
}
func (m *SubsetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubsetSpec.Unmarshal(m, b)
}
func (m *SubsetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubsetSpec.Marshal(b, m, deterministic)
}
func (m *SubsetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubsetSpec.Merge(m, src)
}
func (m *SubsetSpec) XXX_Size() int {
	return xxx_messageInfo_SubsetSpec.Size(m)
}
func (m *SubsetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SubsetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SubsetSpec proto.InternalMessageInfo

func (m *SubsetSpec) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

type Selector struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbea30dcd5183d4d, []int{1}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*SubsetSpec)(nil), "options.gloo.solo.io.SubsetSpec")
	proto.RegisterType((*Selector)(nil), "options.gloo.solo.io.Selector")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/subset_spec.proto", fileDescriptor_dbea30dcd5183d4d)
}

var fileDescriptor_dbea30dcd5183d4d = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0xf3, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x8b, 0x4b, 0x93,
	0x8a, 0x53, 0x4b, 0xe2, 0x8b, 0x0b, 0x52, 0x93, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0xa0, 0x52, 0x7a, 0x20, 0xe5, 0x7a, 0x20, 0x93, 0xf4, 0x32, 0xf3, 0xa5, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x0a, 0xf4, 0x41, 0x2c, 0x88, 0x5a, 0x29, 0xa1, 0xd4, 0x8a, 0x12, 0x88, 0x60, 0x6a,
	0x45, 0x09, 0x44, 0x4c, 0xc9, 0x8b, 0x8b, 0x2b, 0x18, 0x6c, 0x68, 0x70, 0x41, 0x6a, 0xb2, 0x90,
	0x0d, 0x17, 0x67, 0x71, 0x6a, 0x4e, 0x6a, 0x72, 0x49, 0x7e, 0x51, 0xb1, 0x04, 0xa3, 0x02, 0xb3,
	0x06, 0xb7, 0x91, 0x9c, 0x1e, 0x36, 0x1b, 0xf4, 0x82, 0xa1, 0xca, 0x82, 0x10, 0x1a, 0x94, 0xe4,
	0xb8, 0x38, 0x60, 0xc2, 0x42, 0x42, 0x5c, 0x2c, 0xd9, 0xa9, 0x95, 0x10, 0x43, 0x38, 0x83, 0xc0,
	0x6c, 0x27, 0xa7, 0x1d, 0x5f, 0x59, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x8c, 0xb2, 0x20, 0xce, 0xf7,
	0x05, 0xd9, 0xe9, 0x68, 0x21, 0x90, 0xc4, 0x06, 0x76, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0xe5, 0x38, 0xad, 0x40, 0x01, 0x00, 0x00,
}

func (this *SubsetSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubsetSpec)
	if !ok {
		that2, ok := that.(SubsetSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Selectors) != len(that1.Selectors) {
		return false
	}
	for i := range this.Selectors {
		if !this.Selectors[i].Equal(that1.Selectors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Selector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Selector)
	if !ok {
		that2, ok := that.(Selector)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if this.Keys[i] != that1.Keys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
