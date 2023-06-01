package token

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/jmoiron/sqlx"
)

func (r *repo) GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	now := time.Now()
	refreshToken, err := r.queries.FetchRefreshToken(ctx, token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, custerr.ErrInvalidToken
		}
		return nil, err
	}
	if refreshToken.Status == constants.RefreshTokenStatusInactive {
		return nil, custerr.ErrInvalidToken
	}
	if refreshToken.ExpiresAt.Before(now) {
		return nil, custerr.ErrTokenExpired
	}
	return refreshToken, nil
}

func (r *repo) AddNewRefreshToken(ctx context.Context, token string, userID int64) error {
	dbtx, err := r.queries.StartTransaction(ctx)
	if err != nil {
		return err
	}
	expiresAt := time.Now().Add(constants.RefreshTokenExpirationTime)
	_, err = r.queries.InsertRefreshToken(ctx, dbtx, userID, token, &expiresAt, constants.RefreshTokenStatusActive)
	if err != nil {
		return err
	}
	err = r.RevokeOldRefreshToken(ctx, dbtx, userID)
	if err != nil {
		return err
	}
	return dbtx.Commit()
}

func (r *repo) RevokeOldRefreshToken(ctx context.Context, dbTx *sqlx.Tx, userID int64) error {
	var err error
	isSelfTransacted := false
	if dbTx == nil {
		dbTx, err = r.queries.StartTransaction(ctx)
		if err != nil {
			return err
		}
		defer dbTx.Rollback()
		isSelfTransacted = true
	}
	toBeInvalidatedToken, err := r.queries.FetchtToBeInvalidRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return err
	}
	for _, token := range toBeInvalidatedToken {
		err = r.queries.UpdateRefreshToken(ctx, dbTx, constants.RefreshTokenStatusInactive, token.ID)
		if err != nil {
			return err
		}
	}
	if isSelfTransacted {
		return dbTx.Commit()
	}
	return err
}

func (r *repo) RegisterLoginToken(ctx context.Context, user *models.User) (*models.UserToken, error) {
	resp := &models.UserToken{
		Token: &models.Token{
			AccessToken:  GenerateToken(user.ID),
			Expiry:       time.Now().Add(constants.ExpirationTime),
			RefreshToken: GenerateToken(user.ID),
		},
		User: user,
	}
	err := r.cache.InsertUserToken(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to cache user token because : %w", err)
	}

	err = r.AddNewRefreshToken(ctx, resp.RefreshToken, resp.User.ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *repo) GetUserByToken(ctx context.Context, token string) (*models.UserToken, error) {
	resp, err := r.cache.GetUserByToken(token)
	if err != nil {
		return nil, err
	}
	if resp.Expiry.Before(time.Now()) {
		return nil, custerr.ErrTokenExpired
	}
	return resp, nil
}

func GenerateToken(userID int64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", userID)))
	h.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (r *repo) RevokeToken(ctx context.Context, token string) error {
	return r.cache.RemoveUserToken(token)
}
