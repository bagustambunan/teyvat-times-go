package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
		// TODO: check expired
		//if uv.DateExpired
		if uv.IsUsed != 0 {
			_ = c.Error(httperror.BadRequestError("User voucher is already used", "USER_VOUCHER_ALREADY_USED"))
			return
		}
		discount = uv.GetVoucherAmount()
		_, updateErr := h.voucherService.UseUserVoucher(uv)
		if updateErr != nil {
			_ = c.Error(updateErr)
			return
		}
	}

	trRes, resErr := h.subscriptionService.AddTransaction(user, trReq, discount)
	if resErr != nil {
		_ = c.Error(resErr)
		return
	}

	helpers.StandardResponse(c, http.StatusCreated, trRes)
}
