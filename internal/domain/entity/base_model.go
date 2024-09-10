package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
