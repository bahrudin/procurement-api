package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Item struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	SKU   string `gorm:"type:varchar(50);uniqueIndex;not null" json:"sku"`
	Name  string `gorm:"type:varchar(150);not null" json:"name"`
	Stock int    `gorm:"default:0;not null" json:"stock"`
	//Price float64 `gorm:"not null" json:"price"`
	Price decimal.Decimal `gorm:"type:decimal(15,2);not null" json:"price"`
	Unit  string          `gorm:"type:varchar(20)" json:"unit"`

	//Unit  string          `json:"unit"`

	// relasi history stock
	StockMovements []StockMovement `json:"-"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
