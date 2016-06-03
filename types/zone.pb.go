// Code generated by protoc-gen-go.
// source: zone.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	zone.proto

It has these top-level messages:
	Record
	Records
*/
package types

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

type Record_Type int32

const (
	Record_UNKNOWN Record_Type = 0
	Record_SOA     Record_Type = 1
	Record_NS      Record_Type = 2
	Record_A       Record_Type = 3
	Record_CNAME   Record_Type = 4
	Record_MX      Record_Type = 5
	Record_TXT     Record_Type = 6
	Record_AAAA    Record_Type = 7
	Record_SRV     Record_Type = 8
	Record_PTR     Record_Type = 9
	Record_ALIAS   Record_Type = 10
)

var Record_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "SOA",
	2:  "NS",
	3:  "A",
	4:  "CNAME",
	5:  "MX",
	6:  "TXT",
	7:  "AAAA",
	8:  "SRV",
	9:  "PTR",
	10: "ALIAS",
}
var Record_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"SOA":     1,
	"NS":      2,
	"A":       3,
	"CNAME":   4,
	"MX":      5,
	"TXT":     6,
	"AAAA":    7,
	"SRV":     8,
	"PTR":     9,
	"ALIAS":   10,
}

func (x Record_Type) String() string {
	return proto.EnumName(Record_Type_name, int32(x))
}
func (Record_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Record struct {
	Name  string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ttl   int32           `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	Type  string          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	State int32           `protobuf:"varint,4,opt,name=state" json:"state,omitempty"`
	Value []*Record_Value `protobuf:"bytes,5,rep,name=value" json:"value,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetValue() []*Record_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type Record_Value struct {
	Record     []string          `protobuf:"bytes,1,rep,name=record" json:"record,omitempty"`
	View       string            `protobuf:"bytes,2,opt,name=view" json:"view,omitempty"`
	Weight     int32             `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Continent  string            `protobuf:"bytes,4,opt,name=continent" json:"continent,omitempty"`
	Country    string            `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	Soa        *Record_Value_Soa `protobuf:"bytes,6,opt,name=soa" json:"soa,omitempty"`
	Preference int32             `protobuf:"varint,7,opt,name=preference" json:"preference,omitempty"`
}

func (m *Record_Value) Reset()                    { *m = Record_Value{} }
func (m *Record_Value) String() string            { return proto.CompactTextString(m) }
func (*Record_Value) ProtoMessage()               {}
func (*Record_Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Record_Value) GetSoa() *Record_Value_Soa {
	if m != nil {
		return m.Soa
	}
	return nil
}

type Record_Value_Soa struct {
	Mname   string `protobuf:"bytes,1,opt,name=mname" json:"mname,omitempty"`
	Nname   string `protobuf:"bytes,2,opt,name=nname" json:"nname,omitempty"`
	Serial  uint32 `protobuf:"varint,3,opt,name=serial" json:"serial,omitempty"`
	Refresh uint32 `protobuf:"varint,4,opt,name=refresh" json:"refresh,omitempty"`
	Retry   uint32 `protobuf:"varint,5,opt,name=retry" json:"retry,omitempty"`
	Expire  uint32 `protobuf:"varint,6,opt,name=expire" json:"expire,omitempty"`
	Minttl  uint32 `protobuf:"varint,7,opt,name=minttl" json:"minttl,omitempty"`
}

func (m *Record_Value_Soa) Reset()                    { *m = Record_Value_Soa{} }
func (m *Record_Value_Soa) String() string            { return proto.CompactTextString(m) }
func (*Record_Value_Soa) ProtoMessage()               {}
func (*Record_Value_Soa) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0, 0} }

