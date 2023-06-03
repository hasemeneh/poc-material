package definition

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

type Intools interface {
	GetPermission(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error)
	AddPermission(ctx context.Context, req *models.Customer, token string) error
	UpdatePermission(ctx context.Context, req *models.Customer, token string) error

	GetRole(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error)
	UpdateRole(ctx context.Context, token string, req *models.InternalRole) error
	AddRole(ctx context.Context, req *models.AddRoleReq, token string) error

	AddPermissionToRole(ctx context.Context, token string, req *models.InternalRolePermission) error
	UpdateRolePermission(ctx context.Context, token string, req *models.InternalRolePermission) error
	GetPermissionByRole(ctx context.Context, roleID int64, token string) ([]*models.Customer, error)

	GetUser(ctx context.Context, queries url.Values, token string) (*pagespecifier.IntoolsResponse, error)
	AddUserAccess(ctx context.Context, req *models.UserAccess, token string) error
	ChangeStatusUserAccess(ctx context.Context, req *models.UserAccess, token string) error
	GetUserAccessByIDInternal(ctx context.Context, userID int64, token string) ([]*models.UserAccess, error)
}
