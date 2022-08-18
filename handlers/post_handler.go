package handlers

import (
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetPosts(c *gin.Context) {
	opt, parsingErr := models.NewGetPostsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}

	postsRes, err := h.postService.GetPosts(opt)
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postsRes)
}

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
