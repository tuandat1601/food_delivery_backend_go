package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type UpdateUserStorage interface {
	GetUser(ctx context.Context,cond map[string]interface{})(*model.User,error)
	UpdateUser(cxt context.Context, cond map[string]interface{},dataupdate *model.User)error
}
type updateUserBiz struct {
	store UpdateUserStorage
}
func NewUpdateUser(store UpdateUserStorage) *updateUserBiz{
	return &updateUserBiz{store}
}

func (biz *updateUserBiz) UpdateUserById(ctx context.Context, id int, dataupdate *model.User) error{

	data, err:= biz.store.GetUser(ctx , map[string]interface{}{"id":id,"is_deleted":false})
	if err!=nil &&data==nil{
		return err
	}
	if err:=biz.store.UpdateUser(ctx, map[string]interface{}{"id":id}, dataupdate);err!=nil{
		return err
	}
	return nil

}




