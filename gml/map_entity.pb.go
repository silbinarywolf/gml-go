// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: map_entity.proto

/*
	Package gml is a generated protocol buffer package.

	It is generated from these files:
		map_entity.proto
		map.proto

	It has these top-level messages:
		MapEntity
		Map
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

type MapEntity struct {
	ObjectIndex int32 `protobuf:"varint,1,opt,name=ObjectIndex,proto3" json:"ObjectIndex,omitempty"`
	X           int64 `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y           int64 `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *MapEntity) Reset()                    { *m = MapEntity{} }
func (m *MapEntity) String() string            { return proto.CompactTextString(m) }
func (*MapEntity) ProtoMessage()               {}
func (*MapEntity) Descriptor() ([]byte, []int) { return fileDescriptorMapEntity, []int{0} }

func (m *MapEntity) GetObjectIndex() int32 {
	if m != nil {
		return m.ObjectIndex
	}
	return 0
}

func (m *MapEntity) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *MapEntity) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*MapEntity)(nil), "gml.MapEntity")
}
func (m *MapEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapEntity) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ObjectIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMapEntity(dAtA, i, uint64(m.ObjectIndex))
	}
	if m.X != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMapEntity(dAtA, i, uint64(m.X))
	}
	if m.Y != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMapEntity(dAtA, i, uint64(m.Y))
	}
	return i, nil
}

func encodeVarintMapEntity(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MapEntity) Size() (n int) {
	var l int
	_ = l
	if m.ObjectIndex != 0 {
		n += 1 + sovMapEntity(uint64(m.ObjectIndex))
	}
	if m.X != 0 {
		n += 1 + sovMapEntity(uint64(m.X))
	}
	if m.Y != 0 {
		n += 1 + sovMapEntity(uint64(m.Y))
	}
	return n
}

func sovMapEntity(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMapEntity(x uint64) (n int) {
	return sovMapEntity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MapEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMapEntity
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
			return fmt.Errorf("proto: MapEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectIndex", wireType)
			}
			m.ObjectIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMapEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ObjectIndex |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			m.X = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMapEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.X |= (int64(b) & 0x7F) << shift
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
					return ErrIntOverflowMapEntity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Y |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMapEntity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMapEntity
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
func skipMapEntity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMapEntity
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
					return 0, ErrIntOverflowMapEntity
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
					return 0, ErrIntOverflowMapEntity
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
				return 0, ErrInvalidLengthMapEntity
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMapEntity
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
				next, err := skipMapEntity(dAtA[start:])
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
	ErrInvalidLengthMapEntity = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMapEntity   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("map_entity.proto", fileDescriptorMapEntity) }

var fileDescriptorMapEntity = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4d, 0x2c, 0x88,
	0x4f, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcf,
	0xcd, 0x51, 0xf2, 0xe4, 0xe2, 0xf4, 0x4d, 0x2c, 0x70, 0x05, 0x8b, 0x0b, 0x29, 0x70, 0x71, 0xfb,
	0x27, 0x65, 0xa5, 0x26, 0x97, 0x78, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0,
	0x06, 0x21, 0x0b, 0x09, 0xf1, 0x70, 0x31, 0x46, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x31,
	0x46, 0x80, 0x78, 0x91, 0x12, 0xcc, 0x10, 0x5e, 0xa4, 0x93, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c, 0x43, 0x12, 0x1b, 0xd8, 0x22,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe9, 0x36, 0x88, 0x7c, 0x00, 0x00, 0x00,
}
