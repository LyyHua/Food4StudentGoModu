package models

import (
	"gorm.io/gorm"
	"time"
)

type Restaurant struct {
	gorm.Model
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Address        string         `json:"address"`
	PhoneNumber    string         `json:"phone_number"`
	Status         Status         `json:"status"`
	LogoUrl        string         `json:"logo_url"`
	BannerUrl      string         `json:"banner_url"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	FoodCategories []FoodCategory `json:"food_categories"`
	Orders         []Order        `json:"orders"`
	Favorites      []Favorite     `json:"-"`
	TotalRating    int            `json:"total_rating"`
	AverageRating  float64        `json:"average_rating"`
}

type Status int

const (
	Pending Status = iota
	Approved
	Banned
	Suspended
	Open
	Closed
)
