package seeders

import (
	"fmt"

	"procurement-api/models"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func SeedSuppliers(db *gorm.DB, total int) error {
	var suppliers []models.Supplier

	for i := 0; i < total; i++ {
		suppliers = append(suppliers, models.Supplier{
			Name:    fmt.Sprintf("PT %s", gofakeit.Company()),
			Email:   fmt.Sprintf("info%d@%s.com", i, gofakeit.DomainName()),
			Address: gofakeit.Address().Address,
			Phone:   gofakeit.Phone(),
		})
	}

	return db.Create(&suppliers).Error
}
