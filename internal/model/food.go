package model

import "food_delivery/common"

type Food struct {
	common.SQLModel
	Name        string       `json:"name" gorm:"column:name"`
	Description string       `json:"description" gorm:"column:description"`
	Image       string       `json:"image" gorm:"column:image"`
	Price       int          `json:"price" gorm:"column:price"`
	MenuId      int          `json:"menu_id" gorm:"column:menu_id"`
}