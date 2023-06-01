package public

import (
	"net/http"
	"strings"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
)

func (p *Public) HandleGetAccessList(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
	token := strings.Replace(r.Header.Get(constants.AuthHeaderKey), constants.AuthPrefix, "", 1)
	resp, err := p.service.Usecase.Authorization.GetAllUserAccess(r.Context(), token)
	if err != nil {
		return response.NewJSONResponse().SetError(err)
	}
	return response.NewJSONResponse().SetData(resp)
}
