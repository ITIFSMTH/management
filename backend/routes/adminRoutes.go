package routes

import (
	"management-backend/controllers/adminController"
	"management-backend/controllers/operatorsControllers/seniorOperatorController"
	"management-backend/middlewares"
	"management-backend/shared"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.RouterGroup) {
	// Create new group "/admin"
	adminGroup := r.Group("/admin")

	// Use middleware Auth (Only Role Admin)
	adminGroup.Use(middlewares.GetJWT().AuthHandler(shared.WorkerRoleAdministrator))

	// Get controller object
	controller := new(adminController.AdminController)
	seniorOperatorController := new(seniorOperatorController.SeniorOperatorConroller)

	// Bind routes
	adminGroup.GET("/operators", controller.GetOperators)
	adminGroup.GET("/operators/statistic", controller.GetOperatorsStatistic)
	adminGroup.GET("/settings", controller.GetSettings)
	adminGroup.GET("/poll/budget", seniorOperatorController.GetBudgetPoll)

	adminGroup.POST("/operator", controller.AddOperator)
	adminGroup.POST("/poll/budget", controller.AddBudgetPoll)

	adminGroup.PATCH("/operator", controller.UpdateOperator)
	adminGroup.PATCH("/settings", controller.UpdateSettings)
	adminGroup.PATCH("/worker", controller.UpdateWorker)
	adminGroup.PATCH("/worker/password", controller.UpdateWorkerPassword)

	adminGroup.DELETE("/operator", controller.RemoveOperator)
}
