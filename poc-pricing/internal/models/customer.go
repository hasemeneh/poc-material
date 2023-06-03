package models

import (
	"database/sql"
)

type Customer struct {
	ID        string       `json:"id,omitempty" db:"id"`
	Address   string       `json:"address" db:"address"`
	City      string       `json:"city" db:"city"`
	State     string       `json:"state" db:"state"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
