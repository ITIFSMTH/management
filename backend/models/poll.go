package models

import (
	"time"

	"gorm.io/gorm"
)

type Poll struct {
	gorm.Model
	StartDate  time.Time `gorm:"not null"`
	PollTypeID uint      `gorm:"not null"`
	PollType   PollType  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
