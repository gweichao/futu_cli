// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Qot_UpdateBroker.proto

package Qot_UpdateBroker

import (
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "pbgo/Common"
	Qot_Common "pbgo/Qot_Common"
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

type S2C struct {
	Security      *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	BrokerAskList []*Qot_Common.Broker `protobuf:"bytes,2,rep,name=brokerAskList" json:"brokerAskList,omitempty"`
	BrokerBidList []*Qot_Common.Broker `protobuf:"bytes,3,rep,name=brokerBidList" json:"brokerBidList,omitempty"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b58d393d5c0052b, []int{0}
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

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetBrokerAskList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerAskList
	}
	return nil
}

func (m *S2C) GetBrokerBidList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerBidList
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
	return fileDescriptor_1b58d393d5c0052b, []int{1}
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
	proto.RegisterType((*S2C)(nil), "Qot_UpdateBroker.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateBroker.Response")
}

func init() { proto.RegisterFile("Qot_UpdateBroker.proto", fileDescriptor_1b58d393d5c0052b) }

var fileDescriptor_1b58d393d5c0052b = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4e, 0xf3, 0x30,
	0x1c, 0x47, 0xe3, 0xa6, 0xfd, 0xbe, 0xe2, 0x82, 0x84, 0x0c, 0x45, 0x51, 0x85, 0x4c, 0xd4, 0x85,
	0x2c, 0x98, 0xca, 0x62, 0x40, 0x6c, 0x24, 0x8c, 0x30, 0xe0, 0xc2, 0xc2, 0x82, 0xda, 0xc4, 0x54,
	0x51, 0x95, 0xd8, 0xb2, 0x9d, 0xa1, 0xb7, 0x80, 0x73, 0x70, 0x91, 0x8e, 0x1d, 0x99, 0x10, 0x4a,
	0x2e, 0x82, 0x9a, 0xa6, 0x10, 0xca, 0xc0, 0xfa, 0x7e, 0xef, 0x25, 0x7f, 0x19, 0x1e, 0xdc, 0x0a,
	0xf3, 0x78, 0x2f, 0xa3, 0x91, 0xe1, 0xbe, 0x12, 0x53, 0xae, 0x88, 0x54, 0xc2, 0x08, 0xb4, 0xbb,
	0xc9, 0x7b, 0xdb, 0x81, 0x48, 0x12, 0x91, 0xae, 0xf6, 0x5e, 0xb9, 0xd7, 0x49, 0xff, 0x15, 0x40,
	0x7b, 0x48, 0x03, 0x34, 0x80, 0x6d, 0xcd, 0xc3, 0x4c, 0xc5, 0x66, 0xe6, 0x00, 0xb7, 0xe1, 0x75,
	0xe8, 0x3e, 0xa9, 0xc9, 0xc3, 0x6a, 0x63, 0x5f, 0x16, 0x3a, 0x87, 0x3b, 0xe3, 0xf2, 0x1f, 0x97,
	0x7a, 0x7a, 0x1d, 0x6b, 0xe3, 0x34, 0x5c, 0xdb, 0xeb, 0x50, 0x54, 0xcf, 0x56, 0x47, 0xb0, 0x9f,
	0xe2, 0x77, 0xe9, 0xc7, 0x51, 0x59, 0xda, 0x7f, 0x95, 0x95, 0xd8, 0x7f, 0x01, 0xb0, 0xcd, 0xb8,
	0x96, 0x22, 0xd5, 0x1c, 0x61, 0xf8, 0x5f, 0x71, 0x73, 0x37, 0x93, 0xbc, 0xbc, 0xb8, 0x75, 0xd1,
	0x3c, 0x39, 0x1b, 0x0c, 0xd8, 0x1a, 0xa2, 0x43, 0xf8, 0x4f, 0x71, 0x73, 0xa3, 0x27, 0x4e, 0xc3,
	0x05, 0xde, 0x96, 0xdf, 0x9c, 0xbf, 0x1f, 0x59, 0xac, 0x62, 0xcb, 0x9a, 0x2b, 0x15, 0x88, 0x88,
	0x3b, 0xb6, 0x0b, 0xbc, 0x56, 0x35, 0xaf, 0x21, 0x3a, 0x86, 0xb6, 0xa6, 0xa1, 0xd3, 0x74, 0x81,
	0xd7, 0xa1, 0x5d, 0xf2, 0xeb, 0xc1, 0x87, 0x34, 0x60, 0x4b, 0xc3, 0xbf, 0x9a, 0xe7, 0x18, 0x2c,
	0x72, 0x0c, 0x3e, 0x72, 0x0c, 0x9e, 0x0b, 0x6c, 0x2d, 0x0a, 0x6c, 0xbd, 0x15, 0xd8, 0x82, 0x7b,
	0xa1, 0x48, 0xc8, 0x53, 0x66, 0x32, 0x22, 0x24, 0x4f, 0x47, 0x32, 0x26, 0x72, 0xfc, 0xd0, 0x95,
	0xe3, 0x89, 0x38, 0xdd, 0xfc, 0xdc, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x09, 0x3c, 0x16,
	0xd2, 0x01, 0x00, 0x00,
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
	if len(m.BrokerBidList) > 0 {
		for iNdEx := len(m.BrokerBidList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BrokerBidList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BrokerAskList) > 0 {
		for iNdEx := len(m.BrokerAskList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BrokerAskList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Security == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("security")
	} else {
		{
			size, err := m.Security.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(size))
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
			i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(m.ErrCode))
	i--
	dAtA[i] = 0x18
	i -= len(m.RetMsg)
	copy(dAtA[i:], m.RetMsg)
	i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(len(m.RetMsg)))
	i--
	dAtA[i] = 0x12
	if m.RetType == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("retType")
	} else {
		i = encodeVarintQot_UpdateBroker(dAtA, i, uint64(*m.RetType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQot_UpdateBroker(dAtA []byte, offset int, v uint64) int {
	offset -= sovQot_UpdateBroker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *S2C) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Security != nil {
		l = m.Security.Size()
		n += 1 + l + sovQot_UpdateBroker(uint64(l))
	}
	if len(m.BrokerAskList) > 0 {
		for _, e := range m.BrokerAskList {
			l = e.Size()
			n += 1 + l + sovQot_UpdateBroker(uint64(l))
		}
	}
	if len(m.BrokerBidList) > 0 {
		for _, e := range m.BrokerBidList {
			l = e.Size()
			n += 1 + l + sovQot_UpdateBroker(uint64(l))
		}
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
		n += 1 + sovQot_UpdateBroker(uint64(*m.RetType))
	}
	l = len(m.RetMsg)
	n += 1 + l + sovQot_UpdateBroker(uint64(l))
	n += 1 + sovQot_UpdateBroker(uint64(m.ErrCode))
	if m.S2C != nil {
		l = m.S2C.Size()
		n += 1 + l + sovQot_UpdateBroker(uint64(l))
	}
	return n
}

func sovQot_UpdateBroker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQot_UpdateBroker(x uint64) (n int) {
	return sovQot_UpdateBroker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return ErrIntOverflowQot_UpdateBroker
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Security", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateBroker
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
				return ErrInvalidLengthQot_UpdateBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Security == nil {
				m.Security = &Qot_Common.Security{}
			}
			if err := m.Security.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BrokerAskList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateBroker
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
				return ErrInvalidLengthQot_UpdateBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BrokerAskList = append(m.BrokerAskList, &Qot_Common.Broker{})
			if err := m.BrokerAskList[len(m.BrokerAskList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BrokerBidList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQot_UpdateBroker
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
				return ErrInvalidLengthQot_UpdateBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BrokerBidList = append(m.BrokerBidList, &Qot_Common.Broker{})
			if err := m.BrokerBidList[len(m.BrokerBidList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQot_UpdateBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("security")
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
				return ErrIntOverflowQot_UpdateBroker
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
					return ErrIntOverflowQot_UpdateBroker
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
					return ErrIntOverflowQot_UpdateBroker
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
				return ErrInvalidLengthQot_UpdateBroker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateBroker
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
					return ErrIntOverflowQot_UpdateBroker
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
					return ErrIntOverflowQot_UpdateBroker
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
				return ErrInvalidLengthQot_UpdateBroker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQot_UpdateBroker
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
			skippy, err := skipQot_UpdateBroker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQot_UpdateBroker
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQot_UpdateBroker
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
func skipQot_UpdateBroker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQot_UpdateBroker
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
					return 0, ErrIntOverflowQot_UpdateBroker
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
					return 0, ErrIntOverflowQot_UpdateBroker
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
				return 0, ErrInvalidLengthQot_UpdateBroker
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQot_UpdateBroker
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQot_UpdateBroker
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
				next, err := skipQot_UpdateBroker(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQot_UpdateBroker
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
	ErrInvalidLengthQot_UpdateBroker = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQot_UpdateBroker   = fmt.Errorf("proto: integer overflow")
)
