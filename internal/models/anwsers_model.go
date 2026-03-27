package models

import (
	"time"

	"github.com/google/uuid"
)

type AnswerModel struct {
	ID            uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	SubjectID     uuid.UUID  `gorm:"type:char(36)" json:"subject_id"`
	SubjectTypeID uuid.UUID  `gorm:"type:char(36)" json:"subject_type_id"`
	Answer        string     `gorm:"type:text" json:"answer"`
	IPAddress     string     `gorm:"type:char(45)" json:"ip_address"`
	Comment       string     `gorm:"type:text" json:"comment"`
	CreatedBy     uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy     uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (AnswerModel) TableName() string {
	return "answers"
}
