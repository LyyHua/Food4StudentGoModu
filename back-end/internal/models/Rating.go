package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Comment string `json:"comment"`
	Stars   int    `json:"stars"`
	UserID  uint   `json:"user_id"`
	OrderID uint   `json:"order_id"`
}
