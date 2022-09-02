package services

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type GiftService interface {
	GetGifts() ([]*models.Gift, error)
	GetGift(gift *models.Gift) (*models.Gift, error)
	GetUnclaimedUserGifts(user *models.User) ([]*models.UserGift, error)
	SaveGiftClaim(user *models.User) (*models.GiftClaim, error)
	GetGiftClaims(opt *models.GetGiftClaimsOption) (*dto.GiftClaimsRes, error)
	GetUserGiftClaims(user *models.User) ([]*models.GiftClaim, error)
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

func (serv *giftService) GetUnclaimedUserGifts(user *models.User) ([]*models.UserGift, error) {
	return serv.giftRepository.FindUnclaimedUserGifts(user)
}

func (serv *giftService) SaveGiftClaim(user *models.User) (*models.GiftClaim, error) {
	userGifts, _ := serv.GetUnclaimedUserGifts(user)
	if len(userGifts) == 0 {
		return nil, httperror.BadRequestError("No unclaimed gifts found", "INVALID_GIFT_CLAIM")
	}

	var giftClaimItems []*models.GiftClaimItem
	for _, ug := range userGifts {
		giftClaimItems = append(giftClaimItems, &models.GiftClaimItem{Gift: ug.Gift})
	}

	gc := &models.GiftClaim{
		UserID:         user.ID,
		AddressID:      user.AddressID,
		StatusID:       1,
		GiftClaimItems: giftClaimItems,
	}

	// TODO: UPDATE USER CLAIM > IS CLAIMED

	return serv.giftRepository.SaveGiftClaim(gc)
}

func (serv *giftService) GetGiftClaims(opt *models.GetGiftClaimsOption) (*dto.GiftClaimsRes, error) {
	return serv.giftRepository.FindGiftClaims(opt)
}

func (serv *giftService) GetUserGiftClaims(user *models.User) ([]*models.GiftClaim, error) {
	opt := &models.GetGiftClaimsOption{UserID: user.ID}
	return serv.giftRepository.FindUserGiftClaims(opt)
}
