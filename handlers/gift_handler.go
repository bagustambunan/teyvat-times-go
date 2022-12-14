package handlers

import (
	"final-project-backend/helpers"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetGifts(c *gin.Context) {
	gifts, fetchErr := h.giftService.GetGifts()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, gifts)
}
func (h *Handler) GetGift(c *gin.Context) {
	giftID, idErr := strconv.Atoi(c.Param("giftID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	gift, fetchErr := h.giftService.GetGift(&models.Gift{ID: giftID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, gift)
}
