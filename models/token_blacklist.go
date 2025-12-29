package models

import "time"

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey"`
	Token     string    `gorm:"type:text;not null"`
	ExpiredAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
