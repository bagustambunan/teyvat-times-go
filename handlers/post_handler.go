package handlers

import (
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetPost(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		_ = c.Error(httperror.BadRequestError("Slug is invalid", "INVALID_SLUG"))
		return
	}

	postRes, err := h.postService.GetPostBySlug(slug)
	if err != nil {
		_ = c.Error(err)
		return
	}

	helpers.StandardResponse(c, http.StatusOK, postRes)
}
