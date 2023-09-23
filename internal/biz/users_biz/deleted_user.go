package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type DeletedUserStorage interface {
	GetUser(ctx context.Context,cond map[string]interface{})(*model.User,error)
	DeletedUser(ctx context.Context,cond map[string]interface{}) error
}
type deletedUser struct{
	store DeletedUserStorage
}
func NewDeletedUser(store DeletedUserStorage) *deletedUser{
	return &deletedUser{store}
}
func (biz *deletedUser) DeletedUseById(ctx context.Context, id int) error{
	data,err:= biz.store.GetUser(ctx,map[string]interface{}{"id":id,"is_deleted":false})
	if err!=nil&& data==nil{
		return err
	}
	if err := biz.store.DeletedUser(ctx, map[string]interface{}{"id":id});err!=nil{
		return err
	} 
	return nil	
}