package shared

import "management-backend/models"

var PollTypeRating = models.PollType{
	ID:   1,
	Poll: "Рейтинг",
}

var PollTypeBudget = models.PollType{
	ID:   2,
	Poll: "Бюджет",
}

var PollTypes = []models.PollType{
	PollTypeRating,
	PollTypeBudget,
}
