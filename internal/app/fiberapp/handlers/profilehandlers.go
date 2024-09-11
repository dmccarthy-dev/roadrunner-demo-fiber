package handlers

import (
	"fmt"
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp/models"
	"github.com/goccy/go-json"
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

		agent := fiber.Get("http://" + c.Hostname() + "/api/v2/purchase-history/" + id)
		statusCode, body, errs := agent.Bytes()
		if len(errs) > 0 {
			fmt.Println(statusCode, errs)
		}
		var parsedHistory models.PurchaseHistory
		err := json.Unmarshal(body, &parsedHistory)

		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
		}

		return c.Status(fiber.StatusOK).JSON(models.Profile{
			Id:            id,
			Name:          "Jane Smith",
			Email:         "jsmith@example.com",
			PastPurchases: parsedHistory,
		})
	}
}
