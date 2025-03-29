package models

import "gorm.io/gorm"

type FoodItem struct {
	gorm.Model
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	BasePrice      string      `json:"base_price"`
	PhotoUrl       string      `json:"photo_url"`
	Variations     []Variation `json:"variations"`
	FoodCategoryId uint        `json:"-"`
}
