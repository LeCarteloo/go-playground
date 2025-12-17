package orders

import (
	"context"

	repo "go_playground/internal/adapters/postgresql/sqlc"
	"go_playground/internal/apperrors"

	"github.com/jackc/pgx/v5"
)

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo,
		db,
	}
}

func (s *svc) ListOrders(ctx context.Context) ([]repo.Order, error) {
	return s.repo.ListOrders(ctx)
}

func (s *svc) CreateOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerID == 0 {
		return repo.Order{}, apperrors.ErrInvalidCustomerID
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, apperrors.ErrNoOrderItems
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	order, err := qtx.CreateOrder(ctx, tempOrder.CustomerID)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range tempOrder.Items {
		product, err := qtx.GetProductById(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, apperrors.ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, apperrors.ErrInsufficientProductQuantity
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			PriceCents: product.PriceCents,
		})
		if err != nil {
			return repo.Order{}, err
		}

		// TODO: Update product quantity
	}

	tx.Commit(ctx)

	return order, nil
}
