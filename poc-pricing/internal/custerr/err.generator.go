package custerr

import (
	"fmt"
	"net/http"

	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
)

func NewInvalidInput(message string) error {
	return fmt.Errorf("invalid input : %w", response.NewError(message, http.StatusPreconditionRequired))
}
