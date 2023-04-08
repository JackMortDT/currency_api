package app

import (
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
	routes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
