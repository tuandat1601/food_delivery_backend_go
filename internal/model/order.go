package model

import "time"

type orders struct {
	Id         int          `json:"id" gorm:"column:id"`
	UserId     int          `json:"user_id" gorm:"column:user_id"`
	OrderTime  *time.Time   `json:"order_id" gorm:"column:order_time"`
}
type OrderItem struct {
	Id         int        `json:"id" gorm:"column:id"`
	OrderId    int        `json:"order_id" gorm:"column:order_id"`
	FoodId     int        `json:"food_id" gorm:"column:food_id"`
	Quantity   int        `json:"quantity" gorm:"column:quantity"`
	TotalPrice float64    `json:"total_price" gorm:"column:total_price"`
	
}