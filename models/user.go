package models

import "gorm.io/gorm"

type User struct {
	gorm.Model   `json:"-"`
	ID           int      `json:"userID" gorm:"primaryKey"`
	RoleID       int      `json:"roleID"`
	Role         *Role    `json:"-"`
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Name         string   `json:"name"`
	Phone        string   `json:"phone"`
	AddressID    int      `json:"addressID"`
	Address      *Address `json:"-"`
	ReferralCode string   `json:"referralCode"`
	ProfilePicID int      `json:"profilePicID"`
	ProfilePic   *Image   `json:"-"`
	Password     string   `json:"-"`
}
