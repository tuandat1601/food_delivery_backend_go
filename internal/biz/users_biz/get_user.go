package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type GetUserStorage interface {
	GetUser(cxt context.Context, cond map[string]interface{}) (*model.User, error)
}
type getUserBiz struct {
	store GetUserStorage
}

func NewGetUser(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store}
}
func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*model.User, error) {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id":id,"is_deleted": false})
	if err != nil {
		return nil, err
	}
	return data, nil
}
