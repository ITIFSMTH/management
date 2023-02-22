package models

import (
	"time"

	"gorm.io/gorm"
)

type Timeout struct {
	gorm.Model
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:""`
	ShiftID   uint
}
