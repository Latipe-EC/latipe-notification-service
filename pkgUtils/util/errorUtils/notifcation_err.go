package errorUtils

import "latipe-notification-service/pkgUtils/util/response"

var (
	ErrInternalServer = &responses.Error{
		Code:      500,
		ErrorCode: "GE001",
		Message:   "Internal server error",
	}

	ErrBadRequest = &responses.Error{
		Code:      400,
		ErrorCode: "GE002",
		Message:   "Bad request",
	}

	ErrPermissionDenied = &responses.Error{
		Code:      403,
		ErrorCode: "GE003",
		Message:   "Permission denied",
	}

	ErrNotFound = &responses.Error{
		Code:      404,
		ErrorCode: "GE004",
		Message:   "Not found",
	}

	ErrAlreadyExists = &responses.Error{
		Code:      409,
		ErrorCode: "GE005",
		Message:   "Already exists",
	}

	ErrUnauthenticated = &responses.Error{
		Code:      401,
		ErrorCode: "GE006",
		Message:   "Unauthorized",
	}

	ErrInvalidCredentials = &responses.Error{
		Code:      401,
		ErrorCode: "GE007",
		Message:   "Invalid credentials",
	}

	ErrNotFoundRecord = &responses.Error{
		Code:      404,
		ErrorCode: "GE008",
		Message:   "Record does not exist",
	}

	ErrInvalidParameters = &responses.Error{
		Code:      400,
		ErrorCode: "GE009",
		Message:   "Invalid parameters",
	}

	ErrTooManyRequest = &responses.Error{
		Code:      429,
		ErrorCode: "GE010",
		Message:   "Too Many Requests",
	}

	ErrInvalidFilter = &responses.Error{
		Code:      400,
		ErrorCode: "GE011",
		Message:   "Invalid filters",
	}

	ErrInvalidDatetimeParameters = &responses.Error{
		Code:      400,
		ErrorCode: "GE012",
		Message:   "Invalid datetime parameters (may be start date is after end date or before current date)",
	}

	ErrParseDatetimeParameters = &responses.Error{
		Code:      400,
		ErrorCode: "GE013",
		Message:   "Failed to parse datetime parameters with the given layout (yyyy-MM-dd HH:mm:ss)",
	}

	ErrNotificationNotFoundOrInActive = &responses.Error{
		Code:      404,
		ErrorCode: "GE014",
		Message:   "Notification not found or inactive",
	}
	ErrDeviceAlreadyRegistered = &responses.Error{
		Code:      400,
		ErrorCode: "GE014",
		Message:   "Device already registered",
	}
)
