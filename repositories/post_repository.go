package repositories

import "C"
import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

type PostRepository interface {
	FindPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error)
	FindPost(post *models.Post) (*models.Post, error)
	FindPostBySlug(slug string) (*models.Post, error)
	Save(post *models.Post) (*models.Post, int, error)
	FindUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error)
	SaveUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error)
	FindActivity(act *models.UserPostActivities) (*models.UserPostActivities, error)
	SaveActivity(act *models.UserPostActivities) (*models.UserPostActivities, error)
	IncreaseViewsActivity(act *models.UserPostActivities) (*models.UserPostActivities, error)
	UpdateActivity(act *models.UserPostActivities) (*models.UserPostActivities, error)
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

func (repo *postRepository) FindPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error) {
	var posts []*models.Post
	result := repo.db.
		Table("posts")

	if opt.Category != 0 {
		result = result.
			Where("post_category_id = ?", opt.Category)
	}
	if opt.Tier != 0 {
		result = result.
			Where("post_tier_id = ?", opt.Tier)
	}
	if opt.S != "" {
		result = result.
			Where("title ILIKE ?", "%"+opt.S+"%")
	}
	orderQuery := fmt.Sprintf("%s %s", opt.SortBy, opt.SortOrder)
	result = result.
		Order(orderQuery).
		Find(&posts)
	totalData := int(result.RowsAffected)

	result = result.
		Limit(opt.Limit).
		Offset((opt.Page - 1) * opt.Limit)

	postsRes := new(dto.GetPostsRes).FromPosts(posts)
	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	postsRes.SetValues(int(result.RowsAffected), opt.Limit, opt.Page, totalPage, totalData)
	return postsRes, result.Error
}

func (repo *postRepository) FindPost(post *models.Post) (*models.Post, error) {
	result := repo.db.
		First(&post)
	return post, result.Error
}

func (repo *postRepository) FindPostBySlug(slug string) (*models.Post, error) {
	// TODO: FIX THIS

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

func (repo *postRepository) FindUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error) {
	result := repo.db.
		Where("user_id = ?", unlock.UserID).
		Where("post_id = ?", unlock.PostID).
		First(&unlock)
	return unlock, result.Error
}

func (repo *postRepository) SaveUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error) {
	result := repo.db.
		Create(unlock)
	return unlock, result.Error
}

func (repo *postRepository) FindActivity(act *models.UserPostActivities) (*models.UserPostActivities, error) {
	result := repo.db.
		Where("user_id = ?", act.UserID).
		Where("post_id = ?", act.PostID).
		First(&act)
	return act, result.Error
}

func (repo *postRepository) SaveActivity(act *models.UserPostActivities) (*models.UserPostActivities, error) {
	result := repo.db.
		Create(act)
	return act, result.Error
}

func (repo *postRepository) IncreaseViewsActivity(act *models.UserPostActivities) (*models.UserPostActivities, error) {
	result := repo.db.
		Model(&act).
		UpdateColumn("views_count", gorm.Expr("views_count + ?", 1))
	return act, result.Error
}

func (repo *postRepository) UpdateActivity(act *models.UserPostActivities) (*models.UserPostActivities, error) {
	result := repo.db.
		Model(&act).
		UpdateColumn("is_liked", act.IsLiked).
		UpdateColumn("is_shared", act.IsShared)
	return act, result.Error
}

func (repo *postRepository) Save(post *models.Post) (*models.Post, int, error) {
	result := repo.db.
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(post)
	return post, int(result.RowsAffected), result.Error
}
