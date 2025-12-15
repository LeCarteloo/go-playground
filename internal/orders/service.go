package orders

import (
	"context"

	repo "go_playground/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListOrders(ctx context.Context) ([]repo.Order, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo,
	}
}

func (s *svc) ListOrders(ctx context.Context) ([]repo.Order, error) {
	return s.repo.ListOrders(ctx)
}
