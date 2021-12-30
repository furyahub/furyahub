// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rate/rate.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ExchangeRate struct {
	Denom string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Value github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"value"`
}

func (m *ExchangeRate) Reset()         { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()    {}
func (*ExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_466953fd5f8ee6df, []int{0}
}
func (m *ExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeRate.Merge(m, src)
}
func (m *ExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *ExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeRate proto.InternalMessageInfo

func (m *ExchangeRate) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type EraExchangeRate struct {
	Denom string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Era   uint32                                 `protobuf:"varint,2,opt,name=era,proto3" json:"era,omitempty"`
	Value github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"value"`
}

func (m *EraExchangeRate) Reset()         { *m = EraExchangeRate{} }
func (m *EraExchangeRate) String() string { return proto.CompactTextString(m) }
func (*EraExchangeRate) ProtoMessage()    {}
func (*EraExchangeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_466953fd5f8ee6df, []int{1}
}
func (m *EraExchangeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EraExchangeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EraExchangeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EraExchangeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EraExchangeRate.Merge(m, src)
}
func (m *EraExchangeRate) XXX_Size() int {
	return m.Size()
}
func (m *EraExchangeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_EraExchangeRate.DiscardUnknown(m)
}

var xxx_messageInfo_EraExchangeRate proto.InternalMessageInfo

func (m *EraExchangeRate) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EraExchangeRate) GetEra() uint32 {
	if m != nil {
		return m.Era
	}
	return 0
}

func init() {
	proto.RegisterType((*ExchangeRate)(nil), "stafiprotocol.stafihub.rate.ExchangeRate")
	proto.RegisterType((*EraExchangeRate)(nil), "stafiprotocol.stafihub.rate.EraExchangeRate")
}

func init() { proto.RegisterFile("rate/rate.proto", fileDescriptor_466953fd5f8ee6df) }

var fileDescriptor_466953fd5f8ee6df = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2c, 0x49,
	0xd5, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xd2, 0xc5, 0x25, 0x89, 0x69, 0x99,
	0x60, 0x76, 0x72, 0x7e, 0x8e, 0x1e, 0x98, 0x97, 0x51, 0x9a, 0xa4, 0x07, 0x52, 0x22, 0x25, 0x92,
	0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd3, 0x07, 0xb1, 0x20, 0x5a, 0x94, 0xb2, 0xb8, 0x78, 0x5c, 0x2b,
	0x92, 0x33, 0x12, 0xf3, 0xd2, 0x53, 0x83, 0x12, 0x4b, 0x52, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x52,
	0xf3, 0xf2, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x17, 0x2e, 0xd6,
	0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x90, 0xa8, 0x93, 0xde, 0x89, 0x7b, 0xf2, 0x0c, 0xb7,
	0xee, 0xc9, 0xab, 0xa5, 0x67, 0x96, 0x80, 0xcc, 0x4f, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0x86, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x2e,
	0xa9, 0xc9, 0x41, 0x10, 0xcd, 0x4a, 0xf5, 0x5c, 0xfc, 0xae, 0x45, 0x89, 0x44, 0x58, 0x27, 0xc0,
	0xc5, 0x9c, 0x5a, 0x94, 0x08, 0xb6, 0x8c, 0x37, 0x08, 0xc4, 0x44, 0x38, 0x80, 0x99, 0x02, 0x07,
	0x38, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x1e, 0x92, 0x41,
	0x28, 0x81, 0xa8, 0x0f, 0x0b, 0x44, 0xfd, 0x0a, 0x70, 0x48, 0x43, 0x0c, 0x4d, 0x62, 0x03, 0xcb,
	0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x0c, 0xc8, 0x37, 0x83, 0x01, 0x00, 0x00,
}

func (m *ExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRate(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRate(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EraExchangeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EraExchangeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EraExchangeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRate(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Era != 0 {
		i = encodeVarintRate(dAtA, i, uint64(m.Era))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRate(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRate(dAtA []byte, offset int, v uint64) int {
	offset -= sovRate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRate(uint64(l))
	}
	l = m.Value.Size()
	n += 1 + l + sovRate(uint64(l))
	return n
}

func (m *EraExchangeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRate(uint64(l))
	}
	if m.Era != 0 {
		n += 1 + sovRate(uint64(m.Era))
	}
	l = m.Value.Size()
	n += 1 + l + sovRate(uint64(l))
	return n
}

func sovRate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRate(x uint64) (n int) {
	return sovRate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRate
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
			return fmt.Errorf("proto: ExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRate
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
				return ErrInvalidLengthRate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRate
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
				return ErrInvalidLengthRate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EraExchangeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRate
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
			return fmt.Errorf("proto: EraExchangeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EraExchangeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRate
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
				return ErrInvalidLengthRate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Era", wireType)
			}
			m.Era = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Era |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRate
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
				return ErrInvalidLengthRate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRate
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
					return 0, ErrIntOverflowRate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRate
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
				return 0, ErrInvalidLengthRate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRate = fmt.Errorf("proto: unexpected end of group")
)
