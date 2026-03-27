package models

import (
	"time"

	"github.com/google/uuid"
)

type AnswerItemModel struct {
	ID            uuid.UUID `gorm:"primaryKey" json:"id"`
	AnswerID      string    `json:"answer_id"`
	SubjectID     string    `json:"subject_id"`
	SubjectItemID string    `json:"subject_item_id"`
	Answer        string    `json:"answer"`
	CreatedBy     uuid.UUID `json:"created_by"`
	UpdatedBy     uuid.UUID `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (AnswerItemModel) TableName() string {
	return "answer_items"
}
