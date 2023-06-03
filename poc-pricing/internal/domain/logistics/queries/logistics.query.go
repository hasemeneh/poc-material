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
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
	FetchLogisticDataByPlaceAndType(ctx context.Context, city, state, logisticType string) (*models.Logistics, error)
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

func (r *resource) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.GetMaster().BeginTxx(ctx, nil)

}

const selectLogisticDataByPlaceAndType = `
SELECT
	id,
	type,
	capacity,
	cost,
	city,
	state,
	created_at,
	updated_at
FROM logistics 
WHERE type = ? AND city = ? AND state = ?`

func (r *resource) FetchLogisticDataByPlaceAndType(ctx context.Context, city, state, logisticType string) (*models.Logistics, error) {
	resp := &models.Logistics{}
	err := r.db.GetSlave().GetContext(ctx, &resp, selectLogisticDataByPlaceAndType,
		logisticType, city, state)
	return resp, err
}
