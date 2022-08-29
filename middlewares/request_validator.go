package middlewares

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"reflect"
)

func RequestValidator(model any) gin.HandlerFunc {
	return func(c *gin.Context) {
		modelPtr := reflect.ValueOf(model).Elem()
		modelPtr.Set(reflect.Zero(modelPtr.Type()))
		c.Set("payload", new(any))

		if err := c.ShouldBindJSON(&model); err != nil {
			badRequest := httperror.BadRequestError(err.Error(), "")
			c.AbortWithStatusJSON(badRequest.StatusCode, badRequest)
			return
		}

		c.Set("payload", model)
		c.Next()
	}
}
