package services

import (
	"errors"
	"todo-fiber/internal/dto"
	"todo-fiber/internal/repositories"
	"todo-fiber/internal/utils"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// เช็คว่ารหัสผ่านถูกต้องหรือไม่ (ในแอปจริงควรใช้การเข้ารหัสและตรวจสอบรหัสผ่านที่ปลอดภัย)
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	claims := dto.AccessTokenClaimsRequest{
		UserID:        user.ID.String(),
		DepartmentID:  user.DepartmentID.String(),
		Role:          user.Role,
		Secret:        "your-secret-key",
		ExpiresMinute: 15,
	}

	token, err := utils.GenerateAccessToken(claims)

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *AuthService) RefreshToken(req *dto.RefreshTokenRequest) (string, error) {
	// ในแอปจริงควรตรวจสอบ refresh token และสร้าง access token ใหม่
	// ตัวอย่างนี้จะสร้าง access token ใหม่โดยไม่ตรวจสอบ refresh token จริงๆ
	claims := dto.AccessTokenClaimsRequest{
		UserID:        "user-id",
		DepartmentID:  "department-id",
		Role:          "user-role",
		Secret:        "your-secret-key",
		ExpiresMinute: 15,
	}

	token, err := utils.GenerateAccessToken(claims)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
