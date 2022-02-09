// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rvote/proposal.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// ProposalStatus enumerates the valid statuses of a proposal.
type ProposalStatus int32

const (
	StatusInitiated ProposalStatus = 0
	StatusApproved  ProposalStatus = 1
	StatusExpired   ProposalStatus = 2
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_INITIATED",
	1: "PROPOSAL_STATUS_APPROVED",
	2: "PROPOSAL_STATUS_EXPIRED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_INITIATED": 0,
	"PROPOSAL_STATUS_APPROVED":  1,
	"PROPOSAL_STATUS_EXPIRED":   2,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a18424b8936b52ae, []int{0}
}

type Proposal struct {
	Content     *types.Any     `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Status      ProposalStatus `protobuf:"varint,2,opt,name=status,proto3,enum=stafihub.stafihub.rvote.ProposalStatus" json:"status,omitempty" yaml:"prop_status"`
	Voted       []string       `protobuf:"bytes,3,rep,name=voted,proto3" json:"voted,omitempty" yaml:"voted"`
	StartBlock  int64          `protobuf:"varint,4,opt,name=startBlock,proto3" json:"startBlock,omitempty" yaml:"start_block"`
	ExpireBlock int64          `protobuf:"varint,5,opt,name=expireBlock,proto3" json:"expireBlock,omitempty" yaml:"expire_block"`
}

func (m *Proposal) Reset()      { *m = Proposal{} }
func (*Proposal) ProtoMessage() {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18424b8936b52ae, []int{0}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("stafihub.stafihub.rvote.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterType((*Proposal)(nil), "stafihub.stafihub.rvote.Proposal")
}

func init() { proto.RegisterFile("rvote/proposal.proto", fileDescriptor_a18424b8936b52ae) }

var fileDescriptor_a18424b8936b52ae = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0x87, 0x7d, 0x49, 0xda, 0xc2, 0x05, 0x42, 0xb8, 0x46, 0xc4, 0xf5, 0x70, 0xb1, 0x3c, 0x40,
	0x54, 0x09, 0x1b, 0x05, 0x09, 0x89, 0x32, 0xd9, 0xc4, 0x83, 0x25, 0x44, 0x2c, 0x27, 0xfc, 0x11,
	0x8b, 0xe5, 0x24, 0xae, 0xb1, 0x48, 0x7c, 0x96, 0x7d, 0xa9, 0x9a, 0x8d, 0x11, 0x65, 0xe2, 0x0b,
	0x44, 0x42, 0x62, 0x65, 0xe4, 0x0b, 0xb0, 0x55, 0x4c, 0x1d, 0x3b, 0x45, 0x34, 0x59, 0x98, 0xf3,
	0x09, 0x90, 0xef, 0x6a, 0x08, 0x91, 0xba, 0xbd, 0xf6, 0xef, 0x79, 0xde, 0xf7, 0xb5, 0xef, 0x60,
	0x2d, 0x39, 0x21, 0xd4, 0xd7, 0xe2, 0x84, 0xc4, 0x24, 0xf5, 0x46, 0x6a, 0x9c, 0x10, 0x4a, 0x50,
	0x3d, 0xa5, 0xde, 0x71, 0xf8, 0x7e, 0xd2, 0x57, 0xff, 0x16, 0x8c, 0x93, 0x6a, 0x01, 0x09, 0x08,
	0x63, 0xb4, 0xac, 0xe2, 0xb8, 0x74, 0x30, 0x20, 0xe9, 0x98, 0xa4, 0x2e, 0x0f, 0xf8, 0x43, 0x1e,
	0x05, 0x84, 0x04, 0x23, 0x36, 0x80, 0x92, 0xfe, 0xe4, 0x58, 0xf3, 0xa2, 0x29, 0x8f, 0x94, 0x1f,
	0x05, 0x78, 0xc3, 0xbe, 0x9a, 0x8b, 0x9e, 0xc1, 0xbd, 0x01, 0x89, 0xa8, 0x1f, 0x51, 0x11, 0xc8,
	0xa0, 0x59, 0x6e, 0xd5, 0x54, 0x6e, 0xaa, 0xb9, 0xa9, 0xea, 0xd1, 0xd4, 0x28, 0xff, 0xfc, 0xfe,
	0x70, 0xef, 0x39, 0x07, 0x9d, 0xdc, 0x40, 0x6f, 0xe0, 0x6e, 0x4a, 0x3d, 0x3a, 0x49, 0xc5, 0x82,
	0x0c, 0x9a, 0x95, 0xd6, 0x03, 0xf5, 0x9a, 0xfd, 0xd5, 0x7c, 0x5e, 0x97, 0xe1, 0xc6, 0xbd, 0xf5,
	0xa2, 0x81, 0xa6, 0xde, 0x78, 0x74, 0xa4, 0x64, 0x7f, 0xc0, 0xe5, 0x5d, 0x14, 0xe7, 0xaa, 0x1d,
	0xba, 0x0f, 0x77, 0x32, 0x6d, 0x28, 0x16, 0xe5, 0x62, 0xf3, 0xa6, 0x51, 0x5d, 0x2f, 0x1a, 0xb7,
	0x38, 0xce, 0x5e, 0x2b, 0x0e, 0x8f, 0xd1, 0x13, 0x08, 0x53, 0xea, 0x25, 0xd4, 0x18, 0x91, 0xc1,
	0x07, 0xb1, 0x24, 0x83, 0x66, 0x71, 0xb3, 0x37, 0xcb, 0xdc, 0x7e, 0x16, 0x2a, 0xce, 0x06, 0x89,
	0x9e, 0xc2, 0xb2, 0x7f, 0x1a, 0x87, 0x89, 0xcf, 0xc5, 0x1d, 0x26, 0xd6, 0xd7, 0x8b, 0xc6, 0x3e,
	0x17, 0x79, 0x98, 0x9b, 0x9b, 0xec, 0x51, 0xe9, 0xf7, 0x97, 0x06, 0x38, 0xfc, 0x06, 0x60, 0xe5,
	0xff, 0x6f, 0x42, 0x2d, 0x78, 0x60, 0x3b, 0x1d, 0xbb, 0xd3, 0xd5, 0x5f, 0xb8, 0xdd, 0x9e, 0xde,
	0x7b, 0xd5, 0x75, 0xad, 0x97, 0x56, 0xcf, 0xd2, 0x7b, 0x66, 0xbb, 0x2a, 0x48, 0xfb, 0xb3, 0xb9,
	0x7c, 0x87, 0xa3, 0x56, 0x14, 0xd2, 0xd0, 0xcb, 0xf6, 0x7f, 0x04, 0xc5, 0x6d, 0x47, 0xb7, 0x6d,
	0xa7, 0xf3, 0xda, 0x6c, 0x57, 0x81, 0x84, 0x66, 0x73, 0xb9, 0xc2, 0x15, 0x3d, 0x8e, 0x13, 0x72,
	0xe2, 0x0f, 0x91, 0x0a, 0xeb, 0xdb, 0x86, 0xf9, 0xd6, 0xb6, 0x1c, 0xb3, 0x5d, 0x2d, 0x48, 0x77,
	0x67, 0x73, 0xf9, 0x36, 0x17, 0x4c, 0xb6, 0xf2, 0x50, 0x2a, 0x7d, 0xfa, 0x8a, 0x05, 0xc3, 0x3e,
	0xbb, 0xc4, 0xc2, 0xc5, 0x25, 0x16, 0x3e, 0x2e, 0xb1, 0x70, 0xb6, 0xc4, 0xe0, 0x7c, 0x89, 0xc1,
	0xaf, 0x25, 0x06, 0x9f, 0x57, 0x58, 0x38, 0x5f, 0x61, 0xe1, 0x62, 0x85, 0x85, 0x77, 0x87, 0x41,
	0x48, 0xb3, 0x23, 0x1b, 0x90, 0xb1, 0x96, 0x9f, 0xdf, 0xbf, 0xe2, 0x54, 0xe3, 0x57, 0x96, 0x4e,
	0x63, 0x3f, 0xed, 0xef, 0xb2, 0xeb, 0xf1, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x8d,
	0x10, 0x96, 0xc8, 0x02, 0x00, 0x00,
}

func (this *Proposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Proposal)
	if !ok {
		that2, ok := that.(Proposal)
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
	if !this.Content.Equal(that1.Content) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if len(this.Voted) != len(that1.Voted) {
		return false
	}
	for i := range this.Voted {
		if this.Voted[i] != that1.Voted[i] {
			return false
		}
	}
	if this.StartBlock != that1.StartBlock {
		return false
	}
	if this.ExpireBlock != that1.ExpireBlock {
		return false
	}
	return true
}
func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireBlock != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ExpireBlock))
		i--
		dAtA[i] = 0x28
	}
	if m.StartBlock != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Voted) > 0 {
		for iNdEx := len(m.Voted) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Voted[iNdEx])
			copy(dAtA[i:], m.Voted[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Voted[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovProposal(uint64(m.Status))
	}
	if len(m.Voted) > 0 {
		for _, s := range m.Voted {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if m.StartBlock != 0 {
		n += 1 + sovProposal(uint64(m.StartBlock))
	}
	if m.ExpireBlock != 0 {
		n += 1 + sovProposal(uint64(m.ExpireBlock))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voted = append(m.Voted, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireBlock", wireType)
			}
			m.ExpireBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
