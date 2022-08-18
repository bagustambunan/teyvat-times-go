package repositories

import "C"
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
	// TEST VIA RAW
	//query := "SELECT posts.id,posts.post_tier_id,posts.post_category_id,posts.title,posts.content,posts.slug,posts.summary,posts.img_thumbnail_id,posts.img_content_id,posts.created_by_id,posts.updated_by_id,"
	//query += " C.id AS PostCategory__id,C.created_at AS PostCategory__created_at,C.updated_at AS PostCategory__updated_at,C.deleted_at AS PostCategory__deleted_at,C.name AS PostCategory__name"
	//query += " FROM posts LEFT JOIN post_categories C ON posts.post_category_id = C.id AND C.deleted_at IS NULL"
	//
	//post := &models.Post{}
	//result := repo.db.
	//	Raw(query).
	//	Scan(&post)
	//return post, result.Error

	// GET SQL QUERY
	//post := &models.Post{}
	//sql := repo.db.
	//	ToSQL(func(tx *gorm.DB) *gorm.DB {
	//		return tx.
	//			//Joins("PostTier").
	//			Joins("PostCategory").
	//			Where("slug = ?", slug).
	//			First(&post)
	//
	//	})
	//fmt.Println("================ THE SQL ==================")
	//fmt.Println(sql)
	//return nil, nil

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
