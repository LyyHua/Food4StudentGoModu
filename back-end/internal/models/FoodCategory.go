package models

import "gorm.io/gorm"

type FoodCategory struct {
	gorm.Model
	Name         string     `json:"name"`
	FoodItems    []FoodItem `json:"food_items"`
	RestaurantID uint       `json:"-"`
}
