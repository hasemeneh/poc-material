package intools

import (
	"context"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-login/internal/custerr"
)

func (m *module) IsValidUser(ctx context.Context, token string, permissionCode string) error {
	userToken, err := m.user.GetUserByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("[GetAllUserAccess] failed to get user by token : %w", err)
	}
	permissions, err := m.permission.GetAccessByUserTokenForInternal(ctx, userToken)
	if err != nil {
		return fmt.Errorf("[GetAllUserAccess] failed to get permissions by token : %w", err)
	}
	for _, permission := range permissions {
		if permission.Code == permissionCode {
			return nil
		}
	}
	return custerr.ErrInvalidAccess
}
