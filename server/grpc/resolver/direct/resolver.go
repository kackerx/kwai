package direct

import "google.golang.org/grpc/resolver"

type directResolver struct {
}

func (d *directResolver) ResolveNow(options resolver.ResolveNowOptions) {
	// TODO implement me
	panic("implement me")
}

func (d *directResolver) Close() {
	// TODO implement me
	panic("implement me")
}

func newDirectResolver() *directResolver {
	return &directResolver{}
}
