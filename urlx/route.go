package urlx

import (
	"djan-go/urlx/internal/route"
)

type Router struct {
	root    *route.Node
	routers map[*route.Node]any // warning: 线程不安全
}

func NewRouter() *Router {
	return &Router{
		root:    route.NewRoot(),
		routers: make(map[*route.Node]any),
	}
}

func (r *Router) Find(path string) any {
	n := r.root.Find(path)
	if n == nil {
		return nil
	}
	return r.routers[n]
}

func (r *Router) Register(path string, inter any) {
	n := r.root.Insert(path)
	r.routers[n] = inter
}
