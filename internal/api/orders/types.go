package orders

import (
	"context"

	repo "go_playground/internal/adapters/postgresql/sqlc"

	"github.com/jackc/pgx/v5"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]repo.Order, error)
	CreateOrder(ctx context.Context, customerID int64) (repo.Order, error)
	WithTx(tx pgx.Tx) *repo.Queries
}

type orderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type createOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []orderItem `json:"items"`
}

type Service interface {
	ListOrders(ctx context.Context) ([]repo.Order, error)
	CreateOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}
