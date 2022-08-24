package dto

import (
	"final-project-backend/models"
)

type GetPostRes struct {
	ID           int                  `json:"postID" gorm:"primaryKey"`
	PostTier     *models.PostTier     `json:"postTier"`
	PostCategory *models.PostCategory `json:"postCategory"`
	Title        string               `json:"title"`
	Content      string               `json:"content"`
	Slug         string               `json:"slug"`
	Summary      string               `json:"summary"`
	ImgThumbnail *models.Image        `json:"imgThumbnail"`
	ImgContent   *models.Image        `json:"imgContent"`
	CreatedBy    *models.User         `json:"createdBy"`
	UpdatedBy    *models.User         `json:"updatedBy"`
	CreatedAt    string               `json:"createdAt"`
	UpdatedAt    string               `json:"updatedAt"`
}

func (_ *GetPostRes) FromPost(p *models.Post) *GetPostRes {
	return &GetPostRes{
		ID:           p.ID,
		PostTier:     p.PostTier,
		PostCategory: p.PostCategory,
		Title:        p.Title,
		Content:      p.Content,
		Slug:         p.Slug,
		Summary:      p.Summary,
		ImgThumbnail: p.ImgThumbnail,
		ImgContent:   p.ImgContent,
		CreatedBy:    p.CreatedBy,
		UpdatedBy:    p.UpdatedBy,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
