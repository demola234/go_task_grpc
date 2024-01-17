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
		TaskName:        req.GetTaskName(),
		TaskDescription: req.GetTaskDescription(),
		DueDate:         req.GetTaskCreatedAt().AsTime(),
		CreatedAt:       time.Now(),
	}

	task, err := server.store.CreateTasks(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create task: %v", err)
	}

	rsp := &pb.CreateTaskResponse{
		Task: &pb.{
			TaskName:        task.TaskName,
			TaskDescription: task.TaskDescription,
			TaskDeadline:    task.DueDate.Format("2006-01-02 15:04:05"),
			TaskStatus:      task.TaskStatus,
			TaskId:          task.ID,
		},
	}
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
