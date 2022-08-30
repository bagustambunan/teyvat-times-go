package handlers

import (
	"final-project-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) TestPing(c *gin.Context) {
	helpers.StandardResponse(c, http.StatusOK, "pong, but via test handler")
}
