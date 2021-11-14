// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: proto/statesync.proto

package spec

import (
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

type RequestSync_State int32

const (
	RequestSync_UNKNOWN RequestSync_State = 0
	RequestSync_ON      RequestSync_State = 1
	RequestSync_OFF     RequestSync_State = 2
)

// Enum value maps for RequestSync_State.
var (
	RequestSync_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "ON",
		2: "OFF",
	}
	RequestSync_State_value = map[string]int32{
		"UNKNOWN": 0,
		"ON":      1,
		"OFF":     2,
	}
)

func (x RequestSync_State) Enum() *RequestSync_State {
	p := new(RequestSync_State)
	*p = x
	return p
}

func (x RequestSync_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestSync_State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_statesync_proto_enumTypes[0].Descriptor()
}

func (RequestSync_State) Type() protoreflect.EnumType {
	return &file_proto_statesync_proto_enumTypes[0]
}

func (x RequestSync_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestSync_State.Descriptor instead.
func (RequestSync_State) EnumDescriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{0, 0}
}

type ResponseSync_Response int32

const (
	ResponseSync_UNKNOWN ResponseSync_Response = 0
	ResponseSync_SUCCESS ResponseSync_Response = 1
	ResponseSync_FAILED  ResponseSync_Response = 2
)

// Enum value maps for ResponseSync_Response.
var (
	ResponseSync_Response_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILED",
	}
	ResponseSync_Response_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILED":  2,
	}
)

func (x ResponseSync_Response) Enum() *ResponseSync_Response {
	p := new(ResponseSync_Response)
	*p = x
	return p
}

func (x ResponseSync_Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseSync_Response) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_statesync_proto_enumTypes[1].Descriptor()
}

func (ResponseSync_Response) Type() protoreflect.EnumType {
	return &file_proto_statesync_proto_enumTypes[1]
}

func (x ResponseSync_Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseSync_Response.Descriptor instead.
func (ResponseSync_Response) EnumDescriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{1, 0}
}

type Stations_StationId int32

const (
	Stations_unknown       Stations_StationId = 0
	Stations_motoyawata_s1 Stations_StationId = 1
	Stations_motoyawata_s2 Stations_StationId = 2
	Stations_motoyawata_s3 Stations_StationId = 3
	Stations_motoyawata_s4 Stations_StationId = 4
	Stations_motoyawata_s5 Stations_StationId = 5
	Stations_motoyawata_s6 Stations_StationId = 6
	Stations_iwamotocho_s1 Stations_StationId = 11
	Stations_iwamotocho_s2 Stations_StationId = 12
	Stations_iwamotocho_s4 Stations_StationId = 13
	Stations_iwamotocho_b1 Stations_StationId = 14
	Stations_iwamotocho_b2 Stations_StationId = 15
	Stations_iwamotocho_b3 Stations_StationId = 16
	Stations_iwamotocho_b4 Stations_StationId = 17
	Stations_kudanshita_s5 Stations_StationId = 21
	Stations_kudanshita_s6 Stations_StationId = 22
	Stations_sasazuka_b1   Stations_StationId = 31
	Stations_sasazuka_b2   Stations_StationId = 32
	Stations_sasazuka_s1   Stations_StationId = 33
	Stations_sasazuka_s2   Stations_StationId = 34
	Stations_sasazuka_s3   Stations_StationId = 35
	Stations_sasazuka_s4   Stations_StationId = 36
	Stations_sasazuka_s5   Stations_StationId = 37
	Stations_meidaimae_s1  Stations_StationId = 41
	Stations_meidaimae_s2  Stations_StationId = 42
	Stations_chofu_s1      Stations_StationId = 51
	Stations_chofu_s2      Stations_StationId = 52
	Stations_chofu_s3      Stations_StationId = 53
	Stations_chofu_s4      Stations_StationId = 54
	Stations_chofu_s5      Stations_StationId = 55
	Stations_chofu_s6      Stations_StationId = 56
	Stations_chofu_b1      Stations_StationId = 57
	Stations_chofu_b2      Stations_StationId = 58
	Stations_chofu_b3      Stations_StationId = 59
	Stations_chofu_b4      Stations_StationId = 60
	Stations_chofu_b5      Stations_StationId = 61
	Stations_kitano_b1     Stations_StationId = 71
	Stations_kitano_b2     Stations_StationId = 72
	Stations_kitano_b3     Stations_StationId = 73
	Stations_kitano_b4     Stations_StationId = 74
	Stations_kitano_s1     Stations_StationId = 75
	Stations_kitano_s2     Stations_StationId = 76
	Stations_kitano_s3     Stations_StationId = 77
	Stations_kitano_s4     Stations_StationId = 78
	Stations_kitano_s5     Stations_StationId = 79
	Stations_kitano_s6     Stations_StationId = 80
	Stations_takao_s1      Stations_StationId = 91
	Stations_takao_s2      Stations_StationId = 92
)

