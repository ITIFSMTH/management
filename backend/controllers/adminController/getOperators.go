package adminController

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns all operators
func (*AdminController) GetOperators(c *gin.Context) {
	// Get DB
	db := db.GetDB()

	// Get all operators
	var operators []models.Operator
	db.Model(&models.Operator{}).
		Preload("Worker").
		Preload("Worker.Role").
		Find(&operators)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"operators": responses.CreateOperatorsResponse(&operators)},
	})
}
