package repository

import (
	"back-end/internal/models"
	"gorm.io/gorm"
)

type DatabaseRepo interface {
	Connection() *gorm.DB
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetAllRestaurants() ([]*models.Restaurant, error)
	GetRestaurant(id uint) (*models.Restaurant, error)
}
