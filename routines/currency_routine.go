package routines

import (
	"currency_api/services"
	"log"
	"os"
	"strconv"
	"time"
)

func ExecuteApiRequest() {
	log.Println("Inicio de proceso de la goroutine")
	intervalRequest := os.Getenv("INTERVAL_REQUEST")
	interval, err := strconv.Atoi(intervalRequest)
	if err != nil {
		panic("Interval is not set correctly on environment variables")
	}
	go func() {
		for {
			_, err := services.RatesService.RequestCurrencyRates()
			if err != nil {
				log.Println("Ocurrio un error a la hora de ejecutar currency rates")
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
}
