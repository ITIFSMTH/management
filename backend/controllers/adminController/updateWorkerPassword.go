package adminController

import (
	"management-backend/controllers/authController"
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateWorkerPasswordRequestBody struct {
	WorkerID uint   `json:"worker_id" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=80"`
}

// Update worker password
func (*AdminController) UpdateWorkerPassword(c *gin.Context) {
	// Parse request body
	var requestBody UpdateWorkerPasswordRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Create a password hash
	hash, err := authController.GetPasswordHash(requestBody.Password)
	if err != nil {
		// If creating hash errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Update a worker password
	if res := db.Model(&models.Worker{}).Where(requestBody.WorkerID).Updates(&models.Worker{
		Password: hash,
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
