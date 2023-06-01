package webservice

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/hasemeneh/poc-material/poc-login/pkg/router"
)

type WebHandler struct {
	router             router.Registrator
	registrators       []WebRegistrator
	server             *http.Server
	port               string
	httpListenAndServe httpListenAndServe
}

func NewWebService(port string, routerRegistrator router.Registrator, registrators ...WebRegistrator) WebService {
	return &WebHandler{
		router:             routerRegistrator,
		port:               port,
		registrators:       registrators,
		httpListenAndServe: http.ListenAndServe,
		server: &http.Server{
			Addr:    port,
			Handler: routerRegistrator,
		},
	}
}

func (w *WebHandler) Run() error {
	if w.port == "" {
		return errors.New("empty port defined")
	}
	for _, registrator := range w.registrators {
		registrator.Register(w.router)
	}

	return w.server.ListenAndServe()
}
func (w *WebHandler) GracefulStop() error {
	tenSecond := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), tenSecond)
	defer cancel()
	return w.server.Shutdown(ctx)
}

type httpListenAndServe func(addr string, handler http.Handler) error
