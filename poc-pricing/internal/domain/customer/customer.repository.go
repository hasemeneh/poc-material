package customer

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/domain/customer/cache"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/domain/customer/queries"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/repository"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/redis"
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

func New(o *Args) repository.Customer {
	return &repo{
		queries: queries.NewQuery(o.DB),
		cache:   cache.NewCache(o.Redis),
	}
}

func (r *repo) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.queries.StartTransaction(ctx)

}
