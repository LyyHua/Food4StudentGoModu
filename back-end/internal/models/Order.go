package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderStatus     OrderStatus `json:"order_status"`
	OrderItems      []OrderItem `json:"order_items"`
	Note            string      `json:"note"`
	TotalPrice      int         `json:"total_price"`
	ShippingAddress string      `json:"shipping_address"`
	PhoneNumber     string      `json:"phone_number"`
	Orderer         string      `json:"orderer"`
	UserID          uint        `json:"orderer_id"`
	RestaurantID    uint        `json:"restaurant_id"`
}

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	Accepted
	Delivering
	Finished
	Rejected
)
