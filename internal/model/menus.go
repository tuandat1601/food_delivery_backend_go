package model

import (
	"food_delivery/common"
	
)

type Menus struct {
	common.SQLModel
	Name         string   `json:"name" gorm:"column:name"`
	Description  string   `json:"description" gorm:"column:description"`
	RestaurantID int      `json:"restaurant_id" gorm:"column:restaurant_id"`

}