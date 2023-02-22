package responses

const (
	// Global errors
	ErrorAlreadyExists = "already_exists"
	ErrorNotExists     = "not_exists"
	ErrorBadData       = "bad_data"
	ErrorServer        = "server_error"

	// Bot errors
	ErrorNotToday            = "not_today"
	ErrorAlreadyOnShift      = "already_on_shift"
	ErrorNoShift             = "no_shift"
	ErrorNoTimeout           = "no_timeout"
	ErrorWrongAnswer         = "wrong_captcha_answer"
	ErrorFirstProvideCaptcha = "first_provide_captcha"

	// Auth errors
	ErrorAuth          = "no_permission"
	ErrorIncorrectData = "incorrect_data"
)
