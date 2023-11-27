package entity

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UpdatedAt *time.Time
	CreatedAt *time.Time `gorm:"index"`

	IsActive  bool
	IsDeleted *bool
}
