package restaurantbiz
import (
	"context"
	
	"food_delivery/internal/model"
)
	 
type RestaurantUpdateRepository interface {
	GetRestaurant(ctx context.Context, cond map[string]interface{})(*model.Restaurant, error)
	UpdateRestaurant(ctx context.Context,cond map[string]interface{}, data *model.Restaurant) error
	
}
type restaurantUpdate struct {
	store RestaurantUpdateRepository 
}

func NewUpdateRestaurant(store RestaurantUpdateRepository  ) *restaurantUpdate {
	return &restaurantUpdate{store}
}
func (res *restaurantUpdate) UpdateNewRestaurant(ctx context.Context,user_id, res_id int, data *model.Restaurant) error {
	data, err:= res.store.GetRestaurant(ctx , map[string]interface{}{"id":res_id,"user_id":user_id,"is_deleted":false})
	if err!=nil &&data==nil{
		return err
	}
	return res.store.UpdateRestaurant(ctx,map[string]interface{}{"id": res_id,"user_id":user_id,"is_deleted": false}, data)
}
