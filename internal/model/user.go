package model

import "food_delivery/common"

type User struct {
	common.SQLModel
	UserName string   `json:"username" gorm:"column:username" `
	Email    string   `json:"email" gorm:"column:email"`
	Password string   `json:"password" gorm:"column:password" `
	Address  string   `json:"address" gorm:"column:address" `
	Phone    string   `json:"phone" gorm:"column:phone"`
	Role     string   `json:"role" gorm:"column:role" `
	IsDeleted  bool   `json:"is_deleted" gorm:"column:is_deleted"`
}