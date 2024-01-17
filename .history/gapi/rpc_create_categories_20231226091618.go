package gapi

import (
	"context"

	"github.com/demola234/todogrpc/val"
	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
	"github.com/demola234/todogrpc/utils"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	violations := validateCreateCategory(req)
	if violations != nil {
		return nil, invalidArgErr(violations)
	}


	

	user, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		
	return rsp, nil
}

func validateCreateCategory(req *pb.CreateCategoryRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	
}
