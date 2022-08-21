package services

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type SubscriptionService interface {
	AddTransaction(user *models.User, req *dto.TransactionReq, discount int) (*dto.TransactionRes, error)
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

func (serv *subscriptionService) AddTransaction(user *models.User, req *dto.TransactionReq, discount int) (*dto.TransactionRes, error) {
	subscription, subErr := serv.subscriptionRepository.FindSubscription(&models.Subscription{ID: req.SubscriptionID})
	if subErr != nil {
		return nil, httperror.BadRequestError("Invalid subscription", "INVALID_SUBSCRIPTION")
	}

	netTotal := subscription.Price - discount
	if netTotal < 0 {
		netTotal = 0
	}
	transaction := &models.Transaction{
		UserID:         user.ID,
		SubscriptionID: req.SubscriptionID,
		StatusID:       1,
		GrossTotal:     subscription.Price,
		NetTotal:       netTotal,
		UserVoucherID:  req.UserVoucherID,
	}

	insertedTr, insertErr := serv.subscriptionRepository.SaveTransaction(transaction)
	return new(dto.TransactionRes).FromTransaction(insertedTr), insertErr
}
