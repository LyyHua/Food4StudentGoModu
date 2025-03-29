package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Price            int    `json:"price"`
	Quantity         int    `json:"quantity"`
	FoodItemPhotoUrl string `json:"food_item_photo_url"`
	Variations       string `json:"variations"`
	OrderID          uint   `json:"-"`
}
