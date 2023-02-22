package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateWorkerRequestBody struct {
	WorkerID uint   `json:"worker_id" binding:"required"`
	Login    string `json:"login" binding:"omitempty,min=3,max=20"`
	RoleID   uint   `json:"role_id" binding:"omitempty,min=2,max=3"`
}

// Update worker login
func (*AdminController) UpdateWorker(c *gin.Context) {
	// Parse request body
	var requestBody UpdateWorkerRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Update worker
	if res := db.Model(&models.Worker{}).Where(requestBody.WorkerID).Updates(&models.Worker{
		Login:  requestBody.Login,
		RoleID: requestBody.RoleID,
	}); res.Error != nil || res.RowsAffected == 0 {
		// If updating failed, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Send response
	c.Status(http.StatusNoContent)
}
