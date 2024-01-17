package gapi

import (
	"context"
	"time"

	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	arg := db.CreateTasksParams{
		TaskName:   req.GetTitle(),
		TaskDescription: ,
		
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}

func (server *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.EmptyTaskRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}

func (server *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {

	arg := db.UpdateTasksParams{
		TaskName: req.GetTitle(),
		DueDate:  time.Now(),
	}

	err := server.store.UpdateTasks(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update task: %v", err)
	}

	rsp := &pb.UpdateTaskResponse{}

	return rsp, nil
}

func (server *Server) GetTask(ctx context.Context, req *pb.EmptyGetRequest) (*pb.GetAllTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
