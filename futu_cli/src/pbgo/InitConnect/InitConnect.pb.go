// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: InitConnect.proto

package InitConnect

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "pbgo/Common"
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

type C2S struct {
	ClientVer  int32  `protobuf:"varint,1,req,name=clientVer" json:"clientVer"`
	ClientID   string `protobuf:"bytes,2,req,name=clientID" json:"clientID"`
	RecvNotify bool   `protobuf:"varint,3,opt,name=recvNotify" json:"recvNotify"`
	//如果通信要加密，首先得在FutuOpenD和客户端都配置RSA密钥，不配置始终不加密
	//如果配置了RSA密钥且指定的加密算法不为PacketEncAlgo_None则加密(即便这里不设置，配置了RSA密钥，也会采用默认加密方式)，默认采用FTAES_ECB算法
	PacketEncAlgo int32 `protobuf:"varint,4,opt,name=packetEncAlgo" json:"packetEncAlgo"`
	PushProtoFmt  int32 `protobuf:"varint,5,opt,name=pushProtoFmt" json:"pushProtoFmt"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e68d7dcc5c46c8, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return m.Size()
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetClientVer() int32 {
	if m != nil {
		return m.ClientVer
	}
	return 0
}

func (m *C2S) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *C2S) GetRecvNotify() bool {
	if m != nil {
		return m.RecvNotify
	}
	return false
}

func (m *C2S) GetPacketEncAlgo() int32 {
	if m != nil {
		return m.PacketEncAlgo
	}
	return 0
}

func (m *C2S) GetPushProtoFmt() int32 {
	if m != nil {
		return m.PushProtoFmt
	}
	return 0
}

type S2C struct {
	ServerVer         int32  `protobuf:"varint,1,req,name=serverVer" json:"serverVer"`
	LoginUserID       uint64 `protobuf:"varint,2,req,name=loginUserID" json:"loginUserID"`
	ConnID            uint64 `protobuf:"varint,3,req,name=connID" json:"connID"`
	ConnAESKey        string `protobuf:"bytes,4,req,name=connAESKey" json:"connAESKey"`
	KeepAliveInterval int32  `protobuf:"varint,5,req,name=keepAliveInterval" json:"keepAliveInterval"`
	AesCBCiv          string `protobuf:"bytes,6,opt,name=aesCBCiv" json:"aesCBCiv"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e68d7dcc5c46c8, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return m.Size()
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetServerVer() int32 {
	if m != nil {
		return m.ServerVer
	}
	return 0
}

func (m *S2C) GetLoginUserID() uint64 {
	if m != nil {
		return m.LoginUserID
	}
	return 0
}

func (m *S2C) GetConnID() uint64 {
	if m != nil {
		return m.ConnID
	}
	return 0
}

func (m *S2C) GetConnAESKey() string {
	if m != nil {
		return m.ConnAESKey
	}
	return ""
}

func (m *S2C) GetKeepAliveInterval() int32 {
	if m != nil {
		return m.KeepAliveInterval
	}
	return 0
}

func (m *S2C) GetAesCBCiv() string {
	if m != nil {
		return m.AesCBCiv
	}
	return ""
}

type Request struct {
	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e68d7dcc5c46c8, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType *int32 `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg"`
	ErrCode int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode"`
	S2C     *S2C   `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_70e68d7dcc5c46c8, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "InitConnect.C2S")
	proto.RegisterType((*S2C)(nil), "InitConnect.S2C")
	proto.RegisterType((*Request)(nil), "InitConnect.Request")
	proto.RegisterType((*Response)(nil), "InitConnect.Response")
}

func init() { proto.RegisterFile("InitConnect.proto", fileDescriptor_70e68d7dcc5c46c8) }

var fileDescriptor_70e68d7dcc5c46c8 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xeb, 0x26, 0xdd, 0x3a, 0x77, 0x48, 0x9b, 0xb9, 0x58, 0x08, 0x85, 0x28, 0x42, 0xa8,
	0x42, 0x5a, 0x98, 0x22, 0x4e, 0x5c, 0x50, 0xeb, 0x0d, 0xa9, 0x42, 0x20, 0x94, 0x00, 0x07, 0x6e,
	0x9d, 0x79, 0x2b, 0xd1, 0x52, 0xdb, 0xd8, 0x6e, 0xa4, 0xfe, 0x0f, 0x1c, 0xf8, 0xaf, 0xd8, 0x71,
	0x47, 0x4e, 0x08, 0xb5, 0x47, 0xfe, 0x09, 0xe4, 0x36, 0xed, 0xbc, 0xf5, 0x96, 0xfc, 0xbe, 0x2f,
	0x79, 0xef, 0xfb, 0x1e, 0x3e, 0x1e, 0x89, 0xd2, 0x32, 0x29, 0x04, 0x70, 0x9b, 0x2a, 0x2d, 0xad,
	0x24, 0x3d, 0x0f, 0x3d, 0x3a, 0x64, 0x72, 0x3a, 0x95, 0x62, 0x2d, 0x25, 0xbf, 0x10, 0x0e, 0x58,
	0x56, 0x90, 0x04, 0x1f, 0xf0, 0xaa, 0x04, 0x61, 0x3f, 0x83, 0xa6, 0x28, 0x6e, 0xf7, 0x3b, 0xc3,
	0xf0, 0xfa, 0xcf, 0x93, 0x56, 0x7e, 0x8b, 0x49, 0x8c, 0xbb, 0xeb, 0x97, 0xd1, 0x19, 0x6d, 0xc7,
	0xed, 0xfe, 0x41, 0x63, 0xd9, 0x52, 0xf2, 0x14, 0x63, 0x0d, 0xbc, 0x7e, 0x2f, 0x6d, 0x79, 0x39,
	0xa7, 0x41, 0x8c, 0xfa, 0xdd, 0xc6, 0xe3, 0x71, 0xf2, 0x1c, 0x3f, 0x50, 0x63, 0x7e, 0x05, 0xf6,
	0x5c, 0xf0, 0x41, 0x35, 0x91, 0x34, 0x8c, 0xd1, 0x76, 0xde, 0x5d, 0x89, 0xf4, 0xf1, 0xa1, 0x9a,
	0x99, 0x6f, 0x1f, 0xdc, 0xb2, 0x6f, 0xa6, 0x96, 0x76, 0x3c, 0xeb, 0x1d, 0x25, 0xf9, 0x87, 0x70,
	0x50, 0x64, 0xcc, 0x25, 0x31, 0xa0, 0x6b, 0xd0, 0x3b, 0x49, 0xb6, 0x98, 0x3c, 0xc3, 0xbd, 0x4a,
	0x4e, 0x4a, 0xf1, 0xc9, 0x80, 0x6e, 0xc2, 0x84, 0x8d, 0xcb, 0x17, 0xc8, 0x63, 0xbc, 0xc7, 0xa5,
	0x10, 0xa3, 0x33, 0x1a, 0x78, 0x96, 0x86, 0xb9, 0xb4, 0xee, 0x69, 0x70, 0x5e, 0xbc, 0x85, 0x39,
	0x0d, 0xbd, 0x46, 0x3c, 0x4e, 0x32, 0x7c, 0x7c, 0x05, 0xa0, 0x06, 0x55, 0x59, 0xc3, 0x48, 0x58,
	0xd0, 0xf5, 0xb8, 0xa2, 0x1d, 0x6f, 0xaf, 0x5d, 0xd9, 0x35, 0x3d, 0x06, 0xc3, 0x86, 0xac, 0xac,
	0xe9, 0x5e, 0x8c, 0x6e, 0x9b, 0xde, 0xd0, 0xe4, 0x04, 0xef, 0xe7, 0xf0, 0x7d, 0x06, 0xc6, 0x92,
	0x04, 0x07, 0x3c, 0x33, 0xab, 0xa8, 0xbd, 0xec, 0x28, 0xf5, 0xcf, 0xcf, 0xb2, 0x22, 0x77, 0x62,
	0xf2, 0x03, 0xe1, 0x6e, 0x0e, 0x46, 0x49, 0x61, 0x80, 0x44, 0x78, 0x5f, 0x83, 0xfd, 0x38, 0x57,
	0xb0, 0xee, 0xe7, 0x55, 0x78, 0xf2, 0xf2, 0xf4, 0x34, 0xdf, 0x40, 0x97, 0x5a, 0x83, 0x7d, 0x67,
	0x26, 0xb4, 0xed, 0xcd, 0x6e, 0x98, 0xfb, 0x1a, 0xb4, 0x66, 0xf2, 0x2b, 0xac, 0x0e, 0xbc, 0x49,
	0xb1, 0x81, 0x6e, 0x1d, 0x93, 0xf1, 0xd5, 0x4d, 0xef, 0xaf, 0x53, 0x64, 0x2c, 0x77, 0xe2, 0xf0,
	0xf5, 0xf5, 0x22, 0x42, 0x37, 0x8b, 0x08, 0xfd, 0x5d, 0x44, 0xe8, 0xe7, 0x32, 0x6a, 0xdd, 0x2c,
	0xa3, 0xd6, 0xef, 0x65, 0xd4, 0xc2, 0x0f, 0xb9, 0x9c, 0xa6, 0x97, 0x33, 0x3b, 0x4b, 0xa5, 0x02,
	0x31, 0x56, 0x65, 0xaa, 0x2e, 0xbe, 0x1c, 0xa9, 0x8b, 0x89, 0x7c, 0xe1, 0xfd, 0xe9, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0b, 0x8d, 0x32, 0x55, 0xe5, 0x02, 0x00, 0x00,
}

func (m *C2S) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *C2S) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *C2S) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i = encodeVarintInitConnect(dAtA, i, uint64(m.PushProtoFmt))
	i--
	dAtA[i] = 0x28
	i = encodeVarintInitConnect(dAtA, i, uint64(m.PacketEncAlgo))
	i--
	dAtA[i] = 0x20
	i--
	if m.RecvNotify {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i -= len(m.ClientID)
	copy(dAtA[i:], m.ClientID)
	i = encodeVarintInitConnect(dAtA, i, uint64(len(m.ClientID)))
	i--
	dAtA[i] = 0x12
	i = encodeVarintInitConnect(dAtA, i, uint64(m.ClientVer))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *S2C) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S2C) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *S2C) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.AesCBCiv)
	copy(dAtA[i:], m.AesCBCiv)
	i = encodeVarintInitConnect(dAtA, i, uint64(len(m.AesCBCiv)))
	i--
	dAtA[i] = 0x32
	i = encodeVarintInitConnect(dAtA, i, uint64(m.KeepAliveInterval))
	i--
	dAtA[i] = 0x28
	i -= len(m.ConnAESKey)
	copy(dAtA[i:], m.ConnAESKey)
	i = encodeVarintInitConnect(dAtA, i, uint64(len(m.ConnAESKey)))
	i--
	dAtA[i] = 0x22
	i = encodeVarintInitConnect(dAtA, i, uint64(m.ConnID))
	i--
	dAtA[i] = 0x18
	i = encodeVarintInitConnect(dAtA, i, uint64(m.LoginUserID))
	i--
	dAtA[i] = 0x10
	i = encodeVarintInitConnect(dAtA, i, uint64(m.ServerVer))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.C2S == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("c2s")
	} else {
		{
			size, err := m.C2S.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInitConnect(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.S2C != nil {
		{
			size, err := m.S2C.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInitConnect(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintInitConnect(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintInitConnect(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintInitConnect(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintInitConnect(dAtA []byte, offset int, v uint64) int {
	offset -= sovInitConnect(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *C2S) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInitConnect(uint64(m.ClientVer))
	l = len(m.ClientID)
	n += 1 + l + sovInitConnect(uint64(l))
	n += 2
	n += 1 + sovInitConnect(uint64(m.PacketEncAlgo))
	n += 1 + sovInitConnect(uint64(m.PushProtoFmt))
	return n
}

func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovInitConnect(uint64(m.ServerVer))
	n += 1 + sovInitConnect(uint64(m.LoginUserID))
	n += 1 + sovInitConnect(uint64(m.ConnID))
	l = len(m.ConnAESKey)
	n += 1 + l + sovInitConnect(uint64(l))
	n += 1 + sovInitConnect(uint64(m.KeepAliveInterval))
	l = len(m.AesCBCiv)
	n += 1 + l + sovInitConnect(uint64(l))
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.C2S != nil {
		l = m.C2S.Size()
		n += 1 + l + sovInitConnect(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RetType != nil {
		n += 1 + sovInitConnect(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovInitConnect(uint64(l))
	n += 1 + sovInitConnect(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovInitConnect(uint64(l))
	}
	return n
}

func sovInitConnect(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInitConnect(x uint64) (n int) {
	return sovInitConnect(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *C2S) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInitConnect
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: C2S: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: C2S: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientVer", wireType)
			}
			m.ClientVer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientVer |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RecvNotify = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketEncAlgo", wireType)
			}
			m.PacketEncAlgo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketEncAlgo |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PushProtoFmt", wireType)
			}
			m.PushProtoFmt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PushProtoFmt |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInitConnect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("clientVer")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("clientID")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *S2C) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInitConnect
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S2C: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S2C: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerVer", wireType)
			}
			m.ServerVer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerVer |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginUserID", wireType)
			}
			m.LoginUserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoginUserID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnID", wireType)
			}
			m.ConnID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConnID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnAESKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnAESKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeepAliveInterval", wireType)
			}
			m.KeepAliveInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeepAliveInterval |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AesCBCiv", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AesCBCiv = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInitConnect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("serverVer")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("loginUserID")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("connID")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("connAESKey")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("keepAliveInterval")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Request) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInitConnect
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C2S", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.C2S == nil {
				m.C2S = &C2S{}
			}
			if err := m.C2S.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipInitConnect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("c2s")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInitConnect
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetType", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RetType = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RetMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S2C", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInitConnect
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInitConnect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.S2C == nil {
				m.S2C = &S2C{}
			}
			if err := m.S2C.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInitConnect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthInitConnect
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInitConnect(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInitConnect
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInitConnect
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthInitConnect
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthInitConnect
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInitConnect
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipInitConnect(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthInitConnect
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthInitConnect = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInitConnect   = fmt.Errorf("proto: integer overflow")
)
