package middlewares

import (
	"os"
	"procurement-api/config"
	"procurement-api/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
		}

		tokenStr := strings.Replace(auth, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"message": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			return c.Status(401).JSON(fiber.Map{"message": "Invalid token claims"})
		}

		userID := uint(claims["user_id"].(float64))

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"message": "User not found"})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
