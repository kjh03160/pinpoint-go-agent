// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: v1/Annotation.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PIntStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntValue    int32                   `protobuf:"varint,1,opt,name=intValue,proto3" json:"intValue,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
}

func (x *PIntStringValue) Reset() {
	*x = PIntStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntStringValue) ProtoMessage() {}

func (x *PIntStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntStringValue.ProtoReflect.Descriptor instead.
func (*PIntStringValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{0}
}

func (x *PIntStringValue) GetIntValue() int32 {
	if x != nil {
		return x.IntValue
	}
	return 0
}

func (x *PIntStringValue) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

type PIntStringStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntValue     int32                   `protobuf:"varint,1,opt,name=intValue,proto3" json:"intValue,omitempty"`
	StringValue1 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue1,proto3" json:"stringValue1,omitempty"`
	StringValue2 *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=stringValue2,proto3" json:"stringValue2,omitempty"`
}

func (x *PIntStringStringValue) Reset() {
	*x = PIntStringStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntStringStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntStringStringValue) ProtoMessage() {}

func (x *PIntStringStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntStringStringValue.ProtoReflect.Descriptor instead.
func (*PIntStringStringValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{1}
}

func (x *PIntStringStringValue) GetIntValue() int32 {
	if x != nil {
		return x.IntValue
	}
	return 0
}

func (x *PIntStringStringValue) GetStringValue1() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue1
	}
	return nil
}

func (x *PIntStringStringValue) GetStringValue2() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue2
	}
	return nil
}

type PLongIntIntByteByteStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LongValue   int64                   `protobuf:"varint,1,opt,name=longValue,proto3" json:"longValue,omitempty"`
	IntValue1   int32                   `protobuf:"varint,2,opt,name=intValue1,proto3" json:"intValue1,omitempty"`
	IntValue2   int32                   `protobuf:"varint,3,opt,name=intValue2,proto3" json:"intValue2,omitempty"`
	ByteValue1  int32                   `protobuf:"zigzag32,4,opt,name=byteValue1,proto3" json:"byteValue1,omitempty"`
	ByteValue2  int32                   `protobuf:"zigzag32,5,opt,name=byteValue2,proto3" json:"byteValue2,omitempty"`
	StringValue *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
}

func (x *PLongIntIntByteByteStringValue) Reset() {
	*x = PLongIntIntByteByteStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLongIntIntByteByteStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLongIntIntByteByteStringValue) ProtoMessage() {}

func (x *PLongIntIntByteByteStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLongIntIntByteByteStringValue.ProtoReflect.Descriptor instead.
func (*PLongIntIntByteByteStringValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{2}
}

func (x *PLongIntIntByteByteStringValue) GetLongValue() int64 {
	if x != nil {
		return x.LongValue
	}
	return 0
}

func (x *PLongIntIntByteByteStringValue) GetIntValue1() int32 {
	if x != nil {
		return x.IntValue1
	}
	return 0
}

func (x *PLongIntIntByteByteStringValue) GetIntValue2() int32 {
	if x != nil {
		return x.IntValue2
	}
	return 0
}

func (x *PLongIntIntByteByteStringValue) GetByteValue1() int32 {
	if x != nil {
		return x.ByteValue1
	}
	return 0
}

func (x *PLongIntIntByteByteStringValue) GetByteValue2() int32 {
	if x != nil {
		return x.ByteValue2
	}
	return 0
}

func (x *PLongIntIntByteByteStringValue) GetStringValue() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue
	}
	return nil
}

type PIntBooleanIntBooleanValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntValue1  int32 `protobuf:"varint,1,opt,name=intValue1,proto3" json:"intValue1,omitempty"`
	BoolValue1 bool  `protobuf:"varint,2,opt,name=boolValue1,proto3" json:"boolValue1,omitempty"`
	IntValue2  int32 `protobuf:"varint,3,opt,name=intValue2,proto3" json:"intValue2,omitempty"`
	BoolValue2 bool  `protobuf:"varint,4,opt,name=boolValue2,proto3" json:"boolValue2,omitempty"`
}

func (x *PIntBooleanIntBooleanValue) Reset() {
	*x = PIntBooleanIntBooleanValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PIntBooleanIntBooleanValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PIntBooleanIntBooleanValue) ProtoMessage() {}

func (x *PIntBooleanIntBooleanValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PIntBooleanIntBooleanValue.ProtoReflect.Descriptor instead.
func (*PIntBooleanIntBooleanValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{3}
}

func (x *PIntBooleanIntBooleanValue) GetIntValue1() int32 {
	if x != nil {
		return x.IntValue1
	}
	return 0
}

func (x *PIntBooleanIntBooleanValue) GetBoolValue1() bool {
	if x != nil {
		return x.BoolValue1
	}
	return false
}

func (x *PIntBooleanIntBooleanValue) GetIntValue2() int32 {
	if x != nil {
		return x.IntValue2
	}
	return 0
}

func (x *PIntBooleanIntBooleanValue) GetBoolValue2() bool {
	if x != nil {
		return x.BoolValue2
	}
	return false
}

type PStringStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringValue1 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=stringValue1,proto3" json:"stringValue1,omitempty"`
	StringValue2 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=stringValue2,proto3" json:"stringValue2,omitempty"`
}

func (x *PStringStringValue) Reset() {
	*x = PStringStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PStringStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PStringStringValue) ProtoMessage() {}

func (x *PStringStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PStringStringValue.ProtoReflect.Descriptor instead.
func (*PStringStringValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{4}
}

func (x *PStringStringValue) GetStringValue1() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue1
	}
	return nil
}

func (x *PStringStringValue) GetStringValue2() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValue2
	}
	return nil
}

type PAnnotationValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Field:
	//	*PAnnotationValue_StringValue
	//	*PAnnotationValue_BoolValue
	//	*PAnnotationValue_IntValue
	//	*PAnnotationValue_LongValue
	//	*PAnnotationValue_ShortValue
	//	*PAnnotationValue_DoubleValue
	//	*PAnnotationValue_BinaryValue
	//	*PAnnotationValue_ByteValue
	//	*PAnnotationValue_IntStringValue
	//	*PAnnotationValue_StringStringValue
	//	*PAnnotationValue_IntStringStringValue
	//	*PAnnotationValue_LongIntIntByteByteStringValue
	//	*PAnnotationValue_IntBooleanIntBooleanValue
	Field isPAnnotationValue_Field `protobuf_oneof:"field"`
}

func (x *PAnnotationValue) Reset() {
	*x = PAnnotationValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PAnnotationValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PAnnotationValue) ProtoMessage() {}

func (x *PAnnotationValue) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PAnnotationValue.ProtoReflect.Descriptor instead.
func (*PAnnotationValue) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{5}
}

