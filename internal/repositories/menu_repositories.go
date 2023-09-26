package repositories

import (
	"context"
	"errors"
	"food_delivery/internal/model"
)

func (s *sqlStore) CreateMenu(ctx context.Context, data *model.Menu) error{
	var existing model.Menu
	if err := s.db.Table("menus").Where(map[string]interface{}{"name": data.Name, "restaurant_id":data.RestaurantID,"is_deleted": false,}).First(&existing).Error; err == nil {
		return errors.New("This Menu is already")
	}
	if err:=s.db.Table("menus").Create(&data).Error;err!=nil{
		return err
	}
	return nil
}