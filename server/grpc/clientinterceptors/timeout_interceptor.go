package clientinterceptors

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func TimeoutInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		st := time.Now()
		fmt.Println("timeout before", st)
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println("timeout after", reply, time.Since(st))

		return err
	}
}
