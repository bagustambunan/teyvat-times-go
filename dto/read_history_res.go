package dto

import "final-project-backend/models"

type ReadHistoryRes struct {
	Count      int                        `json:"count"`
	Limit      int                        `json:"limit"`
	Page       int                        `json:"page"`
	TotalPage  int                        `json:"totalPage"`
	TotalData  int                        `json:"totalData"`
	Activities []*models.UserPostActivity `json:"activities"`
}

func (_ *ReadHistoryRes) FromActivities(activities []*models.UserPostActivity) *ReadHistoryRes {
	return &ReadHistoryRes{
		Activities: activities,
	}
}

func (res *ReadHistoryRes) SetValues(count int, limit int, page int, totalPage int, totalData int) {
	res.Count = count
	res.Limit = limit
	res.Page = page
	res.TotalPage = totalPage
	res.TotalData = totalData
}
