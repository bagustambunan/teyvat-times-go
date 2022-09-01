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
	FindReadingHistory(user *models.User, opt *models.ReadHistoryOption) (*dto.ReadHistoryRes, error)
	FindPostBySlug(slug string) (*models.Post, error)
	Save(post *models.Post) (*models.Post, int, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePost(post *models.Post) error
	FindUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error)
	SaveUnlock(unlock *models.PostUnlock) (*models.PostUnlock, error)
	FindActivity(act *models.UserPostActivity) (*models.UserPostActivity, error)
	SaveActivity(act *models.UserPostActivity) (*models.UserPostActivity, error)
	IncreaseViewsActivity(act *models.UserPostActivity) (*models.UserPostActivity, error)
	UpdateActivity(act *models.UserPostActivity) (*models.UserPostActivity, error)
	FindUserLatestSubscription(user *models.User) (*models.UserSubscription, error)
	FindTiers() ([]*models.PostTier, error)
	FindCategories() ([]*models.PostCategory, error)
	FindTier(tier *models.PostTier) (*models.PostTier, error)
	FindCategory(category *models.PostCategory) (*models.PostCategory, error)
	SaveCategory(category *models.PostCategory) (*models.PostCategory, error)
	UpdateCategory(category *models.PostCategory) (*models.PostCategory, error)
	FindTrending() (any, error)
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
		Table("posts").
		Joins("PostTier").
		Joins("PostCategory").
		Joins("ImgThumbnail").
		Joins("ImgContent").
		Joins("CreatedBy").
		Joins("UpdatedBy")

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
		Offset((opt.Page - 1) * opt.Limit).
		Find(&posts)

	postsRes := new(dto.GetPostsRes).FromPosts(posts)
	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	postsRes.SetValues(int(result.RowsAffected), opt.Limit, opt.Page, totalPage, totalData)
	return postsRes, result.Error
}

func (repo *postRepository) FindPost(post *models.Post) (*models.Post, error) {
	result := repo.db.
		Joins("PostTier").
		Joins("PostCategory").
		Joins("ImgThumbnail").
		Joins("ImgContent").
		Joins("CreatedBy").
		Joins("UpdatedBy").
		First(&post)
	return post, result.Error
}

func (repo *postRepository) FindReadingHistory(user *models.User, opt *models.ReadHistoryOption) (*dto.ReadHistoryRes, error) {
	var activities []*models.UserPostActivity
	result := repo.db.
		Preload("Post.PostTier").
		Preload("Post.PostCategory").
		Preload("Post.ImgThumbnail").
		Preload("Post.ImgContent").
		Preload("Post.CreatedBy").
		Preload("Post.UpdatedBy").
		Preload(clause.Associations).
		Where("user_id", user.ID).
		Order("updated_at DESC").
		Find(&activities)
	totalData := int(result.RowsAffected)

	result = result.
		Limit(opt.Limit).
		Offset((opt.Page - 1) * opt.Limit).
		Find(&activities)

	readHistoryRes := new(dto.ReadHistoryRes).FromActivities(activities)
	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	readHistoryRes.SetValues(int(result.RowsAffected), opt.Limit, opt.Page, totalPage, totalData)
	return readHistoryRes, result.Error
}

