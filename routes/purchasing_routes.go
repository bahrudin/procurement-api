package routes

import (
	"procurement-api/controllers"
	"procurement-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PurchasingRoutes(router fiber.Router) {

	purchases := router.Group("/purchases", middlewares.JWTProtected())

	purchases.Post("/", controllers.CreatePurchase)
	purchases.Get("/", controllers.GetPurchases)
	purchases.Get("/:id", controllers.GetPurchaseByID)
}
