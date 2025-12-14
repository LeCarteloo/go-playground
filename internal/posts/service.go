package posts

import "context"

type Service interface {
	ListPosts(ctx context.Context) error
}

type svc struct {
	// repository
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListPosts(ctx context.Context) error {
	return nil
}
