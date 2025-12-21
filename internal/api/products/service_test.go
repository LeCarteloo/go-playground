package products

import (
	"context"
	"errors"
	"testing"

	repo "go_playground/internal/adapters/postgresql/sqlc"
	"go_playground/internal/apperrors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockRepo struct {
	listProducts   func(ctx context.Context) ([]repo.Product, error)
	getProductById func(ctx context.Context, id int64) (repo.Product, error)
}

func (m *mockRepo) ListProducts(ctx context.Context) ([]repo.Product, error) {
	if m.listProducts != nil {
		return m.listProducts(ctx)
	}
	return nil, nil
}

func (m *mockRepo) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	if m.getProductById != nil {
		return m.getProductById(ctx, id)
	}
	return repo.Product{}, nil
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

		mock := &mockRepo{
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
		mockService := &mockRepo{
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
		mockService := &mockRepo{
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

func TestGetProductById(t *testing.T) {
	ctx := context.Background()

	t.Run("successfully return a product", func(t *testing.T) {
		expectedProduct := repo.Product{
			ID:         1,
			Name:       "Product 1",
			PriceCents: 1099,
			Quantity:   10,
			CreatedAt:  pgtype.Timestamptz{},
		}

		mock := &mockRepo{
			getProductById: func(ctx context.Context, id int64) (repo.Product, error) {
				return expectedProduct, nil
			},
		}

		service := NewService(mock)
		product, err := service.GetProductById(ctx, 1)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if product.ID != expectedProduct.ID || product.Name != expectedProduct.Name {
			t.Errorf("expected product %+v, got %+v", expectedProduct, product)
		}
	})

	t.Run("returns ProductNotFound error when product not found", func(t *testing.T) {
		mock := &mockRepo{
			getProductById: func(ctx context.Context, id int64) (repo.Product, error) {
				return repo.Product{}, pgx.ErrNoRows
			},
		}

		service := NewService(mock)
		_, error := service.GetProductById(ctx, 1)

		if !errors.Is(error, apperrors.ErrProductNotFound) {
			t.Fatalf("expected ProductNotFound error, got %v", error)
		}
	})

	t.Run("gracefully handle database error", func(t *testing.T) {
		mock := &mockRepo{
			getProductById: func(ctx context.Context, id int64) (repo.Product, error) {
				return repo.Product{}, errors.New("database error")
			},
		}

		service := NewService(mock)
		_, err := service.GetProductById(ctx, 1)

		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
