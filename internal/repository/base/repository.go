package base

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io.github.virsavik/alchemist/internal/repository/entity"
	"io.github.virsavik/alchemist/pkg/common/pagination"
)

// Repository is an object that has implemented some data access method
type Repository[T entity.Entity] struct {
	db       *sql.DB
	newQuery func(mods ...qm.QueryMod) entity.Query[T]
}

// New function initialize a new Repository
func New[T entity.Entity](db *sql.DB, newQuery func(mods ...qm.QueryMod) entity.Query[T]) Repository[T] {
	return Repository[T]{
		db:       db,
		newQuery: newQuery,
	}
}

func (i Repository[T]) All(ctx context.Context, input QueryInput, pgnInput pagination.Input) ([]T, int64, error) {
	mods := []qm.QueryMod{}

	if input != nil {
		mods = append(mods, input.ToQueryMods()...)
	}

	var total int64
	if pgnInput.WithTotal {
		var err error
		total, err = i.newQuery(mods...).Count(ctx, i.db)
		if err != nil {
			return nil, 0, errors.WithStack(err)
		}
	}

	offset, limit := pgnInput.ToOffsetLimit()
	if pgnInput.Index > 0 {
		mods = append(mods, qm.Offset(offset))
	}

	if pgnInput.Size > 0 {
		mods = append(mods, qm.Limit(limit))
	}

	entities, err := i.newQuery(mods...).All(ctx, i.db)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	return entities, total, nil
}

func (i Repository[T]) Insert(ctx context.Context, o T) error {
	if err := o.Insert(ctx, i.db, boil.Infer()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (i Repository[T]) Update(ctx context.Context, o T) error {
	rowAff, err := o.Update(ctx, i.db, boil.Infer())
	if err != nil {
		return errors.WithStack(err)
	}

	if rowAff < 1 {
		return errors.New("cannot update")
	}

	return nil
}
