// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/circuit/v1/tx.proto

package types

import (
	fmt "fmt"
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

// MsgAuthorizeCircuitBreaker defines the Msg/AuthorizeCircuitBreaker request type.
type MsgAuthorizeCircuitBreaker struct {
	// granter is the granter of the circuit breaker permissions and must have
	// LEVEL_SUPER_ADMIN.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the account authorized with the provided permissions.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// permissions are the circuit breaker permissions that the grantee receives.
	// These will overwrite any existing permissions. LEVEL_NONE_UNSPECIFIED can
	// be specified to revoke all permissions.
	Permissions *Permissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *MsgAuthorizeCircuitBreaker) Reset()         { *m = MsgAuthorizeCircuitBreaker{} }
func (m *MsgAuthorizeCircuitBreaker) String() string { return proto.CompactTextString(m) }
func (*MsgAuthorizeCircuitBreaker) ProtoMessage()    {}
func (*MsgAuthorizeCircuitBreaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{0}
}
func (m *MsgAuthorizeCircuitBreaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthorizeCircuitBreaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthorizeCircuitBreaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthorizeCircuitBreaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthorizeCircuitBreaker.Merge(m, src)
}
func (m *MsgAuthorizeCircuitBreaker) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthorizeCircuitBreaker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthorizeCircuitBreaker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthorizeCircuitBreaker proto.InternalMessageInfo

func (m *MsgAuthorizeCircuitBreaker) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgAuthorizeCircuitBreaker) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *MsgAuthorizeCircuitBreaker) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// MsgAuthorizeCircuitBreakerResponse defines the Msg/AuthorizeCircuitBreaker response type.
type MsgAuthorizeCircuitBreakerResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgAuthorizeCircuitBreakerResponse) Reset()         { *m = MsgAuthorizeCircuitBreakerResponse{} }
func (m *MsgAuthorizeCircuitBreakerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthorizeCircuitBreakerResponse) ProtoMessage()    {}
func (*MsgAuthorizeCircuitBreakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{1}
}
func (m *MsgAuthorizeCircuitBreakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthorizeCircuitBreakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthorizeCircuitBreakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthorizeCircuitBreakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthorizeCircuitBreakerResponse.Merge(m, src)
}
func (m *MsgAuthorizeCircuitBreakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthorizeCircuitBreakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthorizeCircuitBreakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthorizeCircuitBreakerResponse proto.InternalMessageInfo

func (m *MsgAuthorizeCircuitBreakerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// MsgTripCircuitBreaker defines the Msg/TripCircuitBreaker request type.
type MsgTripCircuitBreaker struct {
	// authority is the account authorized to trip the circuit breaker.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// msg_type_urls specifies a list of type URLs to immediately stop processing.
	// IF IT IS LEFT EMPTY, ALL MSG PROCESSING WILL STOP IMMEDIATELY.
	// This value is validated against the authority's permissions and if the
	// authority does not have permissions to trip the specified msg type URLs
	// (or all URLs), the operation will fail.
	MsgTypeUrls []string `protobuf:"bytes,2,rep,name=msg_type_urls,json=msgTypeUrls,proto3" json:"msg_type_urls,omitempty"`
}

func (m *MsgTripCircuitBreaker) Reset()         { *m = MsgTripCircuitBreaker{} }
func (m *MsgTripCircuitBreaker) String() string { return proto.CompactTextString(m) }
func (*MsgTripCircuitBreaker) ProtoMessage()    {}
func (*MsgTripCircuitBreaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{2}
}
func (m *MsgTripCircuitBreaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTripCircuitBreaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTripCircuitBreaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTripCircuitBreaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTripCircuitBreaker.Merge(m, src)
}
func (m *MsgTripCircuitBreaker) XXX_Size() int {
	return m.Size()
}
func (m *MsgTripCircuitBreaker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTripCircuitBreaker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTripCircuitBreaker proto.InternalMessageInfo

func (m *MsgTripCircuitBreaker) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgTripCircuitBreaker) GetMsgTypeUrls() []string {
	if m != nil {
		return m.MsgTypeUrls
	}
	return nil
}

// MsgTripCircuitBreakerResponse defines the Msg/TripCircuitBreaker response type.
type MsgTripCircuitBreakerResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgTripCircuitBreakerResponse) Reset()         { *m = MsgTripCircuitBreakerResponse{} }
func (m *MsgTripCircuitBreakerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTripCircuitBreakerResponse) ProtoMessage()    {}
func (*MsgTripCircuitBreakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{3}
}
func (m *MsgTripCircuitBreakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTripCircuitBreakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTripCircuitBreakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTripCircuitBreakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTripCircuitBreakerResponse.Merge(m, src)
}
func (m *MsgTripCircuitBreakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTripCircuitBreakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTripCircuitBreakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTripCircuitBreakerResponse proto.InternalMessageInfo

func (m *MsgTripCircuitBreakerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// MsgResetCircuitBreaker defines the Msg/ResetCircuitBreaker request type.
type MsgResetCircuitBreaker struct {
	// authority is the account authorized to trip or reset the circuit breaker.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// msg_type_urls specifies a list of Msg type URLs to resume processing. If
	// it is left empty all Msg processing for type URLs that the account is
	// authorized to trip will resume.
	MsgTypeUrls []string `protobuf:"bytes,3,rep,name=msg_type_urls,json=msgTypeUrls,proto3" json:"msg_type_urls,omitempty"`
}

func (m *MsgResetCircuitBreaker) Reset()         { *m = MsgResetCircuitBreaker{} }
func (m *MsgResetCircuitBreaker) String() string { return proto.CompactTextString(m) }
func (*MsgResetCircuitBreaker) ProtoMessage()    {}
func (*MsgResetCircuitBreaker) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{4}
}
func (m *MsgResetCircuitBreaker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgResetCircuitBreaker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgResetCircuitBreaker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgResetCircuitBreaker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgResetCircuitBreaker.Merge(m, src)
}
func (m *MsgResetCircuitBreaker) XXX_Size() int {
	return m.Size()
}
func (m *MsgResetCircuitBreaker) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgResetCircuitBreaker.DiscardUnknown(m)
}

var xxx_messageInfo_MsgResetCircuitBreaker proto.InternalMessageInfo

func (m *MsgResetCircuitBreaker) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgResetCircuitBreaker) GetMsgTypeUrls() []string {
	if m != nil {
		return m.MsgTypeUrls
	}
	return nil
}

