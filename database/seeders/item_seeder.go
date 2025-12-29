package seeders

import (
	"errors"
	"fmt"
	"strings"

	"procurement-api/models"
	"procurement-api/utils"

	"github.com/shopspring/decimal"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func SeedItems(db *gorm.DB, total int) error {
	if total <= 0 {
		return errors.New("total must be greater than 0")
	}

	units := []string{"pcs", "box", "kg", "unit"}

	for i := 0; i < total; i++ {

		tx := db.Begin()

		item := models.Item{
			SKU:   fmt.Sprintf("SKU-%d-%s", i+1, strings.ToUpper(gofakeit.LetterN(4))),
			Name:  gofakeit.ProductName(),
			Stock: 0, // stok aktual diisi via movement
			//Price: float64(gofakeit.Number(50_000, 15_000_000)),
			Price: decimal.NewFromInt(
				int64(gofakeit.Number(50_000, 15_000_000)),
			),
			Unit: gofakeit.RandomString(units),
		}

		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}

		// ===== INITIAL STOCK (OPSIONAL) =====
		initialStock := gofakeit.Number(0, 100)

		if initialStock > 0 {
			if err := utils.AdjustStock(
				tx,
				item.ID,
				initialStock,
				utils.RefTypeInitial,
				0, // initial stock tidak punya reference
				"Initial stock from seeder",
			); err != nil {
				tx.Rollback()
				return err
			}
		}

		tx.Commit()
	}

	return nil
}
