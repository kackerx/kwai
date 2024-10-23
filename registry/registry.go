package registry

import "context"

type Registrar interface {
	Register(ctx context.Context, svc *ServiceInstance) (err error)
	Deregister(ctx context.Context, svc *ServiceInstance) (err error)
}

type Discovery interface {
	GetService(ctx context.Context, serviceName string) ([]*ServiceInstance, error)
	Watch(ctx context.Context, serviceName string) (Watcher, error)
}

type Watcher interface {
	// Next
	// 1, 第一次监听时, 列表不为空, 返回实例服务列表
	// 2, 服务实例发生变化也返回服务实例列表
	// 3, 阻塞超时
	Next() ([]*ServiceInstance, error)
	Stop() error
}

type ServiceInstance struct {
	ID        string
	Name      string
	Metadata  map[string]string
	Version   string
	Endpoints []string
}
