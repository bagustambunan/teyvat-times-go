package dto

type TransactionReq struct {
	SubscriptionID int `json:"subscriptionID" binding:"required"`
	UserVoucherID  int `json:"userVoucherID"`
}