func (repo *postRepository) FindPostBySlug(slug string) (*models.Post, error) {
	post := &models.Post{}
	result := repo.db.
		Joins("PostTier").
		Joins("PostCategory").
		Joins("ImgThumbnail").
		Joins("ImgContent").
		Joins("CreatedBy").
		Joins("UpdatedBy").
		Where("slug = ?", slug).
		First(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	postWithInfo, err := repo.FindPostActivityInfo(post)
	return postWithInfo, err
}

func (repo *postRepository) FindPostActivityInfo(post *models.Post) (*models.Post, error) {
	var info *models.PostInfo

	result1 := repo.db.
		Raw("SELECT COUNT(*) AS total_like FROM user_post_activities WHERE post_id = ? AND is_liked = ?", post.ID, 1).
		Scan(&info)
	if result1.Error != nil {
		return nil, result1.Error
	}
	result2 := repo.db.
		Raw("SELECT COUNT(*) AS total_share FROM user_post_activities WHERE post_id = ? AND is_shared = ?", post.ID, 1).
		Scan(&info)
	if result2.Error != nil {
		return nil, result2.Error
	}

	post.TotalLike = info.TotalLike
	post.TotalShare = info.TotalShare
	return post, nil
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

func (repo *postRepository) FindActivity(act *models.UserPostActivity) (*models.UserPostActivity, error) {
	result := repo.db.
		Where("user_id = ?", act.UserID).
		Where("post_id = ?", act.PostID).
		First(&act)
	return act, result.Error
}

func (repo *postRepository) SaveActivity(act *models.UserPostActivity) (*models.UserPostActivity, error) {
	result := repo.db.
		Select("UserID", "PostID").
		Create(act)
	return act, result.Error
}

func (repo *postRepository) IncreaseViewsActivity(act *models.UserPostActivity) (*models.UserPostActivity, error) {
	result := repo.db.
		Model(&act).
		Update("views_count", gorm.Expr("views_count + ?", 1)).
		Update("updated_at", gorm.Expr("now()"))
	return act, result.Error
}

func (repo *postRepository) UpdateActivity(act *models.UserPostActivity) (*models.UserPostActivity, error) {
	result := repo.db.
		Model(&act).
		Update("is_liked", act.IsLiked).
		Update("is_shared", act.IsShared).
		Update("updated_at", gorm.Expr("now()"))
	return act, result.Error
}

func (repo *postRepository) Save(post *models.Post) (*models.Post, int, error) {
	result := repo.db.
		Clauses(clause.OnConflict{DoNothing: true}).
		Select(
			"PostTierID",
			"PostTier",
			"PostCategoryID",
			"PostCategory",
			"Title",
			"Content",
			"Slug",
			"Summary",
			"CreatedByID",
			"CreatedBy",
			"UpdatedByID",
			"UpdatedBy",
		).
		Create(post)
	return post, int(result.RowsAffected), result.Error
}
func (repo *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	result := repo.db.
		Model(&post).
		UpdateColumn("post_tier_id", post.PostTierID).
		UpdateColumn("post_category_id", post.PostCategoryID).
		UpdateColumn("title", post.Title).
		UpdateColumn("content", post.Content).
		UpdateColumn("summary", post.Summary)
	return post, result.Error
}
func (repo *postRepository) DeletePost(post *models.Post) error {
	result := repo.db.
		Delete(&post)
	return result.Error
}

func (repo *postRepository) FindUserLatestSubscription(user *models.User) (*models.UserSubscription, error) {
	us := &models.UserSubscription{}
	result := repo.db.
		Where("user_id = ?", user.ID).
		Last(&us)
	return us, result.Error
}

func (repo *postRepository) FindTiers() ([]*models.PostTier, error) {
	var tiers []*models.PostTier
	result := repo.db.
		Find(&tiers)
	return tiers, result.Error
}

func (repo *postRepository) FindTier(tier *models.PostTier) (*models.PostTier, error) {
	result := repo.db.
		First(&tier)
	return tier, result.Error
}

func (repo *postRepository) FindCategories() ([]*models.PostCategory, error) {
	var categories []*models.PostCategory
	result := repo.db.
		Find(&categories)
	return categories, result.Error
}

func (repo *postRepository) FindCategory(category *models.PostCategory) (*models.PostCategory, error) {
	result := repo.db.
		First(&category)
	return category, result.Error
}

func (repo *postRepository) SaveCategory(category *models.PostCategory) (*models.PostCategory, error) {
	result := repo.db.
		Create(category)
	return category, result.Error
}
func (repo *postRepository) UpdateCategory(category *models.PostCategory) (*models.PostCategory, error) {
	result := repo.db.
		Model(&category).
		UpdateColumn("name", category.Name).
		UpdateColumn("color", category.Color)
	return category, result.Error
}

func (repo *postRepository) FindTrending() (any, error) {
	var activities []*models.UserPostActivity
	result := repo.db.
		//Preload("Post.PostTier").
		//Preload("Post.PostCategory").
		//Preload("Post.ImgThumbnail").
		//Preload("Post.ImgContent").
		//Preload("Post.CreatedBy").
		//Preload("Post.UpdatedBy").
		//Preload(clause.Associations).
		Group("post_id").
		Order("updated_at DESC").
		Find(&activities)

	return activities, result.Error
}
