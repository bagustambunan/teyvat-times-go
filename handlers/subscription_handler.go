package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddTransaction(c *gin.Context) {
	user := h.GetUserFromToken(c)
	payload, _ := c.Get("payload")
	trReq, _ := payload.(*dto.TransactionReq)

	trRes, resErr := h.subscriptionService.AddTransaction(user, trReq)
	if resErr != nil {
		_ = c.Error(resErr)
		return
	}

	helpers.StandardResponse(c, http.StatusCreated, trRes)
}
