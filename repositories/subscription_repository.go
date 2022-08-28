package repositories

import (
	"final-project-backend/models"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	FindSubscriptions() ([]*models.Subscription, error)
	FindSubscription(subscription *models.Subscription) (*models.Subscription, error)
	SaveUserSubscription(us *models.UserSubscription) (*models.UserSubscription, error)
	SaveTransaction(transaction *models.Transaction) (*models.Transaction, error)
	FindUserLatestSubscription(user *models.User) (*models.UserSubscription, error)
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

func (repo *subscriptionRepository) FindSubscriptions() ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	result := repo.db.
		Find(&subscriptions)
	return subscriptions, result.Error
}

func (repo *subscriptionRepository) FindSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	result := repo.db.
		First(&subscription)
	return subscription, result.Error
}

func (repo *subscriptionRepository) FindUserLatestSubscription(user *models.User) (*models.UserSubscription, error) {
	us := &models.UserSubscription{}
	result := repo.db.
		Where("user_id = ?", user.ID).
		Last(&us)
	return us, result.Error
}

func (repo *subscriptionRepository) SaveUserSubscription(us *models.UserSubscription) (*models.UserSubscription, error) {
	result := repo.db.
		Create(us)
	return us, result.Error
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
