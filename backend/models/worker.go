package models

import (
	"errors"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Login    string     `gorm:"unique;not null"`
	Password string     `gorm:"not null"`
	ThemeID  uint       `gorm:"not null"`
	RoleID   uint       `gorm:"not null"`
	Theme    Theme      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role     WorkerRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (w *Worker) BeforeUpdate(tx *gorm.DB) (err error) {
	if w.RoleID == 1 {
		return errors.New("You can't update administrator role")
	}

	return
}
