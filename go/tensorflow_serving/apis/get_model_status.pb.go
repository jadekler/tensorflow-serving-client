// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/get_model_status.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	util "github.com/jadekler/tensorflow-serving-client/v2/go/tensorflow_serving/util"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// States that map to ManagerState enum in
// tensorflow_serving/core/servable_state.h
type ModelVersionStatus_State int32

const (
	// Default value.
	ModelVersionStatus_UNKNOWN ModelVersionStatus_State = 0
	// The manager is tracking this servable, but has not initiated any action
	// pertaining to it.
	ModelVersionStatus_START ModelVersionStatus_State = 10
	// The manager has decided to load this servable. In particular, checks
	// around resource availability and other aspects have passed, and the
	// manager is about to invoke the loader's Load() method.
	ModelVersionStatus_LOADING ModelVersionStatus_State = 20
	// The manager has successfully loaded this servable and made it available
	// for serving (i.e. GetServableHandle(id) will succeed). To avoid races,
	// this state is not reported until *after* the servable is made
	// available.
	ModelVersionStatus_AVAILABLE ModelVersionStatus_State = 30
	// The manager has decided to make this servable unavailable, and unload
	// it. To avoid races, this state is reported *before* the servable is
	// made unavailable.
	ModelVersionStatus_UNLOADING ModelVersionStatus_State = 40
	// This servable has reached the end of its journey in the manager. Either
	// it loaded and ultimately unloaded successfully, or it hit an error at
	// some point in its lifecycle.
	ModelVersionStatus_END ModelVersionStatus_State = 50
)

var ModelVersionStatus_State_name = map[int32]string{
	0:  "UNKNOWN",
	10: "START",
	20: "LOADING",
	30: "AVAILABLE",
	40: "UNLOADING",
	50: "END",
}

var ModelVersionStatus_State_value = map[string]int32{
	"UNKNOWN":   0,
	"START":     10,
	"LOADING":   20,
	"AVAILABLE": 30,
	"UNLOADING": 40,
	"END":       50,
}

func (x ModelVersionStatus_State) String() string {
	return proto.EnumName(ModelVersionStatus_State_name, int32(x))
}

func (ModelVersionStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1efd090da1a85b62, []int{1, 0}
}

// GetModelStatusRequest contains a ModelSpec indicating the model for which
// to get status.
type GetModelStatusRequest struct {
	// Model Specification. If version is not specified, information about all
	// versions of the model will be returned. If a version is specified, the
	// status of only that version will be returned.
	ModelSpec            *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetModelStatusRequest) Reset()         { *m = GetModelStatusRequest{} }
func (m *GetModelStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusRequest) ProtoMessage()    {}
func (*GetModelStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1efd090da1a85b62, []int{0}
}

