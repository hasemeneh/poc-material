package repository

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

type Role interface {
	GetUserValidRole(ctx context.Context, userToken *models.UserToken) ([]*models.Role, error)
	GetRoleForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error)
	AddRole(ctx context.Context, dbTx database.DBTX, req *models.InternalRole) (int64, error)
	UpdateRole(ctx context.Context, req *models.InternalRole) error
	FetchUserAccessByIDInternal(ctx context.Context, userID int64) ([]*models.UserAccess, error)
	AddUserAccess(ctx context.Context, req *models.UserAccess) error
	ChangeStatusUserAccess(ctx context.Context, req *models.UserAccess) error
	RevokeUserRole(ctx context.Context, userToken *models.UserToken) error
}
