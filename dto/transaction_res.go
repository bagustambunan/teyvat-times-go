package dto

import "git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"

type TransactionRes struct {
	ID             int `json:"transactionID" gorm:"primaryKey"`
	UserID         int `json:"userID"`
	SubscriptionID int `json:"subscriptionID"`
	StatusID       int `json:"statusID"`
	GrossTotal     int `json:"grossTotal"`
	NetTotal       int `json:"netTotal"`
	UserVoucherID  int `json:"userVoucherID"`
}

func (_ *TransactionRes) FromTransaction(tr *models.Transaction) *TransactionRes {
	return &TransactionRes{
		ID:             tr.ID,
		UserID:         tr.UserID,
		SubscriptionID: tr.SubscriptionID,
		StatusID:       tr.StatusID,
		GrossTotal:     tr.GrossTotal,
		NetTotal:       tr.NetTotal,
		UserVoucherID:  tr.UserVoucherID,
	}
}
