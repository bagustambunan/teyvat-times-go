package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) PubGetOverviewPost(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		_ = c.Error(httperror.BadRequestError("Slug is invalid", "INVALID_SLUG"))
		return
	}
	fetchedPost, fetchErr := h.postService.GetPostBySlug(slug)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	postRes := new(dto.PostOverviewRes).FromPost(fetchedPost)
	helpers.StandardResponse(c, http.StatusOK, postRes)
}

func (h *Handler) PubReadPost(c *gin.Context) {
	user := h.GetUserFromToken(c)
	slug := c.Param("slug")
	if slug == "" {
		_ = c.Error(httperror.BadRequestError("Slug is invalid", "INVALID_SLUG"))
		return
	}
	fetchedPost, fetchErr := h.postService.GetPostBySlug(slug)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}

	//accessErr := h.postService.CanUserAccessThisPost(user, fetchedPost)
	//if accessErr != nil {
	//	_ = c.Error(accessErr)
	//	return
	//}

	if fetchedPost.PostTierID != 1 {
		// UNLOCK
		_, unlockErr := h.postService.UnlockAPost(user, fetchedPost)
		if unlockErr != nil {
			_ = c.Error(unlockErr)
			return
		}
		// DECREASE MORA
		_, updateErr := h.userService.UpdateUserMora(user, -fetchedPost.GetMoraRequired())
		if updateErr != nil {
			_ = c.Error(updateErr)
			return
		}
	}

	_, actErr := h.postService.AddActivity(user, fetchedPost)
	if actErr != nil {
		_ = c.Error(actErr)
		return
	}

	postRes := new(dto.GetPostRes).FromPost(fetchedPost)
	helpers.StandardResponse(c, http.StatusOK, postRes)
}

func (h *Handler) PubGetActivity(c *gin.Context) {
	user := h.GetUserFromToken(c)
	postID, idErr := strconv.Atoi(c.Param("postID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	fetchedPost, fetchErr := h.postService.GetPost(&models.Post{ID: postID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	accessErr := h.postService.CanUserAccessThisPost(user, fetchedPost)
	if accessErr != nil {
		_ = c.Error(accessErr)
		return
	}
	act, actErr := h.postService.GetActivity(user, fetchedPost)
	if actErr != nil {
		_ = c.Error(actErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, act)
}

func (h *Handler) PubPostActivity(c *gin.Context) {
	user := h.GetUserFromToken(c)
	postID, idErr := strconv.Atoi(c.Param("postID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	fetchedPost, fetchErr := h.postService.GetPost(&models.Post{ID: postID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}

	accessErr := h.postService.CanUserAccessThisPost(user, fetchedPost)
	if accessErr != nil {
		_ = c.Error(accessErr)
		return
	}

	payload, _ := c.Get("payload")
	actReq, _ := payload.(*dto.ActivityReq)
	act, actErr := h.postService.UpdateActivity(user, fetchedPost, actReq)
	if actErr != nil {
		_ = c.Error(actErr)
		return
	}

	helpers.StandardResponse(c, http.StatusOK, act)
}

func (h *Handler) GetPosts(c *gin.Context) {
	opt, parsingErr := models.NewGetPostsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}

	postsRes, fetchErr := h.postService.GetPosts(opt)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postsRes)
}

func (h *Handler) PubGetReadingHistory(c *gin.Context) {
	user := h.GetUserFromToken(c)
	opt, parsingErr := models.NewReadHistoryOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}
	postsRes, fetchErr := h.postService.GetReadingHistory(user, opt)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postsRes)
}
