// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1alpha1/events.proto

package ecocredit

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

// EventCreateClass is an event emitted when a credit class is created.
type EventCreateClass struct {
	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty" yaml:"class_id"`
	// designer is the designer of the credit class.
	Designer string `protobuf:"bytes,2,opt,name=designer,proto3" json:"designer,omitempty"`
}

func (m *EventCreateClass) Reset()         { *m = EventCreateClass{} }
func (m *EventCreateClass) String() string { return proto.CompactTextString(m) }
func (*EventCreateClass) ProtoMessage()    {}
func (*EventCreateClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b6a013b00aef3af, []int{0}
}
func (m *EventCreateClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateClass.Merge(m, src)
}
func (m *EventCreateClass) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateClass) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateClass.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateClass proto.InternalMessageInfo

func (m *EventCreateClass) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *EventCreateClass) GetDesigner() string {
	if m != nil {
		return m.Designer
	}
	return ""
}

// EventCreateBatch is an event emitted when a credit batch is created.
type EventCreateBatch struct {
	// class_id is the unique ID of credit class.
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty" yaml:"class_id"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty" yaml:"batch_denom"`
	// issuer is the account address of the issuer of the credit batch.
	Issuer string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// total_units is the total number of units in the credit batch.
	TotalUnits string `protobuf:"bytes,4,opt,name=total_units,json=totalUnits,proto3" json:"total_units,omitempty" yaml:"total_units"`
}

func (m *EventCreateBatch) Reset()         { *m = EventCreateBatch{} }
func (m *EventCreateBatch) String() string { return proto.CompactTextString(m) }
func (*EventCreateBatch) ProtoMessage()    {}
func (*EventCreateBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b6a013b00aef3af, []int{1}
}
func (m *EventCreateBatch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateBatch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateBatch.Merge(m, src)
}
func (m *EventCreateBatch) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateBatch.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateBatch proto.InternalMessageInfo

func (m *EventCreateBatch) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *EventCreateBatch) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *EventCreateBatch) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *EventCreateBatch) GetTotalUnits() string {
	if m != nil {
		return m.TotalUnits
	}
	return ""
}

// EventReceive is an event emitted when credits are received either upon
// creation of a new batch or upon transfer. Each batch_denom created or
// transferred will result in a separate EventReceive for easy indexing.
type EventReceive struct {
	// sender is the sender of the credits in the case that this event is the
	// result of a transfer. It will not be set when credits are received at
	// initial issuance.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// recipient is the recipient of the credits
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,3,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty" yaml:"batch_denom"`
	// units is the decimal number of both tradable and retired credits received.
	Units string `protobuf:"bytes,4,opt,name=units,proto3" json:"units,omitempty"`
}

func (m *EventReceive) Reset()         { *m = EventReceive{} }
func (m *EventReceive) String() string { return proto.CompactTextString(m) }
func (*EventReceive) ProtoMessage()    {}
func (*EventReceive) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b6a013b00aef3af, []int{2}
}
func (m *EventReceive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventReceive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventReceive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventReceive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReceive.Merge(m, src)
}
func (m *EventReceive) XXX_Size() int {
	return m.Size()
}
func (m *EventReceive) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReceive.DiscardUnknown(m)
}

var xxx_messageInfo_EventReceive proto.InternalMessageInfo

func (m *EventReceive) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventReceive) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *EventReceive) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *EventReceive) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

// EventRetire is an event emitted when credits are retired. An separate event
// is emitted for each batch_denom in the case where credits from multiple
// batches have been retired at once for easy indexing.
type EventRetire struct {
	// retirer is the account which has done the "retiring". This will be the
	// account receiving credits in the case that credits were retired upon
	// issuance using Msg/CreateBatch or retired upon transfer using Msg/Send.
	Retirer string `protobuf:"bytes,1,opt,name=retirer,proto3" json:"retirer,omitempty"`
	// batch_denom is the unique ID of credit batch.
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty" yaml:"batch_denom"`
	// units is the decimal number of credits that have been retired.
	Units string `protobuf:"bytes,3,opt,name=units,proto3" json:"units,omitempty"`
}

