// Code generated by protoc-gen-go.
// source: gateway.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReqMsgServer struct {
	Cmd    uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
}

func (m *ReqMsgServer) Reset()                    { *m = ReqMsgServer{} }
func (m *ReqMsgServer) String() string            { return proto.CompactTextString(m) }
func (*ReqMsgServer) ProtoMessage()               {}
func (*ReqMsgServer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type SelectMsgServerForClient struct {
	Cmd    uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	Addr   string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
}

func (m *SelectMsgServerForClient) Reset()                    { *m = SelectMsgServerForClient{} }
func (m *SelectMsgServerForClient) String() string            { return proto.CompactTextString(m) }
func (*SelectMsgServerForClient) ProtoMessage()               {}
func (*SelectMsgServerForClient) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func init() {
	proto.RegisterType((*ReqMsgServer)(nil), "protocol.ReqMsgServer")
	proto.RegisterType((*SelectMsgServerForClient)(nil), "protocol.SelectMsgServerForClient")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x4a, 0x16, 0x5c, 0x3c, 0x41, 0xa9, 0x85, 0xbe, 0xc5, 0xe9, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45,
	0x42, 0x02, 0x5c, 0xcc, 0xc9, 0xb9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x20, 0xa6,
	0x90, 0x18, 0x17, 0x5b, 0x72, 0x6e, 0x4a, 0x70, 0x49, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x94, 0xa7, 0x14, 0xc1, 0x25, 0x11, 0x9c, 0x9a, 0x93, 0x9a, 0x5c, 0x02, 0xd7, 0xec, 0x96,
	0x5f, 0xe4, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0x42, 0xbc, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x89, 0x29,
	0x29, 0x45, 0x12, 0xcc, 0x60, 0x51, 0x30, 0x3b, 0x89, 0x0d, 0xec, 0x3a, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x41, 0x57, 0x43, 0x29, 0xb5, 0x00, 0x00, 0x00,
}