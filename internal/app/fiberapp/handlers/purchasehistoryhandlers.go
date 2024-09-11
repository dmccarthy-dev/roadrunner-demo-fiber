package handlers

import (
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type PurchaseHistoryHandlers struct {
}

func NewPurchaseHistoryHandlers() *PurchaseHistoryHandlers {
	return &PurchaseHistoryHandlers{}
}

func (h *PurchaseHistoryHandlers) HandleGetIndividual() fiber.Handler {
	return func(c *fiber.Ctx) error {

		time.Sleep(5 * time.Second)

		id := c.Params("id", "not found")

		return c.Status(fiber.StatusOK).JSON(models.PurchaseHistory{
			Id: id,
			History: []models.PurchaseHistoryEntry{
				{
					Amount: 100,
					Item:   "Laptop",
				},
				{
					Amount: 20,
					Item:   "Mouse",
				},
			},
		})
	}
}
