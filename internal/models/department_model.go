package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentModel struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string     `gorm:"type:char(100);unique;not null" json:"name"`
	CreatedBy uuid.UUID  `gorm:"type:char(36)" json:"created_by"`
	UpdatedBy uuid.UUID  `gorm:"type:char(36)" json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (d *DepartmentModel) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}

func (DepartmentModel) TableName() string {
	return "departments"
}
