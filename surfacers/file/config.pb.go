// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/file/config.proto

/*
Package file is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/surfacers/file/config.proto

It has these top-level messages:
	SurfacerConf
*/
package file

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

type SurfacerConf struct {
	// Where to write the results. If left unset, file surfacer writes to the
	// standard output.
	FilePath         *string `protobuf:"bytes,1,opt,name=file_path,json=filePath" json:"file_path,omitempty"`
	Prefix           *string `protobuf:"bytes,2,opt,name=prefix,def=cloudprober" json:"prefix,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SurfacerConf) Reset()                    { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string            { return proto.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()               {}
func (*SurfacerConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_SurfacerConf_Prefix string = "cloudprober"

func (m *SurfacerConf) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *SurfacerConf) GetPrefix() string {
	if m != nil && m.Prefix != nil {
		return *m.Prefix
	}
	return Default_SurfacerConf_Prefix
}

func init() {
	proto.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.file.SurfacerConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/file/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x2f, 0x2e, 0x2d, 0x4a, 0x4b, 0x4c,
	0x4e, 0x2d, 0x2a, 0xd6, 0x4f, 0xcb, 0x04, 0x49, 0xe5, 0xe7, 0xa5, 0x65, 0xa6, 0xeb, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x22, 0x29, 0xd4, 0x83, 0x29, 0xd4, 0x03, 0xa9, 0x53, 0x0a, 0xe0,
	0xe2, 0x09, 0x86, 0x0a, 0x38, 0xe7, 0xe7, 0xa5, 0x09, 0x49, 0x73, 0x71, 0x82, 0xc4, 0xe3, 0x0b,
	0x12, 0x4b, 0x32, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0x40, 0x02, 0x01, 0x89, 0x25,
	0x19, 0x42, 0xca, 0x5c, 0x6c, 0x05, 0x45, 0xa9, 0x69, 0x99, 0x15, 0x12, 0x4c, 0x20, 0x19, 0x2b,
	0x6e, 0x24, 0x73, 0x83, 0xa0, 0x52, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x0b, 0x91, 0x94,
	0xa8, 0x00, 0x00, 0x00,
}
