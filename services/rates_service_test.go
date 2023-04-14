package services

import (
	"currency_api/client"
	"currency_api/command"
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	apiRequest                   func() (*http.Response, error_utils.MessageErr)
)

type repoMock struct{}
type clientMock struct{}

func (m *clientMock) ApiRequest() (*http.Response, error_utils.MessageErr) {
	return apiRequest()
}

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

func TestRequestCurrencyRates(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	createOrUpdate = func(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
		return nil
	}
	client.CurrencyClient = &clientMock{}
	apiRequest = func() (*http.Response, error_utils.MessageErr) {
		rate := command.Rate{
			Meta: command.Meta{
				LastUpdated: "2023-04-13T23:59:59Z",
			},
			Data: command.Data{
				"USD": {
					Code:  "USD",
					Value: 1,
				},
				"MXN": {
					Code:  "MXN",
					Value: 18.17,
				},
			},
		}
		body, _ := json.Marshal(rate)

		response := httptest.NewRecorder()
		response.WriteHeader(http.StatusOK)
		response.Write(body)
		return response.Result(), nil
	}

	_, err := RatesService.RequestCurrencyRates()
	assert.NoError(t, err)
}

func TestRequestCurrencyRates_WithErrorFromApiRequest(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	createOrUpdate = func(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
		return nil
	}
	client.CurrencyClient = &clientMock{}
	apiRequest = func() (*http.Response, error_utils.MessageErr) {
		return nil, error_utils.NewInternalServerError("Error on currency API")
	}

	_, err := RatesService.RequestCurrencyRates()
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message(), "Error on currency API")
}

func TestRequestCurrencyRates_WithErrorOnResponseFromApiRequest(t *testing.T) {
	domain.CurrencyRepo = &repoMock{}
	createOrUpdate = func(currencyRate *domain.CurrencyRate) error_utils.MessageErr {
		return nil
	}
	client.CurrencyClient = &clientMock{}
	apiRequest = func() (*http.Response, error_utils.MessageErr) {
		response := httptest.NewRecorder()
		response.WriteHeader(http.StatusInternalServerError)
		return response.Result(), nil
	}

	_, err := RatesService.RequestCurrencyRates()
	assert.NotNil(t, err)
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

	err := saveCurrencyResponse(rate)
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

	err := saveCurrencyResponse(rate)
	assert.EqualValues(t, err.Message(), "Fail to insert into database")

}
