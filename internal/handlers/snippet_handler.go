package handlers

import (
	"github.com/LeanAlbu/goFiberWebsite/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var snippets = []models.Snippet{}

func GetSnippet(c *fiber.Ctx) error{
	return c.Status(fiber.StatusOK).JSON(snippets)
}


func CreateSnippet(c *fiber.Ctx) error{
	snippet := new(models.Snippet)

	if err := c.BodyParser(snippet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao processar o JSON",
		})
	}
	snippet.ID = uuid.NewString()
	snippets = append(snippets, *snippet)
	return c.Status(fiber.StatusCreated).JSON(snippet)
}

