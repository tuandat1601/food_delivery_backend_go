package repositories

import (
	"context"
	"food_delivery/internal/model"


	"food_delivery/internal/usescases"
)

// type RestaurantCreationRepository interface {
// 	CreateRestaurant(ctx context.Context, data *model.Restaurants) error
// }

type restaurantUsecase struct {
	store restaurantusecase.
}

func NewRestaurantUsecase(store RestaurantCreationRepository ) *restaurantUsecase {
	return &restaurantUsecase{store}
}
func (res *restaurantUsecase) CreateNewRestaurant(ctx context.Context, data *model.Restaurants) error {
	return res.store.CreateRestaurant(ctx, data)
}

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *model.Restaurants) error{
	if err:=s.db.Table("restaurants").Create(&data).Error;err!=nil{
		return err
	}
	return nil
}