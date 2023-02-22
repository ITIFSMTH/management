package routes

import (
	"management-backend/controllers/userController"
	"management-backend/middlewares"
	"management-backend/shared"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	// Create new group "/user"
	userGroup := r.Group("/user")

	// Use middleware Auth (For all users)
	userGroup.Use(middlewares.GetJWT().AuthHandler(shared.WorkerRoles...))

	// Get controller object
	controller := new(userController.UserController)

	// Bind routes
	userGroup.GET("/theme", controller.GetTheme)

	userGroup.PATCH("/theme", controller.EditTheme)
}
