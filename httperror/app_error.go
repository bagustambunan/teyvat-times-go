package httperror

import (
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (err AppError) Error() string {
	return err.Message
}

func BadRequestError(message string, code string) AppError {
	if code == "" {
		code = "BAD_REQUEST"
	}
	return AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func InternalServerError(message string) AppError {
	return AppError{
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func ForbiddenError() AppError {
	return AppError{
		Code:       "FORBIDDEN_ERROR",
		Message:    "Forbidden error",
		StatusCode: http.StatusForbidden,
	}
}

func UnauthorizedError() AppError {
	return AppError{
		Code:       "UNAUTHORIZED_ERROR",
		Message:    "Unauthorized error",
		StatusCode: http.StatusUnauthorized,
	}
}

func NotFoundError() AppError {
	return AppError{
		Code:       "NOT_FOUND_ERROR",
		Message:    "Not found error",
		StatusCode: http.StatusNotFound,
	}
}
