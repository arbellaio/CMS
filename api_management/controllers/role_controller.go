package controllers

import (
	"api_management/database"
	"api_management/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(ctx *fiber.Ctx) error {
	var roles []models.Role
	database.DB.Find(&roles)
	return ctx.JSON(roles)
}

func CreateRole(ctx *fiber.Ctx) error {
	var role models.Role

	if err := ctx.BodyParser(&role); err != nil {
		return err
	}
	database.DB.Create(&role)
	return ctx.JSON(role)
}

func GetRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var role models.Role

	database.DB.Find(&role, id)
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
