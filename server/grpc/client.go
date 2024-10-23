package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"kwai/registry"
	"kwai/server/grpc/clientinterceptors"
	"kwai/server/grpc/resolver/discovery"
)

type BalanceType string

const (
	BalanceTypeRoundRobin = "round_robin"
)

type clientOptions struct {
	endpoint string

	timeout time.Duration

	// 服务发现接口
	discovery registry.Discovery

	// grpc.DialOption
	dialOpts []grpc.DialOption

	// 客户端拦截器
	unaryInts []grpc.UnaryClientInterceptor

	// 负载均衡
	balanceType BalanceType
}

type ClientOption func(opt *clientOptions)

func WithEndpoint(endpoint string) ClientOption {
	return func(opt *clientOptions) {
		opt.endpoint = endpoint
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(opt *clientOptions) {
		opt.timeout = timeout
	}
}

func WithDiscovery(discovery registry.Discovery) ClientOption {
	return func(opt *clientOptions) {
		opt.discovery = discovery
	}
}

func WithClientUnaryInterceptor(in ...grpc.UnaryClientInterceptor) ClientOption {
	return func(opt *clientOptions) {
		opt.unaryInts = in
	}
}

func WithClientDialOptions(dialOpts ...grpc.DialOption) ClientOption {
	return func(opt *clientOptions) {
		opt.dialOpts = dialOpts
	}
}

func WithBalanceType(balanceTYpe BalanceType) ClientOption {
	return func(opt *clientOptions) {
		opt.balanceType = balanceTYpe
	}
}

func DialWithInsecure(ctx context.Context, opts ...ClientOption) (conn *grpc.ClientConn, err error) {
	return dial(ctx, true, opts...)
}

func dial(ctx context.Context, insecure bool, opts ...ClientOption) (conn *grpc.ClientConn, err error) {
	cliOpt := &clientOptions{
		timeout:     time.Millisecond * 3000,
		balanceType: BalanceTypeRoundRobin,
	}
	for _, opt := range opts {
		opt(cliOpt)
	}

	ints := []grpc.UnaryClientInterceptor{
		clientinterceptors.TimeoutInterceptor(cliOpt.timeout),
	}

	if len(cliOpt.unaryInts) > 0 {
		ints = append(ints, cliOpt.unaryInts...)
	}

	grpcOpts := []grpc.DialOption{
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"` + string(cliOpt.balanceType) + `"}`),
		grpc.WithChainUnaryInterceptor(ints...),
	}

	if len(cliOpt.dialOpts) > 0 {
		grpcOpts = append(grpcOpts, cliOpt.dialOpts...)
	}

	// 服务发现
	if cliOpt.discovery != nil {
		grpcOpts = append(grpcOpts, grpc.WithResolvers(
			discovery.NewDiscoveryBuilder(cliOpt.discovery),
		))
	}

	if insecure {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	return grpc.DialContext(ctx, cliOpt.endpoint, grpcOpts...)
}
