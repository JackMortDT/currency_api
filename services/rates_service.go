package services

import (
	"currency_api/client"
	"currency_api/command"
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var RatesService ratesServiceInterface = &ratesService{}

type ratesService struct{}

type ratesServiceInterface interface {
	GetCurrencyRates(string, time.Time, time.Time) ([]domain.CurrencyRate, error_utils.MessageErr)
	RequestCurrencyRates() (*command.Rate, error_utils.MessageErr)
	saveCurrencyResponse(*command.Rate) error_utils.MessageErr
}

func (rS *ratesService) GetCurrencyRates(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
	currencies, err := domain.CurrencyRepo.GetByCurrencyAndBetweenDates(currency, finit, fend)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func (rS *ratesService) RequestCurrencyRates() (*command.Rate, error_utils.MessageErr) {
	start := time.Now()
	resp, requestErr := client.ApiRequest()
	if requestErr != nil {
		CallRecordService.SaveFailCallRecord(start)
		return nil, requestErr
	}
	duration := time.Since(start).Milliseconds()

	var rates command.Rate
	err := json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		errorMessage := fmt.Sprintf("Error decoding API response: %s", err)
		log.Print(errorMessage)

		return nil, error_utils.NewInternalServerError(errorMessage)

	}

	if saveError := RatesService.saveCurrencyResponse(&rates); saveError != nil {
		return nil, saveError
	}
	CallRecordService.SaveCallRecord(start, duration)

	return &rates, nil
}

func (rS *ratesService) saveCurrencyResponse(rates *command.Rate) error_utils.MessageErr {
	currencyRates := command.ConvertToCurrencyRates(rates)
	for _, currencyRate := range currencyRates {
		if err := domain.CurrencyRepo.CreateOrUpdate(currencyRate); err != nil {
			return err
		}

	}
	return nil
}
