package repository

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

type User interface {
	RegisterUser(ctx context.Context, authID string, password string) error
	GetUserByAuthID(ctx context.Context, authID string) (*models.User, error)
	ValidatePassword(ctx context.Context, user *models.User, password string) error
	RegisterToken(ctx context.Context, user *models.User) (*models.UserToken, error)
	GetUserByToken(ctx context.Context, token string) (*models.UserToken, error)
	GetUserForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error)
	GetUserByUserID(ctx context.Context, userID int64) (*models.User, error)
}
