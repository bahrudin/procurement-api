package controllers

import (
	"strconv"

	"procurement-api/config"
	"procurement-api/models"
	"procurement-api/requests"
	"procurement-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

// =========================
// GET ALL
// =========================
func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)
	return c.JSON(items)
}

// =========================
// GET SINGLE ITEM
// =========================
func GetItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "Invalid item ID")
	}

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return utils.NotFound(c, "Item not found")
	}

	return c.JSON(item)
}

// =========================
// CREATE
// =========================
func CreateItem(c *fiber.Ctx) error {
	var req requests.ItemCreateRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	// Convert price: string -> decimal
	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return utils.BadRequest(c, "Invalid price format")
	}

	item := models.Item{
		Name:  req.Name,
		Price: price,
		Stock: req.Stock,
	}

	config.DB.Create(&item)
	return c.JSON(item)
}

// =========================
// UPDATE
// =========================
func UpdateItem(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Item not found"})
	}

	var req requests.ItemUpdateRequest
	c.BodyParser(&req)

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	if req.Name != "" {
		item.Name = req.Name
	}
	//if req.Price > 0 {
	//	item.Price = req.Price
	//}
	
	if req.Price != nil {
		price, err := decimal.NewFromString(*req.Price)
		if err != nil {
			return utils.BadRequest(c, "Invalid price format")
		}
		item.Price = price
	}

	if req.Stock >= 0 {
		item.Stock = req.Stock
	}

	config.DB.Save(&item)
	return c.JSON(item)
}

// =========================
// DELETE
// =========================
func DeleteItem(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	config.DB.Delete(&models.Item{}, id)
	return c.JSON(fiber.Map{"message": "Item deleted"})
}
