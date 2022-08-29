package handlers

import (
	"git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetVouchers(c *gin.Context) {
	vouchers, fetchErr := h.voucherService.GetVouchers()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, vouchers)
}

func (h *Handler) GetUserVoucherFromCode(c *gin.Context) {
	user := h.GetUserFromToken(c)
	code := c.Param("code")
	uv, fetchErr := h.voucherService.GetUserVoucherFromCode(user, code)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, uv)
}
