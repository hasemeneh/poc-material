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

func (p *Private) HandleGetUser(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	data, err := p.service.Usecase.Intools.GetUser(r.Context(), r.URL.Query(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(data)
}

func (p *Private) HandleGetUserAccess(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest)
	}
	data, err := p.service.Usecase.Intools.GetUserAccessByIDInternal(r.Context(), userID, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(data)
}

func (p *Private) HandleAddUserAccess(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.UserAccess{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.AddUserAccess(r.Context(), req, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}

func (p *Private) HandleUpdateUserAccess(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	req := &models.UserAccess{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Intools.ChangeStatusUserAccess(r.Context(), req, token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Success")
}
