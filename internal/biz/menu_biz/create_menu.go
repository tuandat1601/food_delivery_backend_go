package menubiz

import (
	"context"
	"food_delivery/internal/model"
)

type MenuCreationRepo interface {
	CreateMenu(ctx context.Context, data *model.Menu) error
}
type menuCreate struct{
	store MenuCreationRepo
}
func NewCreateMenu(store MenuCreationRepo) *menuCreate{
return &menuCreate{store}
}
func (biz *menuCreate) CreateNewMenu( ctx context.Context, data *model.Menu) error{
	return biz.store.CreateMenu(ctx,data)
}