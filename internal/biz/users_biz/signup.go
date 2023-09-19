package users_biz

import (
	"context"
	"food_delivery/internal/model"

	
)

type SignUpStorage interface {
	SignUp(ctx context.Context, data *model.User) error
}
type signUpBiz struct {
	store SignUpStorage
}

func NewSignUpUser(store SignUpStorage) *signUpBiz {
	return &signUpBiz{store}
}
func (biz *signUpBiz) SignUpUser(ctx context.Context, data *model.User) error {
	
	if err := biz.store.SignUp(ctx, data); err != nil {
		return err
	}
	return nil
}
