// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/session_service.proto

package tensorflow_serving

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "github.com/jadekler/tensorflow-serving-client/v2/go/tensorflow/core/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SessionRunRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Tensors to be fed in the step. Each feed is a named tensor.
	Feed []*protobuf.NamedTensorProto `protobuf:"bytes,2,rep,name=feed,proto3" json:"feed,omitempty"`
	// Fetches. A list of tensor names. The caller expects a tensor to
	// be returned for each fetch[i] (see RunResponse.tensor). The
	// order of specified fetches does not change the execution order.
	Fetch []string `protobuf:"bytes,3,rep,name=fetch,proto3" json:"fetch,omitempty"`
	// Target Nodes. A list of node names. The named nodes will be run
	// to but their outputs will not be fetched.
	Target []string `protobuf:"bytes,4,rep,name=target,proto3" json:"target,omitempty"`
	// Options for the run call. **Currently ignored.**
	Options              *protobuf.RunOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionRunRequest) Reset()         { *m = SessionRunRequest{} }
func (m *SessionRunRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRunRequest) ProtoMessage()    {}
func (*SessionRunRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6b0c7cbabd9081a, []int{0}
}

func (m *SessionRunRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRunRequest.Unmarshal(m, b)
}
func (m *SessionRunRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRunRequest.Marshal(b, m, deterministic)
}
func (m *SessionRunRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRunRequest.Merge(m, src)
}
func (m *SessionRunRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRunRequest.Size(m)
}
func (m *SessionRunRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRunRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRunRequest proto.InternalMessageInfo

func (m *SessionRunRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *SessionRunRequest) GetFeed() []*protobuf.NamedTensorProto {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *SessionRunRequest) GetFetch() []string {
	if m != nil {
		return m.Fetch
	}
	return nil
}

func (m *SessionRunRequest) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *SessionRunRequest) GetOptions() *protobuf.RunOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type SessionRunResponse struct {
	// Effective Model Specification used for session run.
	ModelSpec *ModelSpec `protobuf:"bytes,3,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// NOTE: The order of the returned tensors may or may not match
	// the fetch order specified in RunRequest.
	Tensor []*protobuf.NamedTensorProto `protobuf:"bytes,1,rep,name=tensor,proto3" json:"tensor,omitempty"`
	// Returned metadata if requested in the options.
	Metadata             *protobuf.RunMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SessionRunResponse) Reset()         { *m = SessionRunResponse{} }
func (m *SessionRunResponse) String() string { return proto.CompactTextString(m) }
func (*SessionRunResponse) ProtoMessage()    {}
func (*SessionRunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6b0c7cbabd9081a, []int{1}
}

func (m *SessionRunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRunResponse.Unmarshal(m, b)
}
func (m *SessionRunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRunResponse.Marshal(b, m, deterministic)
}
func (m *SessionRunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRunResponse.Merge(m, src)
}
func (m *SessionRunResponse) XXX_Size() int {
	return xxx_messageInfo_SessionRunResponse.Size(m)
}
func (m *SessionRunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRunResponse proto.InternalMessageInfo

func (m *SessionRunResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *SessionRunResponse) GetTensor() []*protobuf.NamedTensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *SessionRunResponse) GetMetadata() *protobuf.RunMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionRunRequest)(nil), "tensorflow.serving.SessionRunRequest")
	proto.RegisterType((*SessionRunResponse)(nil), "tensorflow.serving.SessionRunResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/session_service.proto", fileDescriptor_b6b0c7cbabd9081a)
}

var fileDescriptor_b6b0c7cbabd9081a = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xb3, 0x14, 0xf8, 0xff, 0x19, 0x12, 0x13, 0x37, 0x06, 0x37, 0x44, 0x13, 0x82, 0xc1,
	0x90, 0x18, 0x5b, 0x02, 0x1e, 0x3d, 0x79, 0x47, 0xcd, 0xe2, 0xcd, 0x03, 0x29, 0xed, 0x14, 0x9b,
	0xd0, 0xdd, 0xda, 0xdd, 0xea, 0xd3, 0xf9, 0x34, 0xbe, 0x84, 0x47, 0xc3, 0xee, 0x22, 0x04, 0x24,
	0xea, 0xad, 0x33, 0xf3, 0x9b, 0x7e, 0xdf, 0xd7, 0x0e, 0x5c, 0x6a, 0x14, 0x4a, 0x16, 0xc9, 0x42,
	0xbe, 0x4e, 0x15, 0x16, 0x2f, 0xa9, 0x98, 0x07, 0x61, 0x9e, 0xaa, 0x40, 0xa1, 0x52, 0xa9, 0x14,
	0xb6, 0x19, 0xa1, 0x9f, 0x17, 0x52, 0x4b, 0x4a, 0xd7, 0xb8, 0xef, 0xf0, 0xf6, 0xd9, 0xbe, 0x57,
	0x64, 0x32, 0xc6, 0x85, 0x5d, 0x6c, 0xf7, 0xd6, 0x50, 0x10, 0xc9, 0x02, 0x03, 0xd3, 0x9e, 0x95,
	0x49, 0x10, 0x49, 0x91, 0xa4, 0x73, 0x87, 0x5d, 0xec, 0xc5, 0x44, 0x98, 0x61, 0x3c, 0xb5, 0x63,
	0x0b, 0x77, 0xdf, 0x09, 0x1c, 0x4e, 0xac, 0x4d, 0x5e, 0x0a, 0x8e, 0xcf, 0x25, 0x2a, 0x4d, 0xaf,
	0x01, 0x8c, 0xf0, 0x54, 0xe5, 0x18, 0x31, 0xd2, 0x21, 0xfd, 0xe6, 0xf0, 0xd4, 0xdf, 0xf5, 0xed,
	0x8f, 0x97, 0xd4, 0x24, 0xc7, 0x88, 0x37, 0xb2, 0xd5, 0x23, 0x1d, 0x40, 0x35, 0x41, 0x8c, 0x59,
	0xa5, 0xe3, 0xf5, 0x9b, 0xc3, 0x93, 0xcd, 0xbd, 0xdb, 0xa5, 0x83, 0x07, 0x53, 0xdf, 0x2f, 0xf5,
	0xb9, 0x21, 0xe9, 0x11, 0xd4, 0x12, 0xd4, 0xd1, 0x13, 0xf3, 0x3a, 0x5e, 0xbf, 0xc1, 0x6d, 0x41,
	0x5b, 0x50, 0xd7, 0x61, 0x31, 0x47, 0xcd, 0xaa, 0xa6, 0xed, 0x2a, 0x3a, 0x80, 0x7f, 0x32, 0xd7,
	0xa9, 0x14, 0x8a, 0xd5, 0x8c, 0xb5, 0xd6, 0xa6, 0x04, 0x2f, 0xc5, 0x9d, 0x9d, 0xf2, 0x15, 0xd6,
	0x7d, 0x23, 0x40, 0x37, 0x53, 0xaa, 0x5c, 0x0a, 0x85, 0x5b, 0x31, 0xbd, 0x3f, 0xc6, 0xbc, 0x82,
	0xba, 0x45, 0x19, 0xf9, 0x45, 0x50, 0xc7, 0xd2, 0x11, 0xfc, 0xcf, 0x50, 0x87, 0x71, 0xa8, 0x43,
	0x56, 0x31, 0x8a, 0xc7, 0x5b, 0xee, 0xc7, 0x6e, 0xcc, 0xbf, 0xc0, 0x61, 0x06, 0x07, 0xce, 0xfe,
	0xc4, 0x9e, 0x12, 0x7d, 0x04, 0x58, 0x07, 0xa2, 0xbd, 0xef, 0x4c, 0xef, 0xfc, 0xd6, 0xf6, 0xf9,
	0x4f, 0x98, 0xfd, 0x2e, 0x37, 0xde, 0x07, 0x21, 0xb3, 0xba, 0x39, 0x90, 0xd1, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0x1e, 0x34, 0x3d, 0xde, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionServiceClient interface {
	// Runs inference of a given model.
	SessionRun(ctx context.Context, in *SessionRunRequest, opts ...grpc.CallOption) (*SessionRunResponse, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) SessionRun(ctx context.Context, in *SessionRunRequest, opts ...grpc.CallOption) (*SessionRunResponse, error) {
	out := new(SessionRunResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.SessionService/SessionRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
type SessionServiceServer interface {
	// Runs inference of a given model.
	SessionRun(context.Context, *SessionRunRequest) (*SessionRunResponse, error)
}

// UnimplementedSessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (*UnimplementedSessionServiceServer) SessionRun(ctx context.Context, req *SessionRunRequest) (*SessionRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionRun not implemented")
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_SessionRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SessionRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.SessionService/SessionRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SessionRun(ctx, req.(*SessionRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SessionRun",
			Handler:    _SessionService_SessionRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/session_service.proto",
}
