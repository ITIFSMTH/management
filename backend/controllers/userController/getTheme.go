package userController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*UserController) GetTheme(c *gin.Context) {
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

	// Get worker theme
	var worker models.Worker
	db.Model(&models.Worker{}).
		Preload("Theme").
		Find(&worker, id)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"theme": responses.CreateThemeResponse(&worker.Theme)},
	})
}
