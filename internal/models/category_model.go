package models

import (
	"time"

	"github.com/google/uuid"
)

type CategoryModel struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
