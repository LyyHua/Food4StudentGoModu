package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string       `json:"first_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	Address     string       `json:"address"`
	PhoneNumber string       `json:"phone_number"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
	Restaurants []Restaurant `json:"restaurants"`
	Favorites   []Favorite   `json:"favorites"`
	Ratings     []Rating     `json:"ratings"`
	Orders      []Order      `json:"orders"`
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
