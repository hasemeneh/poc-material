package models

import (
	"database/sql"
)

type Logistics struct {
	ID        string       `json:"id,omitempty" db:"id"`
	Type      string       `json:"type" db:"type"`
	Capacity  int64        `json:"capacity" db:"capacity"`
	Cost      int64        `json:"cost" db:"cost"`
	City      string       `json:"city" db:"city"`
	State     string       `json:"state" db:"state"`
	CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
}
