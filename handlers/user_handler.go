package handlers

import (
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetUser(c *gin.Context) {
	user := h.GetUserFromToken(c)
	userID, idErr := strconv.Atoi(c.Param("userID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	if user.ID != userID {
		_ = c.Error(httperror.ForbiddenError())
		return
	}

	userRes, fetchErr := h.userService.GetUser(&models.User{ID: userID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, userRes)
}

func (h *Handler) GetUserDownLines(c *gin.Context) {
	user := h.GetUserFromToken(c)
	userID, idErr := strconv.Atoi(c.Param("userID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	if user.ID != userID {
		_ = c.Error(httperror.ForbiddenError())
		return
	}

	uSpending, fetchErr := h.userService.GetUserDownLines(&models.User{ID: userID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, uSpending)
}
