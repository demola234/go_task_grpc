// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: services_todo.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TodoGrpc_CreateTask_FullMethodName = "/pb.TodoGrpc/CreateTask"
	TodoGrpc_GetTask_FullMethodName    = "/pb.TodoGrpc/GetTask"
	TodoGrpc_DeleteTask_FullMethodName = "/pb.TodoGrpc/DeleteTask"
	TodoGrpc_UpdateTask_FullMethodName = "/pb.TodoGrpc/UpdateTask"
)

// TodoGrpcClient is the client API for TodoGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoGrpcClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	GetTask(ctx context.Context, in *EmptyGetRequest, opts ...grpc.CallOption) (*GetAllTaskResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*EmptyTaskRequest, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
}

type todoGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoGrpcClient(cc grpc.ClientConnInterface) TodoGrpcClient {
	return &todoGrpcClient{cc}
}

func (c *todoGrpcClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TodoGrpc_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGrpcClient) GetTask(ctx context.Context, in *EmptyGetRequest, opts ...grpc.CallOption) (*GetAllTaskResponse, error) {
	out := new(GetAllTaskResponse)
	err := c.cc.Invoke(ctx, TodoGrpc_GetTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGrpcClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*EmptyTaskRequest, error) {
	out := new(EmptyTaskRequest)
	err := c.cc.Invoke(ctx, TodoGrpc_DeleteTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoGrpcClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TodoGrpc_UpdateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoGrpcServer is the server API for TodoGrpc service.
// All implementations must embed UnimplementedTodoGrpcServer
// for forward compatibility
type TodoGrpcServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	GetTask(context.Context, *EmptyGetRequest) (*GetAllTaskResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*EmptyTaskRequest, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	mustEmbedUnimplementedTodoGrpcServer()
}

// UnimplementedTodoGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedTodoGrpcServer struct {
}

func (UnimplementedTodoGrpcServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTodoGrpcServer) GetTask(context.Context, *EmptyGetRequest) (*GetAllTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTodoGrpcServer) DeleteTask(context.Context, *DeleteTaskRequest) (*EmptyTaskRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedTodoGrpcServer) UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedTodoGrpcServer) mustEmbedUnimplementedTodoGrpcServer() {}

// UnsafeTodoGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoGrpcServer will
// result in compilation errors.
type UnsafeTodoGrpcServer interface {
	mustEmbedUnimplementedTodoGrpcServer()
}

func RegisterTodoGrpcServer(s grpc.ServiceRegistrar, srv TodoGrpcServer) {
	s.RegisterService(&TodoGrpc_ServiceDesc, srv)
}

func _TodoGrpc_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGrpcServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoGrpc_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGrpcServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGrpc_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGrpcServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoGrpc_GetTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGrpcServer).GetTask(ctx, req.(*EmptyGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGrpc_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGrpcServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoGrpc_DeleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGrpcServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoGrpc_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoGrpcServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TodoGrpc_UpdateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoGrpcServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoGrpc_ServiceDesc is the grpc.ServiceDesc for TodoGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoGrpc",
	HandlerType: (*TodoGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TodoGrpc_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TodoGrpc_GetTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TodoGrpc_DeleteTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TodoGrpc_UpdateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services_todo.proto",
}