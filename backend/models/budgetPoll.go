package models

import "gorm.io/gorm"

type BudgetPoll struct {
	gorm.Model
	Budget uint `gorm:"not null"`
	PollID uint `gorm:"not null"`
	Poll   Poll `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
