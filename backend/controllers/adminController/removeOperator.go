package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"management-backend/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RemoveOperatorRequestBody struct {
	WorkerID uint `json:"worker_id" binding:"required"`
}

// Create a new operator
func (*AdminController) RemoveOperator(c *gin.Context) {
	// Parse request body
	var requestBody RemoveOperatorRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Get worker
	var worker models.Worker
	if res := db.Model(&worker).First(&worker, requestBody.WorkerID); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get operator
	var operator models.Operator
	if res := db.Model(&operator).Where("worker_id = ?", worker.ID).First(&operator); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	if worker.RoleID == shared.WorkerRoleJuniorOperator.ID {
		db.Unscoped().Where("operator_id = ?", operator.ID).Delete(&models.JuniorOperator{})
	} else if worker.RoleID == shared.WorkerRoleSeniorOperator.ID {
		db.Unscoped().Where("operator_id = ?", operator.ID).Delete(&models.SeniorOperator{})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Remove operator
	db.Unscoped().Delete(&operator)
	db.Unscoped().Delete(&worker)

	// Send response
	c.Status(http.StatusNoContent)
}
