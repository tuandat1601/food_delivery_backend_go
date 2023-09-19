package model

import "food_delivery/common"

type Restaurants struct {
	common.SQLModel
	Name     string          `json:"name" gorm:"column:name"`
	Address  string          `json:"address" gorm:"column:address"`
	Phone    string          `json:"phone" gorm:"column:phone"`
	Email    string          `json:"email" gorm:"column:email"`
	OperatingHours []common.OperatingHours `json:"operation_hours" gorm:"column:operation_hours"`
}