func (m *PAnnotationValue) GetField() isPAnnotationValue_Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (x *PAnnotationValue) GetStringValue() string {
	if x, ok := x.GetField().(*PAnnotationValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *PAnnotationValue) GetBoolValue() bool {
	if x, ok := x.GetField().(*PAnnotationValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *PAnnotationValue) GetIntValue() int32 {
	if x, ok := x.GetField().(*PAnnotationValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *PAnnotationValue) GetLongValue() int64 {
	if x, ok := x.GetField().(*PAnnotationValue_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *PAnnotationValue) GetShortValue() int32 {
	if x, ok := x.GetField().(*PAnnotationValue_ShortValue); ok {
		return x.ShortValue
	}
	return 0
}

func (x *PAnnotationValue) GetDoubleValue() float64 {
	if x, ok := x.GetField().(*PAnnotationValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *PAnnotationValue) GetBinaryValue() []byte {
	if x, ok := x.GetField().(*PAnnotationValue_BinaryValue); ok {
		return x.BinaryValue
	}
	return nil
}

func (x *PAnnotationValue) GetByteValue() int32 {
	if x, ok := x.GetField().(*PAnnotationValue_ByteValue); ok {
		return x.ByteValue
	}
	return 0
}

func (x *PAnnotationValue) GetIntStringValue() *PIntStringValue {
	if x, ok := x.GetField().(*PAnnotationValue_IntStringValue); ok {
		return x.IntStringValue
	}
	return nil
}

func (x *PAnnotationValue) GetStringStringValue() *PStringStringValue {
	if x, ok := x.GetField().(*PAnnotationValue_StringStringValue); ok {
		return x.StringStringValue
	}
	return nil
}

func (x *PAnnotationValue) GetIntStringStringValue() *PIntStringStringValue {
	if x, ok := x.GetField().(*PAnnotationValue_IntStringStringValue); ok {
		return x.IntStringStringValue
	}
	return nil
}

func (x *PAnnotationValue) GetLongIntIntByteByteStringValue() *PLongIntIntByteByteStringValue {
	if x, ok := x.GetField().(*PAnnotationValue_LongIntIntByteByteStringValue); ok {
		return x.LongIntIntByteByteStringValue
	}
	return nil
}

func (x *PAnnotationValue) GetIntBooleanIntBooleanValue() *PIntBooleanIntBooleanValue {
	if x, ok := x.GetField().(*PAnnotationValue_IntBooleanIntBooleanValue); ok {
		return x.IntBooleanIntBooleanValue
	}
	return nil
}

type isPAnnotationValue_Field interface {
	isPAnnotationValue_Field()
}

type PAnnotationValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=stringValue,proto3,oneof"`
}

type PAnnotationValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=boolValue,proto3,oneof"`
}

type PAnnotationValue_IntValue struct {
	IntValue int32 `protobuf:"varint,3,opt,name=intValue,proto3,oneof"`
}

type PAnnotationValue_LongValue struct {
	LongValue int64 `protobuf:"varint,4,opt,name=longValue,proto3,oneof"`
}

type PAnnotationValue_ShortValue struct {
	// for compatibility
	ShortValue int32 `protobuf:"zigzag32,5,opt,name=shortValue,proto3,oneof"`
}

type PAnnotationValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=doubleValue,proto3,oneof"`
}

type PAnnotationValue_BinaryValue struct {
	BinaryValue []byte `protobuf:"bytes,7,opt,name=binaryValue,proto3,oneof"`
}

type PAnnotationValue_ByteValue struct {
	// for compatibility
	ByteValue int32 `protobuf:"zigzag32,8,opt,name=byteValue,proto3,oneof"`
}

type PAnnotationValue_IntStringValue struct {
	IntStringValue *PIntStringValue `protobuf:"bytes,9,opt,name=intStringValue,proto3,oneof"`
}

type PAnnotationValue_StringStringValue struct {
	StringStringValue *PStringStringValue `protobuf:"bytes,10,opt,name=stringStringValue,proto3,oneof"`
}

type PAnnotationValue_IntStringStringValue struct {
	IntStringStringValue *PIntStringStringValue `protobuf:"bytes,11,opt,name=intStringStringValue,proto3,oneof"`
}

type PAnnotationValue_LongIntIntByteByteStringValue struct {
	LongIntIntByteByteStringValue *PLongIntIntByteByteStringValue `protobuf:"bytes,12,opt,name=longIntIntByteByteStringValue,proto3,oneof"`
}

type PAnnotationValue_IntBooleanIntBooleanValue struct {
	IntBooleanIntBooleanValue *PIntBooleanIntBooleanValue `protobuf:"bytes,13,opt,name=intBooleanIntBooleanValue,proto3,oneof"`
}

func (*PAnnotationValue_StringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_BoolValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_LongValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_ShortValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_DoubleValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_BinaryValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_ByteValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_StringStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntStringStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_LongIntIntByteByteStringValue) isPAnnotationValue_Field() {}

func (*PAnnotationValue_IntBooleanIntBooleanValue) isPAnnotationValue_Field() {}

type PAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   int32             `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *PAnnotationValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PAnnotation) Reset() {
	*x = PAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_Annotation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PAnnotation) ProtoMessage() {}

func (x *PAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_Annotation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PAnnotation.ProtoReflect.Descriptor instead.
func (*PAnnotation) Descriptor() ([]byte, []int) {
	return file_v1_Annotation_proto_rawDescGZIP(), []int{6}
}

func (x *PAnnotation) GetKey() int32 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *PAnnotation) GetValue() *PAnnotationValue {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_v1_Annotation_proto protoreflect.FileDescriptor

var file_v1_Annotation_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0f, 0x50, 0x49, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x15, 0x50, 0x49, 0x6e,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31,
	0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0x22, 0xfa, 0x01, 0x0a, 0x1e, 0x50, 0x4c, 0x6f, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x49,
	0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x31, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x12,
	0x3e, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x1a, 0x50, 0x49, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x49,
	0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x50,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x31, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x22, 0xcb, 0x05, 0x0a, 0x10, 0x50, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e,
	0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x0a, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11,
	0x48, 0x00, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22,
	0x0a, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52, 0x09, 0x62, 0x79, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x11, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4f, 0x0a,
	0x14, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x49, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x14, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6a,
	0x0a, 0x1d, 0x6c, 0x6f, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x4c, 0x6f, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x1d, 0x6c, 0x6f, 0x6e,
	0x67, 0x49, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5e, 0x0a, 0x19, 0x69, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x49, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x49, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x49, 0x6e,
	0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x19, 0x69, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x49, 0x6e, 0x74, 0x42, 0x6f,
	0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x3b, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x61, 0x76, 0x65, 0x72, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x70, 0x69, 0x6e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_Annotation_proto_rawDescOnce sync.Once
	file_v1_Annotation_proto_rawDescData = file_v1_Annotation_proto_rawDesc
)

func file_v1_Annotation_proto_rawDescGZIP() []byte {
	file_v1_Annotation_proto_rawDescOnce.Do(func() {
		file_v1_Annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_Annotation_proto_rawDescData)
	})
	return file_v1_Annotation_proto_rawDescData
}

var file_v1_Annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_Annotation_proto_goTypes = []interface{}{
	(*PIntStringValue)(nil),                // 0: v1.PIntStringValue
	(*PIntStringStringValue)(nil),          // 1: v1.PIntStringStringValue
	(*PLongIntIntByteByteStringValue)(nil), // 2: v1.PLongIntIntByteByteStringValue
	(*PIntBooleanIntBooleanValue)(nil),     // 3: v1.PIntBooleanIntBooleanValue
	(*PStringStringValue)(nil),             // 4: v1.PStringStringValue
	(*PAnnotationValue)(nil),               // 5: v1.PAnnotationValue
	(*PAnnotation)(nil),                    // 6: v1.PAnnotation
	(*wrapperspb.StringValue)(nil),         // 7: google.protobuf.StringValue
}
var file_v1_Annotation_proto_depIdxs = []int32{
	7,  // 0: v1.PIntStringValue.stringValue:type_name -> google.protobuf.StringValue
	7,  // 1: v1.PIntStringStringValue.stringValue1:type_name -> google.protobuf.StringValue
	7,  // 2: v1.PIntStringStringValue.stringValue2:type_name -> google.protobuf.StringValue
	7,  // 3: v1.PLongIntIntByteByteStringValue.stringValue:type_name -> google.protobuf.StringValue
	7,  // 4: v1.PStringStringValue.stringValue1:type_name -> google.protobuf.StringValue
	7,  // 5: v1.PStringStringValue.stringValue2:type_name -> google.protobuf.StringValue
	0,  // 6: v1.PAnnotationValue.intStringValue:type_name -> v1.PIntStringValue
	4,  // 7: v1.PAnnotationValue.stringStringValue:type_name -> v1.PStringStringValue
	1,  // 8: v1.PAnnotationValue.intStringStringValue:type_name -> v1.PIntStringStringValue
	2,  // 9: v1.PAnnotationValue.longIntIntByteByteStringValue:type_name -> v1.PLongIntIntByteByteStringValue
	3,  // 10: v1.PAnnotationValue.intBooleanIntBooleanValue:type_name -> v1.PIntBooleanIntBooleanValue
	5,  // 11: v1.PAnnotation.value:type_name -> v1.PAnnotationValue
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_v1_Annotation_proto_init() }
func file_v1_Annotation_proto_init() {
	if File_v1_Annotation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_Annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntStringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntStringStringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLongIntIntByteByteStringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PIntBooleanIntBooleanValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PStringStringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PAnnotationValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_Annotation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PAnnotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_v1_Annotation_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*PAnnotationValue_StringValue)(nil),
		(*PAnnotationValue_BoolValue)(nil),
		(*PAnnotationValue_IntValue)(nil),
		(*PAnnotationValue_LongValue)(nil),
		(*PAnnotationValue_ShortValue)(nil),
		(*PAnnotationValue_DoubleValue)(nil),
		(*PAnnotationValue_BinaryValue)(nil),
		(*PAnnotationValue_ByteValue)(nil),
		(*PAnnotationValue_IntStringValue)(nil),
		(*PAnnotationValue_StringStringValue)(nil),
		(*PAnnotationValue_IntStringStringValue)(nil),
		(*PAnnotationValue_LongIntIntByteByteStringValue)(nil),
		(*PAnnotationValue_IntBooleanIntBooleanValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_Annotation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_Annotation_proto_goTypes,
		DependencyIndexes: file_v1_Annotation_proto_depIdxs,
		MessageInfos:      file_v1_Annotation_proto_msgTypes,
	}.Build()
	File_v1_Annotation_proto = out.File
	file_v1_Annotation_proto_rawDesc = nil
	file_v1_Annotation_proto_goTypes = nil
	file_v1_Annotation_proto_depIdxs = nil
}
