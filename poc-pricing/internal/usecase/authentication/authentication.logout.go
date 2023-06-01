package authentication

import (
	"context"
	"fmt"
)

func (m *module) Logout(ctx context.Context, token string) error {
	userToken, err := m.UserRepo.GetUserByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("[GetAllUserAccess] failed to get user by token : %w", err)
	}
	err = m.permission.RevokeUserAccess(ctx, userToken)
	if err != nil {
		return fmt.Errorf("[RevokeUserAccess] failed to revoke user acccess by token : %w", err)
	}
	err = m.role.RevokeUserRole(ctx, userToken)
	if err != nil {
		return fmt.Errorf("[RevokeUserRole] failed to revoke user role by token : %w", err)
	}
	return nil
}
