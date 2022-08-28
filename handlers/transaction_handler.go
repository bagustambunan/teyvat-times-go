package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

	subscription, subErr := h.subscriptionService.GetSubscription(&models.Subscription{ID: trReq.SubscriptionID})
	if subErr != nil {
		_ = c.Error(httperror.BadRequestError("Invalid subscription", "INVALID_SUBSCRIPTION"))
		return
	}

	trRes, resErr := h.transactionService.AddTransaction(user, trReq, discount, subscription)
	if resErr != nil {
		_ = c.Error(resErr)
		return
	}

	helpers.StandardResponse(c, http.StatusCreated, trRes)
}

func (h *Handler) GetTransactions(c *gin.Context) {
	opt, parsingErr := models.NewTransactionsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}

	trsRes, fetchErr := h.transactionService.GetTransactions(opt)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, trsRes)
}

func (h *Handler) GetUserTransactions(c *gin.Context) {
	user := h.GetUserFromToken(c)
	trsRes, fetchErr := h.transactionService.GetUserTransactions(user)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, trsRes)
}

func (h *Handler) GetTransactionStatuses(c *gin.Context) {
	trStatuses, fetchErr := h.transactionService.GetTransactionStatuses()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, trStatuses)
}
