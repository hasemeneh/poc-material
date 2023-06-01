package repository

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
	"github.com/jmoiron/sqlx"
)

type Permission interface {
	AddBulkRolePermission(ctx context.Context, dbTx database.DBTX, req []models.InternalRolePermission, roleID int64) error
	RevokeUserAccess(ctx context.Context, userToken *models.UserToken) error
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
	GetAccessByUserTokenForInternal(ctx context.Context, userToken *models.UserToken) ([]*models.Permission, error)
	GetAccessByUserToken(ctx context.Context, userToken *models.UserToken) (map[int64][]*models.Permission, error)
	GetPermissionForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error)
	AddPermission(ctx context.Context, req *models.Permission) error
	UpdatePermission(ctx context.Context, req *models.Permission) error
	GetPermissionByRoleID(ctx context.Context, RoleID int64) ([]*models.Permission, error)
	ChangeRolePermissionStatus(ctx context.Context, status constants.RolePermissionStatus, RoleID, PermissionID int64) error
	AddRolePermission(ctx context.Context, req *models.InternalRolePermission) error
}
