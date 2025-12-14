package posts

import (
	"context"

	repo "go_playground/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListPosts(ctx context.Context) ([]repo.Post, error)
	GetPostById(ctx context.Context, id int64) (repo.Post, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo,
	}
}

func (s *svc) ListPosts(ctx context.Context) ([]repo.Post, error) {
	return s.repo.ListPosts(ctx)
}

func (s *svc) GetPostById(ctx context.Context, id int64) (repo.Post, error) {
	return s.repo.GetPostById(ctx, id)
}
