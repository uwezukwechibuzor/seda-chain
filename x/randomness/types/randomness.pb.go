// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/randomness/v1/randomness.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type ValidatorVRF struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// vrf_pubkey is the public key of the validator's VRF key pair
	VrfPubkey *types.Any `protobuf:"bytes,2,opt,name=vrf_pubkey,json=vrfPubkey,proto3" json:"vrf_pubkey,omitempty"`
}

func (m *ValidatorVRF) Reset()         { *m = ValidatorVRF{} }
func (m *ValidatorVRF) String() string { return proto.CompactTextString(m) }
func (*ValidatorVRF) ProtoMessage()    {}
func (*ValidatorVRF) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bb7c7510d674163, []int{0}
}
func (m *ValidatorVRF) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorVRF) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorVRF.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorVRF) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorVRF.Merge(m, src)
}
func (m *ValidatorVRF) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorVRF) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorVRF.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorVRF proto.InternalMessageInfo

func (m *ValidatorVRF) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *ValidatorVRF) GetVrfPubkey() *types.Any {
	if m != nil {
		return m.VrfPubkey
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorVRF)(nil), "sedachain.randomness.v1.ValidatorVRF")
}

func init() {
	proto.RegisterFile("sedachain/randomness/v1/randomness.proto", fileDescriptor_5bb7c7510d674163)
}

var fileDescriptor_5bb7c7510d674163 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x4e, 0x4d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4a, 0xcc, 0x4b, 0xc9, 0xcf, 0xcd, 0x4b, 0x2d, 0x2e,
	0xd6, 0x2f, 0x33, 0x44, 0xe2, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0xc3, 0x55, 0xea,
	0x21, 0xc9, 0x95, 0x19, 0x4a, 0x49, 0xa6, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0x95, 0x25,
	0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0xf4, 0x48, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x99,
	0xfa, 0x20, 0x16, 0x54, 0x54, 0x32, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x38, 0x1e, 0x22, 0x01, 0xe1,
	0x40, 0xa4, 0x94, 0x16, 0x31, 0x72, 0xf1, 0x84, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17,
	0x85, 0x05, 0xb9, 0x09, 0x39, 0x73, 0x09, 0xe4, 0x17, 0xa4, 0x16, 0x81, 0xb8, 0xf1, 0x89, 0x29,
	0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x12, 0x97, 0xb6, 0xe8,
	0x8a, 0x40, 0x35, 0x3b, 0x42, 0x64, 0x82, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0x83, 0xf8, 0x61, 0x3a,
	0xa0, 0xc2, 0x42, 0xbe, 0x5c, 0x5c, 0x65, 0x45, 0x69, 0xf1, 0x05, 0xa5, 0x49, 0xd9, 0xa9, 0x95,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x10, 0x67, 0xeb, 0xc1, 0x9c, 0xad, 0xe7,
	0x98, 0x57, 0xe9, 0x24, 0x71, 0x0a, 0x61, 0x68, 0x72, 0x51, 0x65, 0x41, 0x49, 0xbe, 0x5e, 0x40,
	0x69, 0x92, 0x77, 0x6a, 0x65, 0x10, 0x67, 0x59, 0x51, 0x5a, 0x00, 0xd8, 0x00, 0x27, 0xff, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x05, 0x17, 0xd8, 0xec, 0xe4, 0xfc, 0x1c, 0x30, 0x47, 0x17,
	0x12, 0xcc, 0x15, 0xc8, 0x01, 0x5d, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x56, 0x67, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x84, 0x1f, 0x79, 0x1b, 0x8d, 0x01, 0x00, 0x00,
}

func (m *ValidatorVRF) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorVRF) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorVRF) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VrfPubkey != nil {
		{
			size, err := m.VrfPubkey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRandomness(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintRandomness(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRandomness(dAtA []byte, offset int, v uint64) int {
	offset -= sovRandomness(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorVRF) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovRandomness(uint64(l))
	}
	if m.VrfPubkey != nil {
		l = m.VrfPubkey.Size()
		n += 1 + l + sovRandomness(uint64(l))
	}
	return n
}

func sovRandomness(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRandomness(x uint64) (n int) {
	return sovRandomness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorVRF) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRandomness
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
			return fmt.Errorf("proto: ValidatorVRF: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorVRF: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomness
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
				return ErrInvalidLengthRandomness
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRandomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VrfPubkey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRandomness
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
				return ErrInvalidLengthRandomness
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRandomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VrfPubkey == nil {
				m.VrfPubkey = &types.Any{}
			}
			if err := m.VrfPubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRandomness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRandomness
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
func skipRandomness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRandomness
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
					return 0, ErrIntOverflowRandomness
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
					return 0, ErrIntOverflowRandomness
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
				return 0, ErrInvalidLengthRandomness
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRandomness
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRandomness
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRandomness        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRandomness          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRandomness = fmt.Errorf("proto: unexpected end of group")
)