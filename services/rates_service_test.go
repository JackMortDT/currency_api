package services

import (
	"currency_api/command"
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	initialize                   func(*gorm.DB) *gorm.DB
	getAll                       func() ([]domain.CurrencyRate, error_utils.MessageErr)
	createOrUpdate               func(currencyRate *domain.CurrencyRate) error_utils.MessageErr
	getByCurrencyAndBetweenDates func(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr)
)

type repoMock struct{}

func (m *repoMock) Initialize(db *gorm.DB) *gorm.DB {
	return initialize(db)
}

func (m *repoMock) GetAll() ([]domain.CurrencyRate, error_utils.MessageErr) {
	return getAll()
}

func (m *repoMock) CreateOrUpdate(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
	return createOrUpdate(currencyRate)
}

func (m *repoMock) GetByCurrencyAndBetweenDates(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
	return getByCurrencyAndBetweenDates(currency, finit, fend)
}

func TestGetCurrencyRates(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	getByCurrencyAndBetweenDates = func(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
		currencies := []domain.CurrencyRate{
			{Code: "USD", Value: 1, UpdatedAt: time.Now()},
		}
		return currencies, nil
	}

	currency := "USD"
	finit := time.Now()
	fend := time.Now()

	currencies, err := RatesService.GetCurrencyRates(currency, finit, fend)
	assert.NoError(t, err)
	assert.EqualValues(t, len(currencies), 1)
}

func TestGetCurrencyRates_WithError(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	getByCurrencyAndBetweenDates = func(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
		return nil, error_utils.NewBadRequestError("Error to get currencies")
	}

	currency := "USD"
	finit := time.Now()
	fend := time.Now()

	currencies, err := RatesService.GetCurrencyRates(currency, finit, fend)
	assert.Nil(t, currencies)
	assert.EqualValues(t, err.Message(), "Error to get currencies")
}

func TestSaveCurrencyResponse(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	createOrUpdate = func(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
		return nil
	}

	rate := &command.Rate{
		Meta: command.Meta{LastUpdated: ""},
		Data: command.Data{"USD": {Code: "USD", Value: 1}},
	}

	err := RatesService.saveCurrencyResponse(rate)
	assert.NoError(t, err)
}

func TestSaveCurrencyResponse_WithError(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	createOrUpdate = func(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
		return error_utils.NewInternalServerError("Fail to insert into database")
	}

	rate := &command.Rate{
		Meta: command.Meta{LastUpdated: ""},
		Data: command.Data{"USD": {Code: "USD", Value: 1}},
	}

	err := RatesService.saveCurrencyResponse(rate)
	assert.EqualValues(t, err.Message(), "Fail to insert into database")

}
