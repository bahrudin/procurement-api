package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Purchasing struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Date       time.Time `json:"date"`
	SupplierID uint      `gorm:"index" json:"supplier_id"`
	UserID     uint      `gorm:"index" json:"user_id"`
	//GrandTotal float64   `json:"grand_total"`
	GrandTotal decimal.Decimal `gorm:"type:decimal(15,2);default:0" json:"grand_total"`

	Supplier Supplier           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"supplier"`
	User     User               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user"`
	Details  []PurchasingDetail `gorm:"constraint:OnDelete:CASCADE;" json:"details"`

	//Status string `gorm:"index" json:"status"` // DRAFT, APPROVED, RECEIVED

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
