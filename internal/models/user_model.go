package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	DepartmentID uuid.UUID `json:"department_id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	Name         string    `json:"name"`
	Password     string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"enum('user', 'admin');default:'user'" json:"role"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Associations
	Department DepartmentModel `gorm:"foreignKey:DepartmentID" json:"department"`
}

func (UserModel) TableName() string {
	return "users"
}
