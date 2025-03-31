package repository

import "gorm.io/gorm"

type DatabaseRepo interface {
	Connection() *gorm.DB
}
