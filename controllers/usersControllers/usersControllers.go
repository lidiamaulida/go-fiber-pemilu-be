package userscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/models"
	"pemiluApi-go/database"
)

func Getall(c *fiber.Ctx) error {
	var users []models.User
	err := database.DB.Find(&users).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get users"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get all users",
	     "data":   users,
	})
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	err := database.DB.First(&user, id).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get user",
		"data":    user,
	})
}

func Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"message": "succes create user",
		"data":    user,
	})
}