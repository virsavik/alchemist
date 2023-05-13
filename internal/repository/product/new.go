package product

import (
	"context"

	"io.github.virsavik/alchemist/internal/repository/base"
	"io.github.virsavik/alchemist/internal/repository/entity"
	"io.github.virsavik/alchemist/pkg/common/pagination"
)

type Reposity interface {
	All(ctx context.Context, input base.QueryInput, pgnInput pagination.Input) ([]*entity.Product, int64, error)

	Update(ctx context.Context, entity *entity.Product) error
}
