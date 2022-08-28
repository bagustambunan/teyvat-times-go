package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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

func (h *Handler) AddTransaction(c *gin.Context) {
	user := h.GetUserFromToken(c)
	payload, _ := c.Get("payload")
	trReq, _ := payload.(*dto.TransactionReq)

	var discount int
	if trReq.UserVoucherID != 0 {
		uv, err := h.voucherService.GetUserVoucher(&models.UserVoucher{ID: trReq.UserVoucherID})
		if err != nil {
			_ = c.Error(err)
			return
		}
		if uv.UserID != user.ID {
			_ = c.Error(httperror.BadRequestError("Invalid user voucher", "INVALID_USER_VOUCHER"))
			return
		}
		dateExpired, _ := time.Parse("2006-01-02T00:00:00Z", uv.DateExpired)
		if dateExpired.Before(time.Now()) {
			_ = c.Error(httperror.BadRequestError("User voucher expired", "USER_VOUCHER_EXPIRED"))
			return
		}
		if uv.IsUsed != 0 {
			_ = c.Error(httperror.BadRequestError("User voucher is already used", "USER_VOUCHER_ALREADY_USED"))
			return
		}
		discount = uv.GetVoucherAmount()
		//_, updateErr := h.voucherService.UseUserVoucher(uv)
		//if updateErr != nil {
		//	_ = c.Error(updateErr)
		//	return
		//}
	}

	trRes, resErr := h.subscriptionService.AddTransaction(user, trReq, discount)
	if resErr != nil {
		_ = c.Error(resErr)
		return
	}

	helpers.StandardResponse(c, http.StatusCreated, trRes)
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