type Records struct {
	Domain  string    `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Records []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
}

func (m *Records) Reset()                    { *m = Records{} }
func (m *Records) String() string            { return proto.CompactTextString(m) }
func (*Records) ProtoMessage()               {}
func (*Records) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Records) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "types.Record")
	proto.RegisterType((*Record_Value)(nil), "types.Record.Value")
	proto.RegisterType((*Record_Value_Soa)(nil), "types.Record.Value.Soa")
	proto.RegisterType((*Records)(nil), "types.Records")
	proto.RegisterEnum("types.Record_Type", Record_Type_name, Record_Type_value)
}

var fileDescriptor0 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x25, 0x9b, 0x38, 0x69, 0x66, 0xb5, 0x92, 0x65, 0x10, 0x44, 0x15, 0x42, 0xd5, 0x5e, 0x28,
	0x97, 0x1c, 0xca, 0x2f, 0x88, 0x10, 0x07, 0x3e, 0x9a, 0x22, 0xef, 0x52, 0x7a, 0x0d, 0xdb, 0x81,
	0x46, 0xda, 0xb5, 0x57, 0x8e, 0xdb, 0x52, 0xfe, 0x11, 0xff, 0x8e, 0x33, 0x27, 0x3c, 0xe3, 0x44,
	0x14, 0xa9, 0x39, 0xcd, 0x7b, 0x7e, 0x9e, 0x79, 0x6f, 0x1c, 0x80, 0x9f, 0xd6, 0x60, 0xbd, 0x77,
	0xd6, 0x5b, 0x25, 0xfc, 0xdd, 0x1e, 0x87, 0xe5, 0xef, 0x0c, 0x72, 0x8d, 0x1b, 0xeb, 0x2e, 0x95,
	0x82, 0xcc, 0x74, 0x3b, 0xac, 0x92, 0xa3, 0xe4, 0xb8, 0xd4, 0x5c, 0x2b, 0x09, 0xa9, 0xf7, 0xdb,
	0x6a, 0x16, 0x28, 0xa1, 0xa9, 0x24, 0x15, 0xdd, 0xac, 0xd2, 0xa8, 0xa2, 0x5a, 0x3d, 0x01, 0x31,
	0xf8, 0xce, 0x63, 0x95, 0xb1, 0x2e, 0x02, 0xf5, 0x0a, 0xc4, 0x4d, 0xb7, 0xbd, 0xc6, 0x4a, 0x1c,
	0xa5, 0xc7, 0xf3, 0x93, 0xc7, 0x35, 0x4f, 0xac, 0xe3, 0xb4, 0xfa, 0x9c, 0x8e, 0x74, 0x54, 0x1c,
	0xfe, 0x99, 0x81, 0x60, 0x42, 0x3d, 0x85, 0xdc, 0xb1, 0x20, 0xd8, 0x48, 0xc3, 0x80, 0x11, 0xd1,
	0xd8, 0x9b, 0x1e, 0x6f, 0xd9, 0x49, 0x18, 0x4b, 0x35, 0x69, 0x6f, 0xb1, 0xff, 0x7e, 0xe5, 0xd9,
	0x8c, 0xd0, 0x23, 0x52, 0xcf, 0xa1, 0xdc, 0x58, 0xe3, 0x7b, 0x83, 0xc6, 0xb3, 0xa5, 0x52, 0xff,
	0x23, 0x54, 0x05, 0xc5, 0xc6, 0x5e, 0x1b, 0xef, 0xee, 0x82, 0x31, 0x3a, 0x9b, 0x60, 0x30, 0x9c,
	0x0e, 0xb6, 0xab, 0xf2, 0xc0, 0xce, 0x4f, 0x9e, 0x3d, 0x60, 0xb7, 0x5e, 0xd9, 0x4e, 0x93, 0x46,
	0xbd, 0x00, 0xd8, 0x3b, 0xfc, 0x86, 0x0e, 0xcd, 0x06, 0xab, 0x82, 0xc7, 0xdf, 0x63, 0x0e, 0x7f,
	0x25, 0x90, 0x06, 0x31, 0x6d, 0x66, 0x77, 0x6f, 0xa9, 0x11, 0x10, 0x6b, 0x98, 0x8d, 0x69, 0x22,
	0xa0, 0x38, 0x03, 0xba, 0xbe, 0xdb, 0x72, 0x9c, 0x85, 0x1e, 0x11, 0x19, 0x0e, 0x8d, 0x1d, 0x0e,
	0x57, 0x1c, 0x66, 0xa1, 0x27, 0x48, 0x7d, 0x1c, 0x4e, 0x41, 0x16, 0x3a, 0x02, 0xea, 0x83, 0x3f,
	0xf6, 0xbd, 0x43, 0x4e, 0x12, 0xfa, 0x44, 0x44, 0xfc, 0xae, 0x37, 0xf4, 0x9c, 0x45, 0xe4, 0x23,
	0x5a, 0x1a, 0xc8, 0xd6, 0xf4, 0x8a, 0x73, 0x28, 0x3e, 0xb7, 0x1f, 0xda, 0xb3, 0x2f, 0xad, 0x7c,
	0xa4, 0x8a, 0xe0, 0xff, 0xac, 0x91, 0x89, 0xca, 0x61, 0xd6, 0xae, 0xe4, 0x4c, 0x09, 0x48, 0x1a,
	0x99, 0xaa, 0x12, 0xc4, 0x9b, 0xb6, 0x39, 0x7d, 0x2b, 0x33, 0x3a, 0x39, 0xbd, 0x90, 0x82, 0xa4,
	0xeb, 0x8b, 0xb5, 0xcc, 0xd5, 0x01, 0x64, 0x4d, 0xf8, 0x64, 0xc1, 0xb7, 0xf5, 0xb9, 0x3c, 0xa0,
	0xe2, 0xd3, 0x5a, 0xcb, 0x92, 0xee, 0x35, 0x1f, 0xdf, 0x35, 0x2b, 0x09, 0xcb, 0xf7, 0x50, 0xc4,
	0xa5, 0x0e, 0x64, 0xe9, 0xd2, 0xee, 0xba, 0xde, 0x8c, 0xfb, 0x19, 0x91, 0x7a, 0x49, 0x91, 0x59,
	0x12, 0x56, 0x44, 0x3f, 0xcf, 0xe2, 0xbf, 0xd7, 0xd0, 0xd3, 0xe9, 0xd7, 0x9c, 0x7f, 0xe6, 0xd7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x94, 0xe5, 0xd1, 0xfa, 0xda, 0x02, 0x00, 0x00,
}
