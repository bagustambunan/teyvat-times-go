package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GiftRepository interface {
	FindGifts() ([]*models.Gift, error)
	FindGift(gift *models.Gift) (*models.Gift, error)
	FindUnclaimedUserGifts(user *models.User) ([]*models.UserGift, error)
}

type giftRepository struct {
	db *gorm.DB
}

type GRConfig struct {
	DB *gorm.DB
}

func NewGiftRepository(c *GRConfig) GiftRepository {
	return &giftRepository{db: c.DB}
}

func (repo *giftRepository) FindGifts() ([]*models.Gift, error) {
	var gifts []*models.Gift
	result := repo.db.
		Joins("Image").
		Find(&gifts)
	return gifts, result.Error
}
func (repo *giftRepository) FindGift(gift *models.Gift) (*models.Gift, error) {
	result := repo.db.
		Joins("Image").
		First(&gift)
	return gift, result.Error
}

func (repo *giftRepository) FindUnclaimedUserGifts(user *models.User) ([]*models.UserGift, error) {
	var ugs []*models.UserGift
	result := repo.db.
		Preload("Gift.Image").
		Preload(clause.Associations).
		Where("user_id", user.ID).
		Where("is_claimed", 0).
		Find(&ugs)
	return ugs, result.Error
}
