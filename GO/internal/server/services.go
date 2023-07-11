package server

import (
	"context"

	"github.com/osmait/crud/internal/domain"
	storage "github.com/osmait/crud/internal/storage"
	// storage "github.com/osmait/crud/internal/storage"
)

type PostService struct {
	repository storage.PostRepository
}

func NewPostService(postRepostory storage.PostRepository) *PostService {
	return &PostService{
		repository: postRepostory,
	}
}

func (p *PostService) FindAll() []domain.Post {
	return p.repository.FindPost()
}

func (p *PostService) Created(ctx context.Context, post domain.Post) {
	p.repository.Save(ctx, post)
}
