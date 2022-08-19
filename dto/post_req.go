package dto

type PostReq struct {
	PostTierID     int    `json:"postTierID"`
	PostCategoryID int    `json:"postCategoryID"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Summary        string `json:"summary"`
	CreatedByID    int    `json:"createdByID"`
	UpdatedByID    int    `json:"updatedByID"`
}