func (m *EventRetire) Reset()         { *m = EventRetire{} }
func (m *EventRetire) String() string { return proto.CompactTextString(m) }
func (*EventRetire) ProtoMessage()    {}
func (*EventRetire) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b6a013b00aef3af, []int{3}
}
func (m *EventRetire) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRetire) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRetire.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRetire) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRetire.Merge(m, src)
}
func (m *EventRetire) XXX_Size() int {
	return m.Size()
}
func (m *EventRetire) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRetire.DiscardUnknown(m)
}

var xxx_messageInfo_EventRetire proto.InternalMessageInfo

func (m *EventRetire) GetRetirer() string {
	if m != nil {
		return m.Retirer
	}
	return ""
}

func (m *EventRetire) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *EventRetire) GetUnits() string {
	if m != nil {
		return m.Units
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreateClass)(nil), "regen.ecocredit.v1alpha1.EventCreateClass")
	proto.RegisterType((*EventCreateBatch)(nil), "regen.ecocredit.v1alpha1.EventCreateBatch")
	proto.RegisterType((*EventReceive)(nil), "regen.ecocredit.v1alpha1.EventReceive")
	proto.RegisterType((*EventRetire)(nil), "regen.ecocredit.v1alpha1.EventRetire")
}

func init() {
	proto.RegisterFile("regen/ecocredit/v1alpha1/events.proto", fileDescriptor_5b6a013b00aef3af)
}

