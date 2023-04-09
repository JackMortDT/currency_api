package domain

import (
	"gorm.io/gorm"
)

var (
	CurrencyRepo currencyRepoInterface = &currencyRepo{}
)

type currencyRepoInterface interface {
	Initialize(*gorm.DB) *gorm.DB
	CreateOrUpdate(*CurrencyRate) error
}

type currencyRepo struct {
	db *gorm.DB
}

func (cR *currencyRepo) Initialize(db *gorm.DB) *gorm.DB {
	cR.db = db

	return cR.db
}

func (cR *currencyRepo) CreateOrUpdate(currency *CurrencyRate) error {
	result := cR.db.
		Where(CurrencyRate{Code: currency.Code, UpdatedAt: currency.UpdatedAt}).
		Assign(currency).
		FirstOrCreate(currency)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
