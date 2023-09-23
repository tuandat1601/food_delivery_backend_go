package restaurantusecase

import (
	"context"
	"food_delivery/internal/model"

)
type RestaurantCreationRepository interface {
	CreateRestaurant(ctx context.Context, data *model.Restaurants) error
	UpdateRestaurant(ctx context.Context,id int, cond map[string]interface{},dataupdate *model.Restaurants) error
}

