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
	FetchSupplierByID(ctx context.Context, id string) (*models.Supplier, error)
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

const selectSupplierByID = `
SELECT
	id,
	address,
	city,
	state,
	created_at,
	updated_at
FROM supplier 
WHERE id = ?`

func (r *resource) FetchSupplierByID(ctx context.Context, id string) (*models.Supplier, error) {
	resp := &models.Supplier{}
	err := r.db.GetSlave().GetContext(ctx, &resp, selectSupplierByID,
		id)
	return resp, err
}
