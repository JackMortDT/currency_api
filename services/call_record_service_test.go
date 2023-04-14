package services

import (
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	create func(callRecord *domain.CallRecord) error_utils.MessageErr
)

type callRepoMock struct{}

func (m *callRepoMock) Create(callRecord *domain.CallRecord) error_utils.MessageErr {
	return create(callRecord)
}

func (m *callRepoMock) Initialize(db *gorm.DB) *gorm.DB {
	return initialize(db)
}

func TestSaveCallRecord(t *testing.T) {
	domain.CallRecordRepo = &callRepoMock{}
	create = func(callRecord *domain.CallRecord) error_utils.MessageErr {
		return nil
	}
	requestDate := time.Now()
	duration := 2.7

	err := CallRecordService.SaveCallRecord(requestDate, int64(duration))
	assert.NoError(t, err)
}

func TestSaveCallRecord_WithError(t *testing.T) {
	domain.CallRecordRepo = &callRepoMock{}
	create = func(callRecord *domain.CallRecord) error_utils.MessageErr {
		return error_utils.NewInternalServerError("Error on creation on call record")
	}
	requestDate := time.Now()
	duration := 2.7

	err := CallRecordService.SaveCallRecord(requestDate, int64(duration))
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message(), "Error on creation on call record")
}

func TestSaveFailRecord(t *testing.T) {
	domain.CallRecordRepo = &callRepoMock{}
	create = func(callRecord *domain.CallRecord) error_utils.MessageErr {
		return nil
	}
	requestDate := time.Now()

	err := CallRecordService.SaveFailCallRecord(requestDate)
	assert.NoError(t, err)
}

func TestSaveFailRecord_WithError(t *testing.T) {
	domain.CallRecordRepo = &callRepoMock{}
	create = func(callRecord *domain.CallRecord) error_utils.MessageErr {
		return error_utils.NewInternalServerError("Error on creation on call record")
	}
	requestDate := time.Now()

	err := CallRecordService.SaveFailCallRecord(requestDate)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message(), "Error on creation on call record")
}
