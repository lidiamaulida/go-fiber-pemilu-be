package voterscontrollers

import (
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/models"
	"pemiluApi-go/database"
)

func Getall(c *fiber.Ctx) error {
	var voters []models.Voter
	err := database.DB.Preload("User").Preload("Paslon").Find(&voters).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get voters"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get all voters",
	     "data":   voters,
	})
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var voter models.Voter
	err := database.DB.Preload("User").Preload("Paslon").First(&voter, id).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get voter"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get voter",
		"data":    voter,
	})
}

func Create(c *fiber.Ctx) error {
	var voter models.Voter
	if err := c.BodyParser(&voter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&voter).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"message": "succes create voter",
		"data":    voter,
	})
}