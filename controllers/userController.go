package controllers

import (
	"alisafdarirepo/database"
	"alisafdarirepo/models"
	"alisafdarirepo/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	var users []models.User

	result := database.DB.Preload("Role").Find(&users)
	err := result.Error

	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "",
			"data": err,
		})
	}

	fmt.Printf("Retreving All Records")

	return c.JSON(fiber.Map{"status": "success", "message": "Retrieving All Records", "data": users})

}

func CreateUser(ctx *fiber.Ctx) error {

	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hash, _ := util.HashPassword(user.Password)

	user.Password = hash
	database.DB.Create(&user)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})

}

func GetUserById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}
	database.DB.Preload("Role").Find(&user)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "find User by username", "data": user})
}

func UpdateUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	database.DB.Model(&user).Updates(user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Update user by Id", "data": user})

}

func DeleteUser(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	user := models.User{
		ID: uint(id),
	}

	database.DB.Delete(&user)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Deleting user by id", "data": nil})
}
