package models

import (
	"time"

	"github.com/google/uuid"
)

type AnswerModel struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	SubjectID     string    `json:"subject_id"`
	SubjectTypeID string    `json:"subject_type_id"`
	Answer        string    `json:"answer"`
	IPAddress     string    `json:"ip_address"`
	Comment       string    `json:"comment"`
	CreatedBy     uuid.UUID `json:"created_by"`
	UpdatedBy     uuid.UUID `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
