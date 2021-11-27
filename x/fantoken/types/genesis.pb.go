// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/fantoken/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the fantoken module's genesis state
type GenesisState struct {
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Tokens      []FanToken   `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens"`
	BurnedCoins []types.Coin `protobuf:"bytes,3,rep,name=burned_coins,json=burnedCoins,proto3" json:"burned_coins"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a9d02535fd9f192, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTokens() []FanToken {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GenesisState) GetBurnedCoins() []types.Coin {
	if m != nil {
		return m.BurnedCoins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bitsong.fantoken.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("bitsong/fantoken/v1beta1/genesis.proto", fileDescriptor_3a9d02535fd9f192)
}

var fileDescriptor_3a9d02535fd9f192 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x8a, 0x3a, 0xa4, 0x9d, 0x22, 0x86, 0xd0, 0xc1, 0x54, 0x1d, 0xa0, 0x93, 0xad,
	0x96, 0x85, 0x09, 0xa1, 0x22, 0xc1, 0x8a, 0x0a, 0x13, 0x0b, 0xb2, 0x53, 0x37, 0xb5, 0x68, 0x7c,
	0x55, 0xce, 0x41, 0xf0, 0x16, 0x3c, 0x56, 0xc7, 0x0e, 0x0c, 0x4c, 0x08, 0x25, 0x2f, 0x82, 0x12,
	0x27, 0x61, 0xca, 0x66, 0x9d, 0xbf, 0xff, 0xd3, 0x7f, 0xe7, 0x9f, 0x4b, 0x6d, 0x11, 0x4c, 0xcc,
	0xd7, 0xc2, 0x58, 0x78, 0x55, 0x86, 0xbf, 0xcd, 0xa4, 0xb2, 0x62, 0xc6, 0x63, 0x65, 0x14, 0x6a,
	0x64, 0xbb, 0x14, 0x2c, 0x04, 0x61, 0xcd, 0xb1, 0x86, 0x63, 0x35, 0x37, 0x3a, 0x89, 0x21, 0x86,
	0x0a, 0xe2, 0xe5, 0xcb, 0xf1, 0xa3, 0x8b, 0x4e, 0x6f, 0x2b, 0x70, 0x20, 0x8d, 0x00, 0x13, 0x40,
	0x2e, 0x05, 0xaa, 0x96, 0x89, 0x40, 0xd7, 0xff, 0x93, 0x2f, 0xe2, 0x0f, 0xef, 0x5d, 0x95, 0x47,
	0x2b, 0xac, 0x0a, 0xae, 0xfd, 0xfe, 0x4e, 0xa4, 0x22, 0xc1, 0x90, 0x8c, 0xc9, 0x74, 0x30, 0x1f,
	0xb3, 0xae, 0x6a, 0xec, 0xa1, 0xe2, 0x16, 0xc7, 0xfb, 0x9f, 0x33, 0x6f, 0x59, 0xa7, 0x82, 0x1b,
	0xbf, 0x5f, 0x51, 0x18, 0x1e, 0x8d, 0x7b, 0xd3, 0xc1, 0x7c, 0xd2, 0x9d, 0xbf, 0x13, 0xe6, 0xa9,
	0x1c, 0x34, 0x06, 0x97, 0x0b, 0x16, 0xfe, 0x50, 0x66, 0xa9, 0x51, 0xab, 0x97, 0xb2, 0x27, 0x86,
	0xbd, 0xca, 0x73, 0xca, 0xdc, 0x26, 0xac, 0xdc, 0xa4, 0x55, 0xdc, 0x82, 0x6e, 0xe2, 0x03, 0x17,
	0x2a, 0x27, 0xb8, 0x58, 0xee, 0x73, 0x4a, 0x0e, 0x39, 0x25, 0xbf, 0x39, 0x25, 0x9f, 0x05, 0xf5,
	0x0e, 0x05, 0xf5, 0xbe, 0x0b, 0xea, 0x3d, 0x5f, 0xc5, 0xda, 0x6e, 0x32, 0xc9, 0x22, 0x48, 0x78,
	0xdd, 0x0c, 0xd6, 0x6b, 0x1d, 0x69, 0xb1, 0xe5, 0xd1, 0x46, 0x68, 0x93, 0xc0, 0x2a, 0xdb, 0x2a,
	0xe4, 0xef, 0xff, 0xb7, 0xb5, 0x1f, 0x3b, 0x85, 0xb2, 0x5f, 0x5d, 0xec, 0xf2, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0xcd, 0xb8, 0xf1, 0xd4, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnedCoins) > 0 {
		for iNdEx := len(m.BurnedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BurnedCoins) > 0 {
		for _, e := range m.BurnedCoins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, FanToken{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnedCoins = append(m.BurnedCoins, types.Coin{})
			if err := m.BurnedCoins[len(m.BurnedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)