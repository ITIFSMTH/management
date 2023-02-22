package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns settings
func (*AdminController) GetSettings(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get settings
	var settings models.Setting
	db.First(&settings)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"settings": responses.CreateSettingsResponse(&settings)},
	})
}
