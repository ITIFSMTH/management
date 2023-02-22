package models

import "gorm.io/gorm"

type BudgetVote struct {
	gorm.Model
	Budget      uint           `gorm:"not null"`
	VoterID     uint           `gorm:"not null"`
	CandidateID uint           `gorm:"not null"`
	PollID      uint           `gorm:"not null"`
	Voter       SeniorOperator `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Candidate   JuniorOperator `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Poll        Poll           `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
