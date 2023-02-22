package seniorOperatorController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BudgetVote struct {
	Budget      uint `json:"budget" binding:"exists"`
	CandidateID uint `json:"candidate_id" binding:"required"`
}

type BudgetVotesRequestBody struct {
	Votes  []BudgetVote `json:"votes"`
	PollID uint         `json:"poll_id" binding:"required"`
}

func (*SeniorOperatorConroller) AddRatingVotes(c *gin.Context) {
	// Parse request body
	var requestBody BudgetVotesRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Get poll
	var budgtePoll models.BudgetPoll
	if db.Model(&budgtePoll).Preload("Poll").Last(&budgtePoll); budgtePoll.Poll.StartDate.AddDate(0, 0, 1).Unix() < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	// Validate votes
	if err := validateBudgetVotes(&requestBody.Votes, budgtePoll.Budget); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: err.Error(),
		})
		return
	}

	// Get voter worker ID
	id, _ := c.Get("id")

	// Get voter
	var voter models.SeniorOperator
	if res := db.Model(&voter).
		Joins("JOIN operators on operators.id=senior_operators.operator_id").
		Where("operators.worker_id = ?", id.(uint)).
		Select("senior_operator.id").
		First(&voter); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	// Parse votes from request body to []models.RatinVote
	var votes []models.BudgetVote
	for _, vote := range requestBody.Votes {
		votes = append(votes, models.BudgetVote{
			Budget:      vote.Budget,
			CandidateID: vote.CandidateID,
			VoterID:     voter.ID,
			PollID:      requestBody.PollID,
		})
	}

	// Add new votes
	if res := db.Model(&models.BudgetVote{}).Create(&votes); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Send response
	c.Status(http.StatusNoContent)
}
