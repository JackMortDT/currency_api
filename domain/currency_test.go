package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInitialize(t *testing.T) {
	dbResult, db, err := fakeDatabaseConnection()
	assert.NoError(t, err)
	assert.EqualValues(t, dbResult, db)
}

func TestCreateOrUpdate(t *testing.T) {
	db, _, err := fakeDatabaseConnection()
	assert.NoError(t, err)

	err = db.AutoMigrate(&CurrencyRate{})
	assert.NoError(t, err)

	// Test with one register on DB
	currency := &CurrencyRate{
		Code:      "MXN",
		Value:     18.10,
		UpdatedAt: time.Now(),
	}

	err = CurrencyRepo.CreateOrUpdate(currency)
	assert.NoError(t, err)

	var result CurrencyRate
	db.First(&result, "code = ?", "MXN")
	assert.EqualValues(t, result.Code, currency.Code)
	assert.EqualValues(t, result.Value, currency.Value)

	// Test adding other register on DB
	currency = &CurrencyRate{
		Code:      "USD",
		Value:     1,
		UpdatedAt: time.Now(),
	}

	err = CurrencyRepo.CreateOrUpdate(currency)
	assert.NoError(t, err)

	var usdResult CurrencyRate
	db.First(&usdResult, "code = ?", "USD")
	assert.EqualValues(t, usdResult.Code, currency.Code)
	assert.EqualValues(t, usdResult.Value, currency.Value)

	// Test for all registers
	currencies, err := CurrencyRepo.GetAll()
	assert.NoError(t, err)
	assert.EqualValues(t, len(currencies), 2)

	db.Exec("DELETE FROM currency_rates")

}

func fakeDatabaseConnection() (*gorm.DB, *gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=currency_test sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		return nil, nil, err
	}
	dbResult := CurrencyRepo.Initialize(db)
	return dbResult, db, nil
}
