package command_test

import (
	"currency_api/command"
	"currency_api/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var lastUpdated = "2023-04-08T23:59:59Z"

type convertToCurrencyRatesCase struct {
	name string
	rate *command.Rate
	want []*domain.CurrencyRate
}

var convertToCurrencyRatesCases = []convertToCurrencyRatesCase{
	{
		"Transform one result successfully",
		&command.Rate{
			Meta: command.Meta{LastUpdated: lastUpdated},
			Data: command.Data{"MXN": {Code: "MXN", Value: 18.10}},
		},
		[]*domain.CurrencyRate{
			createFakeCurrencyRates("MXN", 18.10, lastUpdated),
		},
	},
	{
		"Transform more than one result successfully",
		&command.Rate{
			Meta: command.Meta{LastUpdated: lastUpdated},
			Data: command.Data{
				"MXN": {Code: "MXN", Value: 18.10},
				"USD": {Code: "USD", Value: 1},
			},
		},
		[]*domain.CurrencyRate{
			createFakeCurrencyRates("MXN", 18.10, lastUpdated),
			createFakeCurrencyRates("USD", 1, lastUpdated),
		},
	},

	{
		"With incorrect date",
		&command.Rate{
			Meta: command.Meta{LastUpdated: "incorrect date"},
			Data: command.Data{"MXN": {Code: "MXN", Value: 18.10}},
		},
		[]*domain.CurrencyRate{
			createFakeCurrencyRates("MXN", 18.10, lastUpdated),
		},
	},
}

func TestConvertToCurrencyRates(t *testing.T) {
	for _, test := range convertToCurrencyRatesCases {
		t.Run(test.name, func(t *testing.T) {
			result := command.ConvertToCurrencyRates(test.rate)
			for index, got := range result {
				want := test.want[index]
				assert.EqualValues(t, got.Code, want.Code)
				assert.EqualValues(t, got.Value, want.Value)
			}
		})
	}
}

func createFakeCurrencyRates(currency string, rate float64, dateStr string) *domain.CurrencyRate {
	date, _ := time.Parse(time.RFC3339, dateStr)

	return &domain.CurrencyRate{
		Code:      currency,
		Value:     rate,
		UpdatedAt: date,
	}
}
