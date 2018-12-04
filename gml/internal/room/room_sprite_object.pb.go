// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room_sprite_object.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RoomSpriteObject struct {
	UUID       string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	X          int32  `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y          int32  `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	Width      int32  `protobuf:"varint,4,opt,name=Width,proto3" json:"Width,omitempty"`
	Height     int32  `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	SpriteName string `protobuf:"bytes,6,opt,name=SpriteName,proto3" json:"SpriteName,omitempty"`
}

func (m *RoomSpriteObject) Reset()                    { *m = RoomSpriteObject{} }
func (m *RoomSpriteObject) String() string            { return proto.CompactTextString(m) }
func (*RoomSpriteObject) ProtoMessage()               {}
func (*RoomSpriteObject) Descriptor() ([]byte, []int) { return fileDescriptorRoomSpriteObject, []int{0} }

func (m *RoomSpriteObject) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *RoomSpriteObject) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RoomSpriteObject) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *RoomSpriteObject) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *RoomSpriteObject) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *RoomSpriteObject) GetSpriteName() string {
	if m != nil {
		return m.SpriteName
	}
	return ""
}

func init() {
	proto.RegisterType((*RoomSpriteObject)(nil), "room.RoomSpriteObject")
}
func (m *RoomSpriteObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomSpriteObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UUID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(len(m.UUID)))
		i += copy(dAtA[i:], m.UUID)
	}
	if m.X != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(m.X))
	}
	if m.Y != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(m.Y))
	}
	if m.Width != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(m.Width))
	}
	if m.Height != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(m.Height))
	}
	if len(m.SpriteName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRoomSpriteObject(dAtA, i, uint64(len(m.SpriteName)))
		i += copy(dAtA[i:], m.SpriteName)
	}
	return i, nil
}

func encodeVarintRoomSpriteObject(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RoomSpriteObject) Size() (n int) {
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovRoomSpriteObject(uint64(l))
	}
	if m.X != 0 {
		n += 1 + sovRoomSpriteObject(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovRoomSpriteObject(uint64(m.Y))
	}
	if m.Width != 0 {
		n += 1 + sovRoomSpriteObject(uint64(m.Width))
	}
	if m.Height != 0 {
		n += 1 + sovRoomSpriteObject(uint64(m.Height))
	}
	l = len(m.SpriteName)
	if l > 0 {
		n += 1 + l + sovRoomSpriteObject(uint64(l))
	}
	return n
}

func sovRoomSpriteObject(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoomSpriteObject(x uint64) (n int) {
	return sovRoomSpriteObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoomSpriteObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomSpriteObject
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
			return fmt.Errorf("proto: RoomSpriteObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomSpriteObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomSpriteObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			m.Y = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpriteName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomSpriteObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomSpriteObject
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpriteName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomSpriteObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomSpriteObject
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
func skipRoomSpriteObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomSpriteObject
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
					return 0, ErrIntOverflowRoomSpriteObject
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
					return 0, ErrIntOverflowRoomSpriteObject
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
				return 0, ErrInvalidLengthRoomSpriteObject
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoomSpriteObject
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
				next, err := skipRoomSpriteObject(dAtA[start:])
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
	ErrInvalidLengthRoomSpriteObject = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomSpriteObject   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("room_sprite_object.proto", fileDescriptorRoomSpriteObject) }

var fileDescriptorRoomSpriteObject = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xca, 0xcf, 0xcf,
	0x8d, 0x2f, 0x2e, 0x28, 0xca, 0x2c, 0x49, 0x8d, 0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xc9, 0x28, 0x4d, 0x60, 0xe4, 0x12, 0x08, 0xca, 0xcf,
	0xcf, 0x0d, 0x06, 0xab, 0xf0, 0x07, 0x2b, 0x10, 0x12, 0xe2, 0x62, 0x09, 0x0d, 0xf5, 0x74, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x78, 0xb8, 0x18, 0x23, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x58, 0x83, 0x18, 0x23, 0x40, 0xbc, 0x48, 0x09, 0x66, 0x08, 0x2f, 0x52, 0x48, 0x84,
	0x8b, 0x35, 0x3c, 0x33, 0xa5, 0x24, 0x43, 0x82, 0x05, 0x2c, 0x02, 0xe1, 0x08, 0x89, 0x71, 0xb1,
	0x79, 0xa4, 0x66, 0xa6, 0x67, 0x94, 0x48, 0xb0, 0x82, 0x85, 0xa1, 0x3c, 0x21, 0x39, 0x2e, 0x2e,
	0x88, 0x6d, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x6c, 0x60, 0x3b, 0x90, 0x44, 0x9c, 0x04, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92,
	0xd8, 0xc0, 0x2e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xca, 0x0e, 0xa9, 0x1b, 0xcd, 0x00,
	0x00, 0x00,
}
