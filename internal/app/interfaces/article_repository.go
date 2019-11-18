package interfaces

import (
	"database/sql"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
)

type articleRepository struct{
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) FindAll() ([]*domain.Article, error) {
	rows, err := ar.db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*domain.Article, 0)
	for rows.Next() {
		article := new(domain.Article)
		err := rows.Scan(&article.ArticleID, &article.UserID, &article.Title, &article.Title, &article.Description, &article.ImageURL, &article.Likes, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
