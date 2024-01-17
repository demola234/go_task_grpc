package gapi

import (


	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/pb"
	"github.com/demola234/todogrpc/utils"
)

// ?? [Server] serves gRPC requests
type Server struct {
	pb.
	config     utils.Config
	store      db.Store
}

// ? NewServer creates a new gRPC server and setup routing
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		store:      store,
		config:     config,
	}
	return server, nil
}
