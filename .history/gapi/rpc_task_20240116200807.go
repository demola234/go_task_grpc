package gapi

import (
	"context"
	"fmt"
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
		TaskStatus:      req.GetTaskStatus(),
		DueDate:         req.GetTaskCreatedAt().AsTime(),
		CreatedAt:       time.Now(),
	}

	task, err := server.store.CreateTasks(ctx, arg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create task: %v", err)
	}

	rsp := &pb.CreateTaskResponse{
		TaskName:        task.TaskName,
		TaskDescription: task.TaskDescription,
		TaskStatus:      task.TaskStatus,
		TaskId:          task.ID,
		TaskDeadline:    task.DueDate.String(),
	}

	status.New(codes.OK, "task created successfully")

	return rsp, nil
}

func (server *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.EmptyTaskRequest, error) {
	err := server.store.DeleteTasks(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot delete task: %v", err)
	}

	rsp := &pb.EmptyTaskRequest{}

	status.New(codes.OK, "task deleted successfully")

	return rsp, nil
}

func (server *Server) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {

	arg := db.UpdateTasksParams{
		TaskName:   req.GetTitle(),
		DueDate:    time.Now(),
		TaskStatus: req.GetTaskStatus(),
		ID:         req.GetId(),
	}

	err := server.store.UpdateTasks(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update task: %v", err)
	}

	rsp := &pb.UpdateTaskResponse{
		Task: &pb.Task{
			TaskName:        req.GetTitle(),
			TaskDescription: req.GetDescription(),
			TaskDeadline:    time.Now().String(),
			TaskStatus:      req.GetTaskStatus(),
			TaskId:          req.GetId(),
		},
	}

	status.New(codes.OK, "task updated successfully")

	return rsp, nil
}

func (server *Server) GetTask(ctx context.Context, req *pb.EmptyGetRequest) (*pb.GetAllTaskResponse, error) {
	tasks, err := server.store.ListTasks(ctx, db.ListTasksParams{
		Limit:  10,
		Offset: 5,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot fetch task: %v", err)
	}

	rsp := &pb.GetAllTaskResponse{
		Tasks: c
	}

	task := fmt.Sprintf("task fetched %v", tasks)

	status.New(codes.OK, task)

	return rsp, nil
}
