package webservice

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	routermock "github.com/hasemeneh/poc-material/poc-login/pkg/router/mocks"
	"github.com/hasemeneh/poc-material/poc-login/pkg/webservice/mocks"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)
	NewWebService("0.0.0.0:8007", mocksRouterRegistrator)
}
func TestRun(t *testing.T) {
	controller := gomock.NewController(t)
	mocksRouterRegistrator := routermock.NewMockRegistrator(controller)
	mocksWebRegistrator := mocks.NewMockWebRegistrator(controller)
	module := &WebHandler{
		registrators: []WebRegistrator{mocksWebRegistrator},
		router:       mocksRouterRegistrator,
		port:         "0.0.0.0:8007",
		httpListenAndServe: func(addr string, handler http.Handler) error {
			return nil
		},
	}
	mocksWebRegistrator.EXPECT().Register(gomock.Any()).Return()
	err := module.Run()
	require.NoError(t, err)
}

func TestRunErrorNoPort(t *testing.T) {
	module := &WebHandler{
		port: "",
		httpListenAndServe: func(addr string, handler http.Handler) error {
			return nil
		},
	}
	err := module.Run()
	require.Error(t, err)
}
