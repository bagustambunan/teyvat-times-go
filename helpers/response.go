package helpers

import (
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
)

func StandardResponse(c *gin.Context, statusCode int, data interface{}) {
	var response = &models.JSON{StatusCode: statusCode, Data: data}
	c.JSON(statusCode, response)
}

func ErrorResponse(c *gin.Context, statusCode int, err interface{}) {
	c.JSON(statusCode, err)
}
