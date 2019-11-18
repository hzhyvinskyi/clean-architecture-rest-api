package interfaces

import (
	"log"
	"net/http"
)

type ArticleInteractor interface {
	FindAll() (string, error)
}

type ArticleController struct {
	ArticleInteractor ArticleInteractor
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
