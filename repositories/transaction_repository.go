package repositories

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/dto"
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"
	"gorm.io/gorm"
	"math"
)

type TransactionRepository interface {
	FindTransactions(opt *models.GetTransactionsOption) (*dto.TransactionsRes, error)
	SaveTransaction(transaction *models.Transaction) (*models.Transaction, error)
	FindTransactionStatuses() ([]*models.TransactionStatus, error)
	FindTransaction(transaction *models.Transaction) (*models.Transaction, error)
	UpdateTransactionStatus(transaction *models.Transaction, statusID int) (*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

type TRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(c *TRConfig) TransactionRepository {
	return &transactionRepository{db: c.DB}
}

func (repo *transactionRepository) FindTransactions(opt *models.GetTransactionsOption) (*dto.TransactionsRes, error) {
	var transactions []*models.Transaction
	result := repo.db.
		Table("transactions").
		Joins("User").
		Joins("Subscription").
		Joins("Status")

	if opt.UserID != 0 {
		result = result.
			Where("user_id = ?", opt.UserID)
	}
	if opt.StatusID != 0 {
		result = result.
			Where("status_id = ?", opt.StatusID)
	}
	result = result.
		Order("created_at DESC").
		Find(&transactions)
	totalData := int(result.RowsAffected)

	result = result.
		Limit(opt.Limit).
		Offset((opt.Page - 1) * opt.Limit).
		Find(&transactions)

	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	trsRes := &dto.TransactionsRes{
		Count:        int(result.RowsAffected),
		Limit:        opt.Limit,
		Page:         opt.Page,
		TotalPage:    totalPage,
		TotalData:    totalData,
		Transactions: transactions,
	}
	return trsRes, result.Error
}

func (repo *transactionRepository) SaveTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := repo.db
	if transaction.UserVoucherID == 0 {
		result = result.
			Select("UserID", "SubscriptionID", "StatusID", "GrossTotal", "NetTotal")
	}
	if transaction.UserVoucherID != 0 {
		result = result.
			Select("UserID", "SubscriptionID", "StatusID", "GrossTotal", "NetTotal", "UserVoucherID")
	}
	result = result.
		Create(transaction)
	return transaction, result.Error
}

func (repo *transactionRepository) FindTransactionStatuses() ([]*models.TransactionStatus, error) {
	var trStatuses []*models.TransactionStatus
	result := repo.db.
		Order("id ASC").
		Find(&trStatuses)
	return trStatuses, result.Error
}

func (repo *transactionRepository) FindTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := repo.db.
		Joins("User").
		Joins("Subscription").
		Joins("Status").
		Joins("UserVoucher").
		First(&transaction)
	return transaction, result.Error
}

func (repo *transactionRepository) UpdateTransactionStatus(transaction *models.Transaction, statusID int) (*models.Transaction, error) {
	result := repo.db.
		//Model(&transaction).
		//UpdateColumn("status_id", statusID).
		Raw("UPDATE transactions SET status_id = ? WHERE deleted_at IS NULL AND id = ?", statusID, transaction.ID).
		Scan(&transaction)
	return transaction, result.Error
}
