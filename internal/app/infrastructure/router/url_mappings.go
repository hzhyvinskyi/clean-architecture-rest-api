package router

import (
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/interfaces"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/usecases"
)

var articleController interfaces.ArticleController

func init() {
	articleRepository := interfaces.NewArticleRepository()
	articleInteractor := usecases.NewArticleInteractor(articleRepository)
	articleController = interfaces.ArticleController{ArticleInteractor: articleInteractor}
}

func mapUrls() {

	router.HandleFunc("/articles", articleController.FindAll).Methods("GET")
}
