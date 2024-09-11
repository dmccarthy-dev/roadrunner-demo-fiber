package handlers

import (
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp/models"
	"github.com/gofiber/fiber/v2"
)

type ProfileHandlers struct {
}

func NewProfileHandlers() *ProfileHandlers {
	return &ProfileHandlers{}
}

func (h *ProfileHandlers) HandleGetIndividual() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id", "not found")

		return c.Status(fiber.StatusOK).JSON(models.Profile{
			Id:    id,
			Name:  "Jane Smith",
			Email: "jsmith@example.com",
		})
	}
}
