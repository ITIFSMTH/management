package shared

import "management-backend/models"

var WorkerRoleAdministrator = models.WorkerRole{ID: 1, Role: "Администратор"}
var WorkerRoleJuniorOperator = models.WorkerRole{ID: 2, Role: "Младший оператор"}
var WorkerRoleSeniorOperator = models.WorkerRole{ID: 3, Role: "Старший оператор"}

var WorkerRoles = []models.WorkerRole{
	WorkerRoleAdministrator,
	WorkerRoleJuniorOperator,
	WorkerRoleSeniorOperator,
}
