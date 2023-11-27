package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Email         string `gorm:"index"`
	Password      string
	OwnerId       *string `gorm:"index"`
	IsSubAccount  *bool
	WorkspaceName string
	WorkspaceUrl  string `gorm:"index"`

	IsEmailVerified *bool

	AccountId     string `gorm:"primaryKey"`
	AccountSecret string `gorm:"index"`

	Workspace []Workspace
}

type AccountPermissions struct {
}
