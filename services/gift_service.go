package services

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/repositories"
)

type GiftService interface {
	GetGifts() ([]*models.Gift, error)
	GetGift(gift *models.Gift) (*models.Gift, error)
}

type giftService struct {
	giftRepository repositories.GiftRepository
}

type GSConfig struct {
	GiftRepository repositories.GiftRepository
}

func NewGiftService(c *GSConfig) GiftService {
	return &giftService{
		giftRepository: c.GiftRepository,
	}
}

func (serv *giftService) GetGifts() ([]*models.Gift, error) {
	return serv.giftRepository.FindGifts()
}
func (serv *giftService) GetGift(gift *models.Gift) (*models.Gift, error) {
	return serv.giftRepository.FindGift(gift)
}
