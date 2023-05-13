package product

import (
	"context"

	"io.github.virsavik/alchemist/internal/model"
	"io.github.virsavik/alchemist/internal/repository"
	"io.github.virsavik/alchemist/pkg/common/pagination"
)

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}

func (i impl) All(ctx context.Context, pgnInput pagination.Input) ([]model.Product, int64, error) {
	entities, total, err := i.repo.Product().All(ctx, nil, pgnInput)
	if err != nil {
		return nil, 0, err
	}

	result := make([]model.Product, len(entities))
	for idx, ent := range entities {
		result[idx] = model.Product{
			ID:        ent.ID,
			Name:      ent.Name,
			Sku:       ent.Sku,
			CreatedAt: ent.CreatedAt,
			UpdatedAt: ent.UpdatedAt,
		}
	}

	return result, total, nil
}
