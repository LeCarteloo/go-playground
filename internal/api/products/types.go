package products

import (
	"context"

	repo "go_playground/internal/adapters/postgresql/sqlc"
)

type ProductRepository interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
}

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
}
