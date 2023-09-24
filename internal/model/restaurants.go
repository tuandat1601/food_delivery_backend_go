package model

import "food_delivery/common"

type Restaurant struct {
	common.SQLModel
	Name     string          `json:"name" gorm:"column:name"`
	Address  string          `json:"address" gorm:"column:address"`
	Phone    string          `json:"phone" gorm:"column:phone"`
	UserId   int             `json:"user_id" gorm:"column:user_id"`
	DayOfWeek   string    `json:"day_of_week" gorm:"column:day_of_week"`
	OpeningTime string   `json:"opening_time" gorm:"column:opening_time"`
	ClosingTime string   `json:"closing_time" gorm:"column:closing_time"`
	IsDeleted   bool     `json:"is_deleted" gorm:"column:is_deleted"`
}
