package interfaces

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
	"log"
	"net/http"
)

type ArticleInteractor interface {
	FindAll(context.Context) ([]*domain.Article, error)
	Find(context.Context, string) (*domain.Article, error)
}

type ArticleController struct {
	ArticleInteractor ArticleInteractor
}

func (ac *ArticleController) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	articles, err := ac.ArticleInteractor.FindAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.NewEncoder(w).Encode(articles); err != nil {
		log.Fatalln(err)
	}
}

func (ac *ArticleController) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(mux.Vars(r)["uuid"])
	article, err := ac.ArticleInteractor.Find(ctx, mux.Vars(r)["uuid"])
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.NewEncoder(w).Encode(article); err != nil {
		log.Fatalln(err)
	}
}
