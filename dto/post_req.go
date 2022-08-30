package dto

type PostReq struct {
	PostTierID     int    `json:"postTierID" binding:"required"`
	PostCategoryID int    `json:"postCategoryID" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Content        string `json:"content" binding:"required"`
	Summary        string `json:"summary" binding:"required"`
}
