// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile/Profile.proto

package v2

import (
	context "context"
	fmt "fmt"
	v21 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v2"
	v2 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/language/agent/v2"
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

type ProfileTaskCommandQuery struct {
	// current sniffer information
	ServiceId  int32 `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	InstanceId int32 `protobuf:"varint,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	// last command timestamp
	LastCommandTime      int64    `protobuf:"varint,3,opt,name=lastCommandTime,proto3" json:"lastCommandTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileTaskCommandQuery) Reset()         { *m = ProfileTaskCommandQuery{} }
func (m *ProfileTaskCommandQuery) String() string { return proto.CompactTextString(m) }
func (*ProfileTaskCommandQuery) ProtoMessage()    {}
func (*ProfileTaskCommandQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{0}
}

func (m *ProfileTaskCommandQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileTaskCommandQuery.Unmarshal(m, b)
}
func (m *ProfileTaskCommandQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileTaskCommandQuery.Marshal(b, m, deterministic)
}
func (m *ProfileTaskCommandQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileTaskCommandQuery.Merge(m, src)
}
func (m *ProfileTaskCommandQuery) XXX_Size() int {
	return xxx_messageInfo_ProfileTaskCommandQuery.Size(m)
}
func (m *ProfileTaskCommandQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileTaskCommandQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileTaskCommandQuery proto.InternalMessageInfo

func (m *ProfileTaskCommandQuery) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ProfileTaskCommandQuery) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *ProfileTaskCommandQuery) GetLastCommandTime() int64 {
	if m != nil {
		return m.LastCommandTime
	}
	return 0
}

// dumped thread snapshot
type ThreadSnapshot struct {
	// profile task id
	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	// dumped segment id
	TraceSegmentId *v2.UniqueId `protobuf:"bytes,2,opt,name=traceSegmentId,proto3" json:"traceSegmentId,omitempty"`
	// dump timestamp
	Time int64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	// snapshot dump sequence, start with zero
	Sequence int32 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// snapshot stack
	Stack                *ThreadStack `protobuf:"bytes,5,opt,name=stack,proto3" json:"stack,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ThreadSnapshot) Reset()         { *m = ThreadSnapshot{} }
func (m *ThreadSnapshot) String() string { return proto.CompactTextString(m) }
func (*ThreadSnapshot) ProtoMessage()    {}
func (*ThreadSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{1}
}

func (m *ThreadSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadSnapshot.Unmarshal(m, b)
}
func (m *ThreadSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadSnapshot.Marshal(b, m, deterministic)
}
func (m *ThreadSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadSnapshot.Merge(m, src)
}
func (m *ThreadSnapshot) XXX_Size() int {
	return xxx_messageInfo_ThreadSnapshot.Size(m)
}
func (m *ThreadSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadSnapshot proto.InternalMessageInfo

func (m *ThreadSnapshot) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ThreadSnapshot) GetTraceSegmentId() *v2.UniqueId {
	if m != nil {
		return m.TraceSegmentId
	}
	return nil
}

func (m *ThreadSnapshot) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *ThreadSnapshot) GetSequence() int32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ThreadSnapshot) GetStack() *ThreadStack {
	if m != nil {
		return m.Stack
	}
	return nil
}

type ThreadStack struct {
	// stack code signature list
	CodeSignatures       []string `protobuf:"bytes,1,rep,name=codeSignatures,proto3" json:"codeSignatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadStack) Reset()         { *m = ThreadStack{} }
func (m *ThreadStack) String() string { return proto.CompactTextString(m) }
func (*ThreadStack) ProtoMessage()    {}
func (*ThreadStack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{2}
}

func (m *ThreadStack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadStack.Unmarshal(m, b)
}
func (m *ThreadStack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadStack.Marshal(b, m, deterministic)
}
func (m *ThreadStack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadStack.Merge(m, src)
}
func (m *ThreadStack) XXX_Size() int {
	return xxx_messageInfo_ThreadStack.Size(m)
}
func (m *ThreadStack) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadStack.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadStack proto.InternalMessageInfo

func (m *ThreadStack) GetCodeSignatures() []string {
	if m != nil {
		return m.CodeSignatures
	}
	return nil
}

// profile task finished report
type ProfileTaskFinishReport struct {
	// current sniffer information
	ServiceId  int32 `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	InstanceId int32 `protobuf:"varint,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	// profile task
	TaskId               string   `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileTaskFinishReport) Reset()         { *m = ProfileTaskFinishReport{} }
