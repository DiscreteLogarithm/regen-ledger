// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1alpha1/types.proto

package ecocredit

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

// ClassInfo represents the high-level on-chain information for a credit class.
type ClassInfo struct {
	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// designer is the designer of the credit class.
	Designer string `protobuf:"bytes,2,opt,name=designer,proto3" json:"designer,omitempty"`
	// issuers are the approved issuers of the credit class.
	Issuers []string `protobuf:"bytes,3,rep,name=issuers,proto3" json:"issuers,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit class.
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *ClassInfo) Reset()         { *m = ClassInfo{} }
func (m *ClassInfo) String() string { return proto.CompactTextString(m) }
func (*ClassInfo) ProtoMessage()    {}
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5342f4dcaeff1a84, []int{0}
}
func (m *ClassInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassInfo.Merge(m, src)
}
func (m *ClassInfo) XXX_Size() int {
	return m.Size()
}
func (m *ClassInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClassInfo proto.InternalMessageInfo

func (m *ClassInfo) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *ClassInfo) GetDesigner() string {
	if m != nil {
		return m.Designer
	}
	return ""
}

func (m *ClassInfo) GetIssuers() []string {
	if m != nil {
		return m.Issuers
	}
	return nil
}

func (m *ClassInfo) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// BatchInfo represents the high-level on-chain information for a credit batch.
type BatchInfo struct {
	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// issuer is the issuer of the credit batch.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// total_units is the total number of units in the credit batch and is immutable.
	TotalUnits string `protobuf:"bytes,4,opt,name=total_units,json=totalUnits,proto3" json:"total_units,omitempty"`
	// metadata is any arbitrary metadata to attached to the credit batch.
	Metadata []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *BatchInfo) Reset()         { *m = BatchInfo{} }
func (m *BatchInfo) String() string { return proto.CompactTextString(m) }
func (*BatchInfo) ProtoMessage()    {}
func (*BatchInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5342f4dcaeff1a84, []int{1}
}
func (m *BatchInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchInfo.Merge(m, src)
}
func (m *BatchInfo) XXX_Size() int {
	return m.Size()
}
func (m *BatchInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BatchInfo proto.InternalMessageInfo

func (m *BatchInfo) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *BatchInfo) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *BatchInfo) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *BatchInfo) GetTotalUnits() string {
	if m != nil {
		return m.TotalUnits
	}
	return ""
}

func (m *BatchInfo) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ClassInfo)(nil), "regen.ecocredit.v1alpha1.ClassInfo")
	proto.RegisterType((*BatchInfo)(nil), "regen.ecocredit.v1alpha1.BatchInfo")
}

func init() {
	proto.RegisterFile("regen/ecocredit/v1alpha1/types.proto", fileDescriptor_5342f4dcaeff1a84)
}

var fileDescriptor_5342f4dcaeff1a84 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x6b, 0x0a, 0x6d, 0x63, 0x98, 0x32, 0x20, 0xc3, 0x60, 0xaa, 0x8a, 0xa1, 0x0b, 0x89,
	0x2a, 0x6e, 0x50, 0x58, 0x2a, 0xb6, 0x4a, 0x2c, 0x2c, 0x95, 0x13, 0xff, 0x24, 0x16, 0x89, 0x1d,
	0xd9, 0x0e, 0x94, 0x5b, 0x70, 0x00, 0x0e, 0xc4, 0xd8, 0x91, 0x11, 0x25, 0x17, 0x41, 0x36, 0x21,
	0xa8, 0x13, 0xe3, 0xf7, 0xfc, 0x9e, 0xde, 0x93, 0x7f, 0x7c, 0xa9, 0x21, 0x03, 0x19, 0x43, 0xaa,
	0x52, 0x0d, 0x5c, 0xd8, 0xf8, 0x79, 0xc1, 0x8a, 0x2a, 0x67, 0x8b, 0xd8, 0xbe, 0x56, 0x60, 0xa2,
	0x4a, 0x2b, 0xab, 0x42, 0xe2, 0x5d, 0x51, 0xef, 0x8a, 0x7e, 0x5d, 0xb3, 0x2d, 0x0e, 0x6e, 0x0a,
	0x66, 0xcc, 0x4a, 0x3e, 0xaa, 0xf0, 0x0c, 0x4f, 0x52, 0x07, 0x1b, 0xc1, 0x09, 0x9a, 0xa2, 0x79,
	0xb0, 0x1e, 0x7b, 0x5e, 0xf1, 0xf0, 0x1c, 0x4f, 0x38, 0x18, 0x91, 0x49, 0xd0, 0xe4, 0xc0, 0x3f,
	0xf5, 0x1c, 0x12, 0x3c, 0x16, 0xc6, 0xd4, 0xa0, 0x0d, 0x19, 0x4e, 0x87, 0x2e, 0xd5, 0xa1, 0x4b,
	0x95, 0x60, 0x19, 0x67, 0x96, 0x91, 0xc3, 0x29, 0x9a, 0x9f, 0xac, 0x7b, 0x9e, 0xbd, 0x23, 0x1c,
	0x2c, 0x99, 0x4d, 0xf3, 0xff, 0xaa, 0x2f, 0xf0, 0x71, 0xe2, 0x7c, 0x1b, 0x0e, 0x52, 0x95, 0x5d,
	0x3b, 0xf6, 0xd2, 0xad, 0x53, 0xc2, 0x53, 0x3c, 0xfa, 0x29, 0x24, 0x43, 0xff, 0xd6, 0x91, 0x0b,
	0x5a, 0x65, 0x59, 0xb1, 0xa9, 0xa5, 0xb0, 0xc6, 0x0f, 0x08, 0xd6, 0xd8, 0x4b, 0xf7, 0x4e, 0xd9,
	0x9b, 0x77, 0xb4, 0x3f, 0x6f, 0x79, 0xf7, 0xd1, 0x50, 0xb4, 0x6b, 0x28, 0xfa, 0x6a, 0x28, 0x7a,
	0x6b, 0xe9, 0x60, 0xd7, 0xd2, 0xc1, 0x67, 0x4b, 0x07, 0x0f, 0x8b, 0x4c, 0xd8, 0xbc, 0x4e, 0xa2,
	0x54, 0x95, 0xb1, 0xff, 0xd7, 0x2b, 0x09, 0xf6, 0x45, 0xe9, 0xa7, 0x8e, 0x0a, 0xe0, 0x19, 0xe8,
	0x78, 0xfb, 0x77, 0x94, 0x64, 0xe4, 0xcf, 0x70, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x26, 0xbd,
	0x7d, 0x8e, 0xae, 0x01, 0x00, 0x00,
}

func (m *ClassInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuers) > 0 {
		for iNdEx := len(m.Issuers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Issuers[iNdEx])
			copy(dAtA[i:], m.Issuers[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Issuers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Designer) > 0 {
		i -= len(m.Designer)
		copy(dAtA[i:], m.Designer)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Designer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BatchInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TotalUnits) > 0 {
		i -= len(m.TotalUnits)
		copy(dAtA[i:], m.TotalUnits)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TotalUnits)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Designer)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Issuers) > 0 {
		for _, s := range m.Issuers {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BatchInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.TotalUnits)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ClassInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Designer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Designer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuers = append(m.Issuers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata[:0], dAtA[iNdEx:postIndex]...)
			if m.Metadata == nil {
				m.Metadata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *BatchInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BatchInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalUnits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalUnits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata[:0], dAtA[iNdEx:postIndex]...)
			if m.Metadata == nil {
				m.Metadata = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
