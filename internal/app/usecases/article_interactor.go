package usecases

import (
	"context"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
)

type ArticleRepository interface {
	FindAll(context.Context) ([]*domain.Article, error)
	Find(context.Context, string) (*domain.Article, error)
}

type articleInteractor struct {
	articleRepository ArticleRepository
}

func NewArticleInteractor(ar ArticleRepository) *articleInteractor {
	return &articleInteractor{articleRepository: ar}
}

func (ai *articleInteractor) FindAll(ctx context.Context) ([]*domain.Article, error) {
	articles, err := ai.articleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (ai *articleInteractor) Find(ctx context.Context, uuid string) (*domain.Article, error) {
	article, err := ai.articleRepository.Find(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return article, err
}
