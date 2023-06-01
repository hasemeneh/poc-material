package public

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hasemeneh/poc-material/poc-login/internal/service"
	routermocks "github.com/hasemeneh/poc-material/poc-login/pkg/router/mocks"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler(&service.Service{})
	require.NotNil(t, handler)
}

type testObject struct {
	suite.Suite
	module                *Public
	Service               *service.Service
	mockRouterRegistrator *routermocks.MockRegistrator
	request               *http.Request
	writer                http.ResponseWriter
}

func (obj *testObject) SetupTest() {
	gomockController := gomock.NewController(obj.T())
	obj.mockRouterRegistrator = routermocks.NewMockRegistrator(gomockController)
	obj.Service = &service.Service{}
	obj.request = &http.Request{
		URL: &url.URL{},
	}
	obj.writer = httptest.NewRecorder()
	obj.module = &Public{
		prefix: "",
	}
}

func Test_public(t *testing.T) {
	suite.Run(t, new(testObject))
}

func (obj *testObject) TestRegister() {
	obj.mockRouterRegistrator.EXPECT().Register(http.MethodGet, "/ping", gomock.Any()).Return()
	obj.module.Register(obj.mockRouterRegistrator)
}

func (obj *testObject) TestPing() {
	response := obj.module.PING(obj.writer, obj.request)
	require.NotNil(obj.T(), response)
}
