package discovery

import (
	"context"
	"errors"
	"strings"
	"time"

	"google.golang.org/grpc/resolver"

	"kwai/registry"
)

const (
	name = "discover"
)

type Option func(builder *discoveryBuilder)

func WithTimeout(timeout time.Duration) Option {
	return func(b *discoveryBuilder) {
		b.timeout = timeout
	}
}

type discoveryBuilder struct {
	discovery registry.Discovery
	timeout   time.Duration
	insecure  bool
}

func (b *discoveryBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var (
		err error
		w   registry.Watcher
	)

	done := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		w, err = b.discovery.Watch(ctx, strings.TrimPrefix(target.URL.Path, "/"))
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(b.timeout):
		err = errors.New("discovery create watcher timeout")
		return nil, err
	}

	r := &discoveryResolver{
		ctx:      ctx,
		w:        w,
		cc:       cc,
		cancel:   cancel,
		insecure: b.insecure,
	}
	go r.watch()

	return r, nil
}

func (d *discoveryBuilder) Scheme() string {
	return name
}

func NewDiscoveryBuilder(discovery registry.Discovery, opts ...Option) *discoveryBuilder {
	b := &discoveryBuilder{discovery: discovery, timeout: 10 * time.Second}
	for _, opt := range opts {
		opt(b)
	}

	return b
}
