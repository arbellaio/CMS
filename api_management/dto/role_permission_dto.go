package dto

type RolePermissionDto struct {
	RoleId      uint   `json:"role_id"`
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
}
