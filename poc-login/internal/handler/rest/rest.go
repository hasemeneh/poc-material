package rest

import (
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/hasemeneh/poc-material/poc-login/internal/handler/rest/private"
	"github.com/hasemeneh/poc-material/poc-login/internal/handler/rest/public"
	"github.com/hasemeneh/poc-material/poc-login/internal/service"
	"github.com/hasemeneh/poc-material/poc-login/pkg/router"
	routeradapter "github.com/hasemeneh/poc-material/poc-login/pkg/router/adapter"
	"github.com/hasemeneh/poc-material/poc-login/pkg/webservice"
)

func (o *builder) Build(s *service.Service) webservice.WebService {
	return o.NewWebService(s.Cfg.RestPort, o.GenerateRouter(), o.GenerateWebRegistrator(s)...)
}
func GetDefaultBuilder() Builder {
	return &builder{
		NewWebService:          webservice.NewWebService,
		GenerateRouter:         getRouter,
		GenerateWebRegistrator: getWebRegistrator,
	}
}
func getRouter() router.Registrator {
	// TODO : add some middleware here later
	module := routeradapter.UseChiRouter()
	module.AddMiddlewareWrapper(
		chimiddleware.Recoverer,
	)

	return module
}
func getWebRegistrator(s *service.Service) []webservice.WebRegistrator {
	resp := []webservice.WebRegistrator{
		public.NewHandler(s),
		private.NewHandler(s),
	}
	return resp
}
