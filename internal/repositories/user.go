package repositories

import (
	"context"
	"errors"
	"food_delivery/internal/model"
	"log"
)

func (s *sqlStore) SignUp(ctx context.Context, data *model.User) error {
	name := data.UserName
	email := data.Email
	var existingUser model.User
	log.Println(name,email)
	if err := s.db.Table("users").Where(map[string]interface{}{"email": email,"is_deleted":false}).First(&existingUser).Error; err == nil {
		return errors.New("Email is already")
	}
	if err := s.db.Table("users").Where(map[string]interface{}{"username": name,"is_deleted":false}).First(&existingUser).Error; err == nil {
		return errors.New("Username is already")
	}

	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
func (s *sqlStore) SignIn(ctx context.Context, data *model.User) error {
	if err := s.db.Table("users").Where(map[string]interface{}{"username": data.UserName,"password":data.Password,"is_deleted":false}).First(&data).Error; err != nil {
		return errors.New("username or password wrong")
	}
	return nil
}
