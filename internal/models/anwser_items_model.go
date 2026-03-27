package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerItemModel struct {
	ID            uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	AnswerID      uuid.UUID  `gorm:"type:char(36)" json:"answer_id"`
	SubjectID     uuid.UUID  `gorm:"type:char(36)" json:"subject_id"`
	SubjectItemID uuid.UUID  `gorm:"type:char(36)" json:"subject_item_id"`
	Answer        string     `gorm:"type:text" json:"answer"`
	CreatedBy     uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy     uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (r *AnswerItemModel) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (AnswerItemModel) TableName() string {
	return "answer_items"
}
