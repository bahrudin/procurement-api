package utils

import (
	"errors"
	"fmt"
	"procurement-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ==========================
// ENUM / KONSTAN STOCK TYPE
// ==========================
const (
	StockTypeIn  = "IN"
	StockTypeOut = "OUT"

	RefTypePurchase   = "PURCHASE"
	RefTypeSale       = "SALE"
	RefTypeAdjustment = "ADJUSTMENT"
	RefTypeOpname     = "OPNAME"
	RefTypeInitial    = "INITIAL"
)

// ==========================
// HELPER ADJUST STOCK
// ==========================

// AdjustStock mengatur stok item
// qty > 0 : tambah stok
// qty < 0 : kurangi stok
func AdjustStock(tx *gorm.DB, itemID uint, qty int, refType string, refID uint, note string) error {
	if qty == 0 {
		return errors.New("quantity cannot be zero")
	}

	// ===== Ambil item & lock row =====
	var item models.Item
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, itemID).Error; err != nil {
		return fmt.Errorf("item not found: %w", err)
	}

	newStock := item.Stock + qty
	if newStock < 0 {
		return fmt.Errorf("stock cannot be negative (item %d)", itemID)
	}

	// ===== Update stok langsung di DB =====
	if err := tx.Model(&models.Item{}).
		Where("id = ?", item.ID).
		UpdateColumn("stock", gorm.Expr("stock + ?", qty)).Error; err != nil {
		return err
	}

	// ===== Buat StockMovement =====
	stockMovement := models.StockMovement{
		ItemID:  item.ID,
		Qty:     qty,
		Type:    StockTypeIn,
		RefType: refType,
		RefID:   refID,
		Note:    note,
	}

	if qty < 0 {
		stockMovement.Type = StockTypeOut
		stockMovement.Qty = -qty
	}

	if err := tx.Create(&stockMovement).Error; err != nil {
		return err
	}

	return nil
}