func (m *ProfileTaskFinishReport) String() string { return proto.CompactTextString(m) }
func (*ProfileTaskFinishReport) ProtoMessage()    {}
func (*ProfileTaskFinishReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed871c17f6550fe9, []int{3}
}

func (m *ProfileTaskFinishReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileTaskFinishReport.Unmarshal(m, b)
}
func (m *ProfileTaskFinishReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileTaskFinishReport.Marshal(b, m, deterministic)
}
func (m *ProfileTaskFinishReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileTaskFinishReport.Merge(m, src)
}
func (m *ProfileTaskFinishReport) XXX_Size() int {
	return xxx_messageInfo_ProfileTaskFinishReport.Size(m)
}
func (m *ProfileTaskFinishReport) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileTaskFinishReport.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileTaskFinishReport proto.InternalMessageInfo

func (m *ProfileTaskFinishReport) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ProfileTaskFinishReport) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *ProfileTaskFinishReport) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterType((*ProfileTaskCommandQuery)(nil), "skywalking.network.protocol.agent.profile.ProfileTaskCommandQuery")
	proto.RegisterType((*ThreadSnapshot)(nil), "skywalking.network.protocol.agent.profile.ThreadSnapshot")
	proto.RegisterType((*ThreadStack)(nil), "skywalking.network.protocol.agent.profile.ThreadStack")
	proto.RegisterType((*ProfileTaskFinishReport)(nil), "skywalking.network.protocol.agent.profile.ProfileTaskFinishReport")
}

func init() { proto.RegisterFile("profile/Profile.proto", fileDescriptor_ed871c17f6550fe9) }

