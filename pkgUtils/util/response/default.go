package responses

var (
	DefaultSuccess = General{
		Code:    200,
		Message: "success",
		Data:    nil,
	}

	DefaultError = General{
		Code:    500,
		Message: "Internal server error",
		Data:    nil,
	}
)
