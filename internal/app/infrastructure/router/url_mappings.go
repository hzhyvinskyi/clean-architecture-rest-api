package router

import (
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/infrastructure/database"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/interfaces"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/usecases"
	"log"
)

var articleController interfaces.ArticleController

func init() {
	db, err := database.Conn()
	if err != nil {
		log.Fatalln(err)
	}

	articleRepository := interfaces.NewArticleRepository(db)
	articleInteractor := usecases.NewArticleInteractor(articleRepository)
	articleController = interfaces.ArticleController{ArticleInteractor: articleInteractor}
}

func mapUrls() {

	router.HandleFunc("/articles", articleController.FindAll).Methods("GET")
}
