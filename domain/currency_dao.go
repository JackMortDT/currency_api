package domain

import (
	"currency_api/utils/error_utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	CurrencyRepo currencyRepoInterface = &currencyRepo{}
)

type currencyRepoInterface interface {
	Initialize(*gorm.DB) *gorm.DB
	GetAll() ([]CurrencyRate, error_utils.MessageErr)
	CreateOrUpdate(*CurrencyRate) error_utils.MessageErr
	GetByCurrencyAndBetweenDates(string, time.Time, time.Time) ([]CurrencyRate, error_utils.MessageErr)
}

type currencyRepo struct {
	db *gorm.DB
}

func (cR *currencyRepo) Initialize(db *gorm.DB) *gorm.DB {
	cR.db = db

	return cR.db
}

func (cR *currencyRepo) GetAll() ([]CurrencyRate, error_utils.MessageErr) {
	var currencies []CurrencyRate
	result := cR.db.Find(&currencies)
	if result.Error != nil {
		error := fmt.Sprintf("Error when trying to get records: %s", result.Error)
		return nil, error_utils.NewInternalServerError(error)
	}
	return currencies, nil
}

func (cR *currencyRepo) CreateOrUpdate(currency *CurrencyRate) error_utils.MessageErr {
	result := cR.db.
		Where(CurrencyRate{Code: currency.Code, UpdatedAt: currency.UpdatedAt}).
		Assign(currency).
		FirstOrCreate(currency)

	if result.Error != nil {
		error := fmt.Sprintf("Error when trying to save or update record record: %s", result.Error)
		return error_utils.NewInternalServerError(error)
	}

	return nil
}

func (cR *currencyRepo) GetByCurrencyAndBetweenDates(currency string, finit, fend time.Time) ([]CurrencyRate, error_utils.MessageErr) {
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
		error := fmt.Sprintf("Error when trying to get records: %s", result.Error)
		return nil, error_utils.NewInternalServerError(error)
	}
	return currencies, nil
}
