package response

import (
	"net/http/httptest"
	"testing"
)

func TestCSVResponseWriteResponse(t *testing.T) {
	resp := NewCSVResponse()
	resp.WriteResponse(httptest.NewRecorder())
}

func TestCSVResponse_SetData(t *testing.T) {
	resp := &CSVResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
	}
	resp.SetData([]byte("test"))
}
func TestCSVResponse_SetName(t *testing.T) {
	resp := &CSVResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
	}
	resp.SetName("test")
}
func TestCSVResponse_SetName_empty(t *testing.T) {
	resp := &CSVResponse{
		BasicResponse: BasicResponse{
			Body:        make([]byte, 0),
			StatusCode:  200,
			ContentType: "text",
		},
	}
	resp.SetName("")
}
