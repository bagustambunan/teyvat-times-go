package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	SaveTransaction(transaction *models.Transaction) (*models.Transaction, error)
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

func (repo *transactionRepository) SaveTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := repo.db
	if transaction.UserVoucherID == 0 {
		result = result.
			Select("UserID", "SubscriptionID", "StatusID", "GrossTotal", "NetTotal")
	}
	result = result.
		Create(transaction)
	return transaction, result.Error
}
