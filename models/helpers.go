package models

import (
	"time"
)

// DbModelNoID is the base model for gorm database structs with a custom ID
type DbModelNoID struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
