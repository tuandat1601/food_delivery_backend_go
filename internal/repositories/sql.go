package repositories

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}
func NewSQLStore( store *gorm.DB) *sqlStore{
	return &sqlStore{store}
}