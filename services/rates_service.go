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

func GetCurrencyRates(currency string, finit, fend time.Time) ([]domain.CurrencyRate, error_utils.MessageErr) {
	currencies, err := domain.CurrencyRepo.GetByCurrencyAndBetweenDates(currency, finit, fend)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}

func RequestCurrencyRates() (*command.Rate, error_utils.MessageErr) {
	start := time.Now()
	resp, requestErr := client.ApiRequest()
	if requestErr != nil {
		saveFailCallRecord(start)
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

	saveCurrencyResponse(&rates)
	saveCallRecord(start, duration)

	return &rates, nil
}

func saveCurrencyResponse(rates *command.Rate) {
	currencyRates := command.ConvertToCurrencyRates(rates)
	for _, currencyRate := range currencyRates {
		domain.CurrencyRepo.CreateOrUpdate(currencyRate)
	}
}

func saveCallRecord(requestDate time.Time, duration int64) {
	domain.CallRecordRepo.Create(&domain.CallRecord{
		RequestDate: requestDate,
		Duration:    duration,
		Sucess:      true,
	})
}

func saveFailCallRecord(requestDate time.Time) {
	domain.CallRecordRepo.Create(&domain.CallRecord{
		RequestDate: requestDate,
		Sucess:      false,
	})

}
