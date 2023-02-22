package responses

import (
	"time"
)

// Object to send response
type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// models.Setting response
type SettingsResponse struct {
	TelegramBotKey string `json:"telegram_bot_key"`
}

// models.Theme response
type ThemeResponse struct {
	ID    uint   `json:"id"`
	Theme string `json:"theme"`
}

// models.WorkerRole response
type WorkerRoleResponse struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

// models.Worker response
type WorkerResponse struct {
	ID         uint               `json:"id"`
	Login      string             `json:"login"`
	WorkerRole WorkerRoleResponse `json:"worker_role"`
}

// Operator response (Creating with models.Operator)
type OperatorResponse struct {
	ID        uint           `json:"id"`
	Worker    WorkerResponse `json:"worker"`
	Telegram  string         `json:"telegram"`
	OnShift   bool           `json:"on_shift"`
	OnTimeout bool           `json:"on_timeout"`
}

// Junior operator response
type JuniorOperatorResponse struct {
	Operator OperatorResponse `json:"operator"`
}

// Senior operator response
type SeniorOperatorResponse struct {
	Operator    OperatorResponse `json:"operator"`
	Honest      float32          `json:"honest"`
	Involvement float32          `json:"involvement"`
	Help        float32          `json:"help"`
	Activity    float32          `json:"activity"`
}

// models.OperatorShift response
type OperatorShiftResponse struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Delays    uint8     `json:"delays"`
}

// models.OperatorTimeout response
type OperatorTimeoutResponse struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Operator shift with timeouts response
type OperatorShiftWithTimeoutsResponse struct {
	OperatorShiftResponse
	Timeouts []OperatorTimeoutResponse `json:"timeouts"`
}

// Operator Statistic response
type OperatorStatisticResponse struct {
	Operator interface{}                         `json:"operator,omitempty"`
	Shifts   []OperatorShiftWithTimeoutsResponse `json:"operator_shifts"`
}

// Operators
type MonthOperatorsStatisticResponse struct {
	Date      time.Time                   `json:"date"`
	Statistic []OperatorStatisticResponse `json:"statistic"`
}

// Rating vote response (models.RatingVote)
type RatingVoteResponse struct {
	ID          uint                   `json:"id"`
	Voter       JuniorOperatorResponse `json:"voter"`
	Candidate   SeniorOperatorResponse `json:"candidate"`
	Honest      uint8                  `json:"honest"`
	Involvement uint8                  `gorm:"involvement"`
	Help        uint8                  `gorm:"help"`
	Activity    uint8                  `gorm:"activity"`
}

// Budget vote response (models.BudgetVote)
type BudgetVoteResponse struct {
	ID        uint                   `json:"id"`
	Voter     SeniorOperatorResponse `json:"voter"`
	Candidate JuniorOperatorResponse `json:"candidate"`
	Budget    uint                   `json:"budget"`
}

// Poll Type Response
type PollTypeResponse struct {
	ID   uint   `json:"id"`
	Poll string `json:"poll"`
}

// Poll response
type PollResponse struct {
	ID        uint             `json:"id"`
	StartDate time.Time        `json:"start_date"`
	Type      PollTypeResponse `json:"type"`
}

// Budget poll response
type BudgetPollResponse struct {
	PollResponse
	Budget uint `json:"budget"`
}
