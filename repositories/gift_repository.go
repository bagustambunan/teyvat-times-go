package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"
	"gorm.io/gorm"
)

type GiftRepository interface {
	FindGifts() ([]*models.Gift, error)
	FindGift(gift *models.Gift) (*models.Gift, error)
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
