package gapi

import (
	"context"

	"github.com/demola234/todogrpc/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {

	_, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "category already exists: %v", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "cannot create category: %v", err)
	}
	rsp := &pb.CreateCategoryResponse{
		Name: req.Name,
	}

	return rsp, nil
}
