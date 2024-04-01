// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/claimsmanager/v1/claimsmanager.proto

package types

import (
	fmt "fmt"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type ClaimType int32

const (
	// Undefined action (per protobuf spec)
	ClaimTypeUndefined    ClaimType = 0
	ClaimTypeLiquidToken  ClaimType = 1
	ClaimTypeOsmosisPool  ClaimType = 2
	ClaimTypeCrescentPool ClaimType = 3
	ClaimTypeSifchainPool ClaimType = 4
	ClaimTypeUmeeToken    ClaimType = 5
)

var ClaimType_name = map[int32]string{
	0: "ClaimTypeUndefined",
	1: "ClaimTypeLiquidToken",
	2: "ClaimTypeOsmosisPool",
	3: "ClaimTypeCrescentPool",
	4: "ClaimTypeSifchainPool",
	5: "ClaimTypeUmeeToken",
}

var ClaimType_value = map[string]int32{
	"ClaimTypeUndefined":    0,
	"ClaimTypeLiquidToken":  1,
	"ClaimTypeOsmosisPool":  2,
	"ClaimTypeCrescentPool": 3,
	"ClaimTypeSifchainPool": 4,
	"ClaimTypeUmeeToken":    5,
}

func (x ClaimType) String() string {
	return proto.EnumName(ClaimType_name, int32(x))
}

func (ClaimType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_086999747d797382, []int{0}
}

// Params holds parameters for the claimsmanager module.
type Params struct {
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_086999747d797382, []int{0}
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

// Claim define the users claim for holding assets within a given zone.
type Claim struct {
	UserAddress   string                                 `protobuf:"bytes,1,opt,name=user_address,json=userAddress,proto3" json:"user_address,omitempty"`
	ChainId       string                                 `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Module        ClaimType                              `protobuf:"varint,3,opt,name=module,proto3,enum=quicksilver.claimsmanager.v1.ClaimType" json:"module,omitempty"`
	SourceChainId string                                 `protobuf:"bytes,4,opt,name=source_chain_id,json=sourceChainId,proto3" json:"source_chain_id,omitempty"`
	XAmount       uint64                                 `protobuf:"varint,5,opt,name=_amount,json=Amount,proto3" json:"_amount,omitempty"` // Deprecated: Do not use.
	Amount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_086999747d797382, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

// Proof defines a type used to cryptographically prove a claim.
type Proof struct {
	Key       []byte           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data      []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ProofOps  *crypto.ProofOps `protobuf:"bytes,3,opt,name=proof_ops,proto3" json:"proof_ops,omitempty"`
	Height    int64            `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	ProofType string           `protobuf:"bytes,5,opt,name=proof_type,proto3" json:"proof_type,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_086999747d797382, []int{2}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("quicksilver.claimsmanager.v1.ClaimType", ClaimType_name, ClaimType_value)
	proto.RegisterType((*Params)(nil), "quicksilver.claimsmanager.v1.Params")
	proto.RegisterType((*Claim)(nil), "quicksilver.claimsmanager.v1.Claim")
	proto.RegisterType((*Proof)(nil), "quicksilver.claimsmanager.v1.Proof")
}

func init() {
	proto.RegisterFile("quicksilver/claimsmanager/v1/claimsmanager.proto", fileDescriptor_086999747d797382)
}

var fileDescriptor_086999747d797382 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x18, 0xb5, 0x13, 0xc7, 0x6d, 0xae, 0x05, 0xa2, 0x53, 0xa9, 0xdc, 0x16, 0x9c, 0x28, 0x43, 0x89,
	0x90, 0x62, 0xd3, 0x32, 0x51, 0x40, 0xa8, 0xc9, 0x14, 0x09, 0xa9, 0x91, 0x1b, 0x84, 0xc4, 0x62,
	0xb9, 0xf6, 0xc5, 0x39, 0x25, 0xbe, 0x73, 0xef, 0xce, 0x11, 0xe1, 0x17, 0x74, 0x64, 0x64, 0x8c,
	0xc4, 0x06, 0x6b, 0xff, 0x00, 0x5b, 0xc7, 0xaa, 0x13, 0x62, 0xa8, 0x50, 0xb2, 0xf0, 0x33, 0x90,
	0xcf, 0x56, 0x9a, 0x30, 0x30, 0xf9, 0xbe, 0xf7, 0xbe, 0xef, 0xde, 0xdd, 0x7b, 0x3e, 0xf0, 0xec,
	0x3c, 0xc1, 0xfe, 0x90, 0xe3, 0xd1, 0x18, 0x31, 0xdb, 0x1f, 0x79, 0x38, 0xe2, 0x91, 0x47, 0xbc,
	0x10, 0x31, 0x7b, 0x7c, 0xb0, 0x0a, 0x58, 0x31, 0xa3, 0x82, 0xc2, 0x47, 0x4b, 0x13, 0xd6, 0x6a,
	0xc3, 0xf8, 0x60, 0x77, 0xc7, 0xa7, 0x3c, 0xa2, 0xdc, 0x95, 0xbd, 0x76, 0x56, 0x64, 0x83, 0xbb,
	0x5b, 0x21, 0x0d, 0x69, 0x86, 0xa7, 0xab, 0x1c, 0x7d, 0x2c, 0x10, 0x09, 0x10, 0x8b, 0x30, 0x11,
	0xb6, 0xcf, 0x26, 0xb1, 0xa0, 0x76, 0xcc, 0x28, 0xed, 0x67, 0x74, 0x1d, 0x02, 0xbd, 0xeb, 0x31,
	0x2f, 0xe2, 0x47, 0xeb, 0x17, 0xd3, 0xaa, 0xf2, 0x65, 0x5a, 0x55, 0xea, 0x3f, 0x0a, 0xa0, 0xd4,
	0x4e, 0x85, 0xe1, 0x4b, 0xb0, 0x99, 0x70, 0xc4, 0x5c, 0x2f, 0x08, 0x18, 0xe2, 0xdc, 0x50, 0x6b,
	0x6a, 0xa3, 0xdc, 0x32, 0x6e, 0x2e, 0x9b, 0x5b, 0xb9, 0xf4, 0x71, 0xc6, 0x9c, 0x0a, 0x86, 0x49,
	0xe8, 0x6c, 0xa4, 0xdd, 0x39, 0x04, 0x77, 0xc0, 0xba, 0x3f, 0xf0, 0x30, 0x71, 0x71, 0x60, 0x14,
	0xd2, 0x41, 0x67, 0x4d, 0xd6, 0x9d, 0x00, 0xbe, 0x01, 0x7a, 0x44, 0x83, 0x64, 0x84, 0x8c, 0x62,
	0x4d, 0x6d, 0xdc, 0x3f, 0x7c, 0x62, 0xfd, 0xef, 0xd2, 0x96, 0x3c, 0x4c, 0x6f, 0x12, 0x23, 0x27,
	0x1f, 0x83, 0xfb, 0xe0, 0x01, 0xa7, 0x09, 0xf3, 0x91, 0xbb, 0x90, 0xd0, 0xa4, 0xc4, 0xbd, 0x0c,
	0x6e, 0xe7, 0x42, 0x7b, 0x60, 0xcd, 0xf5, 0x22, 0x9a, 0x10, 0x61, 0x94, 0x6a, 0x6a, 0x43, 0x6b,
	0x15, 0x0c, 0xd5, 0xd1, 0x8f, 0x25, 0x02, 0x7b, 0x40, 0xcf, 0x39, 0x5d, 0xde, 0xeb, 0xd5, 0xd5,
	0x6d, 0x55, 0xf9, 0x75, 0x5b, 0xdd, 0x0f, 0xb1, 0x18, 0x24, 0x67, 0x96, 0x4f, 0xa3, 0xdc, 0xe1,
	0xfc, 0xd3, 0xe4, 0xc1, 0xd0, 0x16, 0x93, 0x18, 0x71, 0xab, 0x43, 0xc4, 0xcd, 0x65, 0x13, 0xe4,
	0x2e, 0x74, 0x88, 0x70, 0xf2, 0xbd, 0x8e, 0xb4, 0xd4, 0xc7, 0xfa, 0x37, 0x15, 0x94, 0xba, 0xa9,
	0xcf, 0xb0, 0x02, 0x8a, 0x43, 0x34, 0x91, 0xd6, 0x6d, 0x3a, 0xe9, 0x12, 0x42, 0xa0, 0x05, 0x9e,
	0xf0, 0xa4, 0x29, 0x9b, 0x8e, 0x5c, 0xc3, 0x17, 0xa0, 0x2c, 0x63, 0x71, 0x69, 0xcc, 0xa5, 0x29,
	0x1b, 0x87, 0x7b, 0xd6, 0x5d, 0x74, 0x56, 0x16, 0x9d, 0x25, 0xb7, 0x3c, 0x89, 0xb9, 0x73, 0xd7,
	0x0d, 0xb7, 0x81, 0x3e, 0x40, 0x38, 0x1c, 0x08, 0x69, 0x41, 0xd1, 0xc9, 0x2b, 0x68, 0x02, 0x90,
	0x35, 0xa5, 0x27, 0x96, 0xd7, 0x2f, 0x3b, 0x4b, 0x48, 0x16, 0xf8, 0x9f, 0x69, 0x55, 0x79, 0xfa,
	0x5d, 0x05, 0xe5, 0x85, 0xc7, 0x70, 0x1b, 0xc0, 0x45, 0xf1, 0x8e, 0x04, 0xa8, 0x8f, 0x09, 0x0a,
	0x2a, 0x0a, 0x34, 0xc0, 0xd6, 0x02, 0x7f, 0x8b, 0xcf, 0x13, 0x1c, 0xf4, 0xe8, 0x10, 0x91, 0x8a,
	0xba, 0xc2, 0x9c, 0xa4, 0x86, 0x60, 0xde, 0xa5, 0x74, 0x54, 0x29, 0xc0, 0x1d, 0xf0, 0x70, 0xc1,
	0xb4, 0x19, 0xe2, 0x3e, 0x22, 0x42, 0x52, 0xc5, 0x15, 0xea, 0x14, 0xf7, 0x65, 0x8e, 0x92, 0xd2,
	0x56, 0x4f, 0x10, 0x21, 0x94, 0xe9, 0x94, 0x76, 0xb5, 0x8b, 0xaf, 0xa6, 0xd2, 0x7a, 0x7f, 0x35,
	0x33, 0xd5, 0xeb, 0x99, 0xa9, 0xfe, 0x9e, 0x99, 0xea, 0xe7, 0xb9, 0xa9, 0x5c, 0xcf, 0x4d, 0xe5,
	0xe7, 0xdc, 0x54, 0x3e, 0xbc, 0x5e, 0x0a, 0x6e, 0xe9, 0x87, 0x6a, 0x7e, 0xa2, 0x04, 0x2d, 0x03,
	0xf6, 0xc7, 0x7f, 0x9e, 0xa2, 0xcc, 0xf4, 0x4c, 0x97, 0x4f, 0xe2, 0xf9, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbd, 0x79, 0x19, 0x8a, 0xb4, 0x03, 0x00, 0x00,
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
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaimsmanager(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.XAmount != 0 {
		i = encodeVarintClaimsmanager(dAtA, i, uint64(m.XAmount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.SourceChainId) > 0 {
		i -= len(m.SourceChainId)
		copy(dAtA[i:], m.SourceChainId)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.SourceChainId)))
		i--
		dAtA[i] = 0x22
	}
	if m.Module != 0 {
		i = encodeVarintClaimsmanager(dAtA, i, uint64(m.Module))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserAddress) > 0 {
		i -= len(m.UserAddress)
		copy(dAtA[i:], m.UserAddress)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.UserAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProofType) > 0 {
		i -= len(m.ProofType)
		copy(dAtA[i:], m.ProofType)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.ProofType)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Height != 0 {
		i = encodeVarintClaimsmanager(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if m.ProofOps != nil {
		{
			size, err := m.ProofOps.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaimsmanager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintClaimsmanager(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimsmanager(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimsmanager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserAddress)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	if m.Module != 0 {
		n += 1 + sovClaimsmanager(uint64(m.Module))
	}
	l = len(m.SourceChainId)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	if m.XAmount != 0 {
		n += 1 + sovClaimsmanager(uint64(m.XAmount))
	}
	l = m.Amount.Size()
	n += 1 + l + sovClaimsmanager(uint64(l))
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	if m.ProofOps != nil {
		l = m.ProofOps.Size()
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovClaimsmanager(uint64(m.Height))
	}
	l = len(m.ProofType)
	if l > 0 {
		n += 1 + l + sovClaimsmanager(uint64(l))
	}
	return n
}

func sovClaimsmanager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimsmanager(x uint64) (n int) {
	return sovClaimsmanager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimsmanager
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
		default:
			iNdEx = preIndex
			skippy, err := skipClaimsmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimsmanager
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
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimsmanager
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			m.Module = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Module |= ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field XAmount", wireType)
			}
			m.XAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.XAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaimsmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimsmanager
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimsmanager
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ProofOps == nil {
				m.ProofOps = &crypto.ProofOps{}
			}
			if err := m.ProofOps.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimsmanager
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
				return ErrInvalidLengthClaimsmanager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimsmanager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaimsmanager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimsmanager
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
func skipClaimsmanager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimsmanager
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
					return 0, ErrIntOverflowClaimsmanager
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
					return 0, ErrIntOverflowClaimsmanager
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
				return 0, ErrInvalidLengthClaimsmanager
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimsmanager
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimsmanager
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimsmanager        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimsmanager          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimsmanager = fmt.Errorf("proto: unexpected end of group")
)
