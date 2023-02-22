package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"management-backend/shared"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddBudgetPollRequestBody struct {
	Budget uint `json:"budget" binding:"required"`
}

func (*AdminController) AddBudgetPoll(c *gin.Context) {
	// Parse request body
	var requestBody AddBudgetPollRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Check if already exist
	var budgetPollDb models.BudgetPoll
	if res := db.Model(&budgetPollDb).Preload("Poll").Last(&budgetPollDb); res.Error != nil ||
		res.RowsAffected == 0 ||
		budgetPollDb.Poll.StartDate.AddDate(0, 0, 1).Unix() > time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorAlreadyExists,
		})
		return
	}

	// Declare budget poll
	budgetPoll := models.BudgetPoll{
		Budget: requestBody.Budget,
		Poll: models.Poll{
			StartDate: time.Now(),
			PollType:  shared.PollTypeBudget,
		},
	}

	// Create new poll
	if res := db.Model(&budgetPoll).Create(&budgetPoll); res.Error != nil || res.RowsAffected == 0 {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"poll": responses.CreateBudgetPollResponse(&budgetPoll)},
	})
}
