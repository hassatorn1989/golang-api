package repositories

import (
	"todo-fiber/internal/dto"
	"todo-fiber/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetByEmail(email string) (*models.UserModel, error) {
	user := &models.UserModel{}
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetByID(id string) (*models.UserModel, error) {
	user := &models.UserModel{}
	if err := ur.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Create(req *dto.CreateUserRequest) error {
	user := &models.UserModel{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password, // In a real application, ensure this is a securely hashed password
		Role:     req.Role,
	}
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Update(id string, req *dto.UpdateUserRequest) error {
	user := &models.UserModel{}
	if err := ur.db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password
	user.Role = req.Role

	if err := ur.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id string) error {
	if err := ur.db.Where("id = ?", id).Delete(&models.UserModel{}).Error; err != nil {
		return err
	}
	return nil
}
