package restaurantbiz
import (
	"context"
	
	"food_delivery/internal/model"
)
	 
type RestaurantGetRepository interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{})(*model.Restaurant, error)
	
}
type restaurantGet struct {
	store RestaurantGetRepository 
}

func NewGetRestaurant(store RestaurantGetRepository  ) *restaurantGet {
	return &restaurantGet{store}
}
func (res *restaurantGet) GetNewRestaurant(ctx context.Context,user_id , res_id int) (*model.Restaurant, error) {
	return res.store.GetRestaurant(ctx, map[string]interface{}{"id": res_id,"user_id":user_id,"is_deleted": false})
}


