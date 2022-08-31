package services

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
)

type TransactionService interface {
	GetTransactions(opt *models.GetTransactionsOption) (*dto.TransactionsRes, error)
	GetUserTransactions(user *models.User, opt *models.GetTransactionsOption) (*dto.TransactionsRes, error)
	AddTransaction(user *models.User, req *dto.TransactionReq, discount int, subscription *models.Subscription) (*dto.TransactionRes, error)
	GetTransactionStatuses() ([]*models.TransactionStatus, error)
	GetTransaction(transaction *models.Transaction) (*models.Transaction, error)
	ApproveTransaction(transaction *models.Transaction) (*models.Transaction, error)
	RejectTransaction(transaction *models.Transaction) (*models.Transaction, error)
	ProcessPayment(transaction *models.Transaction, req *dto.PaymentReq) (*models.Transaction, error)
	GetUserTotalSpending(user *models.User) (*models.UserSpending, error)
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

func (serv *transactionService) GetUserTransactions(user *models.User, opt *models.GetTransactionsOption) (*dto.TransactionsRes, error) {
	opt.UserID = user.ID
	return serv.transactionRepository.FindTransactions(opt)
}

func (serv *transactionService) AddTransaction(user *models.User, req *dto.TransactionReq, discount int, subscription *models.Subscription) (*dto.TransactionRes, error) {
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

	insertedTr, insertErr := serv.transactionRepository.SaveTransaction(transaction)
	return new(dto.TransactionRes).FromTransaction(insertedTr), insertErr
}

func (serv *transactionService) GetTransactionStatuses() ([]*models.TransactionStatus, error) {
	return serv.transactionRepository.FindTransactionStatuses()
}

func (serv *transactionService) GetTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return serv.transactionRepository.FindTransaction(transaction)
}

func (serv *transactionService) ApproveTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return serv.transactionRepository.UpdateTransactionStatus(transaction, 3)
}

func (serv *transactionService) RejectTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return serv.transactionRepository.UpdateTransactionStatus(transaction, 4)
}

func (serv *transactionService) ProcessPayment(transaction *models.Transaction, req *dto.PaymentReq) (*models.Transaction, error) {
	if transaction.StatusID != 1 {
		return nil, httperror.BadRequestError("This payment is invalid", "INVALID_PAYMENT")
	}
	if req.Amount != transaction.NetTotal {
		return nil, httperror.BadRequestError("Payment amount doesn't match", "INVALID_AMOUNT")
	}
	return serv.transactionRepository.UpdateTransactionStatus(transaction, 2)
}

func (serv *transactionService) GetUserTotalSpending(user *models.User) (*models.UserSpending, error) {
	return serv.transactionRepository.FindUserTotalSpending(user)
}
