package database

import (
	"todo-fiber/internal/models"
	"todo-fiber/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	// Implement seeding logic here, e.g., create initial users, categories, etc.
	// Example:
	// db.Create(&models.UserModel{ID: uuid.New(), Email: "
	// 	Name: "Admin User", Password: "hashed-password", Role: "admin"})

	categories := []models.CategoryModel{
		{
			ID:   uuid.New(),
			Name: "Work",
		},
		{
			ID:   uuid.New(),
			Name: "Personal",
		},
	}

	for _, category := range categories {
		db.FirstOrCreate(&category, models.CategoryModel{Name: category.Name})
	}

	departments := []models.DepartmentModel{
		{
			ID:   uuid.New(),
			Name: "HR",
		},
		{
			ID:   uuid.New(),
			Name: "IT",
		},
	}

	for _, department := range departments {
		db.FirstOrCreate(&department, models.DepartmentModel{Name: department.Name})
	}

	departmentID := departments[0].ID // Assigning the first department to users for seeding

	adminPassword, _ := utils.HashPassword("1234")
	userPassword, _ := utils.HashPassword("1234")

	users := []models.UserModel{
		{
			ID:           uuid.New(),
			Email:        "admin@system.com",
			Name:         "Admin User",
			Password:     adminPassword, // In a real application, ensure this is a securely hashed password
			Role:         "admin",
			DepartmentID: departmentID,
		},
		{
			ID:           uuid.New(),
			Email:        "user@system.com",
			Name:         "Regular User",
			Password:     userPassword, // In a real application, ensure this is a securely hashed password
			Role:         "user",
			DepartmentID: departmentID,
		},
	}
	for _, user := range users {
		db.FirstOrCreate(&user, models.UserModel{Email: user.Email})
	}
}
