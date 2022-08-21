package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	FindUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error)
	UpdateUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

type VRConfig struct {
	DB *gorm.DB
}

func NewVoucherRepository(c *VRConfig) VoucherRepository {
	return &voucherRepository{db: c.DB}
}

func (repo *voucherRepository) FindUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error) {
	result := repo.db.
		Joins("User").
		Joins("Voucher").
		First(&uv)
	return uv, result.Error
}

func (repo *voucherRepository) UpdateUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error) {
	result := repo.db.
		Model(&uv).
		UpdateColumn("is_used", 1)
	return uv, result.Error
}
