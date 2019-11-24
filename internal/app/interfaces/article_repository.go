package interfaces

import (
	"context"
	"database/sql"
	"github.com/hzhyvinskyi/clean-architecture-rest-api/internal/app/domain"
)

type articleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) FindAll(ctx context.Context) ([]*domain.Article, error) {
	rows, err := ar.db.QueryContext(ctx, "SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*domain.Article, 0)
	for rows.Next() {
		article := new(domain.Article)
		err := rows.Scan(&article.ArticleID, &article.UserID, &article.Title, &article.Description, &article.ImageURL, &article.Likes, &article.Status, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (ar *articleRepository) Find(ctx context.Context, uuid string) (*domain.Article, error) {
	row := ar.db.QueryRowContext(ctx, "SELECT * FROM articles WHERE article_id = $1", uuid)
	article := new(domain.Article)
	err := row.Scan(&article.ArticleID, &article.UserID, &article.Title, &article.Description, &article.ImageURL, &article.Likes, &article.Status, &article.CreatedAt, &article.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return article, nil
}
