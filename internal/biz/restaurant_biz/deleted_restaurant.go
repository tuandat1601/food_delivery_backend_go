package restaurantbiz

import (
	"context"
	"food_delivery/internal/model"
)
	 
type RestaurantDeletedRepository interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{})(*model.Restaurant, error)
	DeleteRestaurant(ctx context.Context, cond map[string]interface{})( error)
	
}
type restaurantDelete struct {
	store RestaurantDeletedRepository 
}

func NewDeleteRestaurant(store RestaurantDeletedRepository  ) *restaurantDelete {
	return &restaurantDelete{store}
}
func (res *restaurantDelete) DeleteNewRestaurant(ctx context.Context,user_id , res_id int) ( error) {
	data, err:= res.store.GetRestaurant(ctx , map[string]interface{}{"id":res_id,"user_id":user_id,"is_deleted":false})
	if err!=nil &&data==nil{
		return err
	}
	return res.store.DeleteRestaurant(ctx, map[string]interface{}{"id": res_id,"user_id":user_id,"is_deleted": false})
}


