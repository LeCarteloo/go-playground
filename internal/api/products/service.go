package products

import (
	"context"
	"errors"

	repo "go_playground/internal/adapters/postgresql/sqlc"
	"go_playground/internal/apperrors"

	"github.com/jackc/pgx/v5"
)

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.ListProducts(ctx)
	if err != nil {
		return []repo.Product{}, err
	}

	if products == nil {
		return []repo.Product{}, nil
	}

	return products, nil
}

func (s *svc) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.GetProductById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.Product{}, apperrors.ErrProductNotFound
		}

		return repo.Product{}, err
	}

	return product, nil
}
