package supplier

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (r *repo) FetchSupplierByID(ctx context.Context, authID string) (*models.Supplier, error) {
	return r.queries.FetchSupplierByID(ctx, authID)
}
