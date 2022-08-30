package helpers

import (
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
)

func StandardResponse(c *gin.Context, statusCode int, data interface{}) {
	var response = &models.JSON{StatusCode: statusCode, Data: data}
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Content-Type", "application/json")
	//c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
	//c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Access-Control-Allow-Headers, Access-Control-Allow-Origin, Authorization")
	c.JSON(statusCode, response)
}

func ErrorResponse(c *gin.Context, statusCode int, err interface{}) {
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Content-Type", "application/json")
	//c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
	//c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, Access-Control-Allow-Headers, Access-Control-Allow-Origin, Authorization")
	c.JSON(statusCode, err)
}
