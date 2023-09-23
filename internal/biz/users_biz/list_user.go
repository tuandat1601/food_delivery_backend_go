package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type GetListUserStorage interface {
	GetListUser(cxt context.Context, cond map[string]interface{}) ([]model.User, error)
}
type getListUserBiz struct {
	store GetListUserStorage
}

func NewGetListUser(store GetListUserStorage) *getListUserBiz {
	return &getListUserBiz{store}
}
func (biz *getListUserBiz) ListUser(ctx context.Context, cond map[string]interface{}) ([]model.User, error) {
	data, err := biz.store.GetListUser(ctx, cond)
	if err != nil {
		return nil, err
	}
	return data, nil
}
