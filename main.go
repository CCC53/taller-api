package main

import (
	"log"
	"taller-api/config"
	"taller-api/db"
	"taller-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.GetAllowedOrigins(),
		AllowCredentials: true,
	}))

	db.InitDB()
	routes.InitRoutes(app)

	log.Println("Server on port " + config.GetPort())
	app.Listen(":" + config.GetPort())
}
