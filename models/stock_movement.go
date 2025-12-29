package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	StockIn  = "IN"
	StockOut = "OUT"
)

type StockMovement struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ItemID uint `gorm:"not null;index" json:"item_id"`
	Item   Item `gorm:"foreignKey:ItemID" json:"-"`

	Qty  int    `gorm:"not null" json:"qty"`
	Type string `gorm:"size:10;not null" json:"type"` // IN | OUT

	RefType string `gorm:"size:50;not null" json:"ref_type"`
	RefID   uint   `gorm:"index" json:"ref_id"`

	Note string `gorm:"type:text" json:"note,omitempty"`

	CreatedBy uint           `gorm:"index" json:"created_by"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
