// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room.proto

/*
	Package gml is a generated protocol buffer package.

	It is generated from these files:
		room.proto
		room_object.proto

	It has these top-level messages:
		Room
		RoomObject
*/
package gml

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Room struct {
	Left      int32         `protobuf:"varint,1,opt,name=Left,proto3" json:"Left,omitempty"`
	Right     int32         `protobuf:"varint,2,opt,name=Right,proto3" json:"Right,omitempty"`
	Top       int32         `protobuf:"varint,3,opt,name=Top,proto3" json:"Top,omitempty"`
	Bottom    int32         `protobuf:"varint,4,opt,name=Bottom,proto3" json:"Bottom,omitempty"`
	Instances []*RoomObject `protobuf:"bytes,5,rep,name=instances" json:"instances,omitempty"`
}

func (m *Room) Reset()                    { *m = Room{} }
func (m *Room) String() string            { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()               {}
func (*Room) Descriptor() ([]byte, []int) { return fileDescriptorRoom, []int{0} }

func (m *Room) GetLeft() int32 {
	if m != nil {
		return m.Left
	}
	return 0
}

func (m *Room) GetRight() int32 {
	if m != nil {
		return m.Right
	}
	return 0
}

func (m *Room) GetTop() int32 {
	if m != nil {
		return m.Top
	}
	return 0
}

func (m *Room) GetBottom() int32 {
	if m != nil {
		return m.Bottom
	}
	return 0
}

func (m *Room) GetInstances() []*RoomObject {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
	proto.RegisterType((*Room)(nil), "gml.Room")
}
func (m *Room) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Room) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Left != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoom(dAtA, i, uint64(m.Left))
	}
	if m.Right != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoom(dAtA, i, uint64(m.Right))
	}
	if m.Top != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoom(dAtA, i, uint64(m.Top))
	}
	if m.Bottom != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRoom(dAtA, i, uint64(m.Bottom))
	}
	if len(m.Instances) > 0 {
		for _, msg := range m.Instances {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintRoom(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintRoom(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Room) Size() (n int) {
	var l int
	_ = l
	if m.Left != 0 {
		n += 1 + sovRoom(uint64(m.Left))
	}
	if m.Right != 0 {
		n += 1 + sovRoom(uint64(m.Right))
	}
	if m.Top != 0 {
		n += 1 + sovRoom(uint64(m.Top))
	}
	if m.Bottom != 0 {
		n += 1 + sovRoom(uint64(m.Bottom))
	}
	if len(m.Instances) > 0 {
		for _, e := range m.Instances {
			l = e.Size()
			n += 1 + l + sovRoom(uint64(l))
		}
	}
	return n
}

func sovRoom(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoom(x uint64) (n int) {
	return sovRoom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Room) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoom
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Room: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Room: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			m.Left = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Left |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			m.Right = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Right |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Top", wireType)
			}
			m.Top = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Top |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bottom", wireType)
			}
			m.Bottom = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bottom |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoom
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Instances = append(m.Instances, &RoomObject{})
			if err := m.Instances[len(m.Instances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoom
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
func skipRoom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoom
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
					return 0, ErrIntOverflowRoom
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoom
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRoom
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoom
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRoom(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRoom = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoom   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("room.proto", fileDescriptorRoom) }

var fileDescriptorRoom = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0xcf,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcf, 0xcd, 0x91, 0x12, 0x04, 0x09, 0xc4,
	0xe7, 0x27, 0x65, 0xa5, 0x26, 0x97, 0x40, 0xc4, 0x95, 0x3a, 0x19, 0xb9, 0x58, 0x82, 0xf2, 0xf3,
	0x73, 0x85, 0x84, 0xb8, 0x58, 0x7c, 0x52, 0xd3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xa0, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x26, 0xb0, 0x20, 0x84,
	0x23, 0x24, 0xc0, 0xc5, 0x1c, 0x92, 0x5f, 0x20, 0xc1, 0x0c, 0x16, 0x03, 0x31, 0x85, 0xc4, 0xb8,
	0xd8, 0x9c, 0xf2, 0x4b, 0x4a, 0xf2, 0x73, 0x25, 0x58, 0xc0, 0x82, 0x50, 0x9e, 0x90, 0x2e, 0x17,
	0x67, 0x66, 0x5e, 0x71, 0x49, 0x62, 0x5e, 0x72, 0x6a, 0xb1, 0x04, 0xab, 0x02, 0xb3, 0x06, 0xb7,
	0x11, 0xbf, 0x5e, 0x7a, 0x6e, 0x8e, 0x1e, 0xc8, 0x46, 0x7f, 0xb0, 0x33, 0x82, 0x10, 0x2a, 0x9c,
	0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f,
	0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x8e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xc0, 0xeb,
	0xcc, 0xca, 0x00, 0x00, 0x00,
}
