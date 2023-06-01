package models

import "github.com/hasemeneh/poc-material/poc-login/internal/constants"

type Permission struct {
	ID     int64                      `json:"id,omitempty" db:"id"`
	Code   string                     `json:"code" db:"code"`
	Name   string                     `json:"name" db:"name"`
	RoleID int64                      `json:"-" db:"role_id"`
	Status constants.PermissionStatus `json:"status" db:"status"`
}
