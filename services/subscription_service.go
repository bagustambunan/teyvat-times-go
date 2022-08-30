package services

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"final-project-backend/repositories"
	"time"
)

type SubscriptionService interface {
	GetSubscriptions() ([]*models.Subscription, error)
	GetSubscription(subscription *models.Subscription) (*models.Subscription, error)
	GetUserNewSubscriptionDate(user *models.User) (string, string)
	GetUserActiveSubscription(user *models.User) *models.UserSubscription
	AddUserSubscription(user *models.User, subscription *models.Subscription) (*models.UserSubscription, error)
	GetUserSubscriptions(user *models.User, opt *models.GetUserSubscriptionsOption) (*dto.UserSubscriptionsRes, error)
}

type subscriptionService struct {
	subscriptionRepository repositories.SubscriptionRepository
}

type SSConfig struct {
	SubscriptionRepository repositories.SubscriptionRepository
}

func NewSubscriptionService(c *SSConfig) SubscriptionService {
	return &subscriptionService{
		subscriptionRepository: c.SubscriptionRepository,
	}
}

func (serv *subscriptionService) GetSubscriptions() ([]*models.Subscription, error) {
	return serv.subscriptionRepository.FindSubscriptions()
}

func (serv *subscriptionService) GetSubscription(subscription *models.Subscription) (*models.Subscription, error) {
	return serv.subscriptionRepository.FindSubscription(subscription)
}

func (serv *subscriptionService) GetUserLatestSubscription(user *models.User) *models.UserSubscription {
	us, err := serv.subscriptionRepository.FindUserLatestSubscription(user)
	if err != nil {
		return nil
	}
	return us
}

func (serv *subscriptionService) GetUserNewSubscriptionDate(user *models.User) (string, string) {
	dateStart := time.Now()
	if latestUs := serv.GetUserLatestSubscription(user); latestUs != nil {
		latestUsEnded, _ := time.Parse("2006-01-02T00:00:00Z", latestUs.DateEnded)
		if latestUsEnded.After(dateStart) {
			dateStart = latestUsEnded.AddDate(0, 0, 1)
		}
	}
	dateEnded := dateStart.AddDate(0, 1, 0)
	return dateStart.Format("2006-01-02"), dateEnded.Format("2006-01-02")
}

func (serv *subscriptionService) GetUserActiveSubscription(user *models.User) *models.UserSubscription {
	dateNow := time.Now()
	if latestUs := serv.GetUserLatestSubscription(user); latestUs != nil {
		latestUsEnded, _ := time.Parse("2006-01-02T00:00:00Z", latestUs.DateEnded)
		if latestUsEnded.After(dateNow) {
			return latestUs
		}
	}
	return nil
}

func (serv *subscriptionService) AddUserSubscription(user *models.User, subscription *models.Subscription) (*models.UserSubscription, error) {
	dateStart, dateEnded := serv.GetUserNewSubscriptionDate(user)

	us := &models.UserSubscription{
		UserID:         user.ID,
		SubscriptionID: subscription.ID,
		DateStart:      dateStart,
		DateEnded:      dateEnded,
	}
	return serv.subscriptionRepository.SaveUserSubscription(us)
}

func (serv *subscriptionService) GetUserSubscriptions(user *models.User, opt *models.GetUserSubscriptionsOption) (*dto.UserSubscriptionsRes, error) {
	opt.UserID = user.ID
	return serv.subscriptionRepository.FindUserSubscriptions(opt)
}
