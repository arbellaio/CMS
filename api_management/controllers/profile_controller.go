package controllers

import (
	"api_management/database"
	"api_management/models"
	"api_management/util"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func UpdateProfileInfo(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	cookie := ctx.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.ID = uint(userId)

	database.DB.Model(&user).Updates(user)

	return ctx.JSON(user)
}

func UpdatePassword(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	cookie := ctx.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	if data["confirmed_password"] != data["password"] {
		ctx.Status(400)
		return ctx.JSON(
			fiber.Map{"message": "passwords do not match"})
	}

	user := models.User{}

	user.SetPassword(data["password"])
	database.DB.Model(&user).Where("id = ?", id).Updates(data)

	return ctx.JSON(user)
}
