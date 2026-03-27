package controllers

import (
	"todo-fiber/internal/dto"
	"todo-fiber/internal/services"

	"github.com/gofiber/fiber/v3"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// func (ac *AuthController) Register(ctx fiber.Ctx) error {
// 	var req dto.RegisterRequest
// 	if err := ctx.Bind().Body(&req); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
// 	}
// 	// Implement registration logic here, e.g., create user in the database
// 	return ctx.JSON(fiber.Map{"message": "user registered successfully"})
// }

func (ac *AuthController) Login(ctx fiber.Ctx) error {
	var req dto.LoginRequest
	token, err := ac.AuthService.Login(&req)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
