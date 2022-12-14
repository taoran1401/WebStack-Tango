// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd.proto

package protopb

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

type CmdBase int32

const (
	//错误
	CmdBase_CmdBaseErr CmdBase = 0
	//网络连接
	CmdBase_CmdBaseConn CmdBase = 1
	//版本更新
	CmdBase_CmdBaseVersion CmdBase = 2
	//online list
	CmdBase_CmdBaseOnlineList CmdBase = 3
	//apply vs
	CmdBase_CmdBaseApplyVs CmdBase = 4
	//response vs
	CmdBase_CmdBaseVsResp CmdBase = 5
	//move chess
	CmdBase_CmdBaseChess CmdBase = 6
	//surrender
	CmdBase_CmdBaseSurrender CmdBase = 7
)

var CmdBase_name = map[int32]string{
	0: "CmdBaseErr",
	1: "CmdBaseConn",
	2: "CmdBaseVersion",
	3: "CmdBaseOnlineList",
	4: "CmdBaseApplyVs",
	5: "CmdBaseVsResp",
	6: "CmdBaseChess",
	7: "CmdBaseSurrender",
}

var CmdBase_value = map[string]int32{
	"CmdBaseErr":        0,
	"CmdBaseConn":       1,
	"CmdBaseVersion":    2,
	"CmdBaseOnlineList": 3,
	"CmdBaseApplyVs":    4,
	"CmdBaseVsResp":     5,
	"CmdBaseChess":      6,
	"CmdBaseSurrender":  7,
}

func (x CmdBase) String() string {
	return proto.EnumName(CmdBase_name, int32(x))
}

func (CmdBase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{0}
}

func init() {
	proto.RegisterEnum("protopb.CmdBase", CmdBase_name, CmdBase_value)
}

func init() { proto.RegisterFile("cmd.proto", fileDescriptor_7520252fb01eaf30) }

var fileDescriptor_7520252fb01eaf30 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xce, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x07, 0x53, 0x05, 0x49, 0x5a, 0x4b, 0x18, 0xb9, 0xd8,
	0x9d, 0x73, 0x53, 0x9c, 0x12, 0x8b, 0x53, 0x85, 0xf8, 0xb8, 0xb8, 0xa0, 0x4c, 0xd7, 0xa2, 0x22,
	0x01, 0x06, 0x21, 0x7e, 0x2e, 0x6e, 0x28, 0xdf, 0x39, 0x3f, 0x2f, 0x4f, 0x80, 0x51, 0x48, 0x88,
	0x8b, 0x0f, 0x2a, 0x10, 0x96, 0x5a, 0x54, 0x9c, 0x99, 0x9f, 0x27, 0xc0, 0x24, 0x24, 0xca, 0x25,
	0x08, 0x15, 0xf3, 0xcf, 0xcb, 0xc9, 0xcc, 0x4b, 0xf5, 0xc9, 0x2c, 0x2e, 0x11, 0x60, 0x46, 0x52,
	0xea, 0x58, 0x50, 0x90, 0x53, 0x19, 0x56, 0x2c, 0xc0, 0x22, 0x24, 0xc8, 0xc5, 0x0b, 0xd3, 0x5e,
	0x1c, 0x94, 0x5a, 0x5c, 0x20, 0xc0, 0x2a, 0x24, 0xc0, 0xc5, 0x03, 0xb3, 0x22, 0x23, 0xb5, 0xb8,
	0x58, 0x80, 0x4d, 0x48, 0x84, 0x4b, 0x00, 0x2a, 0x12, 0x5c, 0x5a, 0x54, 0x94, 0x9a, 0x97, 0x92,
	0x5a, 0x24, 0xc0, 0xee, 0x24, 0x12, 0x25, 0xa4, 0xa7, 0xa7, 0x0f, 0x75, 0xb4, 0x35, 0x94, 0x4e,
	0x62, 0x03, 0x33, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xb1, 0xa8, 0x6a, 0xd9, 0x00,
	0x00, 0x00,
}
