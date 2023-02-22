package routes

import (
	"management-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	InitMiddleware(engine)

	ApiGroup := engine.Group("/api")

	RegisterAuthRoutes(ApiGroup)
	RegisterUserRoutes(ApiGroup)
	RegisterAdminRoutes(ApiGroup)
	RegisterOperatorRoutes(ApiGroup)
}

func InitMiddleware(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
}
