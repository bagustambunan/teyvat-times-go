package repositories

import (
	"gorm.io/gorm"
)

type GiftRepository interface {
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
