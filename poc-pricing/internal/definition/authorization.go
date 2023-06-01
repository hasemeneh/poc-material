package definition

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

type Authorization interface {
	GetAllUserAccess(ctx context.Context, token string) ([]models.RolePermission, error)
}
