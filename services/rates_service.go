package services

import (
	"currency_api/command"
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RequestCurrencyRates() (*command.Rate, error_utils.MessageErr) {
	currencyUrl := os.Getenv("CURRENCY_URL")
	apiKey := os.Getenv("API_KEY")
	stringTimeout := os.Getenv("TIMEOUT")
	timeout, err := strconv.Atoi(stringTimeout)
	if err != nil {
		errorMessage := fmt.Sprintf("Error converting timeout from env")
		log.Print(errorMessage)

		return nil, error_utils.NewInternalServerError(errorMessage)
	}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	start := time.Now()
	resp, err := client.Get(currencyUrl + apiKey)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		errorMessage := fmt.Sprintf("Error requesting currency API: %s", err)
		log.Print(errorMessage)
		saveFailCallRecord(start)

		return nil, error_utils.NewServiceUnavailableError(errorMessage)
	}
	defer resp.Body.Close()

	var rates command.Rate
	err = json.NewDecoder(resp.Body).Decode(&rates)
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
