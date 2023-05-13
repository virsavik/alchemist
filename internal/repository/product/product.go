package product

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io.github.virsavik/alchemist/internal/repository/base"
	"io.github.virsavik/alchemist/internal/repository/entity"
)

type impl struct {
	db *sql.DB
	base.Repository[*entity.Product]
}

func New(db *sql.DB) Reposity {
	return impl{
		db: db,
		Repository: base.New(
			db,
			func(mods ...qm.QueryMod) entity.Query[*entity.Product] {
				return entity.NewProductQuery(mods...)
			},
		),
	}
}

type GetProductInput struct {
}

func (input GetProductInput) ToQueryMods() []qm.QueryMod {
	return []qm.QueryMod{}
}
