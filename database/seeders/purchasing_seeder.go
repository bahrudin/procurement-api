package seeders

import (
	"time"

	"procurement-api/models"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func SeedPurchasing(db *gorm.DB, total int) error {
	var suppliers []models.Supplier
	var items []models.Item

	db.Find(&suppliers)
	db.Find(&items)

	for i := 0; i < total; i++ {

		supplier := suppliers[gofakeit.Number(0, len(suppliers)-1)]

		purchase := models.Purchasing{
			SupplierID: supplier.ID,
			UserID:     1,
			Date:       time.Now(),
			GrandTotal: decimal.Zero,
		}

		if err := db.Create(&purchase).Error; err != nil {
			return err
		}

		grandTotal := decimal.Zero
		detailCount := gofakeit.Number(1, 4)

		for j := 0; j < detailCount; j++ {

			item := items[gofakeit.Number(0, len(items)-1)]
			qty := gofakeit.Number(1, 5)

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

			// Update stock
			db.Model(&item).
				UpdateColumn("stock", gorm.Expr("stock + ?", qty))

			grandTotal = grandTotal.Add(subTotal)
		}

		// Update Grand Total
		db.Model(&purchase).
			Update("grand_total", grandTotal)
	}

	return nil
}
