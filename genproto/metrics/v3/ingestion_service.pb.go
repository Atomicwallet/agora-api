// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics/v3/ingestion_service.proto

package metrics

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SubmitResponse_Result int32

const (
	SubmitResponse_OK SubmitResponse_Result = 0
)

var SubmitResponse_Result_name = map[int32]string{
	0: "OK",
}

var SubmitResponse_Result_value = map[string]int32{
	"OK": 0,
}

func (x SubmitResponse_Result) String() string {
	return proto.EnumName(SubmitResponse_Result_name, int32(x))
}

func (SubmitResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_914678a26d61e0c1, []int{1, 0}
}

type SubmitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitRequest) Reset()         { *m = SubmitRequest{} }
func (m *SubmitRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitRequest) ProtoMessage()    {}
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_914678a26d61e0c1, []int{0}
}

func (m *SubmitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRequest.Unmarshal(m, b)
}
func (m *SubmitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRequest.Marshal(b, m, deterministic)
}
func (m *SubmitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRequest.Merge(m, src)
}
func (m *SubmitRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitRequest.Size(m)
}
func (m *SubmitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRequest proto.InternalMessageInfo

type SubmitResponse struct {
	Result               SubmitResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.metrics.v3.SubmitResponse_Result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SubmitResponse) Reset()         { *m = SubmitResponse{} }
func (m *SubmitResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitResponse) ProtoMessage()    {}
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_914678a26d61e0c1, []int{1}
}

func (m *SubmitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponse.Unmarshal(m, b)
}
func (m *SubmitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponse.Marshal(b, m, deterministic)
}
func (m *SubmitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponse.Merge(m, src)
}
func (m *SubmitResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitResponse.Size(m)
}
func (m *SubmitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponse proto.InternalMessageInfo

func (m *SubmitResponse) GetResult() SubmitResponse_Result {
	if m != nil {
		return m.Result
	}
	return SubmitResponse_OK
}

func init() {
	proto.RegisterEnum("kin.agora.metrics.v3.SubmitResponse_Result", SubmitResponse_Result_name, SubmitResponse_Result_value)
	proto.RegisterType((*SubmitRequest)(nil), "kin.agora.metrics.v3.SubmitRequest")
	proto.RegisterType((*SubmitResponse)(nil), "kin.agora.metrics.v3.SubmitResponse")
}

func init() { proto.RegisterFile("metrics/v3/ingestion_service.proto", fileDescriptor_914678a26d61e0c1) }

var fileDescriptor_914678a26d61e0c1 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xad, 0x87, 0x80, 0x03, 0xd6, 0xb2, 0x78, 0x10, 0xf1, 0x20, 0xd1, 0x83, 0x20, 0xdd,
	0x85, 0xe6, 0xe8, 0x4d, 0x41, 0x10, 0x0f, 0x42, 0x7a, 0xf3, 0xa2, 0x49, 0x18, 0xd7, 0xa1, 0xcd,
	0x6c, 0xdc, 0x99, 0x04, 0xfc, 0xf7, 0xc2, 0x36, 0xc5, 0x0a, 0x22, 0xde, 0x96, 0xd9, 0xef, 0x7b,
	0x3c, 0x1e, 0xe4, 0x2d, 0x6a, 0xa4, 0x46, 0xdc, 0x50, 0x38, 0x62, 0x8f, 0xa2, 0x14, 0xf8, 0x45,
	0x30, 0x0e, 0xd4, 0xa0, 0xed, 0x62, 0xd0, 0x60, 0x8e, 0x57, 0xc4, 0xb6, 0xf2, 0x21, 0x56, 0x76,
	0xa4, 0xed, 0x50, 0xe4, 0x47, 0x70, 0xb8, 0xec, 0xeb, 0x96, 0xb4, 0xc4, 0x8f, 0x1e, 0x45, 0x73,
	0x0f, 0xd3, 0xed, 0x41, 0xba, 0xc0, 0x82, 0xe6, 0x0e, 0xb2, 0x88, 0xd2, 0xaf, 0xf5, 0x64, 0x72,
	0x3e, 0xb9, 0x9a, 0x2e, 0xae, 0xed, 0x6f, 0x49, 0xf6, 0xa7, 0x65, 0xcb, 0xa4, 0x94, 0xa3, 0x9a,
	0xcf, 0x20, 0xdb, 0x5c, 0x4c, 0x06, 0xfb, 0x4f, 0x8f, 0xb3, 0xbd, 0xc5, 0x2b, 0x1c, 0x3c, 0x6c,
	0xab, 0x9a, 0x25, 0x64, 0x1b, 0xdf, 0x5c, 0xfc, 0x9d, 0x9e, 0x4a, 0x9e, 0x5e, 0xfe, 0xa7, 0xc2,
	0xed, 0x1b, 0x9c, 0x85, 0xe8, 0x77, 0x50, 0x8f, 0xbc, 0x83, 0x3f, 0xdf, 0x7b, 0xd2, 0xf7, 0xbe,
	0xb6, 0x4d, 0x68, 0xdd, 0x8a, 0x18, 0x9b, 0x20, 0x9f, 0xa2, 0xd8, 0xba, 0x44, 0xcf, 0xab, 0x8e,
	0xe6, 0xc4, 0x8a, 0x91, 0xab, 0xb5, 0xf3, 0xc8, 0x69, 0x46, 0xf7, 0xbd, 0xf4, 0xcd, 0xf8, 0xac,
	0xb3, 0xf4, 0x53, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x64, 0x2f, 0xb7, 0xc9, 0x86, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IngestionClient is the client API for Ingestion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IngestionClient interface {
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
}

type ingestionClient struct {
	cc *grpc.ClientConn
}

func NewIngestionClient(cc *grpc.ClientConn) IngestionClient {
	return &ingestionClient{cc}
}

func (c *ingestionClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.metrics.v3.Ingestion/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngestionServer is the server API for Ingestion service.
type IngestionServer interface {
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
}

// UnimplementedIngestionServer can be embedded to have forward compatible implementations.
type UnimplementedIngestionServer struct {
}

func (*UnimplementedIngestionServer) Submit(ctx context.Context, req *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}

func RegisterIngestionServer(s *grpc.Server, srv IngestionServer) {
	s.RegisterService(&_Ingestion_serviceDesc, srv)
}

func _Ingestion_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngestionServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.agora.metrics.v3.Ingestion/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngestionServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ingestion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kin.agora.metrics.v3.Ingestion",
	HandlerType: (*IngestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Ingestion_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics/v3/ingestion_service.proto",
}