func (m *GetModelStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusRequest.Unmarshal(m, b)
}
func (m *GetModelStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetModelStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusRequest.Merge(m, src)
}
func (m *GetModelStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusRequest.Size(m)
}
func (m *GetModelStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusRequest proto.InternalMessageInfo

func (m *GetModelStatusRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

// Version number, state, and status for a single version of a model.
type ModelVersionStatus struct {
	// Model version.
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Model state.
	State ModelVersionStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=tensorflow.serving.ModelVersionStatus_State" json:"state,omitempty"`
	// Model status.
	Status               *util.StatusProto `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ModelVersionStatus) Reset()         { *m = ModelVersionStatus{} }
func (m *ModelVersionStatus) String() string { return proto.CompactTextString(m) }
func (*ModelVersionStatus) ProtoMessage()    {}
func (*ModelVersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1efd090da1a85b62, []int{1}
}

func (m *ModelVersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelVersionStatus.Unmarshal(m, b)
}
func (m *ModelVersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelVersionStatus.Marshal(b, m, deterministic)
}
func (m *ModelVersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelVersionStatus.Merge(m, src)
}
func (m *ModelVersionStatus) XXX_Size() int {
	return xxx_messageInfo_ModelVersionStatus.Size(m)
}
func (m *ModelVersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelVersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ModelVersionStatus proto.InternalMessageInfo

func (m *ModelVersionStatus) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ModelVersionStatus) GetState() ModelVersionStatus_State {
	if m != nil {
		return m.State
	}
	return ModelVersionStatus_UNKNOWN
}

func (m *ModelVersionStatus) GetStatus() *util.StatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

// Response for ModelStatusRequest on successful run.
type GetModelStatusResponse struct {
	// Version number and status information for applicable model version(s).
	ModelVersionStatus   []*ModelVersionStatus `protobuf:"bytes,1,rep,name=model_version_status,proto3" json:"model_version_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetModelStatusResponse) Reset()         { *m = GetModelStatusResponse{} }
func (m *GetModelStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusResponse) ProtoMessage()    {}
func (*GetModelStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1efd090da1a85b62, []int{2}
}

func (m *GetModelStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusResponse.Unmarshal(m, b)
}
func (m *GetModelStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetModelStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusResponse.Merge(m, src)
}
func (m *GetModelStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusResponse.Size(m)
}
func (m *GetModelStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusResponse proto.InternalMessageInfo

func (m *GetModelStatusResponse) GetModelVersionStatus() []*ModelVersionStatus {
	if m != nil {
		return m.ModelVersionStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.serving.ModelVersionStatus_State", ModelVersionStatus_State_name, ModelVersionStatus_State_value)
	proto.RegisterType((*GetModelStatusRequest)(nil), "tensorflow.serving.GetModelStatusRequest")
	proto.RegisterType((*ModelVersionStatus)(nil), "tensorflow.serving.ModelVersionStatus")
	proto.RegisterType((*GetModelStatusResponse)(nil), "tensorflow.serving.GetModelStatusResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/get_model_status.proto", fileDescriptor_1efd090da1a85b62)
}

var fileDescriptor_1efd090da1a85b62 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xed, 0xca, 0x36, 0xf6, 0x0e, 0xa5, 0x84, 0x29, 0x65, 0xa0, 0x8e, 0x2a, 0xd2, 0x83,
	0x74, 0x50, 0x0f, 0x5e, 0xbc, 0x74, 0x6c, 0x8c, 0xe1, 0xec, 0x24, 0xfb, 0x10, 0xbc, 0x8c, 0x39,
	0x5f, 0x47, 0x61, 0x6b, 0x6a, 0x93, 0xcd, 0xab, 0x7f, 0xb6, 0x47, 0x49, 0xd2, 0xe1, 0x57, 0x07,
	0x9e, 0xda, 0xbc, 0x79, 0x9e, 0x5f, 0x9e, 0x27, 0x01, 0x4f, 0x60, 0xcc, 0x59, 0xfa, 0xb2, 0x64,
	0x6f, 0x53, 0x8e, 0xe9, 0x26, 0x8a, 0x17, 0xcd, 0x59, 0x12, 0xf1, 0xe6, 0x02, 0xc5, 0x74, 0xc5,
	0x9e, 0x71, 0x39, 0xe5, 0x62, 0x26, 0xd6, 0xdc, 0x4b, 0x52, 0x26, 0x18, 0x21, 0x5f, 0x7a, 0x2f,
	0xd3, 0xd7, 0xcf, 0x76, 0x31, 0x94, 0x5f, 0x1b, 0xeb, 0xe7, 0x39, 0xa2, 0xb5, 0x88, 0x96, 0xcd,
	0xef, 0x78, 0x67, 0x0c, 0x87, 0x5d, 0x14, 0x77, 0xd2, 0x37, 0x54, 0x73, 0x8a, 0xaf, 0x6b, 0xe4,
	0x82, 0xdc, 0x00, 0x64, 0x69, 0x12, 0x9c, 0xdb, 0x46, 0xc3, 0x70, 0xab, 0xfe, 0xb1, 0xf7, 0x37,
	0x8c, 0xa7, 0xbd, 0x09, 0xce, 0x69, 0x65, 0xb5, 0xfd, 0x75, 0xde, 0x0b, 0x40, 0xd4, 0xc6, 0x04,
	0x53, 0x1e, 0xb1, 0x58, 0xb3, 0x89, 0x0d, 0xe5, 0x8d, 0x1e, 0x28, 0xa2, 0x49, 0xb7, 0x4b, 0xd2,
	0x82, 0xa2, 0xcc, 0x85, 0x76, 0xa1, 0x61, 0xb8, 0x07, 0xfe, 0xe5, 0xce, 0x93, 0x7e, 0x00, 0x3d,
	0xf9, 0x41, 0xaa, 0xad, 0xe4, 0x1a, 0x4a, 0xba, 0x9b, 0x6d, 0xaa, 0xb8, 0xa7, 0x79, 0x10, 0x6d,
	0xbc, 0x97, 0xe5, 0x69, 0x26, 0x77, 0x86, 0x50, 0x54, 0x20, 0x52, 0x85, 0xf2, 0x38, 0xbc, 0x0d,
	0x07, 0x0f, 0xa1, 0xb5, 0x47, 0x2a, 0x50, 0x1c, 0x8e, 0x02, 0x3a, 0xb2, 0x40, 0xce, 0xfb, 0x83,
	0xa0, 0xdd, 0x0b, 0xbb, 0x56, 0x8d, 0xec, 0x43, 0x25, 0x98, 0x04, 0xbd, 0x7e, 0xd0, 0xea, 0x77,
	0xac, 0x13, 0xb9, 0x1c, 0x87, 0xdb, 0x5d, 0x97, 0x94, 0xc1, 0xec, 0x84, 0x6d, 0xcb, 0x77, 0x04,
	0x1c, 0xfd, 0xbe, 0x59, 0x9e, 0xb0, 0x98, 0x23, 0x79, 0x84, 0x9a, 0xbe, 0xda, 0xac, 0x7c, 0xf6,
	0xe0, 0xb6, 0xd1, 0x30, 0xdd, 0xaa, 0x7f, 0xf1, 0xbf, 0xea, 0x34, 0x97, 0xd1, 0x32, 0x3f, 0x0c,
	0xe3, 0xa9, 0xa4, 0xde, 0xf6, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x04, 0x57, 0xa9, 0xc7, 0x6c,
	0x02, 0x00, 0x00,
}