var fileDescriptor_5b6a013b00aef3af = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x9b, 0xdb, 0x7b, 0xfb, 0x67, 0x7a, 0xe1, 0x5e, 0x62, 0x29, 0xa1, 0x48, 0x2a, 0x01,
	0xc1, 0x8d, 0x09, 0xc5, 0x85, 0xe0, 0x32, 0xd5, 0x85, 0xb8, 0x0b, 0xb8, 0x71, 0x61, 0x49, 0x33,
	0x87, 0x74, 0x30, 0x9d, 0x29, 0x33, 0xd3, 0xa8, 0x6f, 0x21, 0xf8, 0x52, 0x2e, 0x5c, 0x74, 0xe9,
	0xaa, 0x48, 0xfb, 0x06, 0x7d, 0x02, 0x99, 0x49, 0xd2, 0xc6, 0xae, 0xd4, 0xdd, 0xf9, 0xcd, 0xf9,
	0xce, 0x7c, 0xe7, 0x1c, 0x0e, 0x3a, 0xe4, 0x10, 0x03, 0xf5, 0x20, 0x62, 0x11, 0x07, 0x4c, 0xa4,
	0x97, 0xf6, 0xc3, 0x64, 0x3a, 0x0e, 0xfb, 0x1e, 0xa4, 0x40, 0xa5, 0x70, 0xa7, 0x9c, 0x49, 0x66,
	0x5a, 0x5a, 0xe6, 0x6e, 0x64, 0x6e, 0x21, 0xeb, 0xb6, 0x63, 0x16, 0x33, 0x2d, 0xf2, 0x54, 0x94,
	0xe9, 0x9d, 0x5b, 0xf4, 0xff, 0x42, 0xd5, 0x0f, 0x38, 0x84, 0x12, 0x06, 0x49, 0x28, 0x84, 0xe9,
	0xa2, 0x46, 0xa4, 0x82, 0x21, 0xc1, 0x96, 0x71, 0x60, 0x1c, 0x35, 0xfd, 0xbd, 0xf5, 0xa2, 0xf7,
	0xef, 0x31, 0x9c, 0x24, 0x67, 0x4e, 0x91, 0x71, 0x82, 0xba, 0x0e, 0x2f, 0xb1, 0xd9, 0x45, 0x0d,
	0x0c, 0x82, 0xc4, 0x14, 0xb8, 0xf5, 0x4b, 0xe9, 0x83, 0x0d, 0x3b, 0xaf, 0xc6, 0x27, 0x03, 0x3f,
	0x94, 0xd1, 0xf8, 0xdb, 0x06, 0xa7, 0xa8, 0x35, 0x52, 0x85, 0x43, 0x0c, 0x94, 0x4d, 0x32, 0x0f,
	0xbf, 0xb3, 0x5e, 0xf4, 0xcc, 0xac, 0xa4, 0x94, 0x74, 0x02, 0xa4, 0xe9, 0x5c, 0x81, 0xd9, 0x41,
	0x35, 0x22, 0xc4, 0x0c, 0xb8, 0x55, 0xd5, 0x7d, 0xe5, 0xa4, 0x3e, 0x94, 0x4c, 0x86, 0xc9, 0x70,
	0x46, 0x89, 0x14, 0xd6, 0xef, 0xdd, 0x0f, 0x4b, 0x49, 0x27, 0x40, 0x9a, 0xae, 0x35, 0x3c, 0x1b,
	0xe8, 0xaf, 0x1e, 0x27, 0x80, 0x08, 0x48, 0x0a, 0xca, 0x41, 0x00, 0xc5, 0xc0, 0xb3, 0x41, 0x82,
	0x9c, 0xcc, 0x7d, 0xd4, 0xe4, 0x10, 0x91, 0x29, 0x01, 0x2a, 0xf3, 0xa5, 0x6c, 0x1f, 0x76, 0x07,
	0xaa, 0x7e, 0x79, 0xa0, 0x36, 0xfa, 0x53, 0x6a, 0x39, 0xc8, 0xc0, 0x49, 0x51, 0x2b, 0x6f, 0x4a,
	0x12, 0x0e, 0xa6, 0x85, 0xea, 0x5c, 0x47, 0x45, 0x53, 0x05, 0xfe, 0x7c, 0x91, 0x1b, 0xdf, 0x6a,
	0xc9, 0xd7, 0xbf, 0x7a, 0x59, 0xda, 0xc6, 0x7c, 0x69, 0x1b, 0xef, 0x4b, 0xdb, 0x78, 0x5a, 0xd9,
	0x95, 0xf9, 0xca, 0xae, 0xbc, 0xad, 0xec, 0xca, 0x4d, 0x3f, 0x26, 0x72, 0x3c, 0x1b, 0xb9, 0x11,
	0x9b, 0x78, 0xfa, 0x22, 0x8f, 0x29, 0xc8, 0x7b, 0xc6, 0xef, 0x72, 0x4a, 0x00, 0xc7, 0xc0, 0xbd,
	0x87, 0xed, 0x3d, 0x8f, 0x6a, 0xfa, 0x20, 0x4f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xbd,
	0x84, 0x06, 0xe9, 0x02, 0x00, 0x00,
}

func (m *EventCreateClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Designer) > 0 {
		i -= len(m.Designer)
		copy(dAtA[i:], m.Designer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Designer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateBatch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateBatch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateBatch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalUnits) > 0 {
		i -= len(m.TotalUnits)
		copy(dAtA[i:], m.TotalUnits)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TotalUnits)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventReceive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventReceive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventReceive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Units) > 0 {
		i -= len(m.Units)
		copy(dAtA[i:], m.Units)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Units)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRetire) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRetire) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRetire) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Units) > 0 {
		i -= len(m.Units)
		copy(dAtA[i:], m.Units)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Units)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Retirer) > 0 {
		i -= len(m.Retirer)
		copy(dAtA[i:], m.Retirer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Retirer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Designer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCreateBatch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TotalUnits)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventReceive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Units)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRetire) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Retirer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Units)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Designer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateBatch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateBatch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateBatch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalUnits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventReceive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventReceive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventReceive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Units = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRetire) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRetire: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRetire: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retirer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Retirer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Units = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
