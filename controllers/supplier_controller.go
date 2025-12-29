package controllers

import (
	"strconv"

	"procurement-api/config"
	"procurement-api/models"
	"procurement-api/requests"
	"procurement-api/utils"

	"github.com/gofiber/fiber/v2"
)

// =========================
// GET ALL SUPPLIERS
// =========================
func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier

	if err := config.DB.Find(&suppliers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed fetch suppliers",
		})
	}

	return c.JSON(suppliers)
}

// =========================
// GET SUPPLIER BY ID
// =========================
func GetSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "Invalid supplier ID")
	}

	var supplier models.Supplier
	if err := config.DB.First(&supplier, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Supplier not found",
		})
	}

	return c.JSON(supplier)
}

// =========================
// CREATE SUPPLIER
// =========================
func CreateSupplier(c *fiber.Ctx) error {
	var req requests.SupplierCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	supplier := models.Supplier{
		Name: req.Name,
	}

	if err := config.DB.Create(&supplier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed create supplier",
		})
	}

	return c.JSON(supplier)
}

// =========================
// UPDATE SUPPLIER
// =========================
func UpdateSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "Invalid supplier ID")
	}

	var supplier models.Supplier
	if err := config.DB.First(&supplier, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Supplier not found",
		})
	}

	var req requests.SupplierUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	// Update field (partial update)
	if req.Name != "" {
		supplier.Name = req.Name
	}

	if err := config.DB.Save(&supplier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed update supplier",
		})
	}

	return c.JSON(supplier)
}

// =========================
// DELETE SUPPLIER
// =========================
func DeleteSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "Invalid supplier ID")
	}

	if err := config.DB.Delete(&models.Supplier{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed delete supplier",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier deleted",
	})
}
