package app

import (
	"currency_api/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=currency_dev sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	db.AutoMigrate(&domain.CurrencyRate{})
}
