package usecases

import "log"

type ArticleInteractor interface {
	FindAll() (string, error)
}

type articleInteractor struct {
	articleRepository ArticleRepository
}

func NewArticleInteractor(ar ArticleRepository) *articleInteractor {
	return &articleInteractor{articleRepository: ar}
}

func (ai *articleInteractor) FindAll() (string, error) {
	articles, err := ai.articleRepository.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	return articles, nil
}
