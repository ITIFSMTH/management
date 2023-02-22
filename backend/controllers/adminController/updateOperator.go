package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateOperatorRequestBody struct {
	OperatorID uint   `json:"operator_id" binding:"required"`
	Telegram   string `json:"telegram" binding:"required"`
}

// Update operator telegram
func (*AdminController) UpdateOperator(c *gin.Context) {
	// Parse request body
	var requestBody UpdateOperatorRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Update operator telegram
	if res := db.Model(&models.Operator{}).Where(requestBody.OperatorID).Select("Telegram", "TelegramID").Updates(&models.Operator{
		Telegram:   requestBody.Telegram,
		TelegramID: 0,
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
