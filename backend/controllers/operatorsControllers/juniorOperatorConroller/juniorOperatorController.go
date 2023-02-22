package juniorOperatorController

import (
	"errors"
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
)

type JuniorOperatorConroller struct{}

// Check is votes contains all unique operators
func validateRatingVotes(votes *[]RatingVote) error {
	// Get DB
	db := db.GetDB()

	// Get all senior operators and check is vote exist for every operator
	var seniorOperators []models.SeniorOperator
	if res := db.Model(&models.SeniorOperator{}).Find(&seniorOperators); res.Error != nil || res.RowsAffected == 0 {
		return errors.New(responses.ErrorServer)
	}

	// Check is length same
	if len(*votes) != len(seniorOperators) {
		return errors.New(responses.ErrorBadData)
	}

	// For every senior operator from DB
	for _, operator := range seniorOperators {
		if !votesContainsOperator(votes, &operator) {
			return errors.New(responses.ErrorBadData)
		}
	}

	return nil
}

// Check is votes contains operator
func votesContainsOperator(votes *[]RatingVote, operator *models.SeniorOperator) bool {
	for _, vote := range *votes {
		if operator.ID == vote.CandidateID {
			return true
		}
	}

	return false
}
