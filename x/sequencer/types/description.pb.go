// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: furya/sequencer/description.proto

package types

import (
	fmt "fmt"
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

// Description defines a sequencer description.
type Description struct {
	// moniker defines a human-readable name for the sequencer.
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// identity defines an optional identity signature (ex. UPort or Keybase).
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	// website defines an optional website link.
	Website string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	// securityContact defines an optional email for security contact.
	SecurityContact string `protobuf:"bytes,4,opt,name=securityContact,proto3" json:"securityContact,omitempty"`
	// details define other optional details.
	Details string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *Description) Reset()         { *m = Description{} }
func (m *Description) String() string { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()    {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_91de4c32465eb7e7, []int{0}
}
func (m *Description) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Description.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(m, src)
}
func (m *Description) XXX_Size() int {
	return m.Size()
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

func (m *Description) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Description) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Description) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Description) GetSecurityContact() string {
	if m != nil {
		return m.SecurityContact
	}
	return ""
}

func (m *Description) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*Description)(nil), "furyaxyz.furya.sequencer.Description")
}

func init() {
	proto.RegisterFile("furya/sequencer/description.proto", fileDescriptor_91de4c32465eb7e7)
}

var fileDescriptor_91de4c32465eb7e7 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4e, 0x03, 0x31,
	0x10, 0x85, 0xd7, 0xfc, 0x63, 0x0a, 0x24, 0x57, 0x16, 0x85, 0x15, 0x21, 0x21, 0xa5, 0xb2, 0x8b,
	0xdc, 0x00, 0x68, 0x69, 0x28, 0xe9, 0xb2, 0xf6, 0x08, 0x46, 0xb0, 0xf6, 0x62, 0xcf, 0x8a, 0x98,
	0x53, 0x70, 0x05, 0x6e, 0x43, 0x99, 0x92, 0x12, 0xed, 0x5e, 0x04, 0x65, 0x13, 0x9c, 0x88, 0xf2,
	0xb3, 0xdf, 0xf7, 0xa4, 0x79, 0xfc, 0xca, 0xe5, 0x06, 0x7c, 0xc2, 0xe0, 0x4d, 0x82, 0xd7, 0x0e,
	0xbc, 0x85, 0x68, 0x1c, 0x24, 0x1b, 0xb1, 0x25, 0x0c, 0x5e, 0xb7, 0x31, 0x50, 0x10, 0x93, 0x12,
	0x5b, 0xe4, 0x77, 0x5d, 0x40, 0x17, 0xe7, 0xf2, 0x93, 0xf1, 0xb3, 0xdb, 0xad, 0x27, 0x24, 0x3f,
	0x6e, 0x82, 0xc7, 0x67, 0x88, 0x92, 0x4d, 0xd8, 0xf4, 0xf4, 0xfe, 0x0f, 0xc5, 0x05, 0x3f, 0x41,
	0x07, 0x9e, 0x90, 0xb2, 0xdc, 0x1b, 0xbf, 0x0a, 0xaf, 0xac, 0x37, 0xa8, 0x13, 0x12, 0xc8, 0xfd,
	0xb5, 0xb5, 0x41, 0x31, 0xe5, 0xe7, 0x09, 0x6c, 0x17, 0x91, 0xf2, 0x4d, 0xf0, 0x34, 0xb7, 0x24,
	0x0f, 0xc6, 0xc4, 0xff, 0xe7, 0x55, 0x87, 0x03, 0x9a, 0xe3, 0x4b, 0x92, 0x87, 0xeb, 0x8e, 0x0d,
	0x5e, 0xdf, 0x7d, 0xf5, 0x8a, 0x2d, 0x7b, 0xc5, 0x7e, 0x7a, 0xc5, 0x3e, 0x06, 0x55, 0x2d, 0x07,
	0x55, 0x7d, 0x0f, 0xaa, 0x7a, 0x98, 0x3d, 0x22, 0x3d, 0x75, 0xb5, 0xb6, 0xa1, 0x31, 0xbb, 0xa7,
	0x6e, 0xc1, 0x2c, 0x76, 0x06, 0xa2, 0xdc, 0x42, 0xaa, 0x8f, 0xc6, 0x6d, 0x66, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xad, 0x5c, 0xd5, 0x74, 0x44, 0x01, 0x00, 0x00,
}

func (m *Description) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Description) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Description) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintDescription(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintDescription(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintDescription(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintDescription(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintDescription(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDescription(dAtA []byte, offset int, v uint64) int {
	offset -= sovDescription(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Description) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovDescription(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovDescription(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovDescription(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovDescription(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovDescription(uint64(l))
	}
	return n
}

func sovDescription(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDescription(x uint64) (n int) {
	return sovDescription(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Description) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescription
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
			return fmt.Errorf("proto: Description: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Description: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescription
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
				return ErrInvalidLengthDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescription
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
				return ErrInvalidLengthDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescription
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
				return ErrInvalidLengthDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescription
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
				return ErrInvalidLengthDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescription
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
				return ErrInvalidLengthDescription
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDescription
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDescription(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDescription
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
func skipDescription(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDescription
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
					return 0, ErrIntOverflowDescription
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
					return 0, ErrIntOverflowDescription
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
				return 0, ErrInvalidLengthDescription
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDescription
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDescription
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDescription        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDescription          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDescription = fmt.Errorf("proto: unexpected end of group")
)
