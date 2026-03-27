package models

import (
	"time"

	"github.com/google/uuid"
)

type SubjectItemModel struct {
	ID          string     `gorm:"type:char(36);primaryKey" json:"id"`
	CategoryID  string     `gorm:"type:char(36)" json:"category_id"`
	SubjectID   string     `gorm:"type:char(36)" json:"subject_id"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedBy   uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy   uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	Category CategoryModel `gorm:"foreignKey:CategoryID" json:"category"`
	Subject  SubjectModel  `gorm:"foreignKey:SubjectID" json:"subject"`
}

func (SubjectItemModel) TableName() string {
	return "subject_items"
}
