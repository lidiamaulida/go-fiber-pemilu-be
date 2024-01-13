package paslonscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/models"
	"pemiluApi-go/database"
)

func Create(c *fiber.Ctx) error {
	var paslon models.Paslon
	if err := c.BodyParser(&paslon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&paslon).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"message": "succes create paslon",
		"data":    paslon,
	})
}

func Getall(c *fiber.Ctx) error {
	var paslons []models.Paslon
	err := database.DB.Preload("Partais").Find(&paslons).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get paslons"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get all paslons",
	     "data":   paslons,
	})
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var paslon models.Paslon
	err := database.DB.Preload("Partais").First(&paslon, id).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get paslon"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get paslon",
		"data":    paslon,
	})
}