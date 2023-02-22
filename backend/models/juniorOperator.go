package models

import "gorm.io/gorm"

type JuniorOperator struct {
	gorm.Model
	OperatorID uint     `gorm:"not null"`
	Operator   Operator `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
