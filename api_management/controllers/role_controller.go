package controllers

import (
	"api_management/database"
	"api_management/dto"
	"api_management/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(ctx *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Find(&roles)
	return ctx.JSON(roles)
}

//role creation changed to role_permission_dto

func CreateRoleWithPermissions(ctx *fiber.Ctx) error {
	var rolePermissionDto dto.RolePermissionDto

	if err := ctx.BodyParser(&rolePermissionDto); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(rolePermissionDto.Permissions))

	for i, permissionId := range rolePermissionDto.Permissions {
		permissions[i] = models.Permission{
			PermissionId: uint(permissionId),
		}
	}

	role := models.Role{
		RoleId:      rolePermissionDto.RoleId,
		Name:        rolePermissionDto.Name,
		Permissions: permissions,
	}

	database.DB.Create(&role)
	return ctx.JSON(role)
}

func CreateRole(ctx *fiber.Ctx) error {
	var role models.Role

	if err := ctx.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Create(&role)
	return ctx.JSON(role)
}

func GetRoleWithPermissions(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var role models.Role

	database.DB.Preload("Permissions").Find(&role, id)
	return ctx.JSON(role)
}

func GetRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var role models.Role

	database.DB.Find(&role, id)
	return ctx.JSON(role)
}

func UpdateRoleWithPermissions(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var rolePermissionDto dto.RolePermissionDto
	rolePermissionDto.Id = uint(id)

	if err := ctx.BodyParser(&rolePermissionDto); err != nil {
		return err
	}

	permissions := make([]models.Permission, len(rolePermissionDto.Permissions))

	for i, permissionId := range rolePermissionDto.Permissions {
		permissions[i] = models.Permission{
			PermissionId: uint(permissionId),
		}
	}

	//updating new role permissions and removing old permissions
	database.DB.Table("role_permissions").Where("role_id", id).Delete(rolePermissionDto)

	role := models.Role{
		RoleId:      rolePermissionDto.RoleId,
		Name:        rolePermissionDto.Name,
		Permissions: permissions,
	}
	role.ID = rolePermissionDto.Id
	database.DB.Model(&role).Updates(role)
	return ctx.JSON(role)
}

func UpdateRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var role models.Role
	role.ID = uint(id)
	if err := ctx.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)
	return ctx.JSON(role)
}

func DeleteRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var role models.Role
	database.DB.Find(&role, id)
	database.DB.Delete(&role)
	return ctx.JSON(fiber.Map{
		"message": "role deleted",
	})
}
