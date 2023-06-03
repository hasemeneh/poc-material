package productprice

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
)

func (r *repo) LockProductBySKU(ctx context.Context, dbtx database.DBTX, sku string) ([]*models.ProductPrice, error) {
	return r.queries.FetchProductBySKU(ctx, dbtx, sku)
}

func (r *repo) UpdateProductPrice(ctx context.Context, dbtx database.DBTX, v *models.ProductPrice) error {
	return r.queries.UpdateProductPrice(ctx, dbtx, v)
}
