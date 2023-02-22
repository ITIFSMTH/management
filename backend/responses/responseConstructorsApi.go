package responses

import (
	"management-backend/db"
	"management-backend/models"
	"management-backend/shared"
	"time"
)

// Create a settings response
func CreateSettingsResponse(settings *models.Setting) SettingsResponse {
	return SettingsResponse{
		TelegramBotKey: settings.TelegramBotKey,
	}
}

// Create a theme response
func CreateThemeResponse(theme *models.Theme) ThemeResponse {
	return ThemeResponse{
		ID:    theme.ID,
		Theme: theme.Theme,
	}
}

// Create a worker role response
func CreateWorkerRoleResponse(workerRole *models.WorkerRole) WorkerRoleResponse {
	return WorkerRoleResponse{
		ID:   workerRole.ID,
		Role: workerRole.Role,
	}
}

// Create a worker response
func CreateWorkerResponse(worker *models.Worker) WorkerResponse {
	return WorkerResponse{
		ID:         worker.ID,
		Login:      worker.Login,
		WorkerRole: CreateWorkerRoleResponse(&worker.Role),
	}
}

// Create a operator response
func CreateOperatorResponse(operator *models.Operator) OperatorResponse {
	// Get operator shifts && timeouts
	db.GetDB().Model(operator).Preload("Worker").Preload("Worker.Role").Preload("Shifts").Preload("Shifts.Timeouts").First(&operator)

	onShift := len(operator.Shifts) > 0 && operator.Shifts[len(operator.Shifts)-1].EndDate.IsZero()
	onTimeout := onShift && len(operator.Shifts[len(operator.Shifts)-1].Timeouts) > 0 && operator.Shifts[len(operator.Shifts)-1].Timeouts[len(operator.Shifts[len(operator.Shifts)-1].Timeouts)-1].EndDate.IsZero()

	return OperatorResponse{
		ID:        operator.ID,
		Worker:    CreateWorkerResponse(&operator.Worker),
		Telegram:  operator.Telegram,
		OnShift:   onShift,
		OnTimeout: onTimeout,
	}
}

func CreateOperatorsResponse(operators *[]models.Operator) []OperatorResponse {
	operatorsResponse := []OperatorResponse{}

	for _, operator := range *operators {
		operatorsResponse = append(operatorsResponse, CreateOperatorResponse(&operator))
	}

	return operatorsResponse
}

// Create a junior operator response
func CreateJuniorOperatorResponse(juniorOperator *models.JuniorOperator) JuniorOperatorResponse {
	return JuniorOperatorResponse{
		Operator: CreateOperatorResponse(&juniorOperator.Operator),
	}
}

// Create a senior operator response
func CreateSeniorOperatorResponse(seniorOperator *models.SeniorOperator) SeniorOperatorResponse {
	return SeniorOperatorResponse{
		Operator:    CreateOperatorResponse(&seniorOperator.Operator),
		Honest:      seniorOperator.Honest,
		Involvement: seniorOperator.Involvement,
		Help:        seniorOperator.Help,
		Activity:    seniorOperator.Activity,
	}
}

// Create a operator shift response
func CreateOperatorShiftResponse(shift *models.Shift) OperatorShiftResponse {
	return OperatorShiftResponse{
		StartDate: shift.StartDate,
		EndDate:   shift.EndDate,
	}
}

// Create a operator timeout response
func CreateOperatorTimeoutResponse(timeout *models.Timeout) OperatorTimeoutResponse {
	return OperatorTimeoutResponse{
		StartDate: timeout.StartDate,
		EndDate:   timeout.EndDate,
	}
}

// Create a operator timeouts response
func CreateOperatorTimeoutsResponse(timeouts *[]models.Timeout) []OperatorTimeoutResponse {
	timeoutsResponse := []OperatorTimeoutResponse{}

	for _, timeout := range *timeouts {
		timeoutsResponse = append(timeoutsResponse, OperatorTimeoutResponse{
			StartDate: timeout.StartDate,
			EndDate:   timeout.EndDate,
		})
	}

	return timeoutsResponse
}

// Create a operator shift response with timeouts
func CreateOperatorShiftWithTimeoutsResponse(shift *models.Shift) OperatorShiftWithTimeoutsResponse {
	return OperatorShiftWithTimeoutsResponse{
		OperatorShiftResponse: CreateOperatorShiftResponse(shift),
		Timeouts:              CreateOperatorTimeoutsResponse(&shift.Timeouts),
	}
}

// Create a operator shift response with timeouts
func CreateOperatorShiftsWithTimeoutsResponse(shifts *[]models.Shift) []OperatorShiftWithTimeoutsResponse {
	operatorShiftsResponse := []OperatorShiftWithTimeoutsResponse{}

	for _, shift := range *shifts {
		operatorShiftsResponse = append(operatorShiftsResponse, CreateOperatorShiftWithTimeoutsResponse(&shift))
	}

	return operatorShiftsResponse
}

