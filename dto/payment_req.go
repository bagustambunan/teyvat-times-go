package dto

type PaymentReq struct {
	TransactionID int `json:"transactionID" binding:"required"`
	Amount        int `json:"amount"`
}
