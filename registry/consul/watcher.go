package consul

import (
	"context"

	"kwai/registry"
)

// watcher 多个watcher负责监听ss的通知, ss用来维护si和通知观察者
// 或者理解为, 每一个观察者持有观察的资源set, 当set发生变化时, set通知每一个观察者
type watcher struct {
	event chan struct{}
	set   *serviceSet

	// for cancel
	ctx    context.Context
	cancel context.CancelFunc
}

func (w *watcher) Next() (services []*registry.ServiceInstance, err error) {
	select {
	case <-w.ctx.Done():
		err = w.ctx.Err()
		return
	case <-w.event:
	}

	// watcher需要访问ss的方法和事件通知
	ss, ok := w.set.services.Load().([]*registry.ServiceInstance)

	if ok {
		services = append(services, ss...)
	}
	return
}

func (w *watcher) Stop() error {
	w.cancel()
	w.set.lock.Lock()
	defer w.set.lock.Unlock()
	delete(w.set.watcher, w)
	return nil
}