// Enum value maps for Stations_StationId.
var (
	Stations_StationId_name = map[int32]string{
		0:  "unknown",
		1:  "motoyawata_s1",
		2:  "motoyawata_s2",
		3:  "motoyawata_s3",
		4:  "motoyawata_s4",
		5:  "motoyawata_s5",
		6:  "motoyawata_s6",
		11: "iwamotocho_s1",
		12: "iwamotocho_s2",
		13: "iwamotocho_s4",
		14: "iwamotocho_b1",
		15: "iwamotocho_b2",
		16: "iwamotocho_b3",
		17: "iwamotocho_b4",
		21: "kudanshita_s5",
		22: "kudanshita_s6",
		31: "sasazuka_b1",
		32: "sasazuka_b2",
		33: "sasazuka_s1",
		34: "sasazuka_s2",
		35: "sasazuka_s3",
		36: "sasazuka_s4",
		37: "sasazuka_s5",
		41: "meidaimae_s1",
		42: "meidaimae_s2",
		51: "chofu_s1",
		52: "chofu_s2",
		53: "chofu_s3",
		54: "chofu_s4",
		55: "chofu_s5",
		56: "chofu_s6",
		57: "chofu_b1",
		58: "chofu_b2",
		59: "chofu_b3",
		60: "chofu_b4",
		61: "chofu_b5",
		71: "kitano_b1",
		72: "kitano_b2",
		73: "kitano_b3",
		74: "kitano_b4",
		75: "kitano_s1",
		76: "kitano_s2",
		77: "kitano_s3",
		78: "kitano_s4",
		79: "kitano_s5",
		80: "kitano_s6",
		91: "takao_s1",
		92: "takao_s2",
	}
	Stations_StationId_value = map[string]int32{
		"unknown":       0,
		"motoyawata_s1": 1,
		"motoyawata_s2": 2,
		"motoyawata_s3": 3,
		"motoyawata_s4": 4,
		"motoyawata_s5": 5,
		"motoyawata_s6": 6,
		"iwamotocho_s1": 11,
		"iwamotocho_s2": 12,
		"iwamotocho_s4": 13,
		"iwamotocho_b1": 14,
		"iwamotocho_b2": 15,
		"iwamotocho_b3": 16,
		"iwamotocho_b4": 17,
		"kudanshita_s5": 21,
		"kudanshita_s6": 22,
		"sasazuka_b1":   31,
		"sasazuka_b2":   32,
		"sasazuka_s1":   33,
		"sasazuka_s2":   34,
		"sasazuka_s3":   35,
		"sasazuka_s4":   36,
		"sasazuka_s5":   37,
		"meidaimae_s1":  41,
		"meidaimae_s2":  42,
		"chofu_s1":      51,
		"chofu_s2":      52,
		"chofu_s3":      53,
		"chofu_s4":      54,
		"chofu_s5":      55,
		"chofu_s6":      56,
		"chofu_b1":      57,
		"chofu_b2":      58,
		"chofu_b3":      59,
		"chofu_b4":      60,
		"chofu_b5":      61,
		"kitano_b1":     71,
		"kitano_b2":     72,
		"kitano_b3":     73,
		"kitano_b4":     74,
		"kitano_s1":     75,
		"kitano_s2":     76,
		"kitano_s3":     77,
		"kitano_s4":     78,
		"kitano_s5":     79,
		"kitano_s6":     80,
		"takao_s1":      91,
		"takao_s2":      92,
	}
)

func (x Stations_StationId) Enum() *Stations_StationId {
	p := new(Stations_StationId)
	*p = x
	return p
}

func (x Stations_StationId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Stations_StationId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_statesync_proto_enumTypes[2].Descriptor()
}

func (Stations_StationId) Type() protoreflect.EnumType {
	return &file_proto_statesync_proto_enumTypes[2]
}

func (x Stations_StationId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Stations_StationId.Descriptor instead.
func (Stations_StationId) EnumDescriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{2, 0}
}

type RequestSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station *Stations         `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
	State   RequestSync_State `protobuf:"varint,2,opt,name=state,proto3,enum=RequestSync_State" json:"state,omitempty"`
}

func (x *RequestSync) Reset() {
	*x = RequestSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statesync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSync) ProtoMessage() {}

func (x *RequestSync) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statesync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSync.ProtoReflect.Descriptor instead.
func (*RequestSync) Descriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{0}
}

func (x *RequestSync) GetStation() *Stations {
	if x != nil {
		return x.Station
	}
	return nil
}

func (x *RequestSync) GetState() RequestSync_State {
	if x != nil {
		return x.State
	}
	return RequestSync_UNKNOWN
}

type ResponseSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response ResponseSync_Response `protobuf:"varint,1,opt,name=response,proto3,enum=ResponseSync_Response" json:"response,omitempty"`
}

func (x *ResponseSync) Reset() {
	*x = ResponseSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statesync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSync) ProtoMessage() {}

func (x *ResponseSync) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statesync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSync.ProtoReflect.Descriptor instead.
func (*ResponseSync) Descriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseSync) GetResponse() ResponseSync_Response {
	if x != nil {
		return x.Response
	}
	return ResponseSync_UNKNOWN
}

type Stations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationId Stations_StationId `protobuf:"varint,1,opt,name=stationId,proto3,enum=Stations_StationId" json:"stationId,omitempty"`
}

func (x *Stations) Reset() {
	*x = Stations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_statesync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stations) ProtoMessage() {}

func (x *Stations) ProtoReflect() protoreflect.Message {
	mi := &file_proto_statesync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stations.ProtoReflect.Descriptor instead.
func (*Stations) Descriptor() ([]byte, []int) {
	return file_proto_statesync_proto_rawDescGZIP(), []int{2}
}

func (x *Stations) GetStationId() Stations_StationId {
	if x != nil {
		return x.StationId
	}
	return Stations_unknown
}

var File_proto_statesync_proto protoreflect.FileDescriptor

var file_proto_statesync_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x22, 0x74, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x32, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x22, 0xdc, 0x06, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x31, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x9c, 0x06, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x6d, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x77, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x31, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6d, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x77, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x32, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x6d, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x77, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x33, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x6d, 0x6f, 0x74, 0x6f, 0x79,
	0x61, 0x77, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x34, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x6d, 0x6f,
	0x74, 0x6f, 0x79, 0x61, 0x77, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x35, 0x10, 0x05, 0x12, 0x11, 0x0a,
	0x0d, 0x6d, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x77, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x36, 0x10, 0x06,
	0x12, 0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x6f, 0x5f, 0x73,
	0x31, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68,
	0x6f, 0x5f, 0x73, 0x32, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74,
	0x6f, 0x63, 0x68, 0x6f, 0x5f, 0x73, 0x34, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61,
	0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x6f, 0x5f, 0x62, 0x31, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d,
	0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x6f, 0x5f, 0x62, 0x32, 0x10, 0x0f, 0x12,
	0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x6f, 0x5f, 0x62, 0x33,
	0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x69, 0x77, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x6f,
	0x5f, 0x62, 0x34, 0x10, 0x11, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x75, 0x64, 0x61, 0x6e, 0x73, 0x68,
	0x69, 0x74, 0x61, 0x5f, 0x73, 0x35, 0x10, 0x15, 0x12, 0x11, 0x0a, 0x0d, 0x6b, 0x75, 0x64, 0x61,
	0x6e, 0x73, 0x68, 0x69, 0x74, 0x61, 0x5f, 0x73, 0x36, 0x10, 0x16, 0x12, 0x0f, 0x0a, 0x0b, 0x73,
	0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x62, 0x31, 0x10, 0x1f, 0x12, 0x0f, 0x0a, 0x0b,
	0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x62, 0x32, 0x10, 0x20, 0x12, 0x0f, 0x0a,
	0x0b, 0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x73, 0x31, 0x10, 0x21, 0x12, 0x0f,
	0x0a, 0x0b, 0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x73, 0x32, 0x10, 0x22, 0x12,
	0x0f, 0x0a, 0x0b, 0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x73, 0x33, 0x10, 0x23,
	0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x73, 0x34, 0x10,
	0x24, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x61, 0x73, 0x61, 0x7a, 0x75, 0x6b, 0x61, 0x5f, 0x73, 0x35,
	0x10, 0x25, 0x12, 0x10, 0x0a, 0x0c, 0x6d, 0x65, 0x69, 0x64, 0x61, 0x69, 0x6d, 0x61, 0x65, 0x5f,
	0x73, 0x31, 0x10, 0x29, 0x12, 0x10, 0x0a, 0x0c, 0x6d, 0x65, 0x69, 0x64, 0x61, 0x69, 0x6d, 0x61,
	0x65, 0x5f, 0x73, 0x32, 0x10, 0x2a, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f,
	0x73, 0x31, 0x10, 0x33, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x73, 0x32,
	0x10, 0x34, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x73, 0x33, 0x10, 0x35,
	0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x73, 0x34, 0x10, 0x36, 0x12, 0x0c,
	0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x73, 0x35, 0x10, 0x37, 0x12, 0x0c, 0x0a, 0x08,
	0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x73, 0x36, 0x10, 0x38, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68,
	0x6f, 0x66, 0x75, 0x5f, 0x62, 0x31, 0x10, 0x39, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66,
	0x75, 0x5f, 0x62, 0x32, 0x10, 0x3a, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f,
	0x62, 0x33, 0x10, 0x3b, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x34,
	0x10, 0x3c, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x35, 0x10, 0x3d,
	0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x62, 0x31, 0x10, 0x47, 0x12,
	0x0d, 0x0a, 0x09, 0x6b, 0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x62, 0x32, 0x10, 0x48, 0x12, 0x0d,
	0x0a, 0x09, 0x6b, 0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x62, 0x33, 0x10, 0x49, 0x12, 0x0d, 0x0a,
	0x09, 0x6b, 0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x62, 0x34, 0x10, 0x4a, 0x12, 0x0d, 0x0a, 0x09,
	0x6b, 0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x73, 0x31, 0x10, 0x4b, 0x12, 0x0d, 0x0a, 0x09, 0x6b,
	0x69, 0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x73, 0x32, 0x10, 0x4c, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x69,
	0x74, 0x61, 0x6e, 0x6f, 0x5f, 0x73, 0x33, 0x10, 0x4d, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x69, 0x74,
	0x61, 0x6e, 0x6f, 0x5f, 0x73, 0x34, 0x10, 0x4e, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x69, 0x74, 0x61,
	0x6e, 0x6f, 0x5f, 0x73, 0x35, 0x10, 0x4f, 0x12, 0x0d, 0x0a, 0x09, 0x6b, 0x69, 0x74, 0x61, 0x6e,
	0x6f, 0x5f, 0x73, 0x36, 0x10, 0x50, 0x12, 0x0c, 0x0a, 0x08, 0x74, 0x61, 0x6b, 0x61, 0x6f, 0x5f,
	0x73, 0x31, 0x10, 0x5b, 0x12, 0x0c, 0x0a, 0x08, 0x74, 0x61, 0x6b, 0x61, 0x6f, 0x5f, 0x73, 0x32,
	0x10, 0x5c, 0x32, 0x3c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x31, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x32, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x1a,
	0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_statesync_proto_rawDescOnce sync.Once
	file_proto_statesync_proto_rawDescData = file_proto_statesync_proto_rawDesc
)

func file_proto_statesync_proto_rawDescGZIP() []byte {
	file_proto_statesync_proto_rawDescOnce.Do(func() {
		file_proto_statesync_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_statesync_proto_rawDescData)
	})
	return file_proto_statesync_proto_rawDescData
}

var file_proto_statesync_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_statesync_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_statesync_proto_goTypes = []interface{}{
	(RequestSync_State)(0),     // 0: RequestSync.State
	(ResponseSync_Response)(0), // 1: ResponseSync.Response
	(Stations_StationId)(0),    // 2: Stations.StationId
	(*RequestSync)(nil),        // 3: RequestSync
	(*ResponseSync)(nil),       // 4: ResponseSync
	(*Stations)(nil),           // 5: Stations
}
var file_proto_statesync_proto_depIdxs = []int32{
	5, // 0: RequestSync.station:type_name -> Stations
	0, // 1: RequestSync.state:type_name -> RequestSync.State
	1, // 2: ResponseSync.response:type_name -> ResponseSync.Response
	2, // 3: Stations.stationId:type_name -> Stations.StationId
	3, // 4: Control.Command2Internal:input_type -> RequestSync
	4, // 5: Control.Command2Internal:output_type -> ResponseSync
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_statesync_proto_init() }
func file_proto_statesync_proto_init() {
	if File_proto_statesync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_statesync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSync); i {
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
		file_proto_statesync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSync); i {
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
		file_proto_statesync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stations); i {
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
			RawDescriptor: file_proto_statesync_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_statesync_proto_goTypes,
		DependencyIndexes: file_proto_statesync_proto_depIdxs,
		EnumInfos:         file_proto_statesync_proto_enumTypes,
		MessageInfos:      file_proto_statesync_proto_msgTypes,
	}.Build()
	File_proto_statesync_proto = out.File
	file_proto_statesync_proto_rawDesc = nil
	file_proto_statesync_proto_goTypes = nil
	file_proto_statesync_proto_depIdxs = nil
}
