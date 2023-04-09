package app

import (
	"currency_api/domain"

	"gorm.io/gorm"
)

func startRepositories(db *gorm.DB) {
	domain.CurrencyRepo.Initialize(db)
	domain.CallRecordRepo.Initialize(db)
}
