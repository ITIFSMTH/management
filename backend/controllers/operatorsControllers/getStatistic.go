package operatorsController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns operator statistic
func (*OperatorsController) GetOperatorStatistic(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get worker ID
	id, _ := c.Get("id")

	// Get operator
	var operator models.Operator
	db.Model(&models.Operator{}).
		Joins("JOIN workers on workers.id=operators.worker_id").
		Preload("Worker.Role").
		Preload("Shifts").
		Preload("Shifts.Timeouts").
		First(&operator, "workers.id = ?", id) // Select *

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"statistic": responses.CreateMonthsOperatorsStatisticResponse(&[]models.Operator{operator})},
	})
}
