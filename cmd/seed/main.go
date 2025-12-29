package main

import (
	"fmt"
	"log"
	"os"

	"procurement-api/config"
	"procurement-api/database/seeders"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
)

func main() {

	// ===============================
	//  ENV
	// ===============================
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Gagal load .env:", err)
	}

	if os.Getenv("APP_ENV") == "production" {
		log.Fatal("Seeder tidak boleh dijalankan di production")
	}

	// ===============================
	// CONNECT DATABASE
	// ===============================
	config.ConnectDB()

	// ===============================
	// INIT FAKER
	// ===============================
	gofakeit.Seed(0)

	fmt.Println("===================================")
	fmt.Println("START DATABASE SEEDER")
	fmt.Println("===================================")

	runSeeder("users", func() error {
		return seeders.SeedUsers(config.DB, 5)
	})

	runSeeder("suppliers", func() error {
		return seeders.SeedSuppliers(config.DB, 10)
	})

	runSeeder("items", func() error {
		return seeders.SeedItems(config.DB, 20)
	})

	runSeeder("purchasing", func() error {
		return seeders.SeedPurchasing(config.DB, 15)
	})

	runSeeder("purchasing_details", func() error {
		return seeders.SeedPurchasingDetails(config.DB)
	})

	runSeeder("stock_movements", func() error {
		return seeders.SeedStockMovementsFromPurchasing(config.DB)
	})

	fmt.Println("===================================")
	fmt.Println("ALL SEEDERS COMPLETED SUCCESSFULLY")
	fmt.Println("===================================")
}

// ===============================
// HELPER RUN SEEDER
// ===============================
func runSeeder(name string, fn func() error) {

	fmt.Printf("Seeding %-25s : ", name)

	if err := fn(); err != nil {
		fmt.Println("FAILED")
		log.Fatalf("Seeder [%s] error: %v", name, err)
	}

	fmt.Println("SUCCESS")
}
