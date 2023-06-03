package logistics

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
)

func (r *repo) FetchLogisticDataByPlaceAndType(ctx context.Context, city, state, logisticType string) (*models.Logistics, error) {
	return r.queries.FetchLogisticDataByPlaceAndType(ctx, city, state, logisticType)
}
