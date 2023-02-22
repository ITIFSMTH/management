package seniorOperatorController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (*SeniorOperatorConroller) GetBudgetPoll(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get last budget poll
	var budgetPoll models.BudgetPoll
	db.Model(&budgetPoll).Preload("Poll").Preload("Poll.PollType").Last(&budgetPoll)

	// Check is poll ended (1 day after start)
	if budgetPoll.Poll.StartDate.AddDate(0, 0, 1).Unix() < time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Error: responses.ErrorNotExists,
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"poll": responses.CreateBudgetPollResponse(&budgetPoll)},
	})
}
