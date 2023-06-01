package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
	"github.com/stretchr/testify/require"
)

func TestPathRecorder(t *testing.T) {
	resp := pathRecorder("/test", func(w http.ResponseWriter, r *http.Request) response.HTTPResponse {
		return nil
	})
	req := &http.Request{
		Header: make(http.Header),
	}
	resp(httptest.NewRecorder(), req)
	require.Equal(t, "/test", req.Header.Get("router-path"))
}
