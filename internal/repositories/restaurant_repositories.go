package repositories

import (
	"context"
	"errors"
	"food_delivery/internal/model"
)
	 


func (s *sqlStore) CreateRestaurant(ctx context.Context,  data *model.Restaurant) error{
	var existing model.Restaurant
	if err := s.db.Table("restaurants").Where(map[string]interface{}{"name": data.Name, "user_id":data.UserId,"is_deleted": false,}).First(&existing).Error; err == nil {
		return errors.New("This restaurant is already")
	}
	if err:=s.db.Table("restaurants").Create(&data).Error;err!=nil{
		return err
	}
	return nil
}
func (s *sqlStore) GetRestaurant(ctx context.Context,  cond map[string]interface{})(*model.Restaurant, error){
	var existing model.Restaurant
	if err := s.db.Table("restaurants").Where(cond).First(&existing).Error; err != nil {
		return nil, err
	}

	return &existing,nil
}
func (s *sqlStore) UpdateRestaurant(ctx context.Context,  cond map[string]interface{}, data *model.Restaurant)( error){

	if err := s.db.Table("restaurants").Where(cond).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
func (s *sqlStore) DeleteRestaurant(ctx context.Context,  cond map[string]interface{})( error){

	if err := s.db.Table("restaurants").Where(cond).Updates(map[string]interface{}{"is_deleted":true}).Error; err != nil {
		return err
	}

	return nil
}