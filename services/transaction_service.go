package services

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
	"time"
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
	GetUserThisMonthSpending(user *models.User) (*models.UserSpending, error)
	CheckGiftReward(user *models.User, lastTrNetTotal int) error
	CheckVoucherReward(lastTrNetTotal int, user *models.User, referrerUser *models.User) error
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

func (serv *transactionService) GetUserThisMonthSpending(user *models.User) (*models.UserSpending, error) {
	return serv.transactionRepository.FindUserThisMonthSpending(user)
}

func (serv *transactionService) CheckGiftReward(user *models.User, lastTrNetTotal int) error {
	thisMonthSpending, _ := serv.GetUserThisMonthSpending(user)
	lastSpend := thisMonthSpending.TotalSpending - lastTrNetTotal
	currentSpend := thisMonthSpending.TotalSpending
	if lastSpend <= 50000 && currentSpend > 50000 {
		return serv.GiveGiftToUser(user, &models.Gift{ID: 1})
	}
	if lastSpend <= 150000 && currentSpend > 150000 {
		return serv.GiveGiftToUser(user, &models.Gift{ID: 2})
	}
	if lastSpend <= 200000 && currentSpend > 200000 {
		return serv.GiveGiftToUser(user, &models.Gift{ID: 3})
	}
	return nil
}

func (serv *transactionService) GiveGiftToUser(user *models.User, gift *models.Gift) error {
	_, ugErr := serv.transactionRepository.SaveUserGift(&models.UserGift{
		UserID:    user.ID,
		GiftID:    gift.ID,
		IsClaimed: 0,
	})
	return ugErr
}

func (serv *transactionService) CheckVoucherReward(lastTrNetTotal int, user *models.User, referrerUser *models.User) error {
	uSpending, _ := serv.GetUserTotalSpending(user)
	lastSpend := uSpending.TotalSpending - lastTrNetTotal
	currentSpend := uSpending.TotalSpending
	if lastSpend <= 100000 && currentSpend > 100000 {
		return serv.GiveVoucherToUser(referrerUser, &models.Voucher{ID: 1})
	}
	if lastSpend <= 200000 && currentSpend > 200000 {
		return serv.GiveVoucherToUser(referrerUser, &models.Voucher{ID: 2})
	}
	if lastSpend <= 250000 && currentSpend > 250000 {
		return serv.GiveVoucherToUser(referrerUser, &models.Voucher{ID: 3})
	}
	return nil
}

func (serv *transactionService) GiveVoucherToUser(referrerUser *models.User, voucher *models.Voucher) error {
	dateNow := time.Now()
	dateExpired := dateNow.AddDate(0, 1, 0)
	_, uvErr := serv.transactionRepository.SaveUserVoucher(
		&models.UserVoucher{
			UserID:      referrerUser.ID,
			VoucherID:   voucher.ID,
			DateExpired: dateExpired.Format("2006-01-02"),
			IsUsed:      0,
		},
	)
	return uvErr
}
