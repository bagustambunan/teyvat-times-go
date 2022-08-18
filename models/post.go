package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model     `json:"-"`
	ID             int           `json:"postID" gorm:"primaryKey"`
	PostTierID     int           `json:"postTierID"`
	PostTier       *PostTier     `json:"-"`
	PostCategoryID int           `json:"postCategoryID"`
	PostCategory   *PostCategory `json:"_"`
	Title          string        `json:"title"`
	Content        string        `json:"content"`
	Slug           string        `json:"slug"`
	Summary        string        `json:"summary"`
	ImgThumbnailID int           `json:"imgThumbnailID"`
	ImgThumbnail   *Image        `json:"-"`
	ImgContentID   int           `json:"imgContentID"`
	ImgContent     *Image        `json:"-"`
	CreatedByID    int           `json:"createdByID"`
	CreatedBy      *User         `json:"-"`
	UpdatedById    int           `json:"updatedByID"`
	UpdatedBy      *User         `json:"-"`
}
