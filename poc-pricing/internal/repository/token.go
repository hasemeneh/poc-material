package repository

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"

	"github.com/jmoiron/sqlx"
)

type Token interface {
	GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error)
	AddNewRefreshToken(ctx context.Context, token string, userID int64) error
	RevokeOldRefreshToken(ctx context.Context, dbTx *sqlx.Tx, userID int64) error
	RegisterLoginToken(ctx context.Context, user *models.User) (*models.UserToken, error)
	GetUserByToken(ctx context.Context, token string) (*models.UserToken, error)
	RevokeToken(ctx context.Context, token string) error
}
