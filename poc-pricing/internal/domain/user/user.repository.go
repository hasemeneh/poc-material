package user

import (
	"github.com/hasemeneh/poc-material/poc-pricing/internal/domain/user/cache"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/domain/user/queries"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/repository"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
)

type repo struct {
	queries queries.Resource
	cache   cache.Resource
}
type Args struct {
	Redis redis.RedisInterface
	DB    *database.Store
}

func New(o *Args) repository.User {
	return &repo{
		queries: queries.NewUserQuery(o.DB),
		cache:   cache.NewUserCache(o.Redis),
	}
}
