package models

import "gorm.io/gorm"

type UserPostActivity struct {
	gorm.Model `json:"-"`
	ID         int   `json:"userPostActivityID" gorm:"primaryKey"`
	UserID     int   `json:"userID"`
	User       *User `json:"user"`
	PostID     int   `json:"postID"`
	Post       *Post `json:"post"`
	IsLiked    int   `json:"isLiked"`
	IsShared   int   `json:"isShared"`
	ViewsCount int   `json:"viewsCount"`
}
