package models

import "github.com/shopspring/decimal"

type PurchasingDetail struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	PurchasingID uint `gorm:"index" json:"purchasing_id"`
	ItemID       uint `gorm:"index" json:"item_id"`

	Qty int `json:"qty"`

	Price    decimal.Decimal `gorm:"type:decimal(18,2)" json:"price"`
	SubTotal decimal.Decimal `gorm:"type:decimal(18,2)" json:"sub_total"`

	Purchasing Purchasing `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Item       Item       `gorm:"foreignKey:ItemID" json:"item"`
}
