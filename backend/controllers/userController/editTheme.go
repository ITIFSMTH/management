package userController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EditThemeRequestBody struct {
	ThemeID uint `json:"theme_id" binding:"required"`
}

func (*UserController) EditTheme(c *gin.Context) {
	// Get request data
	var requestBody EditThemeRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get worker ID
	id, exist := c.Get("id")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Update worker theme
	if res := db.Model(&models.Worker{}).Where(id).Updates(&models.Worker{
		ThemeID: requestBody.ThemeID,
	}); res.Error != nil || res.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get theme
	var theme models.Theme
	db.Model(&theme).Find(&theme, requestBody.ThemeID)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"theme": responses.CreateThemeResponse(&theme)},
	})
}
