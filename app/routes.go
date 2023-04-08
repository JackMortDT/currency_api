package app

import (
	"currency_api/controller"
	"net/http"
)

func routes() {
	http.HandleFunc("/currencies", controller.GetCurrencyRates)
}
