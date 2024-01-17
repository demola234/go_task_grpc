package gapi

import (
	"context"

	"github.com/demola234/todogrpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAllTaskResponse(ctx context.Context, req *pb.EmptyGetRequest) (*pb.GetAllTaskResponse, error) {

	arg := db.GetAllTaskParams{
	tasks, err := server.store.ListTasks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get tasks: %v", err)
	}
	rsp := &pb.GetAllTaskResponse{
		Tasks: convertTasks(tasks),
	}
	return rsp, nil
}
