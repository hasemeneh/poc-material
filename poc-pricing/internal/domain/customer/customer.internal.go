package customer

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (r *repo) FetchUserDataByAuthID(ctx context.Context, authID string) (*models.Customer, error) {
	return r.queries.FetchUserDataByAuthID(ctx, authID)
}
