package controller

import (
	"currency_api/services"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
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
	fendStr := r.URL.Query().Get("fend")
	var finit, fend time.Time
	var err error

	if finitStr != "" {
		finit, err = time.Parse("2006-01-02T15:04:05", finitStr)
		if err != nil {
			http.Error(w, "Invalid finit format", http.StatusBadRequest)
			return
		}
	}

	if fendStr != "" {
		fend, err = time.Parse("2006-01-02T15:04:05", fendStr)
		if err != nil {
			http.Error(w, "Invalid fend format", http.StatusBadRequest)
			return
		}
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
	fmt.Fprintf(w, "Datos consultados")
}
