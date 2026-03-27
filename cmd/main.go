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

	// Migrate database
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// seed data
	database.SeedDB(db)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
