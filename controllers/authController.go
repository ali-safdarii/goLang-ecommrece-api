package controllers

import (
	"alisafdarirepo/database"
	"alisafdarirepo/models"
	"alisafdarirepo/util"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

//------------------------

// Register /*----------------------------------------------------------*/

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message":
		"Review your input", "data": err})
	}

	hash, err := util.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error",
			"message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message":
		"Review your input", "data": err})
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message":
		"user with this email not found", "data": nil})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {

		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "incorrect password"})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "something wet wrong"})
	}

	return c.JSON(fiber.Map{"status": "success", "message":
	"Login Successfully", "data": fiber.Map{
		"user":  user,
		"token": token,
	}})

}

func User(c *fiber.Ctx) error {

	var users []models.User

	result := database.DB.Find(&users)
	err := result.Error

	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "",
			"data": err,
		})
	}

	fmt.Printf("Retreving All Records")

	return c.JSON(fiber.Map{"status": "success", "message": "Retrieving All Records", "data": users})

}
