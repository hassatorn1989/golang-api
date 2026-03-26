package database

import (
	"log"

	"todo-fiber/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("automigrate error:", err)
	}

	return db
}
