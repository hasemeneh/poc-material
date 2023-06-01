package authorization

import (
	"context"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (m *module) GetAllUserAccess(ctx context.Context, token string) ([]models.RolePermission, error) {
	resp := make([]models.RolePermission, 0)
	userToken, err := m.user.GetUserByToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("[GetAllUserAccess] failed to get user by token : %w", err)
	}
	permissions, err := m.permission.GetAccessByUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("[GetAllUserAccess] failed to get permissions by token : %w", err)
	}
	roles, err := m.role.GetUserValidRole(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("[GetAllUserAccess] failed to get role by token : %w", err)
	}
	for _, role := range roles {
		permission := permissions[role.ID]
		resp = append(resp, models.RolePermission{
			Role:        role,
			Permissions: permission,
		})
	}
	return resp, nil
}
