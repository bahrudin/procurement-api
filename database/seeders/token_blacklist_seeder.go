package seeders

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"

	"procurement-api/models"
)

// SeedTokenBlacklist
// Membuat data token blacklist (JWT revoked)
func SeedTokenBlacklist(db *gorm.DB, total int) error {

	gofakeit.Seed(0)

	for i := 0; i < total; i++ {

		// Simulasi JWT token (bukan valid sign, hanya struktur)
		token := gofakeit.UUID() + "." + gofakeit.UUID() + "." + gofakeit.UUID()

		// Expired 1–7 hari ke depan
		expiredAt := time.Now().Add(time.Duration(gofakeit.Number(1, 7)) * 24 * time.Hour)

		blacklist := models.TokenBlacklist{
			Token:     token,
			ExpiredAt: expiredAt,
			CreatedAt: time.Now(),
		}

		if err := db.Create(&blacklist).Error; err != nil {
			return err
		}
	}

	return nil
}
