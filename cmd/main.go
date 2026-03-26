package main

import (
	"log"

	"todo-fiber/internal/config"
	"todo-fiber/internal/database"
	"todo-fiber/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg.DBDSN)

	app := fiber.New()

	routes.SetupRoutes(
		app,
		db,
		cfg.JWTSecret,
		cfg.JWTRefreshSecret,
		cfg.JWTExpireMinute,
		cfg.JWTRefreshExpireDay,
	)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
