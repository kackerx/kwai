package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"kwai/registry"
)

type discoveryResolver struct {
	ctx      context.Context
	w        registry.Watcher
	cc       resolver.ClientConn
	cancel   context.CancelFunc
	insecure bool
}

func (d *discoveryResolver) ResolveNow(options resolver.ResolveNowOptions) {
	// TODO implement me
	panic("implement me")
}

func (d *discoveryResolver) Close() {
	d.cancel()
	if err := d.w.Stop(); err != nil {
		fmt.Println("[resolver] failed to watch stop", err)
		return
	}
}

func (d *discoveryResolver) watch() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
		}

		ins, err := d.w.Next()
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			time.Sleep(time.Second)
			continue
		}

		d.update(ins)
	}
}

func (d *discoveryResolver) update(ins []*registry.ServiceInstance) {
	addrs := make([]resolver.Address, 0)
	endpoints := make(map[string]struct{})

	for _, in := range ins {
		endpoint, err := ParseEndpoint(in.Endpoints, "grpc", !d.insecure)
		if err != nil {
			fmt.Println("failed to parse enpoint", err)
			continue
		}

		if endpoint == "" {
			continue
		}

		if _, ok := endpoints[endpoint]; ok {
			continue
		}

		endpoints[endpoint] = struct{}{}
		addr := resolver.Address{
			Addr:       endpoint,
			ServerName: in.Name,
			// Attributes: parseAttributes(in.Metadata),
		}

		addr.Attributes = addr.Attributes.WithValue("rawServiceInstance", in)
		addrs = append(addrs, addr)
	}

	if len(addrs) == 0 {
		fmt.Println("[resolver] zero endpoint found")
		return
	}

	if err := d.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		fmt.Println("[resolver] update state error", err)
		return
	}

	b, _ := json.Marshal(ins)
	log.Infof("[resolver] update instance: %s", b)
}

func NewEndpoint(scheme, host string, isSecure bool) *url.URL {
	var query string
	if isSecure {
		query = "isSecure=true"
	}
	return &url.URL{Scheme: scheme, Host: host, RawQuery: query}
}

func parseAttributes(md map[string]string) *attributes.Attributes {
	var a *attributes.Attributes
	for k, v := range md {
		if a == nil {
			a = attributes.New(k, v)
		} else {
			a = a.WithValue(k, v)
		}
	}

	return a
}

func ParseEndpoint(endpoints []string, scheme string, isSecure bool) (string, error) {
	for _, e := range endpoints {
		u, err := url.Parse(e)
		if err != nil {
			return "", nil
		}

		if u.Scheme == scheme {
			if IsSecure(u) == isSecure {
				return u.Host, nil
			}
		}
	}

	return "", nil
}

func IsSecure(u *url.URL) bool {
	ok, err := strconv.ParseBool(u.Query().Get("isSecure"))
	if err != nil {
		return false
	}

	return ok
}
