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
	GetUserAccess(authID string) ([]*models.Permission, error)
	RemoveUserAccess(authID string) error
	InsertUserAccess(authID string, tokenExpiry time.Time, permissions []*models.Permission) error
}

func NewCache(cache redis.RedisInterface) Resource {
	return &resource{
		cache: cache,
	}
}

const UserAccessKey = "access:%s"

func (r *resource) GetUserAccess(authID string) ([]*models.Permission, error) {
	resp := make([]*models.Permission, 0)
	isExist, err := r.cache.Exist(fmt.Sprintf(UserAccessKey, authID))
	if err != nil {
		return resp, err
	}
	if !isExist {
		return resp, custerr.ErrNoPermissionFound
	}
	jsonResp, err := r.cache.Get(fmt.Sprintf(UserAccessKey, authID))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(jsonResp), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *resource) InsertUserAccess(authID string, tokenExpiry time.Time, permissions []*models.Permission) error {
	now := time.Now()
	resp, _ := json.Marshal(permissions)
	_, err := r.cache.Set(fmt.Sprintf(UserAccessKey, authID), string(resp), tokenExpiry.Sub(now)+constants.ExpirationMargin)
	return err
}

func (r *resource) RemoveUserAccess(authID string) error {
	_, err := r.cache.Delete(fmt.Sprintf(UserAccessKey, authID))
	return err
}
