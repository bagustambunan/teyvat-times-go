package services

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type TransactionService interface {
	GetTransactions(opt *models.GetTransactionsOption) (*dto.TransactionsRes, error)
	GetUserTransactions(user *models.User) (*dto.TransactionsRes, error)
	AddTransaction(user *models.User, req *dto.TransactionReq, discount int, subscription *models.Subscription) (*dto.TransactionRes, error)
	GetTransactionStatuses() ([]*models.TransactionStatus, error)
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

type TSConfig struct {
	TransactionRepository repositories.TransactionRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository: c.TransactionRepository,
	}
}

func (serv *transactionService) GetTransactions(opt *models.GetTransactionsOption) (*dto.TransactionsRes, error) {
	return serv.transactionRepository.FindTransactions(opt)
}

func (serv *transactionService) GetUserTransactions(user *models.User) (*dto.TransactionsRes, error) {
	return serv.transactionRepository.FindTransactions(&models.GetTransactionsOption{
		UserID:   user.ID,
		StatusID: 0,
		Limit:    10,
		Page:     1,
	})
}

func (serv *transactionService) AddTransaction(user *models.User, req *dto.TransactionReq, discount int, subscription *models.Subscription) (*dto.TransactionRes, error) {
	netTotal := subscription.Price - discount
	if netTotal < 0 {
		netTotal = 0
	}
	transaction := &models.Transaction{
		UserID:         user.ID,
		SubscriptionID: req.SubscriptionID,
		StatusID:       2,
		GrossTotal:     subscription.Price,
		NetTotal:       netTotal,
		UserVoucherID:  req.UserVoucherID,
	}

	insertedTr, insertErr := serv.transactionRepository.SaveTransaction(transaction)
	return new(dto.TransactionRes).FromTransaction(insertedTr), insertErr
}

func (serv *transactionService) GetTransactionStatuses() ([]*models.TransactionStatus, error) {
	return serv.transactionRepository.FindTransactionStatuses()
}
