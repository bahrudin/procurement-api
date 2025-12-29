package routes

import (
	"procurement-api/controllers"
	"procurement-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ItemRoutes(router fiber.Router) {

	items := router.Group("/items")

	items.Get("/", controllers.GetItems)
	items.Get("/:id", controllers.GetItem)

	items.Post("/", middlewares.JWTProtected(), controllers.CreateItem)
	items.Put("/:id", middlewares.JWTProtected(), controllers.UpdateItem)
	items.Delete("/:id", middlewares.JWTProtected(), controllers.DeleteItem)
}
