// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya/rollapp/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the rollapp module's genesis state.
type GenesisState struct {
	Params                             Params                           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	RollappList                        []Rollapp                        `protobuf:"bytes,2,rep,name=rollappList,proto3" json:"rollappList"`
	StateInfoList                      []StateInfo                      `protobuf:"bytes,3,rep,name=stateInfoList,proto3" json:"stateInfoList"`
	LatestStateInfoIndexList           []StateInfoIndex                 `protobuf:"bytes,4,rep,name=latestStateInfoIndexList,proto3" json:"latestStateInfoIndexList"`
	LatestFinalizedStateIndexList      []StateInfoIndex                 `protobuf:"bytes,5,rep,name=latestFinalizedStateIndexList,proto3" json:"latestFinalizedStateIndexList"`
	BlockHeightToFinalizationQueueList []BlockHeightToFinalizationQueue `protobuf:"bytes,6,rep,name=blockHeightToFinalizationQueueList,proto3" json:"blockHeightToFinalizationQueueList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4bf6d3c28914609, []int{0}
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

func (m *GenesisState) GetRollappList() []Rollapp {
	if m != nil {
		return m.RollappList
	}
	return nil
}

func (m *GenesisState) GetStateInfoList() []StateInfo {
	if m != nil {
		return m.StateInfoList
	}
	return nil
}

func (m *GenesisState) GetLatestStateInfoIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestStateInfoIndexList
	}
	return nil
}

func (m *GenesisState) GetLatestFinalizedStateIndexList() []StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndexList
	}
	return nil
}

func (m *GenesisState) GetBlockHeightToFinalizationQueueList() []BlockHeightToFinalizationQueue {
	if m != nil {
		return m.BlockHeightToFinalizationQueueList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "xblackfury.furya.rollapp.GenesisState")
}

func init() { proto.RegisterFile("furya/rollapp/genesis.proto", fileDescriptor_f4bf6d3c28914609) }

var fileDescriptor_f4bf6d3c28914609 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x3f, 0x4f, 0x3a, 0x31,
	0x18, 0xc7, 0xef, 0x7e, 0xfc, 0x19, 0xca, 0xcf, 0xa5, 0x71, 0x20, 0x24, 0x16, 0xc2, 0xa0, 0xb8,
	0xf4, 0x22, 0xee, 0x0e, 0xc4, 0xa8, 0x44, 0x13, 0x15, 0x74, 0x71, 0x31, 0x05, 0xca, 0xd1, 0x78,
	0xb4, 0x17, 0x5a, 0x12, 0x60, 0xf3, 0x1d, 0x38, 0xf8, 0xa2, 0x18, 0x19, 0x9d, 0x8c, 0x81, 0x37,
	0x62, 0xe8, 0x3d, 0x9c, 0x18, 0x14, 0x48, 0x9c, 0xda, 0xe6, 0xf9, 0x3e, 0x9f, 0xcf, 0x77, 0x28,
	0xca, 0xb7, 0x86, 0x5d, 0x2e, 0xb5, 0x50, 0xd2, 0xeb, 0xa9, 0x20, 0x60, 0x61, 0xe8, 0xf9, 0x5c,
	0x72, 0x2d, 0x34, 0x0d, 0x7b, 0xca, 0x28, 0x4c, 0xe2, 0xc0, 0x60, 0x38, 0xa2, 0xf1, 0x83, 0x42,
	0x3a, 0xb7, 0xeb, 0x2b, 0x5f, 0xd9, 0xa8, 0x37, 0xbf, 0x45, 0x5b, 0x39, 0xb2, 0x8a, 0x0d, 0x59,
	0x8f, 0x75, 0x81, 0x9a, 0xfb, 0x41, 0x0b, 0x27, 0x04, 0x8a, 0xab, 0x01, 0x6d, 0x98, 0xe1, 0x8f,
	0x42, 0xb6, 0x41, 0x52, 0x7c, 0x4e, 0xa1, 0xff, 0xe7, 0x51, 0xd9, 0xfa, 0x7c, 0x86, 0x4f, 0x51,
	0x3a, 0xb2, 0x64, 0xdd, 0x82, 0x5b, 0xca, 0x94, 0xf7, 0xe9, 0xfa, 0xf2, 0xf4, 0xc6, 0xa6, 0x2b,
	0xc9, 0xf1, 0x7b, 0xde, 0xa9, 0xc1, 0x2e, 0xbe, 0x46, 0x19, 0x98, 0x5f, 0x09, 0x6d, 0xb2, 0xff,
	0x0a, 0x89, 0x52, 0xa6, 0x7c, 0xb0, 0x09, 0x55, 0x8b, 0x4e, 0x60, 0x2d, 0x13, 0xf0, 0x3d, 0xda,
	0xb1, 0xdd, 0xab, 0xb2, 0xad, 0x2c, 0x32, 0x61, 0x91, 0x87, 0x9b, 0x90, 0xf5, 0xc5, 0x12, 0x40,
	0xbf, 0x53, 0x70, 0x88, 0xb2, 0x01, 0x33, 0x5c, 0x9b, 0x38, 0x57, 0x95, 0x2d, 0x3e, 0xb0, 0x86,
	0xa4, 0x35, 0xd0, 0xad, 0x0d, 0x76, 0x13, 0x34, 0xbf, 0x52, 0xf1, 0x08, 0xed, 0x45, 0xb3, 0x33,
	0x21, 0x59, 0x20, 0x46, 0xbc, 0x05, 0xa1, 0x85, 0x36, 0xf5, 0x07, 0xed, 0x7a, 0x34, 0x7e, 0x75,
	0x51, 0xb1, 0x11, 0xa8, 0xe6, 0xd3, 0x05, 0x17, 0x7e, 0xc7, 0xdc, 0x29, 0x08, 0x32, 0x23, 0x94,
	0xbc, 0xed, 0xf3, 0x3e, 0xb7, 0x0d, 0xd2, 0xb6, 0xc1, 0xc9, 0xa6, 0x06, 0x95, 0xb5, 0x24, 0x68,
	0xb4, 0x85, 0xaf, 0x72, 0x39, 0x9e, 0x12, 0x77, 0x32, 0x25, 0xee, 0xc7, 0x94, 0xb8, 0x2f, 0x33,
	0xe2, 0x4c, 0x66, 0xc4, 0x79, 0x9b, 0x11, 0xe7, 0xe1, 0xc8, 0x17, 0xa6, 0xd3, 0x6f, 0xd0, 0xa6,
	0xea, 0x7a, 0xcb, 0x6d, 0xbe, 0x1e, 0xde, 0x20, 0xfe, 0xdb, 0x66, 0x18, 0x72, 0xdd, 0x48, 0xdb,
	0x7f, 0x7d, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x3d, 0xbc, 0x58, 0x95, 0x03, 0x00, 0x00,
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
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for iNdEx := len(m.BlockHeightToFinalizationQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockHeightToFinalizationQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for iNdEx := len(m.LatestFinalizedStateIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestFinalizedStateIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for iNdEx := len(m.LatestStateInfoIndexList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestStateInfoIndexList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.StateInfoList) > 0 {
		for iNdEx := len(m.StateInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RollappList) > 0 {
		for iNdEx := len(m.RollappList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RollappList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RollappList) > 0 {
		for _, e := range m.RollappList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StateInfoList) > 0 {
		for _, e := range m.StateInfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestStateInfoIndexList) > 0 {
		for _, e := range m.LatestStateInfoIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LatestFinalizedStateIndexList) > 0 {
		for _, e := range m.LatestFinalizedStateIndexList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BlockHeightToFinalizationQueueList) > 0 {
		for _, e := range m.BlockHeightToFinalizationQueueList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field RollappList", wireType)
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
			m.RollappList = append(m.RollappList, Rollapp{})
			if err := m.RollappList[len(m.RollappList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateInfoList", wireType)
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
			m.StateInfoList = append(m.StateInfoList, StateInfo{})
			if err := m.StateInfoList[len(m.StateInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateInfoIndexList", wireType)
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
			m.LatestStateInfoIndexList = append(m.LatestStateInfoIndexList, StateInfoIndex{})
			if err := m.LatestStateInfoIndexList[len(m.LatestStateInfoIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndexList", wireType)
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
			m.LatestFinalizedStateIndexList = append(m.LatestFinalizedStateIndexList, StateInfoIndex{})
			if err := m.LatestFinalizedStateIndexList[len(m.LatestFinalizedStateIndexList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeightToFinalizationQueueList", wireType)
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
			m.BlockHeightToFinalizationQueueList = append(m.BlockHeightToFinalizationQueueList, BlockHeightToFinalizationQueue{})
			if err := m.BlockHeightToFinalizationQueueList[len(m.BlockHeightToFinalizationQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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