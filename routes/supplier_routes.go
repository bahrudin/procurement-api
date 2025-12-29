package routes

import (
	"procurement-api/controllers"
	"procurement-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes(router fiber.Router) {

	suppliers := router.Group("/suppliers", middlewares.JWTProtected())

	suppliers.Post("/", controllers.CreateSupplier)
	suppliers.Get("/", controllers.GetSuppliers)
	suppliers.Get("/:id", controllers.GetSupplier)
	suppliers.Put("/:id", controllers.UpdateSupplier)
	suppliers.Delete("/:id", controllers.DeleteSupplier)
}
