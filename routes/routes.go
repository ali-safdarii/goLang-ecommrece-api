package routes

import (
	"alisafdarirepo/controllers"
	"alisafdarirepo/middlewares"
	"github.com/gofiber/fiber/v2"
)



func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/users", middlewares.AuthRequired(),controllers.User)

}
