package models

import (
	"time"

	"github.com/google/uuid"
)

type SubjectItemModel struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	CategoryID  string    `json:"category_id"`
	SubjectID   string    `json:"subject_id"`
	Description string    `json:"description"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Associations
	Category CategoryModel `gorm:"foreignKey:CategoryID" json:"category"`
	Subject  SubjectModel  `gorm:"foreignKey:SubjectID" json:"subject"`
}

func (SubjectItemModel) TableName() string {
	return "subject_items"
}
