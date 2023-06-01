package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
)

type resource struct {
	cache redis.RedisInterface
}
type Resource interface {
	GetUserRole(authID string) ([]*models.Role, error)
	InsertUserRole(authID string, tokenExpiry time.Time, permissions []*models.Role) error
	RemoveUserRole(authID string) error
}

func NewCache(cache redis.RedisInterface) Resource {
	return &resource{
		cache: cache,
	}
}

const UserRoleKey = "role:%s"

func (r *resource) GetUserRole(authID string) ([]*models.Role, error) {
	resp := make([]*models.Role, 0)
	isExist, err := r.cache.Exist(fmt.Sprintf(UserRoleKey, authID))
	if err != nil {
		return resp, err
	}
	if !isExist {
		return resp, custerr.ErrNoPermissionFound
	}
	jsonResp, err := r.cache.Get(fmt.Sprintf(UserRoleKey, authID))
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal([]byte(jsonResp), &resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *resource) InsertUserRole(authID string, tokenExpiry time.Time, permissions []*models.Role) error {
	now := time.Now()
	resp, _ := json.Marshal(permissions)
	_, err := r.cache.Set(fmt.Sprintf(UserRoleKey, authID), string(resp), tokenExpiry.Sub(now)+constants.ExpirationMargin)
	return err
}
func (r *resource) RemoveUserRole(authID string) error {
	_, err := r.cache.Delete(fmt.Sprintf(UserRoleKey, authID))
	return err
}
