package private

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/response"
)

func (p *Private) HandleGetRoles(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	data, err := p.service.Usecase.Intools.GetRole(r.Context(), r.URL.Query(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(data)
}
func (p *Private) HandleGetPermissionByRole(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	roleID, err := strconv.ParseInt(r.URL.Query().Get("role_id"), 10, 64)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest)
	}
	data, err := p.service.Usecase.Intools.GetPermissionByRole(r.Context(), roleID, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(data)
}

func (p *Private) HandleAddRole(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.AddRoleReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.AddRole(r.Context(), req, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Private) HandleUpdateRole(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.InternalRole{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.UpdateRole(r.Context(), token, req)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Private) HandleAddRolePermission(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.InternalRolePermission{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.AddPermissionToRole(r.Context(), token, req)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Private) HandleUpdateRolePermission(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.InternalRolePermission{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.UpdateRolePermission(r.Context(), token, req)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}
