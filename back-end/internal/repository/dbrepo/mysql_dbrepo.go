package dbrepo

import "gorm.io/gorm"

type MySQLDBRepo struct {
	DB *gorm.DB
}

func (m *MySQLDBRepo) Connection() *gorm.DB {
	return m.DB
}
