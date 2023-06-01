package router

import (
	"net/http"

	"github.com/hasemeneh/poc-material/poc-pricing/pkg/middleware"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
)

type HTTPRouter interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	OPTIONS(path string, handler HandlerFunc)
}
type Registrator interface {
	AddMiddlewareWrapper(wrapper ...middleware.Wrapper) Registrator
	Register(method, path string, handler HandlerFunc)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type httpRouter struct {
	router Registrator
	prefix string
}

const RouterPath = "router-path"

func New(prefix string, router Registrator) HTTPRouter {
	return &httpRouter{
		router: router,
		prefix: prefix,
	}
}

func (h *httpRouter) GET(path string, handler HandlerFunc) {
	h.register(http.MethodGet, path, handler)
}

func (h *httpRouter) POST(path string, handler HandlerFunc) {
	h.register(http.MethodPost, path, handler)
}

func (h *httpRouter) PUT(path string, handler HandlerFunc) {
	h.register(http.MethodPut, path, handler)
}

func (h *httpRouter) OPTIONS(path string, handler HandlerFunc) {
	h.register(http.MethodOptions, path, handler)
}

func (h *httpRouter) register(method, path string, handler HandlerFunc) {
	fullpath := h.prefix + path
	h.router.Register(method, fullpath, pathRecorder(fullpath, handler))
}
func pathRecorder(path string, handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
		r.Header.Set(RouterPath, path)
		return handler(w, r)
	}
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) response.HTTPResponse
