package juniorOperatorController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RatingVote struct {
	Honest      uint8 `json:"honest" binding:"required,min=1,max=10"`
	Involvement uint8 `json:"involvement" binding:"required,min=1,max=10"`
	Help        uint8 `json:"help" binding:"required,min=1,max=10"`
	Activity    uint8 `json:"activity" binding:"required,min=1,max=10"`
	CandidateID uint  `json:"candidate_id" binding:"required"`
}

type RatingVotesRequestBody struct {
	Votes  []RatingVote `json:"votes"`
	PollID uint         `json:"poll_id" binding:"required"`
}

func (*JuniorOperatorConroller) AddRatingVotes(c *gin.Context) {
	// Parse request body
	var requestBody RatingVotesRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Validate votes
	if err := validateRatingVotes(&requestBody.Votes); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: err.Error(),
		})
		return
	}

	// Get poll
	var ratingPoll models.Poll
	if db.Model(&ratingPoll).Last(&ratingPoll); ratingPoll.StartDate.AddDate(0, 0, 1).Unix() < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	// Get voter worker ID
	id, _ := c.Get("id")

	// Get voter
	var voter models.JuniorOperator
	if res := db.Model(&voter).
		Joins("JOIN operators on operators.id=junior_operators.operator_id").
		Where("operators.worker_id = ?", id.(uint)).
		Select("junior_operators.id").
		First(&voter); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	// Parse votes from request body to []models.RatinVote
	var votes []models.RatingVote
	for _, vote := range requestBody.Votes {
		votes = append(votes, models.RatingVote{
			Honest:      vote.Honest,
			Involvement: vote.Involvement,
			Help:        vote.Help,
			Activity:    vote.Activity,
			CandidateID: vote.CandidateID,
			VoterID:     voter.ID,
			PollID:      requestBody.PollID,
		})
	}

	// Add new votes
	if res := db.Model(&models.RatingVote{}).Create(&votes); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Send response
	c.Status(http.StatusNoContent)
}
