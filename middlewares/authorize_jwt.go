package middlewares

import (
	"encoding/json"
	"final-project-backend/config"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httperror.UnauthorizedError()
		}
		return config.Config.JWTSecretKey, nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	s := strings.Split(authHeader, "Bearer ")
	unauthorizedErr := httperror.UnauthorizedError()
	if len(s) < 2 {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	encodedToken := s[1]
	token, err := validateToken(encodedToken)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}

	userJson, _ := json.Marshal(claims["user"])
	var user models.User
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		c.AbortWithStatusJSON(unauthorizedErr.StatusCode, unauthorizedErr)
		return
	}
	c.Set("user", user)
}
