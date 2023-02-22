package models

type WorkerRole struct {
	ID   uint   `gorm:"unique;not null"`
	Role string `gorm:"not null"`
}
