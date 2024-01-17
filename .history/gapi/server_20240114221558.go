package gapi

import (
	"fmt"

	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
	"github.com/demola234/todogrpc/token"
	"github.com/demola234/todogrpc/utils"
)

// ?? [Server] serves gRPC requests
type Server struct {
	pb.UnimplementedTodoGrpcServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

// ? NewServer creates a new gRPC server and setup routing
func TodoGrpcServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %s", err.Error())
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, err
}
