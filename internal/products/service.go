package products

import (
	"context"
	"errors"

	repo "go_playground/internal/adapters/postgresql/sqlc"

	"github.com/jackc/pgx/v5"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.GetProductById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.Product{}, errors.New("product not found")
		}
		return repo.Product{}, err
	}
	return product, nil
}
