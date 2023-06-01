package models

import "github.com/hasemeneh/poc-material/poc-pricing/internal/constants"

type Role struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UserAccess struct {
	RoleID int64                      `json:"role_id" db:"role_id"`
	UserID int64                      `json:"user_id" db:"user_id"`
	Name   string                     `json:"name" db:"name"`
	Status constants.UserAccessStatus `json:"status" db:"status"`
}

type InternalRole struct {
	*Role
	Status constants.RoleStatus `json:"status" db:"status"`
}
type RolePermission struct {
	*Role
	Permissions []*Permission
}

type InternalRolePermission struct {
	RoleID       int64                          `json:"role_id" db:"role_id"`
	PermissionID int64                          `json:"permission_id" db:"permission_id"`
	Status       constants.RolePermissionStatus `json:"status" db:"status"`
}

type AddRoleReq struct {
	Role          InternalRole             `json:"role"`
	PermissionIDs []InternalRolePermission `json:"permissions"`
}
