package models

import "gorm.io/gorm"

type Variation struct {
	gorm.Model
	Name             string            `json:"name"`
	MinSelect        int               `json:"min_select"`
	MaxSelect        int               `json:"maxSelect"`
	VariationOptions []VariationOption `json:"variation_options"`
	FoodItemID       uint              `json:"-"`
}
