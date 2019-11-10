package interfaces

type articleRepository struct{}

func NewArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (ar *articleRepository) FindAll() (string, error) {
	return "Hello from Repo!\n", nil
}