// MsgResetCircuitBreakerResponse defines the Msg/ResetCircuitBreaker response type.
type MsgResetCircuitBreakerResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgResetCircuitBreakerResponse) Reset()         { *m = MsgResetCircuitBreakerResponse{} }
func (m *MsgResetCircuitBreakerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgResetCircuitBreakerResponse) ProtoMessage()    {}
func (*MsgResetCircuitBreakerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02145e57a6fbb1d, []int{5}
}
func (m *MsgResetCircuitBreakerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgResetCircuitBreakerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgResetCircuitBreakerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgResetCircuitBreakerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgResetCircuitBreakerResponse.Merge(m, src)
}
func (m *MsgResetCircuitBreakerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgResetCircuitBreakerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgResetCircuitBreakerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgResetCircuitBreakerResponse proto.InternalMessageInfo

func (m *MsgResetCircuitBreakerResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*MsgAuthorizeCircuitBreaker)(nil), "cosmos.circuit.v1.MsgAuthorizeCircuitBreaker")
	proto.RegisterType((*MsgAuthorizeCircuitBreakerResponse)(nil), "cosmos.circuit.v1.MsgAuthorizeCircuitBreakerResponse")
	proto.RegisterType((*MsgTripCircuitBreaker)(nil), "cosmos.circuit.v1.MsgTripCircuitBreaker")
	proto.RegisterType((*MsgTripCircuitBreakerResponse)(nil), "cosmos.circuit.v1.MsgTripCircuitBreakerResponse")
	proto.RegisterType((*MsgResetCircuitBreaker)(nil), "cosmos.circuit.v1.MsgResetCircuitBreaker")
	proto.RegisterType((*MsgResetCircuitBreakerResponse)(nil), "cosmos.circuit.v1.MsgResetCircuitBreakerResponse")
}

func init() { proto.RegisterFile("cosmos/circuit/v1/tx.proto", fileDescriptor_a02145e57a6fbb1d) }

