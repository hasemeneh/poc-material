package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *repo) RegisterUser(ctx context.Context, authID string, password string) error {
	_, err := r.queries.FetchUserByAuthID(ctx, nil, authID)
	if err == nil {
		return custerr.ErrUserAlreadyRegistered
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	hasehedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.queries.Insert(ctx, nil, &models.User{
		AuthID:   authID,
		Status:   constants.UserStatusActive,
		Password: hasehedPassword,
	})
	return err
}

func (r *repo) GetUserByAuthID(ctx context.Context, authID string) (*models.User, error) {
	user, err := r.queries.FetchUserByAuthID(ctx, nil, authID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, custerr.ErrUserNotFound
		}
		return nil, err
	}
	if user.Status == constants.UserStatusInactive {
		return nil, custerr.ErrInactiveUser
	}
	return user, nil
}

func (r *repo) ValidatePassword(ctx context.Context, user *models.User, password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return custerr.ErrWrongPassword
		}
		return err
	}
	return nil
}
func (r *repo) RevokeToken(ctx context.Context, token string) error {
	return r.cache.RemoveUserToken(token)
}

func (r *repo) GetUserByUserID(ctx context.Context, userID int64) (*models.User, error) {
	user, err := r.queries.FetchUserByID(ctx, nil, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, custerr.ErrUserNotFound
		}
		return nil, err
	}
	if user.Status == constants.UserStatusInactive {
		return nil, custerr.ErrInactiveUser
	}
	return user, nil
}
