package models

import "gorm.io/gorm"

type PostUnlock struct {
	gorm.Model `json:"-"`
	ID         int   `json:"postUnlockID" gorm:"primaryKey"`
	UserID     int   `json:"userID"`
	User       *User `json:"-"`
	PostID     int   `json:"postID"`
	Post       *Post `json:"-"`
}
