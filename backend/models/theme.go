package models

type Theme struct {
	ID    uint   `gorm:"unique;not null"`
	Theme string `gorm:"not null"`
}
