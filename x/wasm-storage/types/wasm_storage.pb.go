// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sedachain/wasm_storage/v1/wasm_storage.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// WasmType is an enum for the type of wasm.
type WasmType int32

const (
	// An unspecified kind of wasm.
	WasmTypeNil WasmType = 0
	// A wasm that is a data request.
	WasmTypeDataRequest WasmType = 1
	// A wasm that is a DR tally.
	WasmTypeTally WasmType = 2
	// A wasm that is an overlay executor.
	WasmTypeDataRequestExecutor WasmType = 3
	// A wasm that is an overlay relayer.
	WasmTypeRelayer WasmType = 4
)

var WasmType_name = map[int32]string{
	0: "WASM_TYPE_UNSPECIFIED",
	1: "WASM_TYPE_DATA_REQUEST",
	2: "WASM_TYPE_TALLY",
	3: "WASM_TYPE_DATA_REQUEST_EXECUTOR",
	4: "WASM_TYPE_RELAYER",
}

var WasmType_value = map[string]int32{
	"WASM_TYPE_UNSPECIFIED":           0,
	"WASM_TYPE_DATA_REQUEST":          1,
	"WASM_TYPE_TALLY":                 2,
	"WASM_TYPE_DATA_REQUEST_EXECUTOR": 3,
	"WASM_TYPE_RELAYER":               4,
}

func (x WasmType) String() string {
	return proto.EnumName(WasmType_name, int32(x))
}

func (WasmType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a4bda463450c942, []int{0}
}

// A Wasm msg.
type Wasm struct {
	Hash     []byte    `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Bytecode []byte    `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
	WasmType WasmType  `protobuf:"varint,3,opt,name=wasm_type,json=wasmType,proto3,enum=sedachain.wasm_storage.v1.WasmType" json:"wasm_type,omitempty"`
	AddedAt  time.Time `protobuf:"bytes,4,opt,name=added_at,json=addedAt,proto3,stdtime" json:"added_at"`
	// prune_height represents the block height till which the Wasm will
	// stay on the chain. At height > prune_height, the wasm should be pruned.
	PruneHeight uint64 `protobuf:"varint,5,opt,name=prune_height,json=pruneHeight,proto3" json:"prune_height,omitempty"`
}

func (m *Wasm) Reset()         { *m = Wasm{} }
func (m *Wasm) String() string { return proto.CompactTextString(m) }
func (*Wasm) ProtoMessage()    {}
func (*Wasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4bda463450c942, []int{0}
}
func (m *Wasm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wasm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wasm.Merge(m, src)
}
func (m *Wasm) XXX_Size() int {
	return m.Size()
}
func (m *Wasm) XXX_DiscardUnknown() {
	xxx_messageInfo_Wasm.DiscardUnknown(m)
}

var xxx_messageInfo_Wasm proto.InternalMessageInfo

func (m *Wasm) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Wasm) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

func (m *Wasm) GetWasmType() WasmType {
	if m != nil {
		return m.WasmType
	}
	return WasmTypeNil
}

func (m *Wasm) GetAddedAt() time.Time {
	if m != nil {
		return m.AddedAt
	}
	return time.Time{}
}

func (m *Wasm) GetPruneHeight() uint64 {
	if m != nil {
		return m.PruneHeight
	}
	return 0
}

