package models

type PollType struct {
	ID   uint   `gorm:"unique;not null"`
	Poll string `gorm:"unique;not null"`
}
