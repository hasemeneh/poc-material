package cache

import (
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
)

type resource struct {
	cache redis.RedisInterface
}
type Resource interface {
}

func NewCache(cache redis.RedisInterface) Resource {
	return &resource{
		cache: cache,
	}
}
