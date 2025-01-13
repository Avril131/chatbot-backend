package models

import (
	"time"

	"gorm.io/gorm"
)

// ID : primary key
type ID struct {
	ID int `json:"id" gorm:"primaryKey"`
}

// Timestamps : CreateTime && UpdateTime
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDeletes : deletedTime or -1
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
