package database

import (
	"todo-fiber/internal/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.DepartmentModel{},
		&models.CategoryModel{},
		&models.UserModel{},
		&models.SubjectModel{},
		&models.SubjectItemModel{},
		&models.SubjectTypeAnswerModel{},
		&models.AnswerModel{},
		&models.AnswerItemModel{},
	)
}
