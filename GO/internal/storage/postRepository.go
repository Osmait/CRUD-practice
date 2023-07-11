package storage

import (
	"context"

	"github.com/osmait/crud/internal/domain"
	filemanager "github.com/osmait/crud/internal/fileManager"
)

type PostRepository struct {
	Postlist    []domain.Post
	filemanager *filemanager.FileManager
}

func NewPostRepository(file *filemanager.FileManager) *PostRepository {
	return &PostRepository{
		filemanager: file,
	}
}

func (p *PostRepository) Save(ctx context.Context, post domain.Post) {
	p.filemanager.Write(post)

}

func (p *PostRepository) FindPost() []domain.Post {
	result := p.filemanager.Get()
	return result

}
