package controller

import (
	"currency_api/services"
	"currency_api/utils"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

func GetCurrencyRates(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	currency := parts[2]
	currencyRegex := regexp.MustCompile(`\d`)
	hasNumber := currencyRegex.MatchString(currency)

	if len(parts) != 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	if currency == "" || len(currency) != 3 || hasNumber {
		http.Error(w, "Invalid currency", http.StatusBadRequest)
		return
	}
	currency = strings.ToUpper(currency)

	finitStr := r.URL.Query().Get("finit")
	finit, parseError := utils.ParseStringToDate(finitStr, "finit")
	if parseError != nil {
		http.Error(w, parseError.Error(), parseError.Status())
		return
	}

	fendStr := r.URL.Query().Get("fend")
	fend, parseError := utils.ParseStringToDate(fendStr, "finit")
	if parseError != nil {
		http.Error(w, parseError.Error(), parseError.Status())
		return
	}

	rates, currencyErr := services.GetCurrencyRates(currency, finit, fend)
	if currencyErr != nil {
		http.Error(w, currencyErr.Error(), currencyErr.Status())
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

func CreateNewCurrencies(w http.ResponseWriter, r *http.Request) {
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
