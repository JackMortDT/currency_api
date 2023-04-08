package controller

import (
	"currency_api/services"
	"encoding/json"
	"net/http"
)

func GetCurrencyRates(w http.ResponseWriter, r *http.Request) {
	rates := services.RequestCurrencyRates()
	jsonBytes, err := json.Marshal(rates)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
