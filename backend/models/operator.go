package models

import "gorm.io/gorm"

type Operator struct {
	gorm.Model
	Telegram   string `gorm:"unique;not null"`
	TelegramID int64
	WorkerID   uint   `gorm:"unique;not null"`
	Worker     Worker `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Shifts     []Shift
}
