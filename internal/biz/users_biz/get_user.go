package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type GetUserStorage interface {
	GetUser(cxt context.Context,cond map[string]interface{})(*model.User,error)
}
type getUserBiz struct {
	store GetUserStorage
}
func NewGetUser(store GetUserStorage) *getUserBiz{
	return &getUserBiz{store}
}
