package responses

var (
	ErrInternalServer = &Error{
		Code:      500,
		ErrorCode: "GE001",
		Message:   "Internal server error",
	}

	ErrBadRequest = &Error{
		Code:      400,
		ErrorCode: "GE002",
		Message:   "Bad request",
	}

	ErrPermissionDenied = &Error{
		Code:      403,
		ErrorCode: "GE003",
		Message:   "Permission denied",
	}

	ErrNotFound = &Error{
		Code:      404,
		ErrorCode: "GE004",
		Message:   "Not found",
	}

	ErrAlreadyExists = &Error{
		Code:      409,
		ErrorCode: "GE005",
		Message:   "Already exists",
	}

	ErrUnauthenticated = &Error{
		Code:      401,
		ErrorCode: "GE006",
		Message:   "Unauthorized",
	}

	ErrInvalidCredentials = &Error{
		Code:      401,
		ErrorCode: "GE007",
		Message:   "Invalid credentials",
	}

	ErrNotFoundRecord = &Error{
		Code:      404,
		ErrorCode: "GE008",
		Message:   "Record does not exist",
	}

	ErrInvalidParameters = &Error{
		Code:      400,
		ErrorCode: "GE009",
		Message:   "Invalid parameters",
	}

	ErrTooManyRequest = &Error{
		Code:      429,
		ErrorCode: "GE010",
		Message:   "Too Many Requests",
	}

	ErrInvalidFilter = &Error{
		Code:      400,
		ErrorCode: "GE011",
		Message:   "Invalid filters",
	}
)
