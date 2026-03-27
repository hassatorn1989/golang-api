package services

import (
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

func (s *AuthService) Login(dto *dto.LoginRequest) (string, error) {
	user, _ := s.userRepo.GetByEmail(dto.Username)
	// เช็คว่าผู้ใช้มีอยู่ในระบบหรือไม่
	if user == nil {
		return "", nil
	}
	// เช็คว่ารหัสผ่านถูกต้องหรือไม่ (ในแอปจริงควรใช้การเข้ารหัสและตรวจสอบรหัสผ่านที่ปลอดภัย)
	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return "", nil
	}

	token, err := utils.GenerateAccessToken(user.ID.String(), user.DepartmentID.String(), user.Role, "your-secret-key", 15) // Example: 15 minutes expiration

	if err != nil {
		return "", err
	}

	return token, nil
}
