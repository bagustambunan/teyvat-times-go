package services

import (
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type VoucherService interface {
	GetUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error)
	UseUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error)
}

type voucherService struct {
	voucherRepository repositories.VoucherRepository
}

type VSConfig struct {
	VoucherRepository repositories.VoucherRepository
}

func NewVoucherService(c *VSConfig) VoucherService {
	return &voucherService{
		voucherRepository: c.VoucherRepository,
	}
}

func (serv *voucherService) GetUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error) {
	fetchedUv, err := serv.voucherRepository.FindUserVoucher(uv)
	if err != nil {
		return nil, httperror.BadRequestError("Invalid user voucher", "INVALID_USER_VOUCHER")
	}
	return fetchedUv, nil
}

func (serv *voucherService) UseUserVoucher(uv *models.UserVoucher) (*models.UserVoucher, error) {
	return serv.voucherRepository.UpdateUserVoucher(uv)
}
