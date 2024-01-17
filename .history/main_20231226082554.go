package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	db "github.com/demola234/todogrpc/db/sqlc"
	"github.com/demola234/todogrpc/gapi"
	"github.com/demola234/todogrpc/pb"
	"github.com/demola234/todogrpc/utils"

	_ "github.com/demola234/todogrpc/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	configs, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config")
	}
	if configs.Enviroment == "developement" {
		fmt.Println("running in developement mode")
	}

	conn, err := sql.Open(configs.DBDriver, configs.DBSource)
	if err != nil {
		fmt.Println("cannot connect to db")
	}

	fmt.Printf("connected to db %s\n", configs.DBDriver)

	store := db.NewStore(conn)

	runGateWayServer(configs, store)
	// runGRPCServer(configs, store)
}

func runGRPCServer(configs utils.Config, store db.Store) {

	server, err := gapi.NewServer(configs, store)
	if err != nil {
		fmt.Println("cannot create server")
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterTodoGrpcServer(gRPCServer, server)

	// Register reflection service on gRPC server.
	reflection.Register(gRPCServer)

	// Listen and serve
	listener, err := net.Listen("tcp", configs.GRPCServerAddress)
	if err != nil {
		fmt.Println("cannot start server")
	}
	fmt.Println("server started at port", configs.GRPCServerAddress)

	err = gRPCServer.Serve(listener)
	if err != nil {
		fmt.Fprintln(log.Writer(), "cannot start server")
	}

}

func runGateWayServer(configs utils.Config, store db.Store) {
	server, err := gapi.NewServer(configs, store)
	if err != nil {
		fmt.Println("cannot create server")
	}

	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	gRPCMux := runtime.NewServeMux(jsonOpt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterTodoGrpcHandlerServer(ctx, gRPCMux, server)
	if err != nil {
		fmt.Println("cannot register gateway")
	}

	mux := http.NewServeMux()
	mux.Handle("/", gRPCMux)

	statikFS, err := fs.New()
	if err != nil {
		fmt.Println("cannot start server")
	}

	fs := http.FileServer(statikFS)
	mux.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	// Listen and serve
	listener, err := net.Listen("tcp", configs.HTTPServerAddress)
	if err != nil {
		fmt.Println("cannot start server")
	}

	fmt.Println("server started at port", configs.HTTPServerAddress)
	handler := http.Handler(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		fmt.Fprintln(log.Writer(), "cannot start server")
	}
}
