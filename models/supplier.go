package models

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"type:varchar(100);index;not null" json:"name"`
	Email   string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Phone   string `gorm:"type:varchar(30)" json:"phone"`
	Address string `gorm:"type:text" json:"address"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
