package model

import (
	"time"
)

type Product struct {
	ID        int64
	Name      string
	Sku       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Product) TableName() string {
	return "product"
}

func (p Product) String() string {
	return "Product{" +
		"id:" + string(rune(p.ID)) + "," +
		"name:" + p.Name + "," +
		"created_at:" + p.CreatedAt.Format(time.RFC3339) +
		"updated_at:" + p.UpdatedAt.Format(time.RFC3339) +
		"}"
}
