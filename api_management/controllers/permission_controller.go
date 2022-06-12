package controllers

import (
	"api_management/database"
	"api_management/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(ctx *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Find(&permissions)
	return ctx.JSON(permissions)
}