// Params to define the max wasm size allowed.
type Params struct {
	MaxWasmSize uint64 `protobuf:"varint,1,opt,name=max_wasm_size,json=maxWasmSize,proto3" json:"max_wasm_size,omitempty"`
	// WasmTTL represents the number of block a wasm's life is extended when it's
	// created or used.
	WasmTTL uint64 `protobuf:"varint,2,opt,name=wasm_ttl,json=wasmTtl,proto3" json:"wasm_ttl,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a4bda463450c942, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxWasmSize() uint64 {
	if m != nil {
		return m.MaxWasmSize
	}
	return 0
}

func (m *Params) GetWasmTTL() uint64 {
	if m != nil {
		return m.WasmTTL
	}
	return 0
}

func init() {
	proto.RegisterEnum("sedachain.wasm_storage.v1.WasmType", WasmType_name, WasmType_value)
	proto.RegisterType((*Wasm)(nil), "sedachain.wasm_storage.v1.Wasm")
	proto.RegisterType((*Params)(nil), "sedachain.wasm_storage.v1.Params")
}

func init() {
	proto.RegisterFile("sedachain/wasm_storage/v1/wasm_storage.proto", fileDescriptor_9a4bda463450c942)
}

var fileDescriptor_9a4bda463450c942 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x3d, 0xa9, 0xbf, 0xc6, 0xdf, 0xa4, 0x25, 0xe9, 0x94, 0x3f, 0xc1, 0x48, 0xb6, 0x09,
	0x52, 0x15, 0x45, 0xd4, 0x56, 0xdb, 0x05, 0x12, 0x1b, 0x70, 0x1a, 0x23, 0x2a, 0x85, 0x92, 0x3a,
	0x8e, 0x4a, 0xba, 0xb1, 0x26, 0xc9, 0x60, 0x5b, 0xb2, 0xeb, 0x60, 0x4f, 0xda, 0xa4, 0x4f, 0x80,
	0xc2, 0xa6, 0x2f, 0x10, 0x09, 0x89, 0x97, 0xe9, 0xb2, 0x4b, 0x16, 0xa8, 0xa0, 0x64, 0xc3, 0x63,
	0xa0, 0x8c, 0x71, 0xa2, 0x4a, 0x65, 0x77, 0xef, 0x9d, 0x73, 0x67, 0x7e, 0x47, 0x73, 0xe0, 0xf3,
	0x98, 0xf4, 0x70, 0xd7, 0xc5, 0xde, 0xa9, 0x76, 0x8e, 0xe3, 0xc0, 0x8e, 0x69, 0x18, 0x61, 0x87,
	0x68, 0x67, 0x3b, 0xb7, 0x7a, 0xb5, 0x1f, 0x85, 0x34, 0x44, 0x8f, 0x17, 0x6a, 0xf5, 0xd6, 0xe9,
	0xd9, 0x8e, 0x78, 0xdf, 0x09, 0x9d, 0x90, 0xa9, 0xb4, 0x79, 0x95, 0x2c, 0x88, 0xb2, 0x13, 0x86,
	0x8e, 0x4f, 0x34, 0xd6, 0x75, 0x06, 0x1f, 0x35, 0xea, 0x05, 0x24, 0xa6, 0x38, 0xe8, 0x27, 0x82,
	0xd2, 0x0f, 0x00, 0xf9, 0x63, 0x1c, 0x07, 0x08, 0x41, 0xde, 0xc5, 0xb1, 0x5b, 0x04, 0x0a, 0x28,
	0xaf, 0x99, 0xac, 0x46, 0x22, 0x14, 0x3a, 0x23, 0x4a, 0xba, 0x61, 0x8f, 0x14, 0x33, 0x6c, 0xbe,
	0xe8, 0xd1, 0x6b, 0xf8, 0x3f, 0x43, 0xa0, 0xa3, 0x3e, 0x29, 0xae, 0x28, 0xa0, 0x7c, 0x6f, 0xf7,
	0x99, 0xfa, 0x4f, 0x3c, 0x75, 0xfe, 0x86, 0x35, 0xea, 0x13, 0x53, 0x38, 0xff, 0x5b, 0xa1, 0x57,
	0x50, 0xc0, 0xbd, 0x1e, 0xe9, 0xd9, 0x98, 0x16, 0x79, 0x05, 0x94, 0x73, 0xbb, 0xa2, 0x9a, 0xe0,
	0xaa, 0x29, 0xae, 0x6a, 0xa5, 0xb8, 0x55, 0xe1, 0xea, 0x46, 0xe6, 0x2e, 0x7f, 0xca, 0xc0, 0xcc,
	0xb2, 0x2d, 0x9d, 0xa2, 0xa7, 0x70, 0xad, 0x1f, 0x0d, 0x4e, 0x89, 0xed, 0x12, 0xcf, 0x71, 0x69,
	0xf1, 0x3f, 0x05, 0x94, 0x79, 0x33, 0xc7, 0x66, 0x6f, 0xd9, 0xa8, 0x74, 0x02, 0x57, 0x1b, 0x38,
	0xc2, 0x41, 0x8c, 0x4a, 0x70, 0x3d, 0xc0, 0x43, 0x3b, 0xe1, 0xf2, 0x2e, 0x08, 0x33, 0xca, 0x9b,
	0xb9, 0x00, 0x0f, 0xe7, 0x6c, 0x4d, 0xef, 0x82, 0xa0, 0x2d, 0x28, 0x24, 0x9e, 0xa8, 0xcf, 0xfc,
	0xf2, 0xd5, 0xdc, 0xf4, 0x46, 0xce, 0x32, 0x76, 0xab, 0x6e, 0x66, 0x19, 0x3a, 0xf5, 0x5f, 0xf2,
	0xbf, 0xbf, 0xca, 0xa0, 0xf2, 0x25, 0x03, 0x85, 0xd4, 0x16, 0xaa, 0xc0, 0x07, 0xc7, 0x7a, 0xf3,
	0x9d, 0x6d, 0xb5, 0x1b, 0x86, 0xdd, 0x3a, 0x6c, 0x36, 0x8c, 0xfd, 0x83, 0x37, 0x07, 0x46, 0xad,
	0xc0, 0x89, 0xf9, 0xf1, 0x44, 0xc9, 0xa5, 0xc2, 0x43, 0xcf, 0x47, 0x7b, 0xf0, 0xe1, 0x52, 0x5b,
	0xd3, 0x2d, 0xdd, 0x36, 0x8d, 0xa3, 0x96, 0xd1, 0xb4, 0x0a, 0x40, 0x7c, 0x34, 0x9e, 0x28, 0x9b,
	0xa9, 0xb8, 0x86, 0x29, 0x36, 0xc9, 0xa7, 0x01, 0x89, 0x29, 0xda, 0x82, 0xf9, 0xe5, 0x92, 0xa5,
	0xd7, 0xeb, 0xed, 0x42, 0x46, 0xdc, 0x18, 0x4f, 0x94, 0xf5, 0x54, 0x6d, 0x61, 0xdf, 0x1f, 0xa1,
	0x1a, 0x94, 0xef, 0xbe, 0xdc, 0x36, 0x3e, 0x18, 0xfb, 0x2d, 0xeb, 0xbd, 0x59, 0x58, 0x11, 0xe5,
	0xf1, 0x44, 0x79, 0x72, 0xc7, 0x2b, 0xc6, 0x90, 0x74, 0x07, 0x34, 0x8c, 0x50, 0x05, 0x6e, 0x2c,
	0x6f, 0x31, 0x8d, 0xba, 0xde, 0x36, 0xcc, 0x02, 0x2f, 0x6e, 0x8e, 0x27, 0x4a, 0x7e, 0xf1, 0x95,
	0xc4, 0xc7, 0x23, 0x12, 0x89, 0xfc, 0xe7, 0x6f, 0x12, 0x57, 0x3d, 0xba, 0x9a, 0x4a, 0xe0, 0x7a,
	0x2a, 0x81, 0x5f, 0x53, 0x09, 0x5c, 0xce, 0x24, 0xee, 0x7a, 0x26, 0x71, 0xdf, 0x67, 0x12, 0x77,
	0xf2, 0xc2, 0xf1, 0xa8, 0x3b, 0xe8, 0xa8, 0xdd, 0x30, 0xd0, 0xe6, 0x01, 0x61, 0x9f, 0xdb, 0x0d,
	0x7d, 0xd6, 0x6c, 0x27, 0xd9, 0x1f, 0xb2, 0xb4, 0x6f, 0xa7, 0xe9, 0x9f, 0x87, 0x2a, 0xee, 0xac,
	0x32, 0xe5, 0xde, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x88, 0x42, 0xe7, 0x24, 0x03, 0x00,
	0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxWasmSize != that1.MaxWasmSize {
		return false
	}
	if this.WasmTTL != that1.WasmTTL {
		return false
	}
	return true
}
func (m *Wasm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wasm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wasm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PruneHeight != 0 {
		i = encodeVarintWasmStorage(dAtA, i, uint64(m.PruneHeight))
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.AddedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.AddedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintWasmStorage(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.WasmType != 0 {
		i = encodeVarintWasmStorage(dAtA, i, uint64(m.WasmType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Bytecode) > 0 {
		i -= len(m.Bytecode)
		copy(dAtA[i:], m.Bytecode)
		i = encodeVarintWasmStorage(dAtA, i, uint64(len(m.Bytecode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintWasmStorage(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WasmTTL != 0 {
		i = encodeVarintWasmStorage(dAtA, i, uint64(m.WasmTTL))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxWasmSize != 0 {
		i = encodeVarintWasmStorage(dAtA, i, uint64(m.MaxWasmSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWasmStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovWasmStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Wasm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovWasmStorage(uint64(l))
	}
	l = len(m.Bytecode)
	if l > 0 {
		n += 1 + l + sovWasmStorage(uint64(l))
	}
	if m.WasmType != 0 {
		n += 1 + sovWasmStorage(uint64(m.WasmType))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.AddedAt)
	n += 1 + l + sovWasmStorage(uint64(l))
	if m.PruneHeight != 0 {
		n += 1 + sovWasmStorage(uint64(m.PruneHeight))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxWasmSize != 0 {
		n += 1 + sovWasmStorage(uint64(m.MaxWasmSize))
	}
	if m.WasmTTL != 0 {
		n += 1 + sovWasmStorage(uint64(m.WasmTTL))
	}
	return n
}

func sovWasmStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWasmStorage(x uint64) (n int) {
	return sovWasmStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wasm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasmStorage
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
			return fmt.Errorf("proto: Wasm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wasm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWasmStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytecode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWasmStorage
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthWasmStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytecode = append(m.Bytecode[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytecode == nil {
				m.Bytecode = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmType", wireType)
			}
			m.WasmType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WasmType |= WasmType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
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
				return ErrInvalidLengthWasmStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWasmStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.AddedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PruneHeight", wireType)
			}
			m.PruneHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PruneHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWasmStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasmStorage
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWasmStorage
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWasmSize", wireType)
			}
			m.MaxWasmSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxWasmSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmTTL", wireType)
			}
			m.WasmTTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWasmStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WasmTTL |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWasmStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWasmStorage
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
func skipWasmStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWasmStorage
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
					return 0, ErrIntOverflowWasmStorage
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
					return 0, ErrIntOverflowWasmStorage
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
				return 0, ErrInvalidLengthWasmStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWasmStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWasmStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWasmStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWasmStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWasmStorage = fmt.Errorf("proto: unexpected end of group")
)
