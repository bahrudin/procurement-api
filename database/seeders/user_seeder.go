package seeders

import (
	"fmt"

	"procurement-api/models"
	"procurement-api/utils"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB, total int) error {

	// Pastikan ada 1 admin default
	hashedAdminPassword, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	admin := models.User{
		Username: "admin",
		Password: hashedAdminPassword,
		Role:     "ADMIN",
	}
	db.FirstOrCreate(&admin, models.User{Username: "admin"})

	var users []models.User

	for i := 0; i < total; i++ {

		hashedPassword, err := utils.HashPassword("password123")
		if err != nil {
			return err
		}

		users = append(users, models.User{
			Username: fmt.Sprintf("user_%d_%s", i, gofakeit.Username()),
			Password: hashedPassword,
			Role:     gofakeit.RandomString([]string{"ADMIN", "STAFF"}),
		})
	}

	return db.Create(&users).Error
}
