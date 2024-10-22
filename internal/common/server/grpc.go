package server

import (
	"net"

	"google.golang.org/grpc"
)

func RunGRPCServer(register func(server *grpc.Server)) {
	server := grpc.NewServer()

	register(server)

	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
