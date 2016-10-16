// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/authenticator.proto
// DO NOT EDIT!

/*
Package internet is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/authenticator.proto
	v2ray.com/core/transport/internet/config.proto

It has these top-level messages:
	AuthenticatorConfig
	NetworkSettings
	StreamConfig
*/
package internet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_loader "v2ray.com/core/common/loader"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticatorConfig struct {
	Name     string                                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Settings *v2ray_core_common_loader.TypedSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
}

func (m *AuthenticatorConfig) Reset()                    { *m = AuthenticatorConfig{} }
func (m *AuthenticatorConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthenticatorConfig) ProtoMessage()               {}
func (*AuthenticatorConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthenticatorConfig) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticatorConfig)(nil), "v2ray.core.transport.internet.AuthenticatorConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/authenticator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x88, 0xac, 0xf1, 0x16, 0x2f, 0x8b, 0x20, 0xac, 0x5e, 0x76, 0x4f, 0x13, 0x58,
	0xf1, 0x01, 0xda, 0xbe, 0x80, 0x54, 0x4f, 0xde, 0x62, 0x3a, 0xd6, 0x80, 0x99, 0x09, 0xd3, 0x51,
	0xe8, 0xdb, 0x8b, 0xad, 0x2d, 0xd5, 0x83, 0xb7, 0x40, 0xe6, 0xfb, 0xfe, 0xcf, 0xdc, 0x7f, 0x1e,
	0xc5, 0x0f, 0x10, 0x38, 0xb9, 0xc0, 0x82, 0x4e, 0xc5, 0x53, 0x9f, 0x59, 0xd4, 0x45, 0x52, 0x14,
	0x42, 0x75, 0xfe, 0x43, 0xdf, 0x90, 0x34, 0x06, 0xaf, 0x2c, 0x90, 0x85, 0x95, 0xed, 0xf5, 0x8c,
	0x09, 0xc2, 0x82, 0xc0, 0x8c, 0x5c, 0xed, 0xff, 0x58, 0x03, 0xa7, 0xc4, 0xe4, 0xde, 0xd9, 0xb7,
	0x28, 0x4e, 0x87, 0x8c, 0x93, 0xe7, 0x96, 0xcc, 0x65, 0xb9, 0xd6, 0xd7, 0x4c, 0xaf, 0xb1, 0xb3,
	0xd6, 0x9c, 0x92, 0x4f, 0xb8, 0x2d, 0x76, 0xc5, 0xe1, 0xbc, 0x19, 0xdf, 0xb6, 0x36, 0x9b, 0x1e,
	0x55, 0x23, 0x75, 0xfd, 0xf6, 0x64, 0x57, 0x1c, 0x2e, 0x8e, 0x7b, 0x58, 0x55, 0x4c, 0x13, 0x30,
	0x4d, 0xc0, 0xd3, 0x90, 0xb1, 0x7d, 0xfc, 0x39, 0x6f, 0x16, 0xb0, 0x2a, 0xcd, 0x4d, 0xe0, 0x04,
	0xff, 0xd6, 0x57, 0xf6, 0x57, 0xd2, 0xc3, 0x77, 0xe8, 0xf3, 0x66, 0xfe, 0x7d, 0x39, 0x1b, 0xcb,
	0xef, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86, 0xf6, 0xf6, 0xdf, 0x3a, 0x01, 0x00, 0x00,
}