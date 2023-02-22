package adminController

import (
	"management-backend/controllers/authController"
	"management-backend/db"
	"management-backend/models"
	"management-backend/responses"
	"management-backend/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddOperatorRequestBody struct {
	Login    string `json:"login" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=8,max=80"`
	RoleID   uint   `json:"role_id" binding:"required,min=2,max=3"`
	Telegram string `json:"telegram" binding:"required"`
}

// Create a new operator
func (*AdminController) AddOperator(c *gin.Context) {
	// Parse request body
	var requestBody AddOperatorRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		// If parse errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Get DB
	db := db.GetDB()

	// Create a password hash
	hash, err := authController.GetPasswordHash(requestBody.Password)
	if err != nil {
		// If creating hash errored, send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Define a new operator
	var operator = models.Operator{
		Telegram: requestBody.Telegram,
		Worker: models.Worker{
			Login:    requestBody.Login,
			Password: hash,
			ThemeID:  shared.ThemeLight.ID,
		},
	}

	// Create a new junior/senior operator
	var creatingRes *gorm.DB
	var juniorOperator models.JuniorOperator
	var seniorOperator models.SeniorOperator
	if requestBody.RoleID == shared.WorkerRoleJuniorOperator.ID {
		operator.Worker.Role = shared.WorkerRoleJuniorOperator
		juniorOperator = models.JuniorOperator{Operator: operator}
		creatingRes = db.Model(&juniorOperator).Create(&juniorOperator)
		operator = juniorOperator.Operator
	} else {
		operator.Worker.Role = shared.WorkerRoleSeniorOperator
		seniorOperator = models.SeniorOperator{Operator: operator}
		creatingRes = db.Model(&seniorOperator).Create(&seniorOperator)
		operator = seniorOperator.Operator
	}

	if creatingRes.Error != nil || creatingRes.RowsAffected == 0 {
		// If creating failed then send response bad data
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Error: responses.ErrorBadData,
		})
		return
	}

	// Send response (new operator)
	c.JSON(http.StatusOK, responses.Response{
		Data: gin.H{"operator": responses.CreateOperatorResponse(&operator)},
	})
}
