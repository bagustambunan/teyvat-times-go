package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetSubscriptions(c *gin.Context) {
	subscriptions, fetchErr := h.subscriptionService.GetSubscriptions()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, subscriptions)
}

func (h *Handler) GetSubscription(c *gin.Context) {
	subscriptionID, idErr := strconv.Atoi(c.Param("subscriptionID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	subscription, fetchErr := h.subscriptionService.GetSubscription(&models.Subscription{ID: subscriptionID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, subscription)
}

func (h *Handler) GetUserNewSubscriptionDate(c *gin.Context) {
	user := h.GetUserFromToken(c)
	dateStart, dateEnded := h.subscriptionService.GetUserNewSubscriptionDate(user)
	newDate := &dto.UserNewSubscriptionDateRes{
		DateStart: dateStart,
		DateEnded: dateEnded,
	}
	helpers.StandardResponse(c, http.StatusOK, newDate)
}

func (h *Handler) GetUserSubscriptions(c *gin.Context) {
	user := h.GetUserFromToken(c)
	opt, parsingErr := models.NewGetUserSubscriptionsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}
	ussRes, fetchErr := h.subscriptionService.GetUserSubscriptions(user, opt)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, ussRes)
}

func (h *Handler) GetUserActiveSubscription(c *gin.Context) {
	user := h.GetUserFromToken(c)
	activeUs := h.subscriptionService.GetUserActiveSubscription(user)
	helpers.StandardResponse(c, http.StatusOK, activeUs)
}
