package routes

import (
	"api_management/controllers"
	"api_management/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	//Profile Controller
	app.Put("/api/users/info", controllers.UpdateProfileInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	//User Controller
	app.Get("/api/user", controllers.User)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)

	//Role Controller
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Post("/api/rolePermissions", controllers.CreateRoleWithPermissions)
	app.Put("/api/rolePermissions/:id", controllers.UpdateRoleWithPermissions)
	app.Get("/api/rolePermissions/:id", controllers.GetRoleWithPermissions)

	//Permission Controller
	app.Get("/api/permissions", controllers.AllPermissions)

	//Product Controller
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)
	app.Get("/api/products", controllers.AllProducts)
	app.Post("/api/product", controllers.CreateProduct)

	app.Get("/api/logout", controllers.LogOut)
}
