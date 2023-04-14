package routines

import (
	"currency_api/services"
	"log"
	"time"
)

func ExecuteApiRequest() {
	log.Println("Inicio de proceso de la goroutine")
	go func() {
		for {
			_, err := services.RatesService.RequestCurrencyRates()
			if err != nil {
				log.Println("Ocurrio un error a la hora de ejecutar currency rates")
			}
			time.Sleep(10 * time.Second)
		}
	}()
}
