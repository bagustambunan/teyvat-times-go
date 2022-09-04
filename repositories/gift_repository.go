package repositories

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

type GiftRepository interface {
	FindGifts() ([]*models.Gift, error)
	FindGift(gift *models.Gift) (*models.Gift, error)
	FindUnclaimedUserGifts(user *models.User) ([]*models.UserGift, error)
	SaveGiftClaim(gc *models.GiftClaim) (*models.GiftClaim, error)
	FindGiftClaims(opt *models.GetGiftClaimsOption) (*dto.GiftClaimsRes, error)
	FindUserGiftClaims(opt *models.GetGiftClaimsOption) ([]*models.GiftClaim, error)
	FindGiftClaim(gc *models.GiftClaim) (*models.GiftClaim, error)
	FindGiftClaimStatuses() ([]*models.GiftClaimStatus, error)
	UpdateUserGiftIsClaimed(ug *models.UserGift, isClaimed int) (*models.UserGift, error)
	UpdateGiftClaimStatus(gc *models.GiftClaim, statusID int) (*models.GiftClaim, error)
	UpdateGiftStock(gift *models.Gift, number int) (*models.Gift, error)
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

func (repo *giftRepository) SaveGiftClaim(gc *models.GiftClaim) (*models.GiftClaim, error) {
	result := repo.db.
		//Select("UserID", "AddressID", "StatusID", "GiftClaimItems").
		Create(gc)
	return gc, result.Error
}

func (repo *giftRepository) FindGiftClaims(opt *models.GetGiftClaimsOption) (*dto.GiftClaimsRes, error) {
	var gcs []*models.GiftClaim
	result := repo.db.
		Model(&models.GiftClaim{}).
		Preload("GiftClaimItems.Gift.Image").
		Preload(clause.Associations)

	if opt.UserID != 0 {
		result = result.
			Where("user_id = ?", opt.UserID)
	}
	if opt.StatusID != 0 {
		result = result.
			Where("status_id = ?", opt.StatusID)
	}

	result = result.
		Order("updated_at DESC").
		Find(&gcs)
	totalData := int(result.RowsAffected)

	if opt.Limit != 0 {
		result = result.
			Limit(opt.Limit).
			Offset((opt.Page - 1) * opt.Limit).
			Find(&gcs)
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	gcsRes := &dto.GiftClaimsRes{
		Count:      int(result.RowsAffected),
		Limit:      opt.Limit,
		Page:       opt.Page,
		TotalPage:  totalPage,
		TotalData:  totalData,
		GiftClaims: gcs,
	}
	return gcsRes, result.Error
}

func (repo *giftRepository) FindUserGiftClaims(opt *models.GetGiftClaimsOption) ([]*models.GiftClaim, error) {
	var gcs []*models.GiftClaim
	result := repo.db.
		Preload("GiftClaimItems.Gift.Image").
		Preload(clause.Associations).
		Where("user_id = ?", opt.UserID).
		Order("updated_at DESC").
		Find(&gcs)
	return gcs, result.Error
}

func (repo *giftRepository) FindGiftClaim(gc *models.GiftClaim) (*models.GiftClaim, error) {
	result := repo.db.
		Preload("GiftClaimItems.Gift.Image").
		Preload(clause.Associations).
		First(&gc)
	return gc, result.Error
}

func (repo *giftRepository) FindGiftClaimStatuses() ([]*models.GiftClaimStatus, error) {
	var gcStatuses []*models.GiftClaimStatus
	result := repo.db.
		Order("id ASC").
		Find(&gcStatuses)
	return gcStatuses, result.Error
}

func (repo *giftRepository) UpdateUserGiftIsClaimed(ug *models.UserGift, isClaimed int) (*models.UserGift, error) {
	result := repo.db.
		Model(&ug).
		Update("is_claimed", isClaimed)
	return ug, result.Error
}

func (repo *giftRepository) UpdateGiftClaimStatus(gc *models.GiftClaim, statusID int) (*models.GiftClaim, error) {
	result := repo.db.
		//Model(&gc).
		//UpdateColumn("status_id", statusID)
		Raw("UPDATE gift_claims SET status_id = ? WHERE deleted_at IS NULL AND id = ?", statusID, gc.ID).
		Scan(&gc)
	return gc, result.Error
}

func (repo *giftRepository) UpdateGiftStock(gift *models.Gift, number int) (*models.Gift, error) {
	result := repo.db.
		Model(&gift).
		Update("stock", gorm.Expr("stock + ?", number))
	return gift, result.Error
}
