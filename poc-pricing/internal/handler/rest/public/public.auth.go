package public

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
)

func (p *Public) HandleAuthLogin(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	authID := r.PostFormValue("username")
	password := r.PostFormValue("password")
	resp, err := p.service.Usecase.Authentication.Login(r.Context(), authID, password)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(resp)
}
func (p *Public) HandleAuthRegister(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	req := &models.RegisterReq{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrorBadRequest).SetMessage(err.Error())
	}
	err = p.service.Usecase.Authentication.Register(r.Context(), req)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Registered")
}

func (p *Public) HandleLogout(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	err := p.service.Usecase.Authentication.Logout(r.Context(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetMessage("Logged out")
}

func (p *Public) HandleRefreshToken(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := r.FormValue("refresh_token")
	resp, err := p.service.Usecase.Authentication.RefreshToken(r.Context(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(resp)
}
