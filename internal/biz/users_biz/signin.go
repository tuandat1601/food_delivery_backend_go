package users_biz

import (
	"context"
	"food_delivery/internal/model"
)

type SignInStorage interface {
	SignIn(ctx context.Context, data *model.User) error
}
type signInBiz struct {
	store SignInStorage
}

func NewSignInUser(store SignInStorage) *signInBiz {
	return &signInBiz{store}
}
func (biz *signInBiz) SignInUser(ctx context.Context, data *model.User) error {

	if err := biz.store.SignIn(ctx, data); err != nil {
		return err
	}
	return nil
}
