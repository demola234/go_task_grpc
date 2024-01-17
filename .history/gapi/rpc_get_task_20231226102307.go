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


func (server *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) error {

	err := server.store.DeleteTasks(ctx, req.Id)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot delete task: %v", err)
	}

	return nil
}


func (server *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	
	arg := db.UpdateTasksParams{
		TaskName:        req.GetTitle(),
		DueDate: ,

		TaskStatus:      req.GetTaskStatus(),
		TaskDeadline:    req.GetTaskDeadline(),
		TaskId:          req.GetTaskId(),
	}

	task, err := server.store.UpdateTasks(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update task: %v", err)
	}
	rsp := &pb.UpdateTaskResponse{
		Task: convertTask(task),
	}
	return rsp, nil
}