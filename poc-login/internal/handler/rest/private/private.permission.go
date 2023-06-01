package private

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/response"
)

func (p *Private) HandleGetPermissions(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	data, err := p.service.Usecase.Intools.GetPermission(r.Context(), r.URL.Query(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(data)
}

func (p *Private) HandleAddPermissions(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.Permission{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.AddPermission(r.Context(), req, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Private) HandleUpdatePermissions(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.Permission{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.UpdatePermission(r.Context(), req, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}
