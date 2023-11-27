package entity

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	BaseEntity
	Id          string  `gorm:"primaryKey"`
	Name        string  `gorm:"index"`
	Description *string `gorm:"null"`

	Projects []Project
}
