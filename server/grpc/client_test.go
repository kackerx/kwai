package grpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"

	pb "kwai/api/proto/user"
	"kwai/registry/consul"
)

func TestDialWithInsecure(t *testing.T) {
	conf := api.DefaultConfig()
	conf.Address = "127.0.0.0:8500"
	conf.Scheme = "http"
	c, err := api.NewClient(conf)
	if err != nil {
		panic(err)
	}

	registry := consul.New(c, consul.WithHealthCheck(true))

	ctx := context.Background()

	conn, err := DialWithInsecure(
		ctx,
		WithDiscovery(registry),
		WithEndpoint("discover:///greeter"),
		WithClientUnaryInterceptor(),
		WithTimeout(time.Second*2),
	)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	client := pb.NewGreeterClient(conn)

	helloResp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "kacker"})
	if err != nil {
		panic(err)
	}

	fmt.Println(helloResp)
}
