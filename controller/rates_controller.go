package controller

import (
	"currency_api/services"
	"encoding/json"
	"net/http"
)

func GetCurrencyRates(w http.ResponseWriter, r *http.Request) {
	rates, err := services.RequestCurrencyRates()
	if err != nil {
		http.Error(w, err.Error(), err.Status())
		return
	}

	jsonBytes, jsonErr := json.Marshal(rates)

	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
