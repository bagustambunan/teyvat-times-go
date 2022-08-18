package dto

import "final-project-backend/models"

type GetPostRes struct {
	ID             int                  `json:"postID" gorm:"primaryKey"`
	PostTier       *models.PostTier     `json:"postTier"`
	PostCategory   *models.PostCategory `json:"postCategory"`
	Title          string               `json:"title"`
	Content        string               `json:"content"`
	Slug           string               `json:"slug"`
	Summary        string               `json:"summary"`
	ImgThumbnailID int                  `json:"imgThumbnailID"`
	ImgContentID   int                  `json:"imgContentID"`
	CreatedByID    int                  `json:"createdByID"`
	UpdatedByID    int                  `json:"updatedByID"`
	//ImgThumbnail *models.Image        `json:"imgThumbnail"`
	//ImgContent   *models.Image        `json:"imgContent"`
	//CreatedBy    *models.User         `json:"createdBy"`
	//UpdatedBy    *models.User         `json:"updatedBy"`
}

func (_ GetPostRes) FromPost(p *models.Post) *GetPostRes {
	return &GetPostRes{
		ID:           p.ID,
		PostTier:     p.PostTier,
		PostCategory: p.PostCategory,
		Title:        p.Title,
		Content:      p.Content,
		Slug:         p.Slug,
		Summary:      p.Summary,
		//ImgThumbnail: p.ImgThumbnail,
		//ImgContent:   p.ImgContent,
		//CreatedBy:    p.CreatedBy,
		//UpdatedBy:    p.UpdatedBy,
		ImgThumbnailID: p.ImgThumbnailID,
		ImgContentID:   p.ImgContentID,
		CreatedByID:    p.CreatedByID,
		UpdatedByID:    p.UpdatedById,
	}
}
