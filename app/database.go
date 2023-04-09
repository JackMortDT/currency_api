package app

import (
	"currency_api/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDatabase() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&domain.CurrencyRate{})
	db.AutoMigrate(&domain.CallRecord{})
}
