package models

import (
	"time"

	"github.com/google/uuid"
)

type SubjectModel struct {
	ID           uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	DepartmentID uuid.UUID  `gorm:"type:char(36)" json:"department_id"`
	Title        string     `gorm:"type:char(100);not null" json:"title"`
	Description  string     `gorm:"type:text" json:"description"`
	Status       string     `gorm:"enum('active', 'inactive');default:'active'" json:"status"`
	StartDate    time.Time  `json:"start_date"`
	EndDate      time.Time  `json:"end_date"`
	Year         int        `json:"year"`
	MaxAnswers   int        `json:"max_answers"`
	CreatedBy    uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy    uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty"`

	// Associations
	Department DepartmentModel `gorm:"foreignKey:DepartmentID" json:"department"`
}

func (SubjectModel) TableName() string {
	return "subjects"
}
