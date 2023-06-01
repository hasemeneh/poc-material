package response

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type CSVResponse struct {
	BasicResponse
	CSVBody CSVBody
	Error   error
}
type CSVBody struct {
	*Error
	Name string
}

func NewCSVResponse() *CSVResponse {
	return &CSVResponse{
		BasicResponse: BasicResponse{
			ContentType: CSVContentType,
			StatusCode:  http.StatusOK,
		},
	}
}

func (r *CSVResponse) SetData(data []byte) *CSVResponse {
	r.Body = data
	return r
}

const defaultNameFormat = "%s.csv"

func (r *CSVResponse) SetName(name string) *CSVResponse {
	if name == "" {
		r.CSVBody.Name = fmt.Sprintf(defaultNameFormat, time.Now().Format(time.RFC1123))
	}
	if !strings.HasSuffix(name, ".csv") {
		r.CSVBody.Name = fmt.Sprintf(defaultNameFormat, name)
	}
	r.CSVBody.Name = name
	return r
}

func (r *CSVResponse) WriteResponse(w http.ResponseWriter) {
	if r.CSVBody.Name == "" {
		r.CSVBody.Name = fmt.Sprintf(defaultNameFormat, time.Now().Format(time.RFC1123))
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=\"%s\"", r.CSVBody.Name))
	r.BasicResponse.WriteResponse(w)
}
