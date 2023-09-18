package config

import "os"

var Secret = "penny"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	return port
}

func GetDSN() string {
	dsn := os.Getenv("DB")
	if dsn == "" {
		dsn = "host=localhost user=carlos password=penny dbname=taller-db port=5432"
	}
	return dsn
}

func GetAllowedOrigins() string {
	allowedOrigins := os.Getenv("ALLOWED")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:4200"
	}
	return allowedOrigins
}
