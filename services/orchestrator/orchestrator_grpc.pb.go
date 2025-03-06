// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: orchestrator.proto

package orchestrator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Orchestrator_GetRunnerInfo_FullMethodName         = "/Orchestrator/GetRunnerInfo"
	Orchestrator_GetNewTask_FullMethodName            = "/Orchestrator/GetNewTask"
	Orchestrator_CloseTask_FullMethodName             = "/Orchestrator/CloseTask"
	Orchestrator_GetWorksOfEvent_FullMethodName       = "/Orchestrator/GetWorksOfEvent"
	Orchestrator_GetWorksDownloadLinks_FullMethodName = "/Orchestrator/GetWorksDownloadLinks"
	Orchestrator_SendCrossCheckReport_FullMethodName  = "/Orchestrator/SendCrossCheckReport"
	Orchestrator_SendDefaultReport_FullMethodName     = "/Orchestrator/SendDefaultReport"
)

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorClient interface {
	GetRunnerInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRunnerInfoResponse, error)
	GetNewTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNewTaskResponse, error)
	CloseTask(ctx context.Context, in *CloseTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetWorksOfEvent(ctx context.Context, in *GetWorksOfEventRequest, opts ...grpc.CallOption) (*GetWorksOfEventResponse, error)
	GetWorksDownloadLinks(ctx context.Context, in *GetWorksDownloadLinksRequest, opts ...grpc.CallOption) (*GetWorksDownloadLinksResponse, error)
	SendCrossCheckReport(ctx context.Context, in *SendCrossCheckReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendDefaultReport(ctx context.Context, in *SendDefaultReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) GetRunnerInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRunnerInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRunnerInfoResponse)
	err := c.cc.Invoke(ctx, Orchestrator_GetRunnerInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetNewTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetNewTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewTaskResponse)
	err := c.cc.Invoke(ctx, Orchestrator_GetNewTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) CloseTask(ctx context.Context, in *CloseTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Orchestrator_CloseTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetWorksOfEvent(ctx context.Context, in *GetWorksOfEventRequest, opts ...grpc.CallOption) (*GetWorksOfEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorksOfEventResponse)
	err := c.cc.Invoke(ctx, Orchestrator_GetWorksOfEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetWorksDownloadLinks(ctx context.Context, in *GetWorksDownloadLinksRequest, opts ...grpc.CallOption) (*GetWorksDownloadLinksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorksDownloadLinksResponse)
	err := c.cc.Invoke(ctx, Orchestrator_GetWorksDownloadLinks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) SendCrossCheckReport(ctx context.Context, in *SendCrossCheckReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Orchestrator_SendCrossCheckReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) SendDefaultReport(ctx context.Context, in *SendDefaultReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Orchestrator_SendDefaultReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
// All implementations must embed UnimplementedOrchestratorServer
// for forward compatibility.
type OrchestratorServer interface {
	GetRunnerInfo(context.Context, *emptypb.Empty) (*GetRunnerInfoResponse, error)
	GetNewTask(context.Context, *emptypb.Empty) (*GetNewTaskResponse, error)
	CloseTask(context.Context, *CloseTaskRequest) (*emptypb.Empty, error)
	GetWorksOfEvent(context.Context, *GetWorksOfEventRequest) (*GetWorksOfEventResponse, error)
	GetWorksDownloadLinks(context.Context, *GetWorksDownloadLinksRequest) (*GetWorksDownloadLinksResponse, error)
	SendCrossCheckReport(context.Context, *SendCrossCheckReportRequest) (*emptypb.Empty, error)
	SendDefaultReport(context.Context, *SendDefaultReportRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrchestratorServer()
}

// UnimplementedOrchestratorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrchestratorServer struct{}

func (UnimplementedOrchestratorServer) GetRunnerInfo(context.Context, *emptypb.Empty) (*GetRunnerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunnerInfo not implemented")
}
func (UnimplementedOrchestratorServer) GetNewTask(context.Context, *emptypb.Empty) (*GetNewTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewTask not implemented")
}
func (UnimplementedOrchestratorServer) CloseTask(context.Context, *CloseTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseTask not implemented")
}
func (UnimplementedOrchestratorServer) GetWorksOfEvent(context.Context, *GetWorksOfEventRequest) (*GetWorksOfEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorksOfEvent not implemented")
}
func (UnimplementedOrchestratorServer) GetWorksDownloadLinks(context.Context, *GetWorksDownloadLinksRequest) (*GetWorksDownloadLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorksDownloadLinks not implemented")
}
func (UnimplementedOrchestratorServer) SendCrossCheckReport(context.Context, *SendCrossCheckReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCrossCheckReport not implemented")
}
func (UnimplementedOrchestratorServer) SendDefaultReport(context.Context, *SendDefaultReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDefaultReport not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}
func (UnimplementedOrchestratorServer) testEmbeddedByValue()                      {}

// UnsafeOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServer will
// result in compilation errors.
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	// If the following call pancis, it indicates UnimplementedOrchestratorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_GetRunnerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetRunnerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_GetRunnerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetRunnerInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetNewTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetNewTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_GetNewTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetNewTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_CloseTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).CloseTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_CloseTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).CloseTask(ctx, req.(*CloseTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetWorksOfEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorksOfEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetWorksOfEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_GetWorksOfEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetWorksOfEvent(ctx, req.(*GetWorksOfEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetWorksDownloadLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorksDownloadLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetWorksDownloadLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_GetWorksDownloadLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetWorksDownloadLinks(ctx, req.(*GetWorksDownloadLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_SendCrossCheckReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCrossCheckReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).SendCrossCheckReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_SendCrossCheckReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).SendCrossCheckReport(ctx, req.(*SendCrossCheckReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_SendDefaultReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDefaultReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).SendDefaultReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_SendDefaultReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).SendDefaultReport(ctx, req.(*SendDefaultReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orchestrator_ServiceDesc is the grpc.ServiceDesc for Orchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRunnerInfo",
			Handler:    _Orchestrator_GetRunnerInfo_Handler,
		},
		{
			MethodName: "GetNewTask",
			Handler:    _Orchestrator_GetNewTask_Handler,
		},
		{
			MethodName: "CloseTask",
			Handler:    _Orchestrator_CloseTask_Handler,
		},
		{
			MethodName: "GetWorksOfEvent",
			Handler:    _Orchestrator_GetWorksOfEvent_Handler,
		},
		{
			MethodName: "GetWorksDownloadLinks",
			Handler:    _Orchestrator_GetWorksDownloadLinks_Handler,
		},
		{
			MethodName: "SendCrossCheckReport",
			Handler:    _Orchestrator_SendCrossCheckReport_Handler,
		},
		{
			MethodName: "SendDefaultReport",
			Handler:    _Orchestrator_SendDefaultReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator.proto",
}
