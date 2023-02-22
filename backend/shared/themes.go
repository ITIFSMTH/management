package shared

import "management-backend/models"

var ThemeLight = models.Theme{ID: 1, Theme: "Светлая"}
var ThemeDark = models.Theme{ID: 2, Theme: "Тёмная"}

var Themes = []models.Theme{
	ThemeLight,
	ThemeDark,
}
