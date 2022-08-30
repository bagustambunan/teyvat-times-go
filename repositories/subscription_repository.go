package repositories

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"gorm.io/gorm"
	"math"
)

type SubscriptionRepository interface {
	FindSubscriptions() ([]*models.Subscription, error)
	FindSubscription(subscription *models.Subscription) (*models.Subscription, error)
	SaveUserSubscription(us *models.UserSubscription) (*models.UserSubscription, error)
	FindUserLatestSubscription(user *models.User) (*models.UserSubscription, error)
	FindUserSubscriptions(opt *models.GetUserSubscriptionsOption) (*dto.UserSubscriptionsRes, error)
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

func (repo *subscriptionRepository) FindUserSubscriptions(opt *models.GetUserSubscriptionsOption) (*dto.UserSubscriptionsRes, error) {
	var uss []*models.UserSubscription
	result := repo.db.
		Joins("User").
		Joins("Subscription")

	if opt.UserID != 0 {
		result = result.
			Where("user_id = ?", opt.UserID)
	}

	result = result.
		Order("created_at DESC").
		Find(&uss)
	totalData := int(result.RowsAffected)

	result = result.
		Limit(opt.Limit).
		Offset((opt.Page - 1) * opt.Limit).
		Find(&uss)

	totalPage := int(math.Ceil(float64(totalData) / float64(opt.Limit)))
	ussRes := &dto.UserSubscriptionsRes{
		Count:             int(result.RowsAffected),
		Limit:             opt.Limit,
		Page:              opt.Page,
		TotalPage:         totalPage,
		TotalData:         totalData,
		UserSubscriptions: uss,
	}
	return ussRes, result.Error
}
