package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var callRecords = []*CallRecord{
	{
		RequestDate: time.Now(),
		Duration:    10,
		Success:     true,
	},
	{
		RequestDate: time.Now(),
		Duration:    60,
		Success:     true,
	},
	{
		RequestDate: time.Now(),
		Duration:    3600,
		Success:     true,
	},
	{
		RequestDate: time.Now(),
		Duration:    5,
		Success:     false,
	},
	{
		RequestDate: time.Now(),
		Duration:    30,
		Success:     false,
	},
}

func TestCallRecordRepo(t *testing.T) {
	crr := &callRecordRepo{}
	db, err := fakeCallRecordConnection(crr)
	assert.NoError(t, err)
	err = db.AutoMigrate(&CallRecord{})
	assert.NoError(t, err)

	t.Run("TestCreate", func(t *testing.T) {
		_, err := fakeCallRecordConnection(crr)
		assert.NoError(t, err)

		for _, callRecord := range callRecords {
			insertErr := crr.Create(callRecord)
			assert.NoError(t, insertErr)

		}
	})

}

func fakeCallRecordConnection(crr *callRecordRepo) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=currency_test sslmode=disable"
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		return nil, err
	}
	dbResult := crr.Initialize(db)
	return dbResult, nil

}
