package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns operators statistic
func (*AdminController) GetOperatorsStatistic(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get all operators
	var operators []models.Operator
	db.Model(&models.Operator{}).
		Preload("Worker").
		Preload("Worker.Role"). // Sort by role
		Preload("Shifts").
		Preload("Shifts.Timeouts").
		Find(&operators)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"statistic": responses.CreateMonthsOperatorsStatisticResponse(&operators)},
	})
}
