package controllers

import (
	"alisafdarirepo/database"
	"alisafdarirepo/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {

	var roles []models.Role

	result := database.DB.Preload("Permission").Find(&roles)
	err := result.Error

	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "",
			"data": err,
		})
	}

	fmt.Printf("Retreving All Records")

	return c.JSON(fiber.Map{"status": "success", "message": "Retrieving All Records", "data": roles})

}

func CreateRole(ctx *fiber.Ctx) error {

	var roleDto fiber.Map
	//role := new(models.Role)

	if err := ctx.BodyParser(&roleDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}
	role := models.Role{
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}
	database.DB.Create(&role)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created Role", "data": role})

}

func GetRoleById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.DB.Preload("Permission").Find(&role)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "find Role by Id", "data": role})
}

func UpdateRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	var roleDto fiber.Map
	if err := ctx.BodyParser(&roleDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)
	role := models.Role{
		Id:         uint(id),
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Update Role by Id", "data": role})

}

func DeleteRole(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Deleting Role by id", "data": nil})
}
