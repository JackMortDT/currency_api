package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var currencies = []*CurrencyRate{
	{
		Code:      "USD",
		Value:     1,
		UpdatedAt: time.Now().AddDate(0, 0, -5)},
	{
		Code:      "EUR",
		Value:     0.9,
		UpdatedAt: time.Now().AddDate(0, 0, -3)},
	{
		Code:      "MXN",
		Value:     18.17,
		UpdatedAt: time.Now().AddDate(0, 0, -2)},
	{
		Code:      "JPY",
		Value:     110,
		UpdatedAt: time.Now().AddDate(0, 0, -7)},
	{
		Code:      "AUD",
		Value:     1.5,
		UpdatedAt: time.Now().AddDate(0, 0, -1)},
}

func TestCurrencyRepo(t *testing.T) {
	cr := &currencyRepo{}
	db, err := fakeDatabaseConnection(cr)
	assert.NoError(t, err)
	err = db.AutoMigrate(&CurrencyRate{})
	assert.NoError(t, err)

	for _, insert := range currencies {
		err := cr.CreateOrUpdate(insert)
		assert.NoError(t, err)
	}

	t.Run("TestInitialize", func(t *testing.T) {
		db, err := fakeDatabaseConnection(cr)
		assert.NoError(t, err)
		assert.NotNil(t, db)
	})

	t.Run("TestCreateOrUpdate", func(t *testing.T) {
		db, err := fakeDatabaseConnection(cr)
		assert.NoError(t, err)

		var result CurrencyRate
		db.First(&result, "code = ?", "MXN")
		assert.EqualValues(t, result.Code, "MXN")
		assert.EqualValues(t, result.Value, 18.17)

		var usdResult CurrencyRate
		db.First(&usdResult, "code = ?", "USD")
		assert.EqualValues(t, usdResult.Code, "USD")
		assert.EqualValues(t, usdResult.Value, 1)

		currencies, err := cr.GetAll()
		assert.NoError(t, err)
		assert.EqualValues(t, len(currencies), 5)
	})

	t.Run("TestGetByCurrencyAndBetweenDates", func(t *testing.T) {
		db, err := fakeDatabaseConnection(cr)
		assert.NoError(t, err)

		currency := "USD"
		finit := time.Now().AddDate(0, 0, -10)
		fend := time.Now()

		// Test 2: ALL currencis and finit
		result, error := cr.GetByCurrencyAndBetweenDates("ALL", finit, fend)
		assert.NoError(t, error)
		assert.NotNil(t, result)
		assert.IsType(t, []CurrencyRate{}, result)
		assert.Greater(t, len(result), 1)
		for _, rate := range result {
			assert.True(t, rate.UpdatedAt.After(finit) || rate.UpdatedAt.Equal(finit))
			assert.True(t, rate.UpdatedAt.Before(fend) || rate.UpdatedAt.Equal(fend))
		}

		// Test 2: USD currency and finit
		result, error = cr.GetByCurrencyAndBetweenDates(currency, finit, time.Time{})
		assert.NoError(t, error)
		assert.NotNil(t, result)
		assert.IsType(t, []CurrencyRate{}, result)
		assert.LessOrEqual(t, len(result), 1)
		if len(result) > 0 {
			assert.True(t, result[0].UpdatedAt.After(finit) || result[0].UpdatedAt.Equal(finit))
		}

		// Test 3: USD currency and fend
		result, error = cr.GetByCurrencyAndBetweenDates(currency, time.Time{}, fend)
		assert.NoError(t, error)
		assert.NotNil(t, result)
		assert.IsType(t, []CurrencyRate{}, result)
		assert.Greater(t, len(result), 0)
		for _, rate := range result {
			assert.True(t, rate.UpdatedAt.Before(fend) || rate.UpdatedAt.Equal(fend))
		}

		// Test 4: USD currency and between dates
		result, error = cr.GetByCurrencyAndBetweenDates(currency, finit, fend)
		assert.NoError(t, error)
		assert.NotNil(t, result)
		assert.IsType(t, []CurrencyRate{}, result)
		assert.Greater(t, len(result), 0)
		for _, rate := range result {
			assert.True(t, rate.UpdatedAt.After(finit) || rate.UpdatedAt.Equal(finit))
			assert.True(t, rate.UpdatedAt.Before(fend) || rate.UpdatedAt.Equal(fend))
		}

		db.Exec("DELETE FROM currency_rates")
	})
}

func fakeDatabaseConnection(cR *currencyRepo) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=currency_test sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		return nil, err
	}
	dbResult := cR.Initialize(db)
	return dbResult, nil
}
