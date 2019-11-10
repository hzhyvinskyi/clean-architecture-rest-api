package interfaces

import (
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/usecases"
	"log"
	"net/http"
)

type ArticleController struct {
	ArticleInteractor usecases.ArticleInteractor
}

func (ac *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	articles, err := ac.ArticleInteractor.FindAll()
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := w.Write([]byte(articles)); err != nil {
		log.Fatalln(err)
	}
}
