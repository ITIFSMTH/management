package routes

import (
	"management-backend/controllers/authController"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	// Create new group "/auth"
	authGroup := r.Group("/auth")

	// Get controller object
	controller := new(authController.AuthController)

	// Bind routes
	authGroup.POST("/login", controller.Login)
}
