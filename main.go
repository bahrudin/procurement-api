package main

import (
	"log"
	"os"

	"procurement-api/config"
	"procurement-api/models"
	"procurement-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	config.ConnectDB()

	// Auto migrate
	if os.Getenv("APP_ENV") != "production" {
		log.Println("Running auto-migrate (non-production)")
		config.DB.AutoMigrate(
			&models.User{},
			&models.Supplier{},
			&models.Item{},
			&models.Purchasing{},
			&models.PurchasingDetail{},
			&models.TokenBlacklist{},
			&models.StockMovement{},
		)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Use(logger.New())

	routes.RegisterRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
