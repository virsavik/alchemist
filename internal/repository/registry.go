package repository

import (
	"database/sql"

	"io.github.virsavik/alchemist/internal/repository/product"
)

type Registry interface {
	Product() product.Reposity
}

type impl struct {
	db          *sql.DB
	productRepo product.Reposity
}

func New(db *sql.DB) Registry {
	return impl{
		db:          db,
		productRepo: product.New(db),
	}
}

func (i impl) Product() product.Reposity {
	return i.productRepo
}
