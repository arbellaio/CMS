package middlewares

import (
	"api_management/util"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(ctx *fiber.Ctx) error {

	cookie := ctx.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return ctx.Next()
}
