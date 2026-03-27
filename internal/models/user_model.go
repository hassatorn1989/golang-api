package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID           uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	DepartmentID uuid.UUID  `gorm:"type:char(36)" json:"department_id"`
	Email        string     `gorm:"type:char(100);unique;not null" json:"email"`
	Name         string     `gorm:"type:char(100);not null" json:"name"`
	Password     string     `gorm:"type:char(255);not null" json:"-"`
	Role         string     `gorm:"enum('user', 'admin');default:'user'" json:"role"`
	CreatedBy    uuid.UUID  `json:"created_by"`
	UpdatedBy    uuid.UUID  `json:"updated_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	Department DepartmentModel `gorm:"foreignKey:DepartmentID" json:"department"`
}

func (UserModel) TableName() string {
	return "users"
}
