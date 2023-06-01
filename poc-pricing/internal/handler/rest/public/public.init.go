package public

import (
	"net/http"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/router"
)

type Public struct {
	service *service.Service
	prefix  string
}

func NewHandler(service *service.Service) *Public {
	return &Public{
		service: service,
		prefix:  "/api",
	}
}
func (p *Public) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/ping", p.PING)
	r.POST("/auth/login", p.HandleAuthLogin)
	r.POST("/auth/logout", p.HandleLogout)
	r.POST("/auth/register", p.HandleAuthRegister)
	r.POST("/auth/refresh", p.HandleRefreshToken)
	r.GET("/access/list", p.HandleGetAccessList)
}

func (p *Public) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
