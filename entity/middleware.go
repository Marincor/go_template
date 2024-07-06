package entity

import (
	"fmt"
	"net/http"
	"slices"
)

type (
	Middleware func(http.Handler) http.Handler
	Router     struct {
		*http.ServeMux
		Chain []Middleware
	}
)

func (r *Router) Use(mx ...Middleware) {
	r.Chain = append(r.Chain, mx...)
}

func (r *Router) Group(fn func(r *Router)) {
	fn(&Router{ServeMux: r.ServeMux, Chain: slices.Clone(r.Chain)})
}

func (r *Router) Get(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodGet, path, fn, mx)
}

func (r *Router) Post(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPost, path, fn, mx)
}

func (r *Router) Patch(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPatch, path, fn, mx)
}

func (r *Router) Put(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodPut, path, fn, mx)
}

func (r *Router) Delete(path string, fn http.HandlerFunc, mx ...Middleware) {
	r.handle(http.MethodDelete, path, fn, mx)
}

func (r *Router) handle(method string, path string, fn http.HandlerFunc, mx []Middleware) {
	r.Handle(fmt.Sprintf("%s %s", method, path), r.wrap(fn, mx))
}

func (r *Router) wrap(fn http.HandlerFunc, mx []Middleware) (out http.Handler) {
	out, mx = http.Handler(fn), append(slices.Clone(r.Chain), mx...)

	slices.Reverse(mx)

	for _, m := range mx {
		out = m(out)
	}

	return
}
