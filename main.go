package main

import (
	"alisafdarirepo/database"
	"alisafdarirepo/routes"
	"github.com/gofiber/fiber/v2"
)

type DB struct {
	Username string
	Password string
	DBName   string
}

func main() {
	app := fiber.New()
	database.Connect()


	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}



}
