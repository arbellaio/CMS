package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleId uint   `json:"role_id"`
	Name   string `json:"name"`

	//Many To Many With Help Of GORM
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
