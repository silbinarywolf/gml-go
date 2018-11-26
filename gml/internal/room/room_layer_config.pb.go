// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room_layer_config.proto

package room

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RoomLayerConfig struct {
	Kind  RoomLayerKind `protobuf:"varint,1,opt,name=Kind,proto3,enum=room.RoomLayerKind" json:"Kind,omitempty"`
	UUID  string        `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name  string        `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Order int32         `protobuf:"varint,4,opt,name=Order,proto3" json:"Order,omitempty"`
	// RoomLayerSprite Only
	HasCollision bool `protobuf:"varint,5,opt,name=HasCollision,proto3" json:"HasCollision,omitempty"`
}

func (m *RoomLayerConfig) Reset()                    { *m = RoomLayerConfig{} }
func (m *RoomLayerConfig) String() string            { return proto.CompactTextString(m) }
func (*RoomLayerConfig) ProtoMessage()               {}
func (*RoomLayerConfig) Descriptor() ([]byte, []int) { return fileDescriptorRoomLayerConfig, []int{0} }

func (m *RoomLayerConfig) GetKind() RoomLayerKind {
	if m != nil {
		return m.Kind
	}
	return RoomLayerKind_None
}

func (m *RoomLayerConfig) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *RoomLayerConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoomLayerConfig) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *RoomLayerConfig) GetHasCollision() bool {
	if m != nil {
		return m.HasCollision
	}
	return false
}

func init() {
	proto.RegisterType((*RoomLayerConfig)(nil), "room.RoomLayerConfig")
}
func (m *RoomLayerConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoomLayerConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomLayerConfig(dAtA, i, uint64(m.Kind))
	}
	if len(m.UUID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoomLayerConfig(dAtA, i, uint64(len(m.UUID)))
		i += copy(dAtA[i:], m.UUID)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomLayerConfig(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Order != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRoomLayerConfig(dAtA, i, uint64(m.Order))
	}
	if m.HasCollision {
		dAtA[i] = 0x28
		i++
		if m.HasCollision {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintRoomLayerConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RoomLayerConfig) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovRoomLayerConfig(uint64(m.Kind))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovRoomLayerConfig(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRoomLayerConfig(uint64(l))
	}
	if m.Order != 0 {
		n += 1 + sovRoomLayerConfig(uint64(m.Order))
	}
	if m.HasCollision {
		n += 2
	}
	return n
}

func sovRoomLayerConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoomLayerConfig(x uint64) (n int) {
	return sovRoomLayerConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RoomLayerConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomLayerConfig
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
			return fmt.Errorf("proto: RoomLayerConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoomLayerConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomLayerConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (RoomLayerKind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomLayerConfig
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
				return ErrInvalidLengthRoomLayerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomLayerConfig
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
				return ErrInvalidLengthRoomLayerConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			m.Order = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomLayerConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Order |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasCollision", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomLayerConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasCollision = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRoomLayerConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomLayerConfig
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
func skipRoomLayerConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomLayerConfig
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
					return 0, ErrIntOverflowRoomLayerConfig
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
					return 0, ErrIntOverflowRoomLayerConfig
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
				return 0, ErrInvalidLengthRoomLayerConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoomLayerConfig
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
				next, err := skipRoomLayerConfig(dAtA[start:])
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
	ErrInvalidLengthRoomLayerConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomLayerConfig   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("room_layer_config.proto", fileDescriptorRoomLayerConfig) }

var fileDescriptorRoomLayerConfig = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xca, 0xcf, 0xcf,
	0x8d, 0xcf, 0x49, 0xac, 0x4c, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x49, 0x48, 0x89, 0x22, 0x49, 0x67, 0x67, 0xe6, 0xa5, 0x40,
	0x24, 0x95, 0xe6, 0x30, 0x72, 0xf1, 0x07, 0xe5, 0xe7, 0xe7, 0xfa, 0x80, 0x24, 0x9c, 0xc1, 0xda,
	0x84, 0xd4, 0xb9, 0x58, 0xbc, 0x33, 0xf3, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x84,
	0xf5, 0x40, 0x3a, 0xf5, 0xe0, 0x8a, 0x40, 0x52, 0x41, 0x60, 0x05, 0x42, 0x42, 0x5c, 0x2c, 0xa1,
	0xa1, 0x9e, 0x2e, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0x36, 0x48, 0xcc, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x19, 0x22, 0x06, 0x62, 0x0b, 0x89, 0x70, 0xb1, 0xfa, 0x17, 0xa5, 0xa4, 0x16,
	0x49, 0xb0, 0x28, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42, 0x4a, 0x5c, 0x3c, 0x1e, 0x89, 0xc5,
	0xce, 0xf9, 0x39, 0x39, 0x99, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0x1c, 0x41,
	0x28, 0x62, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x8c, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x77, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x27, 0xdb, 0xac, 0xef, 0x00, 0x00, 0x00,
}