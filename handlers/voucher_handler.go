package handlers

import (
	"final-project-backend/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
