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
	FetchProductBySKU(ctx context.Context, dbtx database.DBTX, sku string) ([]*models.ProductPrice, error)
	UpdateProductPrice(ctx context.Context, dbtx database.DBTX, v *models.ProductPrice) error
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

func (r *resource) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.GetMaster().BeginTxx(ctx, nil)

}

const selectProductBySKU = `
SELECT
	id,
	supplier_id,
	sku,
	price,
	stock,
	created_at,
	updated_at
FROM product_pricelist 
WHERE sku = ?
ORDER BY price DESC
FOR UPDATE`

func (r *resource) FetchProductBySKU(ctx context.Context, dbtx database.DBTX, sku string) ([]*models.ProductPrice, error) {
	resp := make([]*models.ProductPrice, 0)

	err := r.db.GetSelectContext(dbtx).SelectContext(ctx, &resp, selectProductBySKU,
		sku)
	return resp, err
}

const updateProductBySKU = `
UPDATE product_pricelist
	supplier_id = ?,
	sku = ?,
	price = ?,
	stock = ?
WHERE id = ?
`

func (r *resource) UpdateProductPrice(ctx context.Context, dbtx database.DBTX, v *models.ProductPrice) error {
	_, err := r.db.GetExecContext(dbtx).ExecContext(ctx, updateProductBySKU,
		v.SupplierID,
		v.SKU,
		v.Price,
		v.Stock,
		v.ID)
	return err
}
