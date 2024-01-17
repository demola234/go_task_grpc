package gapi

import (
	"context"

	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetTask(ctx context.Context, req *pb.EmptyGetRequest) (*pb.GetAllTaskResponse, error) {

	arg := db.ListTasksParams{
		Limit:  5,
		Offset: 0,
	}

	tasks, err := server.store.ListTasks(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get tasks: %v", err)
	}
	rsp := &pb.GetAllTaskResponse{
		Tasks: convertTasks(tasks),
	}
	return rsp, nil
}


// DeleteTask
func (server *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	
	arg := db.DeleteTaskParams{
		ID: req.GetTaskId(),
	}

	_, err := server.store.DeleteTask(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot delete task: %v", err)
	}
	rsp := &pb.DeleteTaskResponse{
		TaskId: req.GetTaskId(),
	}
	return rsp, nil
}