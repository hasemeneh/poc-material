package custerr

import (
	"errors"
	"net/http"

	"github.com/hasemeneh/poc-material/poc-pricing/pkg/response"
)

var ErrInactiveUser = response.NewError("user is inactive", http.StatusUnauthorized)
var ErrUserNotFound = response.NewError("user not found", http.StatusNotFound)
var ErrUserAlreadyRegistered = response.NewError("user is already registered", http.StatusConflict)
var ErrInvalidToken = response.NewError("invalid token", http.StatusUnauthorized)
var ErrWrongPassword = response.NewError("wrong password", http.StatusUnauthorized)
var ErrNoPermissionFound = errors.New("no permission found")
var ErrTokenExpired = response.NewError("token expired", http.StatusUnauthorized)
var ErrInvalidAccess = response.NewError("invalid access", http.StatusUnauthorized)
