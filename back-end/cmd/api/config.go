package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	DSN           string
	Domain        string
	JWTSecret     string
	JWTIssuer     string
	JWTAudience   string
	CookieDomain  string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
}

func Load() Config {
	var cfg Config

	// read mysql username and password from command line
	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	flag.StringVar(&cfg.DSN, "dsn", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/food4student?charset=utf8mb4&parseTime=True&loc=Local", mysqlUsername, mysqlPassword), "MySQL connection string")
	flag.StringVar(&cfg.JWTSecret, "jwt-secret", "stevejobsdieofligma", "signing secret")
	flag.StringVar(&cfg.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&cfg.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&cfg.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&cfg.Domain, "domain", "example.com", "domain")
	flag.Parse()

	return cfg
}
