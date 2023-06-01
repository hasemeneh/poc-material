package models

import (
	"time"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
)

type RefreshToken struct {
	ID        int64                        `db:"id" json:"id"`
	UserID    int64                        `db:"user_id" json:"user_id"`
	Token     string                       `db:"token" json:"token"`
	Status    constants.RefreshTokenStatus `db:"status" json:"status"`
	ExpiresAt time.Time                    `db:"expires_at" json:"expires_at"`
}
