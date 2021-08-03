package controllers

import (
	"alisafdarirepo/database"
	"alisafdarirepo/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func FindOrderByUserId(c *fiber.Ctx) error {

	var orders []models.Order

	result := database.DB.Preload("Product").Where("user_id", 2).
		Preload("User").Find(&orders)
	err := result.Error

	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "",
			"data": err,
		})
	}

	fmt.Printf("Retreving All Records")

	return c.JSON(fiber.Map{"status": "success", "message": "Retrieving All Records", "data": orders})

}
