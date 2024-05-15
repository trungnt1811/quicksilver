// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/quicksilver-zone/quicksilver/x/claimsmanager/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSubmitClaim represents a message type for submitting a participation
// claim regarding the given zone (chain).
type MsgSubmitClaim struct {
	UserAddress string          `protobuf:"bytes,1,opt,name=user_address,proto3" json:"user_address,omitempty"`
	Zone        string          `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	SrcZone     string          `protobuf:"bytes,3,opt,name=src_zone,proto3" json:"src_zone,omitempty"`
	ClaimType   types.ClaimType `protobuf:"varint,4,opt,name=claim_type,proto3,enum=quicksilver.claimsmanager.v1.ClaimType" json:"claim_type,omitempty"`
	Proofs      []*types.Proof  `protobuf:"bytes,5,rep,name=proofs,proto3" json:"proofs,omitempty"`
}

func (m *MsgSubmitClaim) Reset()         { *m = MsgSubmitClaim{} }
func (m *MsgSubmitClaim) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitClaim) ProtoMessage()    {}
func (*MsgSubmitClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87e3ea017f90b50, []int{0}
}
func (m *MsgSubmitClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitClaim.Merge(m, src)
}
func (m *MsgSubmitClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitClaim proto.InternalMessageInfo

// MsgSubmitClaimResponse defines the MsgSubmitClaim response type.
type MsgSubmitClaimResponse struct {
}

func (m *MsgSubmitClaimResponse) Reset()         { *m = MsgSubmitClaimResponse{} }
func (m *MsgSubmitClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitClaimResponse) ProtoMessage()    {}
func (*MsgSubmitClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87e3ea017f90b50, []int{1}
}
func (m *MsgSubmitClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitClaimResponse.Merge(m, src)
}
func (m *MsgSubmitClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitClaimResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitClaim)(nil), "quicksilver.participationrewards.v1.MsgSubmitClaim")
	proto.RegisterType((*MsgSubmitClaimResponse)(nil), "quicksilver.participationrewards.v1.MsgSubmitClaimResponse")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/messages.proto", fileDescriptor_b87e3ea017f90b50)
}

var fileDescriptor_b87e3ea017f90b50 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xc7, 0xe3, 0xa4, 0xbf, 0xaa, 0xbf, 0x2b, 0xea, 0x70, 0x2a, 0xc8, 0x58, 0xc8, 0x8d, 0xd2,
	0x81, 0x08, 0x29, 0x36, 0x71, 0x16, 0xd4, 0x80, 0x04, 0xa5, 0xd0, 0xa9, 0x52, 0xe5, 0x32, 0x31,
	0x10, 0x5d, 0x9c, 0xe3, 0x38, 0x61, 0xfb, 0x1d, 0x77, 0x17, 0xd3, 0x32, 0x32, 0x31, 0x21, 0x24,
	0xfe, 0x81, 0xae, 0xec, 0x2c, 0xfc, 0x07, 0x8c, 0x15, 0x2c, 0x8c, 0x28, 0x01, 0x89, 0x3f, 0x03,
	0xf9, 0xec, 0x56, 0xb6, 0x14, 0xa1, 0x08, 0xb6, 0xbb, 0xe7, 0xfb, 0xbc, 0xf7, 0xfd, 0xbe, 0xf7,
	0x8c, 0x82, 0x17, 0x53, 0x1e, 0x3d, 0x57, 0x3c, 0xce, 0xa8, 0xf4, 0x05, 0x91, 0x9a, 0x47, 0x5c,
	0x10, 0xcd, 0x21, 0x95, 0xf4, 0x25, 0x91, 0x13, 0xe5, 0x67, 0x7d, 0x3f, 0xa1, 0x4a, 0x11, 0x46,
	0x95, 0x27, 0x24, 0x68, 0xc0, 0xdb, 0x15, 0xc6, 0x5b, 0xc4, 0x78, 0x59, 0xdf, 0xb9, 0x1a, 0x81,
	0x4a, 0x40, 0x8d, 0x0c, 0xe2, 0x17, 0x97, 0x82, 0x77, 0x36, 0x19, 0x30, 0x28, 0xe2, 0xf9, 0xa9,
	0x8c, 0x5e, 0x63, 0x00, 0x2c, 0xa6, 0x3e, 0x11, 0xdc, 0x27, 0x69, 0x0a, 0xda, 0x64, 0x3c, 0x67,
	0x6e, 0x56, 0x75, 0x46, 0x31, 0xe1, 0x89, 0x4a, 0x48, 0x4a, 0x18, 0x95, 0xb9, 0xc0, 0x5a, 0xa0,
	0x24, 0x06, 0xcb, 0x38, 0x13, 0x12, 0x04, 0x28, 0x12, 0x97, 0x65, 0x3a, 0x6f, 0x9b, 0x68, 0xe3,
	0x40, 0xb1, 0xa3, 0xe9, 0x38, 0xe1, 0xfa, 0x7e, 0x9e, 0x15, 0xdf, 0x46, 0x97, 0xa6, 0x8a, 0xca,
	0x11, 0x99, 0x4c, 0x24, 0x55, 0xca, 0xb6, 0xda, 0x56, 0xf7, 0xff, 0x5d, 0xfb, 0xcb, 0xc7, 0xde,
	0x66, 0xe9, 0xea, 0x5e, 0xf1, 0xe5, 0x48, 0x4b, 0x9e, 0xb2, 0xb0, 0xf6, 0x1a, 0x63, 0xb4, 0xf2,
	0x0a, 0x52, 0x6a, 0x37, 0x73, 0x2a, 0x34, 0x67, 0xec, 0xa0, 0x35, 0x25, 0xa3, 0x91, 0x89, 0xb7,
	0x4c, 0xfc, 0xe2, 0x8e, 0xf7, 0x11, 0x32, 0x66, 0x46, 0xfa, 0x44, 0x50, 0x7b, 0xa5, 0x6d, 0x75,
	0x37, 0x82, 0xeb, 0x5e, 0xb5, 0xe1, 0x75, 0xaf, 0x59, 0xdf, 0x33, 0x32, 0x1f, 0x9d, 0x08, 0x1a,
	0x56, 0x50, 0x3c, 0x44, 0xab, 0x42, 0x02, 0x3c, 0x55, 0xf6, 0x7f, 0xed, 0x56, 0x77, 0x3d, 0xd8,
	0xfe, 0x73, 0x92, 0xc3, 0xfc, 0x6d, 0x58, 0x22, 0x3b, 0x6b, 0x6f, 0x4e, 0xb7, 0x1a, 0xbf, 0x4e,
	0xb7, 0x1a, 0x1d, 0x1b, 0x5d, 0xa9, 0xf7, 0x23, 0xa4, 0x4a, 0x40, 0xaa, 0x68, 0xf0, 0xa1, 0x85,
	0x5a, 0x07, 0x8a, 0xe1, 0x4f, 0x16, 0x5a, 0xaf, 0xf6, 0x6b, 0xe0, 0x2d, 0xb1, 0x1e, 0x5e, 0x3d,
	0xa9, 0x33, 0xfc, 0x0b, 0xe8, 0x5c, 0x49, 0xe7, 0xd6, 0xeb, 0xaf, 0x3f, 0xde, 0x37, 0x83, 0x1d,
	0xeb, 0x46, 0xa7, 0xe7, 0x57, 0xa7, 0xae, 0x8f, 0xcd, 0x8c, 0x17, 0xcd, 0xde, 0x34, 0x00, 0xff,
	0xb4, 0xd0, 0xe5, 0x7d, 0xc8, 0x42, 0x9a, 0x40, 0x46, 0x0f, 0xf3, 0x0d, 0x88, 0x20, 0xde, 0x23,
	0x9a, 0xe0, 0x3b, 0xcb, 0x0a, 0x5a, 0x88, 0x3b, 0x0f, 0xfe, 0x09, 0xbf, 0x70, 0xf6, 0xd0, 0x38,
	0xbb, 0x9b, 0x3b, 0x1b, 0x2e, 0xeb, 0x4c, 0x9a, 0x74, 0xc5, 0x7f, 0x17, 0x41, 0x3c, 0x21, 0x9a,
	0xec, 0x3e, 0xf9, 0x3c, 0x73, 0xad, 0xb3, 0x99, 0x6b, 0x7d, 0x9f, 0xb9, 0xd6, 0xbb, 0xb9, 0xdb,
	0x38, 0x9b, 0xbb, 0x8d, 0x6f, 0x73, 0xb7, 0xf1, 0x78, 0x8f, 0x71, 0xfd, 0x6c, 0x3a, 0xf6, 0x22,
	0x48, 0xaa, 0x05, 0x7a, 0xf9, 0x32, 0xd6, 0x2a, 0x1e, 0x2f, 0xae, 0x96, 0xef, 0x9a, 0x1a, 0xaf,
	0x9a, 0x6a, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd7, 0x14, 0x51, 0x4e, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SubmitClaim(ctx context.Context, in *MsgSubmitClaim, opts ...grpc.CallOption) (*MsgSubmitClaimResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	GovRemoveProtocolData(ctx context.Context, in *MsgGovRemoveProtocolData, opts ...grpc.CallOption) (*MsgGovRemoveProtocolDataResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitClaim(ctx context.Context, in *MsgSubmitClaim, opts ...grpc.CallOption) (*MsgSubmitClaimResponse, error) {
	out := new(MsgSubmitClaimResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.participationrewards.v1.Msg/SubmitClaim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GovRemoveProtocolData(ctx context.Context, in *MsgGovRemoveProtocolData, opts ...grpc.CallOption) (*MsgGovRemoveProtocolDataResponse, error) {
	out := new(MsgGovRemoveProtocolDataResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.participationrewards.v1.Msg/GovRemoveProtocolData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitClaim(context.Context, *MsgSubmitClaim) (*MsgSubmitClaimResponse, error)
	// SignalIntent defines a method for signalling voting intent for one or more
	// validators.
	GovRemoveProtocolData(context.Context, *MsgGovRemoveProtocolData) (*MsgGovRemoveProtocolDataResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitClaim(ctx context.Context, req *MsgSubmitClaim) (*MsgSubmitClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitClaim not implemented")
}
func (*UnimplementedMsgServer) GovRemoveProtocolData(ctx context.Context, req *MsgGovRemoveProtocolData) (*MsgGovRemoveProtocolDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GovRemoveProtocolData not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.participationrewards.v1.Msg/SubmitClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitClaim(ctx, req.(*MsgSubmitClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GovRemoveProtocolData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGovRemoveProtocolData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GovRemoveProtocolData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.participationrewards.v1.Msg/GovRemoveProtocolData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GovRemoveProtocolData(ctx, req.(*MsgGovRemoveProtocolData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.participationrewards.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitClaim",
			Handler:    _Msg_SubmitClaim_Handler,
		},
		{
			MethodName: "GovRemoveProtocolData",
			Handler:    _Msg_GovRemoveProtocolData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/participationrewards/v1/messages.proto",
}

func (m *MsgSubmitClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proofs) > 0 {
		for iNdEx := len(m.Proofs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proofs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessages(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.ClaimType != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ClaimType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SrcZone) > 0 {
		i -= len(m.SrcZone)
		copy(dAtA[i:], m.SrcZone)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.SrcZone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Zone) > 0 {
		i -= len(m.Zone)
		copy(dAtA[i:], m.Zone)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Zone)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserAddress) > 0 {
		i -= len(m.UserAddress)
		copy(dAtA[i:], m.UserAddress)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.UserAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserAddress)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Zone)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.SrcZone)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.ClaimType != 0 {
		n += 1 + sovMessages(uint64(m.ClaimType))
	}
	if len(m.Proofs) > 0 {
		for _, e := range m.Proofs {
			l = e.Size()
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	return n
}

func (m *MsgSubmitClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			m.ClaimType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimType |= types.ClaimType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proofs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proofs = append(m.Proofs, &types.Proof{})
			if err := m.Proofs[len(m.Proofs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSubmitClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSubmitClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
