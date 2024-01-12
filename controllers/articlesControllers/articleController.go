package articlescontrollers

import (
	"github.com/gofiber/fiber/v2"
	"pemiluApi-go/models"
	"pemiluApi-go/database"
)

func Getall(c *fiber.Ctx) error {
	var articles []models.Article
	err := database.DB.Find(&articles).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get articles"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get all article",
		"data":    articles,
	})
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	err := database.DB.First(&article, id).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get articles"})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "succes get article",
		"data":    article,
	})
}
