package models

import "gorm.io/gorm"

type SeniorOperator struct {
	gorm.Model
	Honest      float32  `gorm:"type:REAL"`
	Involvement float32  `gorm:"type:REAL"`
	Help        float32  `gorm:"type:REAL"`
	Activity    float32  `gorm:"type:REAL"`
	OperatorID  uint     `gorm:"not null"`
	Operator    Operator `gorm:"constraint:OnUpdate:CASCADE;OnDelete:SET NULL;"`
}
