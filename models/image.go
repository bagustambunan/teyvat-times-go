package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model `json:"-"`
	ID         int    `json:"imageID" gorm:"primaryKey"`
	Url        string `json:"url"`
	AltText    string `json:"altText"`
}
