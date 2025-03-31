package main

import (
	"back-end/internal/repository"
	"back-end/internal/repository/dbrepo"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type application struct {
	DB   repository.DatabaseRepo
	auth Auth
}

func main() {
	// set application config
	cfg := Load()

	// create application instance with config
	app := application{
		auth: Auth{
			Issuer:        cfg.JWTIssuer,
			Audience:      cfg.JWTAudience,
			Secret:        cfg.JWTSecret,
			TokenExpiry:   time.Minute * 15,
			RefreshExpiry: time.Hour * 24,
			CookiePath:    "/",
			CookieName:    "__Host-refresh_token",
			CookieDomain:  cfg.CookieDomain,
		},
	}

	// connect to the database
	conn, err := cfg.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.MySQLDBRepo{DB: conn}

	// Get the underlying SQL DB and defer its closure
	mysqlDB, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(mysqlDB *sql.DB) {
		err := mysqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(mysqlDB)

	// start the server
	router := gin.Default()
	log.Println("Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
