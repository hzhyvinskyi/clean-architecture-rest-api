package interfaces

import (
	"encoding/json"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
	"log"
	"net/http"
)

type ArticleInteractor interface {
	FindAll() ([]*domain.Article, error)
}

type ArticleController struct {
	ArticleInteractor ArticleInteractor
}

func (ac *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	articles, err := ac.ArticleInteractor.FindAll()
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.NewEncoder(w).Encode(articles); err != nil {
		log.Fatalln(err)
	}
}
