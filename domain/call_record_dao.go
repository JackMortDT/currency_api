package domain

import (
	"currency_api/utils/error_utils"
	"fmt"

	"gorm.io/gorm"
)

var (
	CallRecordRepo callRecordInterface = &callRecordRepo{}
)

type callRecordInterface interface {
	Initialize(*gorm.DB) *gorm.DB
	Create(*CallRecord) error_utils.MessageErr
}

type callRecordRepo struct {
	db *gorm.DB
}

func (crR *callRecordRepo) Initialize(db *gorm.DB) *gorm.DB {
	crR.db = db

	return crR.db
}

func (crR *callRecordRepo) Create(callRecord *CallRecord) error_utils.MessageErr {
	result := crR.db.Create(callRecord)
	if result.Error != nil {
		error := fmt.Sprintf("Error when trying to save record: %s", result.Error)
		return error_utils.NewInternalServerError(error)
	}

	return nil
}
