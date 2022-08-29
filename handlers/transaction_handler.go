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
	opt, parsingErr := models.NewTransactionsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}

	trsRes, fetchErr := h.transactionService.GetUserTransactions(user, opt)
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

func (h *Handler) GetTransactionDetail(c *gin.Context) {
	transactionID, idErr := strconv.Atoi(c.Param("transactionID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	tr, fetchErr := h.transactionService.GetTransaction(&models.Transaction{ID: transactionID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, tr)
}

func (h *Handler) ApproveTransaction(c *gin.Context) {
	transactionID, idErr := strconv.Atoi(c.Param("transactionID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	fetchedTr, fetchTrErr := h.transactionService.GetTransaction(&models.Transaction{ID: transactionID})
	if fetchTrErr != nil {
		_ = c.Error(fetchTrErr)
		return
	}

	updatedTr, updateTrErr := h.transactionService.ApproveTransaction(fetchedTr)
	if updateTrErr != nil {
		_ = c.Error(updateTrErr)
		return
	}

	if fetchedTr.UserVoucherID != 0 {
		_, updateUvErr := h.voucherService.UseUserVoucher(&models.UserVoucher{ID: fetchedTr.UserVoucherID})
		if updateUvErr != nil {
			_ = c.Error(updateUvErr)
			return
		}
	}

	_, updateCoinErr := h.userService.UpdateUserCoins(
		&models.User{ID: fetchedTr.UserID},
		fetchedTr.Subscription.CoinsAmount,
	)
	if updateCoinErr != nil {
		_ = c.Error(updateCoinErr)
		return
	}

	_, usErr := h.subscriptionService.AddUserSubscription(
		&models.User{ID: fetchedTr.UserID},
		&models.Subscription{ID: fetchedTr.SubscriptionID},
	)
	if usErr != nil {
		_ = c.Error(usErr)
		return
	}

	helpers.StandardResponse(c, http.StatusOK, updatedTr)
}

func (h *Handler) RejectTransaction(c *gin.Context) {
	transactionID, idErr := strconv.Atoi(c.Param("transactionID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	fetchedTr, fetchTrErr := h.transactionService.GetTransaction(&models.Transaction{ID: transactionID})
	if fetchTrErr != nil {
		_ = c.Error(fetchTrErr)
		return
	}

	updatedTr, updateTrErr := h.transactionService.RejectTransaction(fetchedTr)
	if updateTrErr != nil {
		_ = c.Error(updateTrErr)
		return
	}

	helpers.StandardResponse(c, http.StatusOK, updatedTr)
}

func (h *Handler) ProcessPayment(c *gin.Context) {
	payload, _ := c.Get("payload")
	paymentReq, _ := payload.(*dto.PaymentReq)

	fetchedTr, fetchTrErr := h.transactionService.GetTransaction(&models.Transaction{ID: paymentReq.TransactionID})
	if fetchTrErr != nil {
		_ = c.Error(fetchTrErr)
		return
	}

	updatedTr, updateTrErr := h.transactionService.ProcessPayment(fetchedTr, paymentReq)
	if updateTrErr != nil {
		_ = c.Error(updateTrErr)
		return
	}

	helpers.StandardResponse(c, http.StatusOK, updatedTr)
}
