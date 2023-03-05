package entities

import (
	"time"

	"gorm.io/gorm"
)

type CreatedBase struct {
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
