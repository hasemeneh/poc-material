package permission

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-login/internal/domain/permission/cache"
	"github.com/hasemeneh/poc-material/poc-login/internal/domain/permission/queries"
	"github.com/hasemeneh/poc-material/poc-login/internal/repository"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/hasemeneh/poc-material/poc-login/pkg/redis"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	queries queries.Resource
	cache   cache.Resource
}
type Args struct {
	Redis redis.RedisInterface
	DB    *database.Store
}

func New(o *Args) repository.Permission {
	return &repo{
		queries: queries.NewQuery(o.DB),
		cache:   cache.NewCache(o.Redis),
	}
}

func (r *repo) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.queries.StartTransaction(ctx)

}
