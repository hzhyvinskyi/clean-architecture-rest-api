package usecases

type ArticleRepository interface {
	FindAll() (string, error)
}
