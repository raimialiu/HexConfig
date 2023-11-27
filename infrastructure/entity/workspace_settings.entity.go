package entity

import "gorm.io/gorm"

type WorkspaceSettings struct {
	gorm.Model
	BaseEntity
	Id           int64 `gorm:"autoIncrement"`
	OutputFormat *string
	InputFormat  *string

	DeletionPolicy string
}
