package models

import "gorm.io/gorm"

type UserPostActivities struct {
	gorm.Model `json:"-"`
	ID         int   `json:"userPostActivityID" gorm:"primaryKey"`
	UserID     int   `json:"userID"`
	User       *User `json:"-"`
	PostID     int   `json:"postID"`
	Post       *Post `json:"-"`
	IsLiked    int   `json:"isLiked"`
	IsShared   int   `json:"isShared"`
	ViewsCount int   `json:"viewsCount"`
}
