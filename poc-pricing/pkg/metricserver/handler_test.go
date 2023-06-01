package metricserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	routermock "github.com/hasemeneh/poc-material/poc-pricing/pkg/router/mocks"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)

	mocksRouterRegistrator.EXPECT().Register(http.MethodGet, "/metrics", gomock.Any()).Return()
	mocksRouterRegistrator.EXPECT().Register(http.MethodGet, "/v1/healthz", gomock.Any()).Return()
	module := &Metric{}

	module.Register(mocksRouterRegistrator)
}
func TestNewHandler(t *testing.T) {
	module := NewHandler()
	require.Equal(t, module.prefix, "")
}
func TestHandleMetrics(t *testing.T) {
	module := NewHandler()
	module.HandlerMetrics(httptest.NewRecorder(), &http.Request{})
}
