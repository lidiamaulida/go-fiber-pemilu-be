package partaiscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/models"
	"pemiluApi-go/database"
)

func Create(c *fiber.Ctx) error {
	var partai models.Partai
	if err := c.BodyParser(&partai); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&partai).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"message": "succes create partai",
		"data":    partai,
	})
}

func Getall(c *fiber.Ctx) error {
	var partai []models.Partai
	err := database.DB.Find(&partai).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get partai"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get all partai",
	     "data":   partai,
	})
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var partai models.Partai
	err := database.DB.First(&partai, id).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get partai"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get partai",
		"data":    partai,
	})
}