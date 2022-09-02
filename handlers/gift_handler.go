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

func (h *Handler) GetUnclaimedUserGifts(c *gin.Context) {
	user := h.GetUserFromToken(c)
	ugs, fetchErr := h.giftService.GetUnclaimedUserGifts(user)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, ugs)
}

func (h *Handler) SaveGiftClaim(c *gin.Context) {
	user := h.GetUserFromToken(c)
	gc, insertErr := h.giftService.SaveGiftClaim(user)
	if insertErr != nil {
		_ = c.Error(insertErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, gc)
}

func (h *Handler) GetUserGiftClaims(c *gin.Context) {
	user := h.GetUserFromToken(c)
	gcs, fetchErr := h.giftService.GetUserGiftClaims(user)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, gcs)
}
