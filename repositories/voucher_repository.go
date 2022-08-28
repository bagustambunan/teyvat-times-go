package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	FindUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error)
	FindUserVoucherFromCode(user *models.User, code string) (*models.UserVoucher, error)
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

func (repo *voucherRepository) FindUserVoucherFromCode(user *models.User, code string) (*models.UserVoucher, error) {
	var uv *models.UserVoucher
	result := repo.db.
		Joins("Voucher").
		Where("code = ?", code).
		Where("user_id = ?", user.ID).
		Where("is_used = ?", 0).
		First(&uv)
	return uv, result.Error
}

func (repo *voucherRepository) UpdateUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error) {
	result := repo.db.
		Model(&uv).
		UpdateColumn("is_used", 1)
	return uv, result.Error
}
