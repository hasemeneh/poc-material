package role

import (
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/role/cache"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/role/queries"
	"github.com/hasemeneh/poc-material/poc-login/internal/repository"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/redis"
)

type repo struct {
	queries queries.Resource
	cache   cache.Resource
}
type Args struct {
	Redis redis.RedisInterface
	DB    *database.Store
}

func New(o *Args) repository.Role {
	return &repo{
		queries: queries.NewQuery(o.DB),
		cache:   cache.NewCache(o.Redis),
	}
}
