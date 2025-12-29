package routes

import (
	"procurement-api/controllers"
	"procurement-api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {

	auth := router.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", middlewares.JWTProtected(), controllers.Logout)
}
