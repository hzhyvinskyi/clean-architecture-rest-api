package router

import (
	"fmt"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/infrastructure/database"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/interfaces"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/usecases"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var articleController interfaces.ArticleController

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := database.NewDB(dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	articleRepository := interfaces.NewArticleRepository(db)
	articleInteractor := usecases.NewArticleInteractor(articleRepository)
	articleController = interfaces.ArticleController{articleInteractor}
}

func mapUrls() {
	router.HandleFunc("/articles", articleController.FindAll).Methods("GET")
}
