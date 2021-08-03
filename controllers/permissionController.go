package controllers

import (
	"alisafdarirepo/database"
	"alisafdarirepo/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(ctx *fiber.Ctx) error {

	var permission []models.Permission

	database.DB.Find(&permission)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Get All permissions", "data": permission})
}
