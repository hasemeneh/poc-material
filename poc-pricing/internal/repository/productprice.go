package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ProductPrice interface {
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}
