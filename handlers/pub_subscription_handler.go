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

// TODO: delete this handler
func (h *Handler) TestAddUserSubscription(c *gin.Context) {
	user := h.GetUserFromToken(c)
	subscription := &models.Subscription{ID: 1}
	us, err := h.subscriptionService.AddUserSubscription(user, subscription)
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusCreated, us)
}
