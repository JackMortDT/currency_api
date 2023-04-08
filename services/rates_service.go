package services

import (
	"currency_api/command"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func RequestCurrencyRates() *command.Rates {
	currencyUrl := os.Getenv("CURRENCY_URL")
	apiKey := os.Getenv("API_KEY")

	resp, err := http.Get(currencyUrl + apiKey)
	if err != nil {
		log.Print(err)
	}

	var rates command.Rates
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		log.Print(err)
	}

	return &rates
}