// Create a operator statistic response
func CreateOperatorStatisticResponse(operator *models.Operator) OperatorStatisticResponse {
	operatorStatisticResponse := OperatorStatisticResponse{
		Shifts: CreateOperatorShiftsWithTimeoutsResponse(&operator.Shifts),
	}

	// Get DB
	db := db.GetDB()

	// Load operator
	if operator.Worker.RoleID == shared.WorkerRoleJuniorOperator.ID {
		var juniorOperator models.JuniorOperator
		db.Model(&juniorOperator).Where("operator_id = ?", operator.ID).Preload("Operator.Worker.Role").First(&juniorOperator)
		operatorStatisticResponse.Operator = CreateJuniorOperatorResponse(&juniorOperator)
	} else if operator.Worker.RoleID == shared.WorkerRoleSeniorOperator.ID {
		var seniorOperator models.SeniorOperator
		db.Model(&seniorOperator).Where("operator_id = ?", operator.ID).Preload("Operator.Worker.Role").First(&seniorOperator)
		operatorStatisticResponse.Operator = CreateSeniorOperatorResponse(&seniorOperator)
	}

	return operatorStatisticResponse
}

// Create a operators statistic response
func CreateOperatorsStatisticResponse(operators *[]models.Operator) []OperatorStatisticResponse {
	operatorsStatisticResponse := []OperatorStatisticResponse{}

	for _, operator := range *operators {
		operatorsStatisticResponse = append(operatorsStatisticResponse, CreateOperatorStatisticResponse(&operator))
	}

	return operatorsStatisticResponse
}

// Create a month operators statistic response
func CreateMonthOperatorsStatisticResponse(date time.Time, operators *[]models.Operator) MonthOperatorsStatisticResponse {
	return MonthOperatorsStatisticResponse{
		Date:      date,
		Statistic: CreateOperatorsStatisticResponse(operators),
	}
}

// Create a months operators statistic response
func CreateMonthsOperatorsStatisticResponse(operators *[]models.Operator) []MonthOperatorsStatisticResponse {
	monthsOperatorsStatisticResponse := []MonthOperatorsStatisticResponse{}

	// Get first shift
	var firstShiftTime time.Time
	db.GetDB().Model(&models.Shift{}).Select("start_date").First(&firstShiftTime)

	// If no one shift exist
	if firstShiftTime.IsZero() {
		return monthsOperatorsStatisticResponse
	}

	// Get array of months
	months := []time.Time{}
	for month := time.Date(firstShiftTime.Year(), firstShiftTime.Month(), 3, 0, 0, 0, 0, firstShiftTime.Location()); ; month = month.AddDate(0, 1, 0) {
		months = append(months, month)

		if month.Format("2006-01") == time.Now().Format("2006-01") {
			break
		}
	}

	// Check is statistic exists
	if len(months) == 0 {
		months = append(months, firstShiftTime)
	}

	// Reverse months
	for i, j := 0, len(months)-1; i < j; i, j = i+1, j-1 {
		months[i], months[j] = months[j], months[i]
	}

	// Get for every month operators shifts
	for _, month := range months {
		// Copy operators
		monthOperators := *operators

		// Get for every operator shifts for current month
		for i, monthOperator := range *operators {
			// Copy operator shifts
			operatorShifts := monthOperator.Shifts

			// Clear operator shifts
			monthOperators[i].Shifts = make([]models.Shift, 0)

			// Get only this month shifts
			for _, operatorShift := range operatorShifts {
				if operatorShift.StartDate.Month() != month.Month() {
					continue
				}
				monthOperators[i].Shifts = append(monthOperators[i].Shifts, operatorShift)
			}
		}

		// Append month response
		monthsOperatorsStatisticResponse = append(monthsOperatorsStatisticResponse, CreateMonthOperatorsStatisticResponse(month, &monthOperators))
	}

	return monthsOperatorsStatisticResponse
}

// Create a rating vote response
func CreateRatingVoteResponse(ratingVote *models.RatingVote) RatingVoteResponse {
	return RatingVoteResponse{
		ID:          ratingVote.ID,
		Voter:       CreateJuniorOperatorResponse(&ratingVote.Voter),
		Candidate:   CreateSeniorOperatorResponse(&ratingVote.Candidate),
		Honest:      ratingVote.Honest,
		Involvement: ratingVote.Involvement,
		Help:        ratingVote.Help,
		Activity:    ratingVote.Activity,
	}
}

// Create a rating vote response
func CreateBudgetVoteResponse(budgetVote *models.BudgetVote) BudgetVoteResponse {
	return BudgetVoteResponse{
		ID:        budgetVote.ID,
		Voter:     CreateSeniorOperatorResponse(&budgetVote.Voter),
		Candidate: CreateJuniorOperatorResponse(&budgetVote.Candidate),
		Budget:    budgetVote.Budget,
	}
}

// Create a poll type response
func CreatePollTypeResponse(pollType *models.PollType) PollTypeResponse {
	return PollTypeResponse{
		ID:   pollType.ID,
		Poll: pollType.Poll,
	}
}

// Create a rating poll response
func CreateRatingPollResponse(poll *models.Poll) PollResponse {
	return PollResponse{
		ID:        poll.ID,
		StartDate: poll.StartDate,
		Type:      CreatePollTypeResponse(&poll.PollType),
	}
}

// Create a budget poll response
func CreateBudgetPollResponse(budgetPoll *models.BudgetPoll) BudgetPollResponse {
	return BudgetPollResponse{
		PollResponse: CreateRatingPollResponse(&budgetPoll.Poll),
		Budget:       budgetPoll.Budget,
	}
}
