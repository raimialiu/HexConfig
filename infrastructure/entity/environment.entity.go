package entity

import "gorm.io/gorm"

type Environment struct {
	gorm.Model
	BaseEntity

	Id          string `gorm:"primaryKey"`
	Name        string
	Description *string
}
