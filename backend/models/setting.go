package models

type Setting struct {
	ID             uint   `gorm:"unique;not null"`
	TelegramBotKey string `gorm:"not null"`
}
