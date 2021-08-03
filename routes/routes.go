package routes

import (
	"alisafdarirepo/controllers"
	"alisafdarirepo/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	/* Crud route for Users */
	app.Get("/api/users", middlewares.AuthRequired(), controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUserById)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	// Crud routes for User Role
	app.Get("/api/roles", controllers.AllRoles)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRoleById)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
	//permission routes
	app.Get("/api/permissions", controllers.AllPermissions)
	// order routes
	app.Get("/api/orders", controllers.FindOrderByUserId)
}
