// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: commons.proto

package edgecluster

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
// The different error types
type Error int32

const (
	// Indicates the operation was successful
	Error_NO_ERROR Error = 0
	// Indicates the operation fail with unknown error
	Error_UNKNOWN Error = 1
	// Indicates the edge cluster already exists
	Error_EDGE_CLUSTER_ALREADY_EXISTS Error = 2
	// Indicates the edge cluster does not exist
	Error_EDGE_CLUSTER_NOT_FOUND Error = 3
	// Indicates the provided values for he operation were invalid
	Error_BAD_REQUEST Error = 4
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "NO_ERROR",
		1: "UNKNOWN",
		2: "EDGE_CLUSTER_ALREADY_EXISTS",
		3: "EDGE_CLUSTER_NOT_FOUND",
		4: "BAD_REQUEST",
	}
	Error_value = map[string]int32{
		"NO_ERROR":                    0,
		"UNKNOWN":                     1,
		"EDGE_CLUSTER_ALREADY_EXISTS": 2,
		"EDGE_CLUSTER_NOT_FOUND":      3,
		"BAD_REQUEST":                 4,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_commons_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{0}
}

//*
// The different sorting direction
type SortingDirection int32

const (
	// Indicates result data must be sorted from low to high sequence
	SortingDirection_ASCENDING SortingDirection = 0
	// Indicates result data must be sorted from high to low sequence
	SortingDirection_DESCENDING SortingDirection = 1
)

// Enum value maps for SortingDirection.
var (
	SortingDirection_name = map[int32]string{
		0: "ASCENDING",
		1: "DESCENDING",
	}
	SortingDirection_value = map[string]int32{
		"ASCENDING":  0,
		"DESCENDING": 1,
	}
)

func (x SortingDirection) Enum() *SortingDirection {
	p := new(SortingDirection)
	*p = x
	return p
}

func (x SortingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_proto_enumTypes[1].Descriptor()
}

func (SortingDirection) Type() protoreflect.EnumType {
	return &file_commons_proto_enumTypes[1]
}

func (x SortingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingDirection.Descriptor instead.
func (SortingDirection) EnumDescriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{1}
}

//*
// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
type EdgeClusterConditionStatus int32

const (
	EdgeClusterConditionStatus_ConditionTrue    EdgeClusterConditionStatus = 0
	EdgeClusterConditionStatus_ConditionFalse   EdgeClusterConditionStatus = 1
	EdgeClusterConditionStatus_ConditionUnknown EdgeClusterConditionStatus = 2
)

// Enum value maps for EdgeClusterConditionStatus.
var (
	EdgeClusterConditionStatus_name = map[int32]string{
		0: "ConditionTrue",
		1: "ConditionFalse",
		2: "ConditionUnknown",
	}
	EdgeClusterConditionStatus_value = map[string]int32{
		"ConditionTrue":    0,
		"ConditionFalse":   1,
		"ConditionUnknown": 2,
	}
)

func (x EdgeClusterConditionStatus) Enum() *EdgeClusterConditionStatus {
	p := new(EdgeClusterConditionStatus)
	*p = x
	return p
}

func (x EdgeClusterConditionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EdgeClusterConditionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_proto_enumTypes[2].Descriptor()
}

func (EdgeClusterConditionStatus) Type() protoreflect.EnumType {
	return &file_commons_proto_enumTypes[2]
}

func (x EdgeClusterConditionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EdgeClusterConditionStatus.Descriptor instead.
func (EdgeClusterConditionStatus) EnumDescriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{2}
}

//*
// The pagination information compatible with graphql-relay connection definition, for more information visit:
// https://facebook.github.io/relay/graphql/connections.htm
type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasFirst  bool   `protobuf:"varint,1,opt,name=hasFirst,proto3" json:"hasFirst,omitempty"`
	First     int32  `protobuf:"varint,2,opt,name=first,proto3" json:"first,omitempty"`
	HasAfter  bool   `protobuf:"varint,3,opt,name=hasAfter,proto3" json:"hasAfter,omitempty"`
	After     string `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	HasLast   bool   `protobuf:"varint,5,opt,name=hasLast,proto3" json:"hasLast,omitempty"`
	Last      int32  `protobuf:"varint,6,opt,name=last,proto3" json:"last,omitempty"`
	HasBefore bool   `protobuf:"varint,7,opt,name=hasBefore,proto3" json:"hasBefore,omitempty"`
	Before    string `protobuf:"bytes,8,opt,name=before,proto3" json:"before,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_commons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetHasFirst() bool {
	if x != nil {
		return x.HasFirst
	}
	return false
}

