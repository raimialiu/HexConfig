package entity

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	BaseEntity

	Id          string `gorm:"primaryKey"`
	Name        string
	Description *string

	Environments []Environment
}
