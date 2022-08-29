package dto

import "git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"

type PostOverviewRes struct {
	ID           int                  `json:"postID" gorm:"primaryKey"`
	PostTier     *models.PostTier     `json:"postTier"`
	PostCategory *models.PostCategory `json:"postCategory"`
	Title        string               `json:"title"`
	Slug         string               `json:"slug"`
	Summary      string               `json:"summary"`
	ImgThumbnail *models.Image        `json:"imgThumbnail"`
	CreatedBy    *models.User         `json:"createdBy"`
	UpdatedBy    *models.User         `json:"updatedBy"`
	CreatedAt    string               `json:"createdAt"`
	UpdatedAt    string               `json:"updatedAt"`
	TotalLike    int                  `json:"totalLike"`
	TotalShare   int                  `json:"totalShare"`
}

func (_ *PostOverviewRes) FromPost(p *models.Post) *PostOverviewRes {
	return &PostOverviewRes{
		ID:           p.ID,
		PostTier:     p.PostTier,
		PostCategory: p.PostCategory,
		Title:        p.Title,
		Slug:         p.Slug,
		Summary:      p.Summary,
		ImgThumbnail: p.ImgThumbnail,
		CreatedBy:    p.CreatedBy,
		UpdatedBy:    p.UpdatedBy,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		TotalLike:    p.TotalLike,
		TotalShare:   p.TotalShare,
	}
}
