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

		profile := models.Profile{
			Id:    id,
			Name:  "Jane Smith",
			Email: "jsmith@example.com",
		}

		if c.Query("include-history") == "true" {
			profile.PastPurchases = getPurchaseHistory(c.Hostname(), id)
		}

		return c.Status(fiber.StatusOK).JSON(profile)
	}
}

func getPurchaseHistory(hostname string, id string) models.PurchaseHistory {
	agent := fiber.Get("http://" + hostname + "/api/v2/purchase-history/" + id)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		fmt.Println(statusCode, errs)
	}
	var parsedHistory models.PurchaseHistory
	err := json.Unmarshal(body, &parsedHistory)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	return parsedHistory
}
