package grpc

import (
	"context"

	pb "kwai/api/proto/user"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

var _ pb.UserServer = &UserServer{}

func (u *UserServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Msg: "hehe" + req.Code}, nil
}

func (u *UserServer) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	// TODO implement me
	panic("implement me")
}
