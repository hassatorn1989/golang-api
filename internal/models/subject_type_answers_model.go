package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectTypeAnswerModel struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	SubjectID uuid.UUID  `gorm:"type:char(36)" json:"subject_id"`
	Answer    string     `gorm:"type:text" json:"answer"`
	CreatedBy uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (a *SubjectTypeAnswerModel) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

func (SubjectTypeAnswerModel) TableName() string {
	return "subject_type_answers"
}
