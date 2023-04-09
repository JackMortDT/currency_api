package services

import (
	"currency_api/command"
	"currency_api/domain"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func RequestCurrencyRates() *command.Rate {
	currencyUrl := os.Getenv("CURRENCY_URL")
	apiKey := os.Getenv("API_KEY")

	resp, err := http.Get(currencyUrl + apiKey)
	if err != nil {
		log.Print(err)
	}

	var rates command.Rate
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		log.Print(err)
	}

	saveCurrencyResponse(&rates)

	return &rates
}

func saveCurrencyResponse(rates *command.Rate) {
	currencyRates := command.ConvertToCurrencyRates(rates)
	for _, currencyRate := range currencyRates {
		domain.CurrencyRepo.CreateOrUpdate(currencyRate)
	}
}
