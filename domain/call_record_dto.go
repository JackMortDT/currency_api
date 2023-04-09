package domain

import (
	"time"

	"gorm.io/gorm"
)

type CallRecord struct {
	gorm.Model
	RequestDate time.Time
	Duration    int64
}
