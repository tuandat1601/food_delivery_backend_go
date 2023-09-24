package restaurantbiz
import (
	"context"
	
	"food_delivery/internal/model"
)
	 
type RestaurantCreationRepository interface {
	CreateRestaurant(ctx context.Context, data *model.Restaurant) error
	
}
type restaurantCreate struct {
	store RestaurantCreationRepository 
}

func NewCreateRestaurant(store RestaurantCreationRepository  ) *restaurantCreate {
	return &restaurantCreate{store}
}
func (res *restaurantCreate) CreateNewRestaurant(ctx context.Context, data *model.Restaurant) error {
	return res.store.CreateRestaurant(ctx, data)
}
