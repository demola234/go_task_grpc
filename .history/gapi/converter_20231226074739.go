package gapi

import (
	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
)

func convertUser(user db.Users) *pb.User {
	return &pb.User{
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}
