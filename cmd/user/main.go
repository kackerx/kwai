package main

import (
	"google.golang.org/grpc"

	pb "kwai/api/proto/user"
	"kwai/internal/common/server"
	"kwai/internal/user"
	"kwai/internal/user/ports"
)

func main() {
	app := user.NewApp("user")

	// serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	serverType := "grpc"
	switch serverType {
	case "http":
	case "grpc":
		server.RunGRPCServer(func(grpcServer *grpc.Server) {
			svc := ports.NewGrpcServer(app)
			pb.RegisterUserServer(grpcServer, svc)
		})
	}
}
