package dto

type CategoryReq struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}