var fileDescriptor_ed871c17f6550fe9 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x76, 0x36, 0xdb, 0xc5, 0xbe, 0x42, 0x57, 0x46, 0x5c, 0x63, 0x10, 0x29, 0x39, 0x2c, 0x11,
	0x34, 0x81, 0x8a, 0x82, 0xd7, 0x15, 0x84, 0x05, 0x91, 0x9a, 0x56, 0x04, 0x6f, 0xe3, 0xe4, 0x99,
	0x86, 0x24, 0x33, 0xd9, 0x99, 0xc9, 0x2e, 0x45, 0x10, 0xbc, 0x78, 0xf2, 0xdf, 0xf0, 0xe2, 0xdf,
	0xe8, 0x41, 0x4c, 0xa6, 0x4d, 0xac, 0x28, 0xad, 0x9e, 0x26, 0xf3, 0xbd, 0x1f, 0xdf, 0x7b, 0xdf,
	0x37, 0x04, 0x6e, 0x55, 0x4a, 0xbe, 0xcf, 0x0a, 0x8c, 0x66, 0xed, 0x19, 0x56, 0x4a, 0x1a, 0x49,
	0xef, 0xeb, 0x7c, 0x75, 0xc5, 0x8a, 0x3c, 0x13, 0x69, 0x28, 0xd0, 0x5c, 0x49, 0x95, 0xb7, 0x11,
	0x2e, 0x8b, 0x90, 0xa5, 0x28, 0x4c, 0x68, 0x0b, 0xbd, 0x9b, 0x5c, 0x96, 0xa5, 0x14, 0x51, 0x7b,
	0xb4, 0x59, 0xde, 0x1d, 0x0b, 0x1a, 0xc5, 0x38, 0x3e, 0xec, 0x87, 0xfc, 0x4f, 0x04, 0x6e, 0x5b,
	0xb2, 0x05, 0xd3, 0xf9, 0x33, 0x59, 0x96, 0x4c, 0x24, 0xaf, 0x6a, 0x54, 0x2b, 0x7a, 0x17, 0x86,
	0x1a, 0xd5, 0x65, 0xc6, 0xf1, 0x3c, 0x71, 0xc9, 0x84, 0x04, 0x83, 0xb8, 0x03, 0xe8, 0x3d, 0x80,
	0x4c, 0x68, 0xc3, 0x44, 0x13, 0x3e, 0x68, 0xc2, 0x3d, 0x84, 0x06, 0x70, 0x5c, 0x30, 0x6d, 0x6c,
	0xc7, 0x45, 0x56, 0xa2, 0xeb, 0x4c, 0x48, 0xe0, 0xc4, 0xdb, 0xb0, 0xff, 0x9d, 0xc0, 0x78, 0xb1,
	0x54, 0xc8, 0x92, 0xb9, 0x60, 0x95, 0x5e, 0x4a, 0x43, 0x4f, 0xe0, 0xc8, 0x30, 0x9d, 0x5b, 0xde,
	0x61, 0x6c, 0x6f, 0x74, 0x01, 0xe3, 0x66, 0x89, 0x39, 0xa6, 0x25, 0x0a, 0x63, 0x89, 0x47, 0xd3,
	0x07, 0xe1, 0xdf, 0x24, 0xb2, 0x1b, 0xbf, 0x16, 0xd9, 0x45, 0x8d, 0xe7, 0x49, 0xbc, 0xd5, 0x83,
	0x52, 0x38, 0x34, 0xdd, 0x7c, 0xcd, 0x37, 0xf5, 0xe0, 0xba, 0xc6, 0x8b, 0x1a, 0x05, 0x47, 0xf7,
	0xb0, 0x59, 0x6e, 0x73, 0xa7, 0x2f, 0x60, 0xa0, 0x0d, 0xe3, 0xb9, 0x3b, 0x68, 0xc8, 0x9f, 0x84,
	0x3b, 0xfb, 0x13, 0xda, 0x3d, 0x7f, 0x56, 0xc7, 0x6d, 0x13, 0xff, 0x31, 0x8c, 0x7a, 0x28, 0x3d,
	0x85, 0x31, 0x97, 0x09, 0xce, 0xb3, 0x54, 0x30, 0x53, 0x2b, 0xd4, 0x2e, 0x99, 0x38, 0xc1, 0x30,
	0xde, 0x42, 0x7d, 0xf9, 0x8b, 0x71, 0xcf, 0x33, 0x91, 0xe9, 0x65, 0x8c, 0x95, 0x54, 0xe6, 0x3f,
	0x8d, 0xeb, 0xb4, 0x77, 0xfa, 0xda, 0x4f, 0xbf, 0x3a, 0x30, 0xea, 0x31, 0xd2, 0x2f, 0x04, 0x4e,
	0x52, 0x34, 0xbf, 0xbf, 0x1e, 0x4d, 0xcf, 0xf6, 0x50, 0xe4, 0x0f, 0xaf, 0xcf, 0xdb, 0xc9, 0xd2,
	0x35, 0xa3, 0x7f, 0x8d, 0x7e, 0x84, 0x63, 0x2e, 0x8b, 0x02, 0xb9, 0xd9, 0xbc, 0xa2, 0xa7, 0xfb,
	0x1b, 0x63, 0x4b, 0xf7, 0x65, 0x0f, 0x08, 0xfd, 0x4c, 0xe0, 0x86, 0x6a, 0xf4, 0xef, 0xfc, 0xf8,
	0x57, 0x21, 0xfa, 0x6e, 0xee, 0x3b, 0xca, 0xd9, 0x07, 0x98, 0x4a, 0x95, 0x86, 0xac, 0x62, 0x7c,
	0x89, 0xfd, 0x5a, 0x56, 0x95, 0x9b, 0xfa, 0x82, 0x89, 0xb4, 0x66, 0x29, 0xae, 0xd9, 0x67, 0xe4,
	0xed, 0x69, 0x97, 0x1a, 0xd9, 0xb4, 0x68, 0x9d, 0x16, 0xad, 0x7f, 0x4c, 0x97, 0xd3, 0x6f, 0x07,
	0xde, 0x3c, 0x5f, 0xbd, 0xb1, 0x3d, 0x5f, 0xb6, 0x89, 0x33, 0x3b, 0xce, 0xbb, 0xa3, 0x66, 0xb0,
	0x47, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x42, 0x80, 0xd5, 0xca, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfileTaskClient is the client API for ProfileTask service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfileTaskClient interface {
	// query all sniffer need to execute profile task commands
	GetProfileTaskCommands(ctx context.Context, in *ProfileTaskCommandQuery, opts ...grpc.CallOption) (*v21.Commands, error)
	// collect dumped thread snapshot
	CollectSnapshot(ctx context.Context, opts ...grpc.CallOption) (ProfileTask_CollectSnapshotClient, error)
	// report profiling task finished
	ReportTaskFinish(ctx context.Context, in *ProfileTaskFinishReport, opts ...grpc.CallOption) (*v21.Commands, error)
}

type profileTaskClient struct {
	cc *grpc.ClientConn
}

func NewProfileTaskClient(cc *grpc.ClientConn) ProfileTaskClient {
	return &profileTaskClient{cc}
}

