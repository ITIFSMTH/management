package juniorOperatorController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (*JuniorOperatorConroller) GetRatingPoll(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get last rating poll
	var ratingPoll models.Poll
	db.Model(&ratingPoll).Preload("PollType").Last(&ratingPoll)

	// Check is poll ended (1 day after start)
	if ratingPoll.StartDate.AddDate(0, 0, 1).Unix() < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"poll": responses.CreateRatingPollResponse(&ratingPoll)},
	})
}
