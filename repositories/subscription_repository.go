package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	FindSubscription(subscription *models.Subscription) (*models.Subscription, error)
	SaveTransaction(transaction *models.Transaction) (*models.Transaction, error)
}

type subscriptionRepository struct {
	db *gorm.DB
}

type SRConfig struct {
	DB *gorm.DB
}

func NewSubscriptionRepository(c *SRConfig) SubscriptionRepository {
	return &subscriptionRepository{db: c.DB}
}

func (repo *subscriptionRepository) FindSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	result := repo.db.
		First(&subscription)
	return subscription, result.Error
}

func (repo *subscriptionRepository) SaveTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := repo.db
	if transaction.UserVoucherID == 0 {
		result = result.
			Select("UserID", "SubscriptionID", "StatusID", "GrossTotal", "NetTotal")
	}
	result = result.
		Create(transaction)
	return transaction, result.Error
}
