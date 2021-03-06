// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proxy/proxy.proto

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

type Proxy struct {
	Creator     string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index       string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	AddressList string `protobuf:"bytes,3,opt,name=addressList,proto3" json:"addressList,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae95c745fc9dd75, []int{0}
}
func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(m, src)
}
func (m *Proxy) XXX_Size() int {
	return m.Size()
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Proxy) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Proxy) GetAddressList() string {
	if m != nil {
		return m.AddressList
	}
	return ""
}

func (m *Proxy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Proxy)(nil), "pchainorg.pibridge.proxy.Proxy")
}

func init() { proto.RegisterFile("proxy/proxy.proto", fileDescriptor_fae95c745fc9dd75) }

var fileDescriptor_fae95c745fc9dd75 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x12, 0x05, 0xc9, 0x19, 0x89,
	0x99, 0x79, 0xf9, 0x45, 0xe9, 0x7a, 0x05, 0x99, 0x49, 0x45, 0x99, 0x29, 0xe9, 0xa9, 0x7a, 0x60,
	0x79, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x22, 0x7d, 0x10, 0x0b, 0xa2, 0x5e, 0xa9, 0x98,
	0x8b, 0x35, 0x00, 0x24, 0x2d, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4,
	0x56, 0x48, 0x30, 0x81, 0xc5, 0x21, 0x1c, 0x21, 0x05, 0x2e, 0xee, 0xc4, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0x62, 0x9f, 0xcc, 0xe2, 0x12, 0x09, 0x66, 0xb0, 0x1c, 0xb2, 0x10, 0xc8, 0x44, 0x28, 0x57,
	0x82, 0x05, 0x62, 0x22, 0x94, 0xeb, 0xe4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0x9f,
	0xe8, 0xe6, 0x17, 0xa5, 0xeb, 0x17, 0x64, 0xea, 0x42, 0xfc, 0xa2, 0x5f, 0x01, 0xf1, 0xad, 0x7e,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x13, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6d, 0xf6, 0x35, 0x86, 0x09, 0x01, 0x00, 0x00,
}

func (m *Proxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AddressList) > 0 {
		i -= len(m.AddressList)
		copy(dAtA[i:], m.AddressList)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.AddressList)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProxy(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	l = len(m.AddressList)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProxy(uint64(l))
	}
	return n
}

func sovProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProxy(x uint64) (n int) {
	return sovProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProxy
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
			return fmt.Errorf("proto: Proxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressList = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProxy
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
				return ErrInvalidLengthProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProxy
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
func skipProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProxy
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
					return 0, ErrIntOverflowProxy
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
					return 0, ErrIntOverflowProxy
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
				return 0, ErrInvalidLengthProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProxy = fmt.Errorf("proto: unexpected end of group")
)
