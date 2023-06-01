package models

import "github.com/hasemeneh/poc-material/poc-pricing/internal/constants"

type User struct {
	ID       int64                `json:"id" db:"id"`
	AuthID   string               `json:"auth_id" db:"auth_id"`
	Password []byte               `json:"-" db:"password"`
	Status   constants.UserStatus `json:"status" db:"status"`
}

type UserToken struct {
	*User
	*Token
}
