package dbrepo

import (
	"back-end/internal/models"
	"gorm.io/gorm"
)

type MySQLDBRepo struct {
	DB *gorm.DB
}

func (m *MySQLDBRepo) Connection() *gorm.DB {
	return m.DB
}

func (m *MySQLDBRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	m.DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return nil, nil
	}
	return &user, nil
}

func (m *MySQLDBRepo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := m.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *MySQLDBRepo) GetAllRestaurants() ([]*models.Restaurant, error) {
	var restaurants []*models.Restaurant
	if err := m.DB.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (m *MySQLDBRepo) GetRestaurant(id uint) (*models.Restaurant, error) {
	var restaurant models.Restaurant

	result := m.DB.Model(&models.Restaurant{}).
		Preload("FoodCategories").
		Preload("FoodCategories.FoodItems").
		Preload("FoodCategories.FoodItems.Variations").
		Preload("FoodCategories.FoodItems.Variations.VariationOptions").
		First(&restaurant, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &restaurant, nil
}
