package domain

import "time"

type CurrencyRate struct {
	ID        uint      `gorm:"primaryKey"`
	Code      string    `gorm:"uniqueIndex:idx_code_updated_at"`
	Value     float64   `gorm:"value"`
	CreatedAt time.Time `gorm:"createdAt"`
	UpdatedAt time.Time `gorm:"uniqueIndex:idx_code_updated_at"`
}
