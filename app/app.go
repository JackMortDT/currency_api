package app

import (
	"currency_api/routines"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var apiKey string
var frequency time.Duration

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

func StartApp() {
	db := connectToDatabase()
	startRepositories(db)
	routes()
	routines.ExecuteApiRequest()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
