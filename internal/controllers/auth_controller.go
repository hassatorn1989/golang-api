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

func (ac *AuthController) Login(ctx fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, err := ac.AuthService.Login(&req)
	if err != nil || token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return ctx.JSON(fiber.Map{"token": token})
}

func (ac *AuthController) RefreshToken(ctx fiber.Ctx) error {
	var req dto.RefreshTokenRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	token, err := ac.AuthService.RefreshToken(&req)
	if err != nil || token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid refresh token"})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
