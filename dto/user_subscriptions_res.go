package dto

import "final-project-backend/models"

type UserSubscriptionsRes struct {
	Count             int                        `json:"count"`
	Limit             int                        `json:"limit"`
	Page              int                        `json:"page"`
	TotalPage         int                        `json:"totalPage"`
	TotalData         int                        `json:"totalData"`
	UserSubscriptions []*models.UserSubscription `json:"userSubscriptions"`
}
