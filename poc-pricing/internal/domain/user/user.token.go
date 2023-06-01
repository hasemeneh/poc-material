package user

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (r *repo) RegisterToken(ctx context.Context, user *models.User) (*models.UserToken, error) {
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
