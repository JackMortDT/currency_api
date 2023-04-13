package client

import (
	"currency_api/utils/error_utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ApiRequest() (*http.Response, error_utils.MessageErr) {
	currencyUrl := os.Getenv("CURRENCY_URL")
	apiKey := os.Getenv("API_KEY")
	stringTimeout := os.Getenv("TIMEOUT")

	timeout, err := strconv.Atoi(stringTimeout)
	if err != nil {
		errorMessage := fmt.Sprintf("Error converting timeout from env")
		log.Print(errorMessage)

		return nil, error_utils.NewInternalServerError(errorMessage)
	}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}

	resp, err := client.Get(currencyUrl + apiKey)

	if err != nil {
		errorMessage := fmt.Sprintf("Error requesting currency API: %s", err)
		log.Print(errorMessage)

		return nil, error_utils.NewServiceUnavailableError(errorMessage)
	}
	return resp, nil
}
