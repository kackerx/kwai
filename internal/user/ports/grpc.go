package ports

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "kwai/api/proto/user"
	"kwai/internal/user"
	"kwai/internal/user/assembler"
)

type GrpcServer struct {
	app *user.Application

	pb.UnimplementedUserServer
}

func (g *GrpcServer) UserDetail(ctx context.Context, req *pb.UserDetailReq) (*pb.UserDetailResp, error) {
	resp := assembler.UserEntity2ToPB(g.app.GetUser(req.GetUid()))
	return resp, nil
}

func NewGrpcServer(app *user.Application) *GrpcServer {
	return &GrpcServer{app: app}
}

func (g *GrpcServer) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	// TODO implement me
	panic("implement me")
}

func (g *GrpcServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	// TODO implement me
	panic("implement me")
}

func (g *GrpcServer) Hehe(ctx context.Context, req *pb.HeheReq) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
