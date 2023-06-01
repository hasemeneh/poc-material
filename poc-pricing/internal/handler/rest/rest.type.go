package rest

import (
	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/router"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/webservice"
)

type Builder interface {
	Build(s *service.Service) webservice.WebService
}
type builder struct {
	NewWebService          newWebServiceFunc
	GenerateRouter         generateRouter
	GenerateWebRegistrator generateWebRegistrator
}
type newWebServiceFunc func(port string, routerRegistrator router.Registrator, registrators ...webservice.WebRegistrator) webservice.WebService
type generateRouter func() router.Registrator
type generateWebRegistrator func(s *service.Service) []webservice.WebRegistrator
