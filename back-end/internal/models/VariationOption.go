package models

import "gorm.io/gorm"

type VariationOption struct {
	gorm.Model
	PriceAdjustment int    `json:"price_adjustment"`
	Name            string `json:"name"`
	VariationID     uint   `json:"-"`
}
