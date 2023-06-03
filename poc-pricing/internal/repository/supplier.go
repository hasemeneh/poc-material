package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Supplier interface {
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
}
