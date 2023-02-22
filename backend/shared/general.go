package shared

import "management-backend/models"

const (
	PW_COST int = 14
)

var Workers = []models.Worker{
	{Login: "admin", Password: "$2a$14$orvjOpUxaUDzOd4Au6tBzOx1h5wbLY0/w8NE4VbSSTeScHrIEdloK", RoleID: WorkerRoleAdministrator.ID, ThemeID: ThemeLight.ID},
}

var Setting = models.Setting{
	ID:             1,
	TelegramBotKey: "",
}
