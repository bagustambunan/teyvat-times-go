package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model `json:"-"`
	ID         int    `json:"addressID" gorm:"primaryKey"`
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}