var fileDescriptor_a02145e57a6fbb1d = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x59, 0x48, 0x54, 0x16, 0x35, 0xb1, 0x89, 0xda, 0x34, 0xb0, 0x21, 0x3d, 0x11, 0x0f,
	0x6d, 0xd0, 0xc4, 0x44, 0x0e, 0x46, 0xf1, 0x4c, 0x62, 0x1a, 0xbc, 0x78, 0x21, 0x58, 0x27, 0xeb,
	0x0a, 0x65, 0x9b, 0x9d, 0x2d, 0x01, 0x8f, 0x3e, 0x81, 0x8f, 0xe0, 0x23, 0xf8, 0x18, 0x1e, 0x39,
	0x7a, 0x34, 0x70, 0xf0, 0x35, 0x4c, 0x29, 0x58, 0xa2, 0x18, 0x4c, 0x3c, 0x4e, 0xff, 0x7f, 0x66,
	0xbe, 0xe9, 0xfe, 0xd4, 0xf2, 0x25, 0x06, 0x12, 0x5d, 0x5f, 0x28, 0x3f, 0x12, 0xda, 0xed, 0x57,
	0x5d, 0x3d, 0x70, 0x42, 0x25, 0xb5, 0x34, 0x76, 0x12, 0xcd, 0x99, 0x69, 0x4e, 0xbf, 0x6a, 0xed,
	0xcf, 0xec, 0x01, 0xf2, 0xd8, 0x1a, 0x20, 0x4f, 0xbc, 0x56, 0x69, 0xc9, 0x9c, 0x61, 0x08, 0x98,
	0xc8, 0xf6, 0x33, 0xa1, 0x56, 0x03, 0xf9, 0x79, 0xa4, 0xef, 0xa4, 0x12, 0x0f, 0x70, 0x91, 0xd8,
	0xea, 0x0a, 0xda, 0x1d, 0x50, 0x86, 0x49, 0xd7, 0xb9, 0x6a, 0xf7, 0x34, 0x28, 0x93, 0x94, 0x49,
	0x25, 0xef, 0xcd, 0xcb, 0x54, 0x01, 0x33, 0xbb, 0xa8, 0x80, 0x71, 0x46, 0x0b, 0x21, 0xa8, 0x40,
	0x20, 0x0a, 0xd9, 0x43, 0x33, 0x57, 0x26, 0x95, 0xc2, 0x21, 0x73, 0x7e, 0x30, 0x3b, 0x97, 0xa9,
	0xcb, 0x5b, 0x6c, 0xa9, 0x6d, 0x3e, 0x7e, 0xbc, 0x1c, 0xcc, 0x37, 0xd9, 0xa7, 0xd4, 0xfe, 0x9d,
	0xd0, 0x03, 0x0c, 0x65, 0x0f, 0x21, 0xe6, 0xc1, 0xc8, 0xf7, 0x01, 0x71, 0x4a, 0xba, 0xe1, 0xcd,
	0x4b, 0x5b, 0xd0, 0xdd, 0x06, 0xf2, 0xa6, 0x12, 0xe1, 0xb7, 0xe3, 0x8a, 0x34, 0xdf, 0x4e, 0xa6,
	0xea, 0xe1, 0xec, 0xbc, 0xf4, 0x83, 0x61, 0xd3, 0xad, 0x00, 0x79, 0x2b, 0xfe, 0x59, 0xad, 0x48,
	0x75, 0xd1, 0xcc, 0x96, 0x73, 0x95, 0xbc, 0x57, 0x08, 0x90, 0x37, 0x87, 0x21, 0x5c, 0xa9, 0x2e,
	0xd6, 0xb6, 0x63, 0xd0, 0xb4, 0xc7, 0x3e, 0xa1, 0xa5, 0xa5, 0xab, 0xfe, 0x40, 0x79, 0x4f, 0xf7,
	0x1a, 0xc8, 0x3d, 0x40, 0xd0, 0xff, 0xc3, 0xcc, 0xad, 0xc6, 0xac, 0x51, 0xb6, 0x7c, 0xd7, 0x6a,
	0xce, 0xfa, 0xf1, 0xeb, 0x98, 0x91, 0xd1, 0x98, 0x91, 0xf7, 0x31, 0x23, 0x4f, 0x13, 0x96, 0x19,
	0x4d, 0x58, 0xe6, 0x6d, 0xc2, 0x32, 0xd7, 0xc5, 0xe4, 0x85, 0xf1, 0xb6, 0xe3, 0x08, 0xe9, 0x0e,
	0xbe, 0x12, 0x37, 0x8d, 0xdb, 0xcd, 0xda, 0x34, 0x6f, 0x47, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x13, 0x37, 0xd5, 0xd8, 0x02, 0x00, 0x00,
}

func (m *MsgAuthorizeCircuitBreaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthorizeCircuitBreaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthorizeCircuitBreaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Permissions != nil {
		{
			size, err := m.Permissions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthorizeCircuitBreakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthorizeCircuitBreakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthorizeCircuitBreakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgTripCircuitBreaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTripCircuitBreaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTripCircuitBreaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrls) > 0 {
		for iNdEx := len(m.MsgTypeUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MsgTypeUrls[iNdEx])
			copy(dAtA[i:], m.MsgTypeUrls[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.MsgTypeUrls[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTripCircuitBreakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTripCircuitBreakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTripCircuitBreakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgResetCircuitBreaker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgResetCircuitBreaker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgResetCircuitBreaker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrls) > 0 {
		for iNdEx := len(m.MsgTypeUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.MsgTypeUrls[iNdEx])
			copy(dAtA[i:], m.MsgTypeUrls[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.MsgTypeUrls[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgResetCircuitBreakerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgResetCircuitBreakerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgResetCircuitBreakerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthorizeCircuitBreaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Permissions != nil {
		l = m.Permissions.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAuthorizeCircuitBreakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *MsgTripCircuitBreaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.MsgTypeUrls) > 0 {
		for _, s := range m.MsgTypeUrls {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgTripCircuitBreakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *MsgResetCircuitBreaker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.MsgTypeUrls) > 0 {
		for _, s := range m.MsgTypeUrls {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgResetCircuitBreakerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthorizeCircuitBreaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAuthorizeCircuitBreaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthorizeCircuitBreaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Permissions == nil {
				m.Permissions = &Permissions{}
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAuthorizeCircuitBreakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAuthorizeCircuitBreakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthorizeCircuitBreakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTripCircuitBreaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTripCircuitBreaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTripCircuitBreaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrls = append(m.MsgTypeUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgTripCircuitBreakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgTripCircuitBreakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTripCircuitBreakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgResetCircuitBreaker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgResetCircuitBreaker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgResetCircuitBreaker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrls = append(m.MsgTypeUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgResetCircuitBreakerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgResetCircuitBreakerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgResetCircuitBreakerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)