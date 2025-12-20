package products

import (
	"context"
	"errors"
	"testing"

	repo "go_playground/internal/adapters/postgresql/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type mockQuerier struct {
	listProducts    func(ctx context.Context) ([]repo.Product, error)
	getProductById  func(ctx context.Context, id int64) (repo.Product, error)
	createOrder     func(ctx context.Context, customerID int64) (repo.Order, error)
	createOrderItem func(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error)
	listOrders      func(ctx context.Context) ([]repo.Order, error)
}

func (m *mockQuerier) ListProducts(ctx context.Context) ([]repo.Product, error) {
	if m.listProducts != nil {
		return m.listProducts(ctx)
	}
	return nil, nil
}

func (m *mockQuerier) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	if m.getProductById != nil {
		return m.getProductById(ctx, id)
	}
	return repo.Product{}, nil
}

func (m *mockQuerier) CreateOrder(ctx context.Context, customerID int64) (repo.Order, error) {
	if m.createOrder != nil {
		return m.createOrder(ctx, customerID)
	}
	return repo.Order{}, nil
}

func (m *mockQuerier) CreateOrderItem(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error) {
	if m.createOrderItem != nil {
		return m.createOrderItem(ctx, arg)
	}
	return repo.OrderItem{}, nil
}

func (m *mockQuerier) ListOrders(ctx context.Context) ([]repo.Order, error) {
	if m.listOrders != nil {
		return m.listOrders(ctx)
	}
	return nil, nil
}

func TestListProducts(t *testing.T) {
	ctx := context.Background()

	t.Run("successfully returns products", func(t *testing.T) {
		expectedProducts := []repo.Product{
			{
				ID:         1,
				Name:       "Product 1",
				PriceCents: 1099,
				Quantity:   10,
				CreatedAt:  pgtype.Timestamptz{},
			},
		}

		mock := &mockQuerier{
			listProducts: func(ctx context.Context) ([]repo.Product, error) {
				return expectedProducts, nil
			},
		}

		service := NewService(mock)
		products, err := service.ListProducts(ctx)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(products) != len(expectedProducts) {
			t.Errorf("expected %d products, got %d", len(expectedProducts), len(products))
		}

		if products[0].ID != 1 || products[0].Name != "Product 1" {
			t.Errorf("got unexpected first product: %+v", products[0])
		}
	})

	t.Run("returns empty slice when no products exists", func(t *testing.T) {
		mockService := &mockQuerier{
			listProducts: func(ctx context.Context) ([]repo.Product, error) {
				return nil, nil
			},
		}

		service := NewService(mockService)
		products, err := service.ListProducts(ctx)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if products == nil {
			t.Errorf("expected empty array, got %v", products)
		}
	})

	t.Run("gracefully handle database error", func(t *testing.T) {
		mockService := &mockQuerier{
			listProducts: func(ctx context.Context) ([]repo.Product, error) {
				return nil, errors.New("database error")
			},
		}

		service := NewService(mockService)
		_, err := service.ListProducts(ctx)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
