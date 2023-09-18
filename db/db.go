package db

import (
	"log"
	"taller-api/config"
	"taller-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = config.GetDSN()
var DB *gorm.DB

func initConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB online")
	}
}

func initMigrations() {
	DB.AutoMigrate(models.Vehicle{}, models.Service{}, models.SparePart{}, models.Employee{})
}

func InitDB() {
	initConnection()
	initMigrations()
}
