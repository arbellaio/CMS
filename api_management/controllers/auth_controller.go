package controllers

import (
	"api_management/database"
	"api_management/models"
	"api_management/util"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	if data["confirmed_password"] != data["password"] {
		ctx.Status(400)
		return ctx.JSON(
			fiber.Map{"message": "passwords do not match"})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message": "not found",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

type Claims struct {
	jwt.StandardClaims
}

func User(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	return ctx.JSON(user)
}

func LogOut(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
