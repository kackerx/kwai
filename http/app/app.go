package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pb "kwai/api/proto/user"
	"kwai/grpc"
)

func main() {
	userSrv := grpc.NewUserServer()
	g := gin.Default()

	pb.RegisterUserServerHTTPServer(userSrv, g)

	// 支持自动生成端口, 定义ip和端口
	_ = g.SetTrustedProxies(nil)
	server := &http.Server{
		Addr:    ":8085",
		Handler: g,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
