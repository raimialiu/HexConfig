package entity

import "gorm.io/gorm"

type AccountSettings struct {
	gorm.Model
	BaseEntity
	Id           int64 `gorm:"autoIncrement"`
	OutputFormat *string
	InputFormat  *string
}
