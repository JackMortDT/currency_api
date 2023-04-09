package services

import (
	"currency_api/command"
	"currency_api/domain"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func RequestCurrencyRates() *command.Rate {
	currencyUrl := os.Getenv("CURRENCY_URL")
	apiKey := os.Getenv("API_KEY")

	start := time.Now()
	resp, err := http.Get(currencyUrl + apiKey)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	var rates command.Rate
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		log.Print(err)
	}

	saveCurrencyResponse(&rates)
	saveCallRecord(start, duration)

	return &rates
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
	})
}
