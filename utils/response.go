package utils

import "github.com/gofiber/fiber/v2"

// =========================
// SUCCESS RESPONSES
// =========================

// OK - 200
func OK(c *fiber.Ctx, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"data": data,
	})
}

// Created - 201
func Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(201).JSON(fiber.Map{
		"data": data,
	})
}

// =========================
// ERROR RESPONSES
// =========================

// BadRequest - 400
func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(400).JSON(fiber.Map{
		"error": msg,
	})
}

// Unauthorized - 401
func Unauthorized(c *fiber.Ctx, msg string) error {
	return c.Status(401).JSON(fiber.Map{
		"error": msg,
	})
}

// Forbidden - 403
func Forbidden(c *fiber.Ctx, msg string) error {
	return c.Status(403).JSON(fiber.Map{
		"error": msg,
	})
}

// NotFound - 404
func NotFound(c *fiber.Ctx, msg string) error {
	return c.Status(404).JSON(fiber.Map{
		"error": msg,
	})
}

// ValidationError - 422
func ValidationError(c *fiber.Ctx, err error) error {
	return c.Status(422).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// ServerError - 500
func ServerError(c *fiber.Ctx, msg string) error {
	// Production: jangan kirim detail error, cukup msg umum
	return c.Status(500).JSON(fiber.Map{
		"error": msg,
	})
}
