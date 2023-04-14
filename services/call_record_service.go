package services

import (
	"currency_api/domain"
	"currency_api/utils/error_utils"
	"time"
)

var (
	CallRecordService callRecordServiceInterface = &callRecordService{}
)

type callRecordService struct{}

type callRecordServiceInterface interface {
	SaveCallRecord(time.Time, int64) error_utils.MessageErr
	SaveFailCallRecord(time.Time) error_utils.MessageErr
}

func (cS *callRecordService) SaveCallRecord(requestDate time.Time, duration int64) error_utils.MessageErr {
	err := domain.CallRecordRepo.Create(&domain.CallRecord{
		RequestDate: requestDate,
		Duration:    duration,
		Success:     true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cS *callRecordService) SaveFailCallRecord(requestDate time.Time) error_utils.MessageErr {
	err := domain.CallRecordRepo.Create(&domain.CallRecord{
		RequestDate: requestDate,
		Success:     false,
	})
	if err != nil {
		return err
	}
	return nil

}
