package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	PermissionId uint   `json:"permission_id"`
	Name         string `json:"name"`
}
