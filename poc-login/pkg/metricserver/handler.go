package metricserver

import (
	"net/http"

	"github.com/hasemeneh/poc-material/poc-login/pkg/response"
	"github.com/hasemeneh/poc-material/poc-login/pkg/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metric struct {
	prefix string
}

func NewHandler() *Metric {
	return &Metric{
		prefix: "",
	}
}
func (p *Metric) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/metrics", p.HandlerMetrics)
	r.GET("/v1/healthz", p.HandlerHealthCheck)
}
func (p *Metric) HandlerMetrics(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	promhttp.Handler().ServeHTTP(w, r)
	return nil
}
func (p *Metric) HandlerHealthCheck(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Success")
}
