package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindPost(post *models.Post) (*models.Post, error)
	FindPostBySlug(slug string) (*models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

type PRConfig struct {
	DB *gorm.DB
}

func NewPostRepository(c *PRConfig) PostRepository {
	return &postRepository{db: c.DB}
}

func (repo *postRepository) FindPost(post *models.Post) (*models.Post, error) {
	result := repo.db.
		First(&post)
	return post, result.Error
}

func (repo *postRepository) FindPostBySlug(slug string) (*models.Post, error) {
	post := &models.Post{}
	result := repo.db.
		Joins("PostTier").
		Joins("PostCategory").
		//Joins("Image").
		//Joins("left join users on users.id = posts.created_by_id").
		//Joins("left join users on users.id = posts.updated_by_id").
		Where("slug = ?", slug).
		First(&post)
	return post, result.Error
}
