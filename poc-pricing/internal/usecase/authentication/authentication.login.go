package authentication

import (
	"context"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (u *module) Login(ctx context.Context, authID, password string) (*models.UserToken, error) {
	user, err := u.UserRepo.GetUserByAuthID(ctx, authID)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to get User : %w", err)
	}
	err = u.UserRepo.ValidatePassword(ctx, user, password)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to validate password : %w", err)
	}
	resp, err := u.token.RegisterLoginToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[authentication.Login] failed to register token : %w", err)
	}
	return resp, nil
}

func (u *module) RefreshToken(ctx context.Context, token string) (*models.UserToken, error) {
	refreshToken, err := u.token.GetRefreshToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("[authentication.RefreshToken] failed to get refresh token : %w", err)
	}
	user, err := u.UserRepo.GetUserByUserID(ctx, refreshToken.UserID)
	if err != nil {
		return nil, fmt.Errorf("[authentication.RefreshToken] failed to get user by id : %w", err)
	}
	resp, err := u.token.RegisterLoginToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[authentication.RefreshToken] failed to register token : %w", err)
	}
	return resp, nil
}