func (c *profileTaskClient) GetProfileTaskCommands(ctx context.Context, in *ProfileTaskCommandQuery, opts ...grpc.CallOption) (*v21.Commands, error) {
	out := new(v21.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.network.protocol.agent.profile.ProfileTask/getProfileTaskCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileTaskClient) CollectSnapshot(ctx context.Context, opts ...grpc.CallOption) (ProfileTask_CollectSnapshotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProfileTask_serviceDesc.Streams[0], "/skywalking.network.protocol.agent.profile.ProfileTask/collectSnapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &profileTaskCollectSnapshotClient{stream}
	return x, nil
}

type ProfileTask_CollectSnapshotClient interface {
	Send(*ThreadSnapshot) error
	CloseAndRecv() (*v21.Commands, error)
	grpc.ClientStream
}

type profileTaskCollectSnapshotClient struct {
	grpc.ClientStream
}

func (x *profileTaskCollectSnapshotClient) Send(m *ThreadSnapshot) error {
	return x.ClientStream.SendMsg(m)
}

func (x *profileTaskCollectSnapshotClient) CloseAndRecv() (*v21.Commands, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(v21.Commands)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *profileTaskClient) ReportTaskFinish(ctx context.Context, in *ProfileTaskFinishReport, opts ...grpc.CallOption) (*v21.Commands, error) {
	out := new(v21.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.network.protocol.agent.profile.ProfileTask/reportTaskFinish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileTaskServer is the server API for ProfileTask service.
type ProfileTaskServer interface {
	// query all sniffer need to execute profile task commands
	GetProfileTaskCommands(context.Context, *ProfileTaskCommandQuery) (*v21.Commands, error)
	// collect dumped thread snapshot
	CollectSnapshot(ProfileTask_CollectSnapshotServer) error
	// report profiling task finished
	ReportTaskFinish(context.Context, *ProfileTaskFinishReport) (*v21.Commands, error)
}

// UnimplementedProfileTaskServer can be embedded to have forward compatible implementations.
type UnimplementedProfileTaskServer struct {
}

func (*UnimplementedProfileTaskServer) GetProfileTaskCommands(ctx context.Context, req *ProfileTaskCommandQuery) (*v21.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileTaskCommands not implemented")
}
func (*UnimplementedProfileTaskServer) CollectSnapshot(srv ProfileTask_CollectSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method CollectSnapshot not implemented")
}
func (*UnimplementedProfileTaskServer) ReportTaskFinish(ctx context.Context, req *ProfileTaskFinishReport) (*v21.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTaskFinish not implemented")
}

func RegisterProfileTaskServer(s *grpc.Server, srv ProfileTaskServer) {
	s.RegisterService(&_ProfileTask_serviceDesc, srv)
}

func _ProfileTask_GetProfileTaskCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileTaskCommandQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileTaskServer).GetProfileTaskCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.network.protocol.agent.profile.ProfileTask/GetProfileTaskCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileTaskServer).GetProfileTaskCommands(ctx, req.(*ProfileTaskCommandQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileTask_CollectSnapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProfileTaskServer).CollectSnapshot(&profileTaskCollectSnapshotServer{stream})
}

type ProfileTask_CollectSnapshotServer interface {
	SendAndClose(*v21.Commands) error
	Recv() (*ThreadSnapshot, error)
	grpc.ServerStream
}

type profileTaskCollectSnapshotServer struct {
	grpc.ServerStream
}

func (x *profileTaskCollectSnapshotServer) SendAndClose(m *v21.Commands) error {
	return x.ServerStream.SendMsg(m)
}

func (x *profileTaskCollectSnapshotServer) Recv() (*ThreadSnapshot, error) {
	m := new(ThreadSnapshot)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProfileTask_ReportTaskFinish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileTaskFinishReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileTaskServer).ReportTaskFinish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.network.protocol.agent.profile.ProfileTask/ReportTaskFinish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileTaskServer).ReportTaskFinish(ctx, req.(*ProfileTaskFinishReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.network.protocol.agent.profile.ProfileTask",
	HandlerType: (*ProfileTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProfileTaskCommands",
			Handler:    _ProfileTask_GetProfileTaskCommands_Handler,
		},
		{
			MethodName: "reportTaskFinish",
			Handler:    _ProfileTask_ReportTaskFinish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "collectSnapshot",
			Handler:       _ProfileTask_CollectSnapshot_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "profile/Profile.proto",
}