package direct

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

func init() {
	// 注册服务发现
	resolver.Register(NewDirectBuilder())
}

type directBuilder struct {
}

// NewDirectBuilder 直连服务发现
// ex:
// direct://<authority>/127.0.0.1:9000,
func NewDirectBuilder() *directBuilder {
	return &directBuilder{}
}

// Build 构造resolver
func (d *directBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	addrs := make([]resolver.Address, 0)
	path := strings.TrimPrefix(target.URL.Path, "/")
	paths := strings.Split(path, ",")

	for _, addr := range paths {
		addrs = append(addrs, resolver.Address{
			Addr:               addr,
			ServerName:         "",
			Attributes:         nil,
			BalancerAttributes: nil,
			Metadata:           nil,
		})
	}

	// grpc建立连接的逻辑
	if err := cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return nil, err
	}

	return newDirectResolver(), nil
}

func (d *directBuilder) Scheme() string {
	return "direct"
}
