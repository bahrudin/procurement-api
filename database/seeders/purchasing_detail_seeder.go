package seeders

import (
	"procurement-api/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// SeedPurchasingDetails membuat detail pembelian berdasarkan purchasing & item
func SeedPurchasingDetails(db *gorm.DB) error {

	var purchases []models.Purchasing
	var items []models.Item

	if err := db.Find(&purchases).Error; err != nil {
		return err
	}

	if err := db.Find(&items).Error; err != nil {
		return err
	}

	for _, purchase := range purchases {

		grandTotal := decimal.Zero
		detailCount := gofakeit.Number(1, 5)

		for i := 0; i < detailCount; i++ {

			item := items[gofakeit.Number(0, len(items)-1)]
			qty := gofakeit.Number(1, 10)

			price := item.Price // decimal.Decimal

			subTotal := decimal.
				NewFromInt(int64(qty)).
				Mul(price)

			detail := models.PurchasingDetail{
				PurchasingID: purchase.ID,
				ItemID:       item.ID,
				Qty:          qty,
				Price:        price,
				SubTotal:     subTotal,
			}

			if err := db.Create(&detail).Error; err != nil {
				return err
			}

			// Update stock IN
			if err := db.Model(&models.Item{}).
				Where("id = ?", item.ID).
				UpdateColumn("stock", gorm.Expr("stock + ?", qty)).Error; err != nil {
				return err
			}

			grandTotal = grandTotal.Add(subTotal)
		}

		// Update GrandTotal
		if err := db.Model(&models.Purchasing{}).
			Where("id = ?", purchase.ID).
			Update("grand_total", grandTotal).Error; err != nil {
			return err
		}
	}

	return nil
}
