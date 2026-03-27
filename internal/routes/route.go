package routes

import (
	// "todo-fiber/internal/middleware"

	"todo-fiber/internal/controllers"
	"todo-fiber/internal/repositories"
	"todo-fiber/internal/services"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(
	app *fiber.App,
	db *gorm.DB,
	jwtSecret string,
	jwtRefreshSecret string,
	jwtExpireMinute int,
	jwtRefreshExpireDay int,
) {
	userRepository := repositories.NewUserRepository(db)
	authController := controllers.NewAuthController(services.NewAuthService(userRepository))
	// authHandler := handlers.NewAuthHandler(
	// 	db,
	// 	jwtSecret,
	// 	jwtRefreshSecret,
	// 	jwtExpireMinute,
	// 	jwtRefreshExpireDay,
	// )
	// todoHandler := handlers.NewTodoHandler(db)

	api := app.Group("/api")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "todo fiber api is running"})
	})

	auth := api.Group("/auth")
	// auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	// auth.Post("/refresh", authController.Refresh)
	// auth.Post("/logout", authController.Logout)

	// protected := api.Group("", middleware.Protected(jwtSecret))
	// protected.Get("/me", authHandler.Me)

	// todos := protected.Group("/todos")
	// todos.Post("/", todoHandler.CreateTodo)
	// todos.Get("/", todoHandler.GetTodos)
	// todos.Get("/:id", todoHandler.GetTodoByID)
	// todos.Put("/:id", todoHandler.UpdateTodo)
	// todos.Delete("/:id", todoHandler.DeleteTodo)
}
