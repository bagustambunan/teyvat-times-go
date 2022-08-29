package middlewares

import (
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}
	firstError := c.Errors[0].Err
	appError, isAppError := firstError.(httperror.AppError)

	if isAppError {
		helpers.ErrorResponse(c, appError.StatusCode, appError)
		return
	}
	serverErr := httperror.InternalServerError(firstError.Error())
	helpers.ErrorResponse(c, serverErr.StatusCode, serverErr)
}
