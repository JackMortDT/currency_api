package domain

import (
	"time"

	"gorm.io/gorm"
)

var (
	CurrencyRepo currencyRepoInterface = &currencyRepo{}
)

type currencyRepoInterface interface {
	Initialize(*gorm.DB) *gorm.DB
	GetAll() ([]CurrencyRate, error)
	CreateOrUpdate(*CurrencyRate) error
	GetByCurrencyAndBetweenDates(string, time.Time, time.Time) []CurrencyRate
}

type currencyRepo struct {
	db *gorm.DB
}

func (cR *currencyRepo) Initialize(db *gorm.DB) *gorm.DB {
	cR.db = db

	return cR.db
}

func (cR *currencyRepo) GetAll() ([]CurrencyRate, error) {
	var currencies []CurrencyRate
	result := cR.db.Find(&currencies)
	if result.Error != nil {
		return nil, result.Error
	}
	return currencies, nil
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

func (cR *currencyRepo) GetByCurrencyAndBetweenDates(currency string, finit, fend time.Time) []CurrencyRate {
	var currencies []CurrencyRate
	query := cR.db

	if currency != "ALL" {
		query = query.Where(CurrencyRate{Code: currency})
	}

	if !finit.IsZero() && fend.IsZero() {
		query = query.Where("updated_at >= ?", finit)
	} else if finit.IsZero() && !fend.IsZero() {
		query = query.Where("updated_at <= ?", fend)
	} else if !finit.IsZero() && !fend.IsZero() {
		query = query.Where("updated_at BETWEEN ? AND ?", finit, fend)
	}

	result := query.Find(&currencies)
	if result.Error != nil {
		//
	}
	return currencies
}
