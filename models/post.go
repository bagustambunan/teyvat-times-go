package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model     `json:"-"`
	ID             int           `json:"postID" gorm:"primaryKey"`
	PostTierID     int           `json:"postTierID"`
	PostTier       *PostTier     `json:"postTier"`
	PostCategoryID int           `json:"postCategoryID"`
	PostCategory   *PostCategory `json:"postCategory"`
	Title          string        `json:"title"`
	Content        string        `json:"content"`
	Slug           string        `json:"slug"`
	Summary        string        `json:"summary"`
	ImgThumbnailID int           `json:"imgThumbnailID"`
	ImgThumbnail   *Image        `json:"imgThumbnail"`
	ImgContentID   int           `json:"imgContentID"`
	ImgContent     *Image        `json:"imgContent"`
	CreatedByID    int           `json:"createdByID"`
	CreatedBy      *User         `json:"createdBy"`
	UpdatedByID    int           `json:"updatedByID"`
	UpdatedBy      *User         `json:"updatedBy"`
	CreatedAt      string        `json:"createdAt"`
	UpdatedAt      string        `json:"updatedAt"`
	TotalLike      int           `json:"totalLike" gorm:"-"`
	TotalShare     int           `json:"totalShare" gorm:"-"`
}

func (p *Post) GetMoraRequired() int {
	return p.PostTier.MoraRequired
}
