package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID       uint `json:"user_id"`
	RestaurantID uint `json:"restaurant_id"`
}
