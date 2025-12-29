package controllers

import (
	"os"
	"strings"
	"time"

	"procurement-api/config"
	"procurement-api/models"
	"procurement-api/requests"
	"procurement-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// =========================
// REGISTER
// =========================
func Register(c *fiber.Ctx) error {
	var req requests.RegisterRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	// Validasi reusable
	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	// Cek duplicate username
	var exist models.User
	if err := config.DB.Where("username = ?", req.Username).First(&exist).Error; err == nil {
		return utils.BadRequest(c, "Username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.ServerError(c, "Failed to hash password")
	}

	// Buat user baru
	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return utils.ServerError(c, "Failed to create user")
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

func Register2(c *fiber.Ctx) error {
	var req requests.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed hash password",
		})
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

// =========================
// LOGIN
// =========================
func Login(c *fiber.Ctx) error {
	var req requests.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return utils.Unauthorized(c, "Invalid credentials")
	}

	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return utils.Unauthorized(c, "Invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

// =========================
// LOGOUT
// =========================
func Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.BadRequest(c, "Missing Authorization header")
	}

	tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return utils.Unauthorized(c, "Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["exp"] == nil {
		return utils.Unauthorized(c, "Invalid token claims")
	}

	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	blacklist := models.TokenBlacklist{
		Token:     tokenStr,
		ExpiredAt: exp,
	}

	config.DB.Create(&blacklist)

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}
