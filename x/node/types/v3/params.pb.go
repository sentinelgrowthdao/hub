// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v3/params.proto

package v3

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type Params struct {
	Deposit             types.Coin                               `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	ActiveDuration      time.Duration                            `protobuf:"bytes,2,opt,name=active_duration,json=activeDuration,proto3,stdduration" json:"active_duration"`
	MinGigabytePrices   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=min_gigabyte_prices,json=minGigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_gigabyte_prices"`
	MinHourlyPrices     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_hourly_prices,json=minHourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_hourly_prices"`
	MaxSessionGigabytes int64                                    `protobuf:"varint,5,opt,name=max_session_gigabytes,json=maxSessionGigabytes,proto3" json:"max_session_gigabytes,omitempty"`
	MinSessionGigabytes int64                                    `protobuf:"varint,6,opt,name=min_session_gigabytes,json=minSessionGigabytes,proto3" json:"min_session_gigabytes,omitempty"`
	MaxSessionHours     int64                                    `protobuf:"varint,7,opt,name=max_session_hours,json=maxSessionHours,proto3" json:"max_session_hours,omitempty"`
	MinSessionHours     int64                                    `protobuf:"varint,8,opt,name=min_session_hours,json=minSessionHours,proto3" json:"min_session_hours,omitempty"`
	StakingShare        cosmossdk_io_math.LegacyDec              `protobuf:"bytes,9,opt,name=staking_share,json=stakingShare,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"staking_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc24098d3255287, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.node.v3.Params")
}

func init() { proto.RegisterFile("sentinel/node/v3/params.proto", fileDescriptor_bbc24098d3255287) }

var fileDescriptor_bbc24098d3255287 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x13, 0xba, 0x75, 0x5b, 0xf8, 0x53, 0x96, 0x81, 0x14, 0x86, 0x70, 0x2b, 0xb8, 0x54,
	0x48, 0xb3, 0x69, 0xcb, 0x85, 0x6b, 0x99, 0xc4, 0x0e, 0x3b, 0x4c, 0x99, 0xc4, 0x81, 0x4b, 0xe4,
	0x24, 0x6e, 0x62, 0xb5, 0xb1, 0xa3, 0xd8, 0x09, 0xad, 0xf8, 0x12, 0x70, 0xe3, 0x23, 0x20, 0x3e,
	0x49, 0x8f, 0x3b, 0x22, 0x0e, 0x1b, 0xb4, 0x5f, 0x04, 0x39, 0x76, 0x68, 0x25, 0x10, 0x27, 0x4e,
	0x71, 0xfc, 0x3e, 0xcf, 0xf3, 0x7b, 0xfd, 0xca, 0x76, 0x9e, 0x08, 0xc2, 0x24, 0x65, 0x64, 0x86,
	0x18, 0x8f, 0x09, 0xaa, 0x46, 0x28, 0xc7, 0x05, 0xce, 0x04, 0xcc, 0x0b, 0x2e, 0xb9, 0x7b, 0xbf,
	0x29, 0x43, 0x55, 0x86, 0xd5, 0xe8, 0x18, 0x44, 0x5c, 0x64, 0x5c, 0xa0, 0x10, 0x0b, 0x82, 0xaa,
	0x41, 0x48, 0x24, 0x1e, 0xa0, 0x88, 0x53, 0xa6, 0x1d, 0xc7, 0x0f, 0x12, 0x9e, 0xf0, 0x7a, 0x89,
	0xd4, 0xca, 0xec, 0x82, 0x84, 0xf3, 0x64, 0x46, 0x50, 0xfd, 0x17, 0x96, 0x13, 0x14, 0x97, 0x05,
	0x96, 0x94, 0x1b, 0xd7, 0xd3, 0x4f, 0xbb, 0x4e, 0xfb, 0xa2, 0x06, 0xbb, 0xaf, 0x9c, 0xbd, 0x98,
	0xe4, 0x5c, 0x50, 0xe9, 0xd9, 0x3d, 0xbb, 0x7f, 0x7b, 0xf8, 0x08, 0x6a, 0x24, 0x54, 0x48, 0x68,
	0x90, 0xf0, 0x35, 0xa7, 0x6c, 0xbc, 0xb3, 0xbc, 0xee, 0x5a, 0x7e, 0xa3, 0x77, 0xcf, 0x9d, 0x0e,
	0x8e, 0x24, 0xad, 0x48, 0xd0, 0xc4, 0x7b, 0xb7, 0x4c, 0x84, 0xe6, 0xc3, 0x86, 0x0f, 0x4f, 0x8d,
	0x60, 0xbc, 0xaf, 0x22, 0x3e, 0xdf, 0x74, 0x6d, 0xff, 0x9e, 0xf6, 0x36, 0x15, 0xf7, 0x83, 0x73,
	0x94, 0x51, 0x16, 0x24, 0x34, 0xc1, 0xe1, 0x42, 0x92, 0x20, 0x2f, 0x68, 0x44, 0x84, 0xd7, 0xea,
	0xb5, 0xfe, 0xdd, 0xd4, 0x0b, 0x95, 0xf8, 0xf5, 0xa6, 0xdb, 0x4f, 0xa8, 0x4c, 0xcb, 0x10, 0x46,
	0x3c, 0x43, 0x66, 0x68, 0xfa, 0x73, 0x22, 0xe2, 0x29, 0x92, 0x8b, 0x9c, 0x88, 0xda, 0x20, 0xfc,
	0xc3, 0x8c, 0xb2, 0x37, 0x06, 0x73, 0x51, 0x53, 0xdc, 0xf7, 0x8e, 0xda, 0x0c, 0x52, 0x5e, 0x16,
	0xb3, 0x45, 0x83, 0xde, 0xf9, 0xff, 0xe8, 0x4e, 0x46, 0xd9, 0x59, 0x0d, 0x31, 0xe0, 0xa1, 0xf3,
	0x30, 0xc3, 0xf3, 0x40, 0x10, 0x21, 0x28, 0xdf, 0x9c, 0x5e, 0x78, 0xbb, 0x3d, 0xbb, 0xdf, 0xf2,
	0x8f, 0x32, 0x3c, 0xbf, 0xd4, 0xb5, 0xa6, 0x63, 0xed, 0xa1, 0xec, 0x2f, 0x9e, 0xb6, 0xf1, 0x50,
	0xf6, 0x87, 0xe7, 0xb9, 0x73, 0xb8, 0xcd, 0x51, 0x07, 0x15, 0xde, 0x5e, 0xad, 0xef, 0x6c, 0x18,
	0xaa, 0x35, 0xad, 0xdd, 0xca, 0xd7, 0xda, 0x7d, 0xa3, 0xfd, 0x9d, 0xad, 0xb5, 0x67, 0xce, 0x5d,
	0x21, 0xf1, 0x94, 0xb2, 0x24, 0x10, 0x29, 0x2e, 0x88, 0x77, 0xd0, 0xb3, 0xfb, 0x07, 0xe3, 0x67,
	0x6a, 0x32, 0xdf, 0xaf, 0xbb, 0x8f, 0xf5, 0x1c, 0x44, 0x3c, 0x85, 0x94, 0xa3, 0x0c, 0xcb, 0x14,
	0x9e, 0x93, 0x04, 0x47, 0x8b, 0x53, 0x12, 0xf9, 0x77, 0x8c, 0xf3, 0x52, 0x19, 0xc7, 0x6f, 0x97,
	0x3f, 0x81, 0xf5, 0x65, 0x05, 0xac, 0xe5, 0x0a, 0xd8, 0x57, 0x2b, 0x60, 0xff, 0x58, 0x01, 0xfb,
	0xe3, 0x1a, 0x58, 0x57, 0x6b, 0x60, 0x7d, 0x5b, 0x03, 0xeb, 0xdd, 0xcb, 0xad, 0x31, 0x37, 0x0f,
	0xe5, 0x84, 0x4f, 0x26, 0x34, 0xa2, 0x78, 0x86, 0xd2, 0x32, 0x44, 0xd5, 0x60, 0x88, 0xe6, 0xfa,
	0x69, 0xd5, 0x53, 0x47, 0xd5, 0x28, 0x6c, 0xd7, 0x97, 0x70, 0xf4, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0xee, 0xf2, 0x72, 0x7b, 0x03, 0x00, 0x00,
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
	{
		size := m.StakingShare.Size()
		i -= size
		if _, err := m.StakingShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.MinSessionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSessionHours))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxSessionHours != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSessionHours))
		i--
		dAtA[i] = 0x38
	}
	if m.MinSessionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSessionGigabytes))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxSessionGigabytes != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSessionGigabytes))
		i--
		dAtA[i] = 0x28
	}
	if len(m.MinHourlyPrices) > 0 {
		for iNdEx := len(m.MinHourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinHourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MinGigabytePrices) > 0 {
		for iNdEx := len(m.MinGigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinGigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ActiveDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.MinGigabytePrices) > 0 {
		for _, e := range m.MinGigabytePrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinHourlyPrices) > 0 {
		for _, e := range m.MinHourlyPrices {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxSessionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MaxSessionGigabytes))
	}
	if m.MinSessionGigabytes != 0 {
		n += 1 + sovParams(uint64(m.MinSessionGigabytes))
	}
	if m.MaxSessionHours != 0 {
		n += 1 + sovParams(uint64(m.MaxSessionHours))
	}
	if m.MinSessionHours != 0 {
		n += 1 + sovParams(uint64(m.MinSessionHours))
	}
	l = m.StakingShare.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinGigabytePrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinGigabytePrices = append(m.MinGigabytePrices, types.Coin{})
			if err := m.MinGigabytePrices[len(m.MinGigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinHourlyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinHourlyPrices = append(m.MinHourlyPrices, types.Coin{})
			if err := m.MinHourlyPrices[len(m.MinHourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSessionGigabytes", wireType)
			}
			m.MaxSessionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSessionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSessionGigabytes", wireType)
			}
			m.MinSessionGigabytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSessionGigabytes |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSessionHours", wireType)
			}
			m.MaxSessionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSessionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSessionHours", wireType)
			}
			m.MinSessionHours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSessionHours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
