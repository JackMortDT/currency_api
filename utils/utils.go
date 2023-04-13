package utils

import (
	"currency_api/utils/error_utils"
	"fmt"
	"time"
)

func ParseStringToDate(date string, fname string) (time.Time, error_utils.MessageErr) {
	var fdate time.Time
	var err error

	if date != "" {
		fdate, err = time.Parse("2006-01-02T15:04:05", date)
		if err != nil {
			error := fmt.Sprintf("Invalid %s format", fname)
			return fdate, error_utils.NewBadRequestError(error)
		}
		return fdate, nil
	}
	return fdate, nil
}
