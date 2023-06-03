package repository

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/jmoiron/sqlx"
)

type Customer interface {
	FetchUserDataByAuthID(ctx context.Context, authID string) (*models.Customer, error)
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}
