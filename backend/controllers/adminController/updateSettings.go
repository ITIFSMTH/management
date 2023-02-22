package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateSettingsRequestBody struct {
	TelegramBotKey string `json:"telegram_bot_key" binding:"required"`
}

// Update telegram bot key
func (*AdminController) UpdateSettings(c *gin.Context) {
	// Parse request body
	var requestBody UpdateSettingsRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Update settings
	if res := db.Model(&models.Setting{}).Updates(&models.Setting{
		TelegramBotKey: requestBody.TelegramBotKey,
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
