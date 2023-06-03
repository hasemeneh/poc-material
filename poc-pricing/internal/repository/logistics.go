package repository

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/jmoiron/sqlx"
)

type Logistics interface {
	FetchLogisticDataByPlaceAndType(ctx context.Context, city, state, logisticType string) (*models.Logistics, error)
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}
