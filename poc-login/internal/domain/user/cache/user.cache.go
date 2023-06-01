package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/redis"
)

type resource struct {
	cache redis.RedisInterface
}
type Resource interface {
	GetUserByToken(token string) (*models.UserToken, error)
	RemoveUserToken(token string) error
	InsertUserToken(u *models.UserToken) error
}

func NewUserCache(cache redis.RedisInterface) Resource {
	return &resource{
		cache: cache,
	}
}

const UserTokenKey = "token:%s"

func (r *resource) GetUserByToken(token string) (*models.UserToken, error) {
	isExist, err := r.cache.Exist(fmt.Sprintf(UserTokenKey, token))
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, custerr.ErrInvalidToken
	}
	jsonResp, err := r.cache.Get(fmt.Sprintf(UserTokenKey, token))
	if err != nil {
		return nil, err
	}
	resp := &models.UserToken{}
	err = json.Unmarshal([]byte(jsonResp), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *resource) InsertUserToken(u *models.UserToken) error {
	now := time.Now()
	resp, _ := json.Marshal(u)
	_, err := r.cache.Set(fmt.Sprintf(UserTokenKey, u.AccessToken), string(resp), u.Expiry.Sub(now)+constants.ExpirationMargin)
	return err
}
func (r *resource) RemoveUserToken(token string) error {
	_, err := r.cache.Delete(fmt.Sprintf(UserTokenKey, token))
	return err
}
