package controller_test

import (
	"currency_api/command"
	"currency_api/controller"
	"currency_api/domain"
	"currency_api/services"
	"currency_api/utils/error_utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	getCurrencyRates     func(string, time.Time, time.Time) ([]domain.CurrencyRate, error_utils.MessageErr)
	requestCurrencyRates func() (*command.Rate, error_utils.MessageErr)
)

type currencyServiceMock struct{}

func (cS *currencyServiceMock) GetCurrencyRates(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
	return getCurrencyRates(currency, finit, fend)
}

func (cS *currencyServiceMock) RequestCurrencyRates() (*command.Rate, error_utils.MessageErr) {
	return requestCurrencyRates()
}

func TestGetCurrency(t *testing.T) {
	services.RatesService = &currencyServiceMock{}
	getCurrencyRates = func(string, time.Time, time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
		return nil, nil
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/currencies/MXN", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusOK)
}

func TestGetCurrency_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/currencies/MXN/USD", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusBadRequest)
}

func TestGetCurrency_InvalidCurrency(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/currencies/MX1", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusBadRequest)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/currencies/MX", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusBadRequest)
}

func TestGetCurrency_InvalidDates(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/currencies/MXN?finit=Invalid", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusBadRequest)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/currencies/MXN?fend=Invalid", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusBadRequest)
}

func TestGetCurrency_ServiceError(t *testing.T) {
	services.RatesService = &currencyServiceMock{}
	getCurrencyRates = func(string, time.Time, time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
		return nil, error_utils.NewNotFoundError("Currency not found")
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/currencies/MXN", nil)

	controller.GetCurrencyRates(w, r)
	assert.EqualValues(t, w.Code, http.StatusNotFound)
}
