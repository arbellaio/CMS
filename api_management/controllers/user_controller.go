package controllers

import (
	"api_management/database"
	"api_management/models"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

// AllUsers adding pagination
func AllUsers(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	limit := 5
	offset := (page - 1) * limit

	var users []models.User

	var total int64

	database.DB.Preload("UserRoles").Offset(offset).Limit(limit).Find(&users)

	return ctx.JSON(fiber.Map{
		"data": users,
		"metadata": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")
	database.DB.Create(&user)
	return ctx.JSON(user)
}

func GetUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var user models.User

	database.DB.Find(&user, id)
	return ctx.JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var user models.User
	user.ID = uint(id)
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var user models.User
	database.DB.Find(&user, id)
	database.DB.Delete(&user)
	return ctx.JSON(fiber.Map{
		"message": "user deleted",
	})
}
