package services

import (
	"final-project-backend/repositories"
)

type GiftService interface {
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
