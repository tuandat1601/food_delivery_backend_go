package repositories

import (
	"context"
	"errors"
	"food_delivery/common"

	"food_delivery/internal/model"

	"log"
)

func (s *sqlStore) SignUp(ctx context.Context, data *model.User) error {
	name := data.UserName
	email := data.Email
	var existingUser model.User
	log.Println(name, email)
	if err := s.db.Table("users").Where(map[string]interface{}{"email": email, "is_deleted": false}).First(&existingUser).Error; err == nil {
		return errors.New("Email is already")
	}
	if err := s.db.Table("users").Where(map[string]interface{}{"username": name, "is_deleted": false}).First(&existingUser).Error; err == nil {
		return errors.New("Username is already")
	}
	HashPassword, err := common.HashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = HashPassword
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) SignIn(ctx context.Context, data *model.User) error {
	password := data.Password
	if err := s.db.Table("users").Where(map[string]interface{}{"username": data.UserName, "is_deleted": false}).First(&data).Error; err != nil {
		return errors.New("username  wrong")
	}
	if !common.CheckPasswordHash(password, data.Password) {
		return errors.New("Wrong password")
	}

	return nil
}
func (s *sqlStore) GetListUser(ctx context.Context, cond map[string]interface{}) ([]model.User, error) {
	var result []model.User

	if err := s.db.Table("users").Where(cond).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
func (s *sqlStore) GetUser(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var result model.User

	if err := s.db.Table("users").Where(cond).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
func (s *sqlStore)DeletedUser(ctx context.Context, cond map[string]interface{})error{
	if err:=s.db.Table("users").Where(cond).Update("is_deleted",true).Error;err!=nil{
		return err
	}
	return nil
}
func (s *sqlStore)UpdateUser(ctx context.Context, cond map[string]interface{}, dataupdate * model.User)error{
	if err:=s.db.Table("users").Where(cond).Updates(dataupdate).Error;err!=nil{
		return err
	}
	return nil
}
