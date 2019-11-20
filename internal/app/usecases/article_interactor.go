package usecases

import (
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
)

type ArticleRepository interface {
	FindAll() ([]*domain.Article, error)
}

type articleInteractor struct {
	articleRepository ArticleRepository
}

func NewArticleInteractor(ar ArticleRepository) *articleInteractor {
	return &articleInteractor{articleRepository: ar}
}

func (ai *articleInteractor) FindAll() ([]*domain.Article, error) {
	articles, err := ai.articleRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
