package main

import (
	"back-end/internal/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

func openDB(dsn string) (*gorm.DB, error) {
	// Create database if it doesn't exist
	if err := createDatabaseIfNotExists(dsn); err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err = db.AutoMigrate(
		&models.User{},
		&models.Restaurant{},
		&models.Favorite{},
		&models.Rating{},
		&models.Order{},
		&models.Variation{},
		&models.FoodCategory{},
		&models.FoodItem{},
		&models.OrderItem{},
		&models.VariationOption{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (cfg *Config) connectToDB() (*gorm.DB, error) {
	connection, err := openDB(cfg.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL!")
	return connection, nil
}

func createDatabaseIfNotExists(dsn string) error {
	// Extract database name from DSN
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid DSN format")
	}

	dbNameParts := strings.Split(parts[1], "?")
	dbName := dbNameParts[0]

	// Create a DSN without the database name
	baseDSN := parts[0] + "/"

	// Connect to MySQL without specifying a database
	db, err := gorm.Open(mysql.Open(baseDSN), &gorm.Config{})
	if err != nil {
		return err
	}

	// Create database if it doesn't exist
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error
	if err != nil {
		return err
	}

	log.Printf("Database '%s' created or already exists", dbName)
	return nil
}
