package command

import (
	"currency_api/domain"
	"log"
	"time"
)

type Rate struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	LastUpdated string `json:"last_updated_at"`
}

type Data map[string]struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

func ConvertToCurrencyRates(rate *Rate) []*domain.CurrencyRate {
	log.Println("Hasta aqui llego D:")
	log.Println(rate)
	var currencyRates []*domain.CurrencyRate

	lastUpdated, err := time.Parse(time.RFC3339, rate.Meta.LastUpdated)
	if err != nil {
		lastUpdated = time.Now()
	}

	for _, data := range rate.Data {
		currencyRate := &domain.CurrencyRate{
			Code:      data.Code,
			Value:     data.Value,
			CreatedAt: time.Now(),
			UpdatedAt: lastUpdated,
		}

		currencyRates = append(currencyRates, currencyRate)
	}

	return currencyRates
}
