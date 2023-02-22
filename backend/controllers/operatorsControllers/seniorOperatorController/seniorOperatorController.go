package seniorOperatorController

import (
	"errors"
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
)

type SeniorOperatorConroller struct{}

// Check is votes contains all unique operators && check is all budget wasted
func validateBudgetVotes(votes *[]BudgetVote, pollBudget uint) error {
	// Get DB
	db := db.GetDB()

	// Get all senior operators and check is vote exist for every operator
	var juniorOperators []models.JuniorOperator
	if res := db.Model(&models.JuniorOperator{}).Find(&juniorOperators); res.Error != nil || res.RowsAffected == 0 {
		return errors.New(responses.ErrorServer)
	}

	// Check is length same
	if len(*votes) != len(juniorOperators) {
		return errors.New(responses.ErrorBadData)
	}

	// Wasted budget variable
	var wastedBudget uint
	// For every senior operator from DB
	for _, vote := range *votes {
		// Add budget for operator to wastedBudget
		wastedBudget += vote.Budget
		// Check is operators contains vote
		if !operatorsContainsVote(&juniorOperators, &vote) {
			return errors.New(responses.ErrorBadData)
		}
	}

	// Check is all budget wasted
	if wastedBudget != pollBudget {
		return errors.New(responses.ErrorBadData)
	}

	return nil
}

// Check is operators contains vote
func operatorsContainsVote(operators *[]models.JuniorOperator, vote *BudgetVote) bool {
	for _, operator := range *operators {
		if operator.ID == vote.CandidateID {
			return true
		}
	}

	return false
}
