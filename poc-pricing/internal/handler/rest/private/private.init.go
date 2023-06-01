package private

import (
	"net/http"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/service"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/router"
)

type Private struct {
	service *service.Service
	prefix  string
}

func NewHandler(service *service.Service) *Private {
	return &Private{
		service: service,
		prefix:  "/internal",
	}
}
func (p *Private) Register(rr router.Registrator) {
	r := router.New(p.prefix, rr)
	r.GET("/ping", p.PING)
	r.GET("/permission", p.HandleGetPermissions)
	r.POST("/permission", p.HandleAddPermissions)
	r.OPTIONS("/permission", p.HandleUpdatePermissions)
	r.GET("/role", p.HandleGetRoles)
	r.POST("/role", p.HandleAddRole)
	r.OPTIONS("/role", p.HandleUpdateRole)
	r.GET("/role/permission", p.HandleGetPermissionByRole)
	r.POST("/role/permission", p.HandleAddRolePermission)
	r.OPTIONS("/role/permission", p.HandleUpdateRolePermission)
	r.GET("/user", p.HandleGetUser)
	r.GET("/user/access", p.HandleGetUserAccess)
	r.POST("/user/access", p.HandleAddUserAccess)
	r.OPTIONS("/user/access", p.HandleUpdateUserAccess)
}

func (p *Private) PING(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	return response.NewJSONResponse().SetMessage("Pong").SetData("Pung")
}
