package seeders

import (
	"procurement-api/models"

	"gorm.io/gorm"
)

// SeedStockMovementsFromPurchasing
// Membuat stock movement IN berdasarkan purchasing detail
func SeedStockMovementsFromPurchasing(db *gorm.DB) error {

	var details []models.PurchasingDetail

	// Ambil semua purchasing detail + item
	if err := db.Preload("Item").Find(&details).Error; err != nil {
		return err
	}

	for _, detail := range details {

		// 1️⃣ Buat stock movement IN
		movement := models.StockMovement{
			ItemID:  detail.ItemID,
			Qty:     detail.Qty,
			Type:    models.StockIn,
			RefType: "PURCHASE",
			RefID:   detail.PurchasingID,
			Note:    "Auto stock in dari purchasing",
		}

		if err := db.Create(&movement).Error; err != nil {
			return err
		}

		// 2️⃣ Update stock item (aggregate)
		if err := db.Model(&models.Item{}).
			Where("id = ?", detail.ItemID).
			Update("stock", gorm.Expr("stock + ?", detail.Qty)).Error; err != nil {
			return err
		}
	}

	return nil
}
