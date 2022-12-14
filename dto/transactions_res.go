package dto

import "final-project-backend/models"

type TransactionsRes struct {
	Count        int                   `json:"count"`
	Limit        int                   `json:"limit"`
	Page         int                   `json:"page"`
	TotalPage    int                   `json:"totalPage"`
	TotalData    int                   `json:"totalData"`
	Transactions []*models.Transaction `json:"transactions"`
}
