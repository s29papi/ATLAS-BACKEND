package worker

const (
	ERR_MAX_NUMBER_LINES_NO           = 90
	ERR_UNEXPECTED_FIELD_NO           = 91
	ERR_MISSING_CURRENCY_SYMBOL_NO    = 92
	ERR_INVALID_LENGTH_WAGERAMOUNT_NO = 93
	ERR_INVALID_AMOUNT_NO             = 94
	ERR_MISSING_REQ_FIELD_NO          = 95
)

var (
	ERR_MAX_NUMBER_LINES           = "Error: Exceeds Maximum Number of lines."
	ERR_UNEXPECTED_FIELD           = "Error: Unexpected field."
	ERR_MISSING_CURRENCY_SYMBOL    = "Error: Missing currency symbol in amount."
	ERR_INVALID_LENGTH_WAGERAMOUNT = "Error: Invalid length of amount field."
	ERR_INVALID_AMOUNT             = "Error: Invalid amount."
	ERR_MISSING_REQ_FIELD          = "Error: missing required fields."
)
