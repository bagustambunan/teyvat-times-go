package models

type UserSpending struct {
	User          *User `json:"user" gorm:"-"`
	TotalSpending int   `json:"totalSpending"`
}
