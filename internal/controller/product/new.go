package product

import (
	"context"

	"io.github.virsavik/alchemist/internal/model"
	"io.github.virsavik/alchemist/pkg/common/pagination"
)

type Controller interface {
	All(ctx context.Context, pgnInput pagination.Input) ([]model.Product, int64, error)
}
