package authController

import (
	"management-backend/db"
	"management-backend/middlewares"
	"management-backend/models"
	"management-backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (*AuthController) Login(c *gin.Context) {
	// Get User Data
	var requestBody LoginRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Get worker
	var worker models.Worker
	db.Model(&worker).
		Where("login = ?", requestBody.Login).
		Preload("Role").
		First(&worker)

	// Check password
	if r := ComparePasswordHash(requestBody.Password, worker.Password); !r {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorIncorrectData,
		})
		return
	}

	// Issue JWT token
	token, _ := middlewares.GetJWT().GenerateToken(
		requestBody.Login,
		worker.Role.ID,
		worker.ID,
	)

	// Send response
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"token": token},
	})
}
