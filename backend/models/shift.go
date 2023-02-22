package models

import (
	"time"

	"gorm.io/gorm"
)

type Shift struct {
	gorm.Model
	StartDate     time.Time `gorm:"not null"`
	EndDate       time.Time
	LastNotify    time.Time
	NextNotify    time.Time
	CaptchaAnswer string
	Delays        uint8
	Timeouts      []Timeout
	OperatorID    uint
}
