package fiberapp

import (
	"github.com/dmccarthy-dev/roadrunner-demo-fiber/internal/app/fiberapp/handlers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Boot() *fiber.App {

	tracer.Start(tracer.WithEnv("my-server-host"))

	app := fiber.New(fiber.Config{
		Immutable:   true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(fibertrace.Middleware())
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	profileHandlers := handlers.NewProfileHandlers()
	purchaseHistoryHandlers := handlers.NewPurchaseHistoryHandlers()

	app.Get("/api/v2/profile/:id", profileHandlers.HandleGetIndividual()).
		Name("get individual profile")

	app.Get("/api/v2/purchase-history/:id", purchaseHistoryHandlers.HandleGetIndividual()).
		Name("get individual purchase history")

	return app
}
