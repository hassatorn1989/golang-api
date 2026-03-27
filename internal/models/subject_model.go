package models

import (
	"time"

	"github.com/google/uuid"
)

type SubjectModel struct {
	ID           uuid.UUID `gorm:"primaryKey" json:"id"`
	DepartmentID uuid.UUID `json:"department_id"`
	Title        string    `gorm:"not null" json:"title"`
	Description  string    `json:"description"`
	Status       string    `gorm:"enum('active', 'inactive');default:'active'" json:"status"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	Year         int       `json:"year"`
	MaxAnswers   int       `json:"max_answers"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Associations
	Department DepartmentModel `gorm:"foreignKey:DepartmentID" json:"department"`
}

func (SubjectModel) TableName() string {
	return "subjects"
}
