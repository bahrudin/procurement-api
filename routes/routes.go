package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// API v1
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Resource
	AuthRoutes(v1)
	ItemRoutes(v1)
	SupplierRoutes(v1)
	PurchasingRoutes(v1)
}
