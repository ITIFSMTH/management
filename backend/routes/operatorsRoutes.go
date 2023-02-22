package routes

import (
	operatorsController "management-backend/controllers/operatorsControllers"
	"management-backend/middlewares"
	"management-backend/shared"

	"github.com/gin-gonic/gin"
)

func RegisterOperatorRoutes(r *gin.RouterGroup) {
	// Create new groups
	operatorsGroup := r.Group("/operator")
	//juniorOperatorGroup := operatorsGroup.Group("/junior/")
	//seniorOperatorGroup := operatorsGroup.Group("/senior/")

	// Use middleware Auth
	operatorsGroup.Use(middlewares.GetJWT().AuthHandler(shared.WorkerRoleJuniorOperator, shared.WorkerRoleSeniorOperator))

	// Get controller object
	operatorsController := new(operatorsController.OperatorsController)
	//juniorOperatorConroller := new(juniorOperatorConroller.JuniorOperatorConroller)
	//seniorOperatorConroller := new(seniorOperatorController.SeniorOperatorConroller)

	// Bind routes
	operatorsGroup.GET("/statistic", operatorsController.GetOperatorStatistic)
}
