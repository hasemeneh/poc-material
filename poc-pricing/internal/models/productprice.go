package models

import "database/sql"

type ProductPrice struct {
	ID         string       `json:"id,omitempty" db:"id"`
	SupplierID string       `json:"supplier_id" db:"supplier_id"`
	Price      int64        `json:"price" db:"price"`
	Stock      int64        `json:"stock" db:"stock"`
	SKU        string       `json:"sku" db:"sku"`
	CreatedAt  sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at" db:"updated_at"`
}
