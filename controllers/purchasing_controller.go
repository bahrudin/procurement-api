package controllers

import (
	"time"

	"procurement-api/config"
	"procurement-api/models"
	"procurement-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ==========================
// REQUEST STRUCT
// ==========================

type PurchaseItemRequest struct {
	ItemID uint `json:"item_id" validate:"required"`
	Qty    int  `json:"qty" validate:"required,gt=0"`
}

type CreatePurchaseRequest struct {
	SupplierID uint                  `json:"supplier_id" validate:"required"`
	Items      []PurchaseItemRequest `json:"items" validate:"required,min=1,dive"`
}

// ==========================
// CREATE PURCHASE
// ==========================

func CreatePurchase(c *fiber.Ctx) error {
	// ===== USER FROM JWT =====
	user, ok := c.Locals("user").(models.User)
	if !ok {
		return utils.Unauthorized(c, "Unauthorized")
	}

	// ===== PARSE REQUEST =====
	var req CreatePurchaseRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid request body")
	}

	if err := utils.ValidateStruct(req); err != nil {
		return utils.ValidationError(c, err)
	}

	// ===== DB TRANSACTION =====
	tx := config.DB.Begin()
	if tx.Error != nil {
		return utils.ServerError(c, "Failed to start transaction")
	}
	defer tx.Rollback()

	// ===== VALIDASI SUPPLIER =====
	var supplier models.Supplier
	if err := tx.First(&supplier, req.SupplierID).Error; err != nil {
		return utils.NotFound(c, "Supplier not found")
	}

	// ===== CREATE HEADER =====
	purchase := models.Purchasing{
		Date:       time.Now(),
		SupplierID: supplier.ID,
		UserID:     user.ID,
		GrandTotal: decimal.Zero,
	}

	if err := tx.Create(&purchase).Error; err != nil {
		return utils.ServerError(c, err.Error())
	}

	grandTotal := decimal.Zero
	itemUsed := make(map[uint]bool) // cegah item duplikat

	// ===== LOOP ITEMS =====
	for _, itemReq := range req.Items {

		// ❗ CEGH DUPLIKASI ITEM
		if itemUsed[itemReq.ItemID] {
			return utils.BadRequest(c, "Duplicate item in purchase items")
		}
		itemUsed[itemReq.ItemID] = true

		// ===== LOCK ITEM ROW =====
		var item models.Item
		//if err := tx.Clauses(gorm.Locking{Strength: "UPDATE"}).
		//	First(&item, itemReq.ItemID).Error; err != nil {
		//	return utils.NotFound(c, "Item not found")
		//}
		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&item, itemReq.ItemID).
			Error; err != nil {
			tx.Rollback()
			return utils.NotFound(c, "Item not found")
		}

		// ===== CALCULATION =====
		qty := decimal.NewFromInt(int64(itemReq.Qty))
		subTotal := item.Price.Mul(qty)
		grandTotal = grandTotal.Add(subTotal)

		// ===== CREATE DETAIL =====
		detail := models.PurchasingDetail{
			PurchasingID: purchase.ID,
			ItemID:       item.ID,
			Qty:          itemReq.Qty,
			Price:        item.Price,
			SubTotal:     subTotal,
		}

		if err := tx.Create(&detail).Error; err != nil {
			return utils.ServerError(c, err.Error())
		}

		// ===== STOCK MOVEMENT (IN) =====
		stock := models.StockMovement{
			ItemID:    item.ID,
			Qty:       itemReq.Qty,
			Type:      utils.StockTypeIn,
			RefType:   utils.RefTypePurchase,
			RefID:     purchase.ID,
			CreatedBy: user.ID,
		}

		if err := tx.Create(&stock).Error; err != nil {
			return utils.ServerError(c, err.Error())
		}

		// ===== UPDATE ITEM STOCK =====
		if err := tx.Model(&item).
			Update("stock", gorm.Expr("stock + ?", itemReq.Qty)).
			Error; err != nil {
			return utils.ServerError(c, err.Error())
		}
	}

	// ===== UPDATE GRAND TOTAL =====
	if err := tx.Model(&purchase).
		Update("grand_total", grandTotal).Error; err != nil {
		return utils.ServerError(c, err.Error())
	}

	// ===== COMMIT =====
	if err := tx.Commit().Error; err != nil {
		return utils.ServerError(c, "Failed to commit transaction")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":     "Purchase created successfully",
		"id":          purchase.ID,
		"grand_total": grandTotal,
	})
}

func GetPurchases(c *fiber.Ctx) error {
	var purchases []models.Purchasing

	if err := config.DB.
		Preload("Supplier").
		Preload("User").
		Find(&purchases).Error; err != nil {
		return utils.ServerError(c, err.Error())
	}

	return c.JSON(purchases)
}

func GetPurchaseByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var purchase models.Purchasing
	if err := config.DB.
		Preload("Supplier").
		Preload("User").
		Preload("Details.Item").
		First(&purchase, id).Error; err != nil {
		return utils.NotFound(c, "Purchase not found")
	}

	return c.JSON(purchase)
}
