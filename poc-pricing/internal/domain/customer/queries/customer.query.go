package queries

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/jmoiron/sqlx"
)

type resource struct {
	db *database.Store
}
type Resource interface {
	FetchUserDataByAuthID(ctx context.Context, authID string) (*models.Customer, error)
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

func (r *resource) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.GetMaster().BeginTxx(ctx, nil)

}

const selectUserDataByAuthID = `
SELECT
	id,
	address,
	city,
	state,
	created_at,
	updated_at
FROM customer c
WHERE id = ?`

func (r *resource) FetchUserDataByAuthID(ctx context.Context, authID string) (*models.Customer, error) {
	resp := &models.Customer{}
	err := r.db.GetSlave().GetContext(ctx, &resp, selectUserDataByAuthID,
		authID)
	return resp, err
}