func (x *Pagination) GetFirst() int32 {
	if x != nil {
		return x.First
	}
	return 0
}

func (x *Pagination) GetHasAfter() bool {
	if x != nil {
		return x.HasAfter
	}
	return false
}

func (x *Pagination) GetAfter() string {
	if x != nil {
		return x.After
	}
	return ""
}

func (x *Pagination) GetHasLast() bool {
	if x != nil {
		return x.HasLast
	}
	return false
}

func (x *Pagination) GetLast() int32 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *Pagination) GetHasBefore() bool {
	if x != nil {
		return x.HasBefore
	}
	return false
}

func (x *Pagination) GetBefore() string {
	if x != nil {
		return x.Before
	}
	return ""
}

//*
// Defines the pair of values that are used to determine how the result data should be sorted.
type SortingOptionPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the field on
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// THe sorting direction
	Direction SortingDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=edgecluster.SortingDirection" json:"direction,omitempty"`
}

func (x *SortingOptionPair) Reset() {
	*x = SortingOptionPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortingOptionPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortingOptionPair) ProtoMessage() {}

func (x *SortingOptionPair) ProtoReflect() protoreflect.Message {
	mi := &file_commons_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortingOptionPair.ProtoReflect.Descriptor instead.
func (*SortingOptionPair) Descriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{1}
}

func (x *SortingOptionPair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SortingOptionPair) GetDirection() SortingDirection {
	if x != nil {
		return x.Direction
	}
	return SortingDirection_ASCENDING
}

//*
// Standard edge cluster object's metadata.
type EdgeClusterObjectMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The namespace
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *EdgeClusterObjectMetadata) Reset() {
	*x = EdgeClusterObjectMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EdgeClusterObjectMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EdgeClusterObjectMetadata) ProtoMessage() {}

func (x *EdgeClusterObjectMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_commons_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EdgeClusterObjectMetadata.ProtoReflect.Descriptor instead.
func (*EdgeClusterObjectMetadata) Descriptor() ([]byte, []int) {
	return file_commons_proto_rawDescGZIP(), []int{2}
}

func (x *EdgeClusterObjectMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EdgeClusterObjectMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EdgeClusterObjectMetadata) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_commons_proto protoreflect.FileDescriptor

var file_commons_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0xd4, 0x01, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x61, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68,
	0x61, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x68, 0x61, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x68, 0x61, 0x73, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x68, 0x61, 0x73, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x22, 0x64, 0x0a, 0x11, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a, 0x19, 0x45, 0x64, 0x67,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0x70, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x45, 0x44, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x02, 0x12, 0x1a, 0x0a,
	0x16, 0x45, 0x44, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x2a, 0x31, 0x0a, 0x10, 0x53, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x2a, 0x59, 0x0a,
	0x1a, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x75, 0x65, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6c, 0x73, 0x65,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_proto_rawDescOnce sync.Once
	file_commons_proto_rawDescData = file_commons_proto_rawDesc
)

func file_commons_proto_rawDescGZIP() []byte {
	file_commons_proto_rawDescOnce.Do(func() {
		file_commons_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_proto_rawDescData)
	})
	return file_commons_proto_rawDescData
}

var file_commons_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_commons_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commons_proto_goTypes = []interface{}{
	(Error)(0),                        // 0: edgecluster.Error
	(SortingDirection)(0),             // 1: edgecluster.SortingDirection
	(EdgeClusterConditionStatus)(0),   // 2: edgecluster.EdgeClusterConditionStatus
	(*Pagination)(nil),                // 3: edgecluster.Pagination
	(*SortingOptionPair)(nil),         // 4: edgecluster.SortingOptionPair
	(*EdgeClusterObjectMetadata)(nil), // 5: edgecluster.EdgeClusterObjectMetadata
}
var file_commons_proto_depIdxs = []int32{
	1, // 0: edgecluster.SortingOptionPair.direction:type_name -> edgecluster.SortingDirection
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_commons_proto_init() }
func file_commons_proto_init() {
	if File_commons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_commons_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortingOptionPair); i {
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
		file_commons_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EdgeClusterObjectMetadata); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commons_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commons_proto_goTypes,
		DependencyIndexes: file_commons_proto_depIdxs,
		EnumInfos:         file_commons_proto_enumTypes,
		MessageInfos:      file_commons_proto_msgTypes,
	}.Build()
	File_commons_proto = out.File
	file_commons_proto_rawDesc = nil
	file_commons_proto_goTypes = nil
	file_commons_proto_depIdxs = nil
}