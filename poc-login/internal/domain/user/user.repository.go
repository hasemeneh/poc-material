package user

import (
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/user/cache"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/user/queries"
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

func New(o *Args) repository.User {
	return &repo{
		queries: queries.NewUserQuery(o.DB),
		cache:   cache.NewUserCache(o.Redis),
	}
}
