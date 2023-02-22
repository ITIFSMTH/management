package models

import "gorm.io/gorm"

type RatingVote struct {
	gorm.Model
	Honest      uint8          `gorm:"not null"`
	Involvement uint8          `gorm:"not null"`
	Help        uint8          `gorm:"not null"`
	Activity    uint8          `gorm:"not null"`
	VoterID     uint           `gorm:"not null"`
	CandidateID uint           `gorm:"not null"`
	PollID      uint           `gorm:"not null"`
	Voter       JuniorOperator `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Candidate   SeniorOperator `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Poll        Poll           `gorm:"constaint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
