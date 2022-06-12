package controllers

import (
	"api_management/database"
	"api_management/models"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

// AllProducts adding pagination
func AllProducts(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	limit := 5
	offset := (page - 1) * limit

	var products []models.Product

	var total int64

	database.DB.Preload("UserRoles").Offset(offset).Limit(limit).Find(&products)

	return ctx.JSON(fiber.Map{
		"data": products,
		"metadata": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func CreateProduct(ctx *fiber.Ctx) error {
	var product models.Product

	if err := ctx.BodyParser(&product); err != nil {
		return err
	}
	database.DB.Create(&product)
	return ctx.JSON(product)
}

func GetProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var product models.Product

	database.DB.Find(&product, id)
	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var product models.Product
	product.ID = uint(id)
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)
	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var user models.Product
	database.DB.Find(&user, id)
	database.DB.Delete(&user)
	return ctx.JSON(fiber.Map{
		"message": "user deleted",
	})
}
