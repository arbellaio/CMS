package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleId uint   `json:"role_id"`
	Name   string `json:"name"`
}
