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

func (h *Handler) GetPost(c *gin.Context) {
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
	postRes := new(dto.GetPostRes).FromPost(fetchedPost)
	helpers.StandardResponse(c, http.StatusOK, postRes)
}

func (h *Handler) AddPost(c *gin.Context) {
	payload, _ := c.Get("payload")
	postReq, _ := payload.(*dto.PostReq)

	postRes, err := h.postService.AddPost(&models.Post{
		PostTierID:     postReq.PostTierID,
		PostCategoryID: postReq.PostCategoryID,
		Title:          postReq.Title,
		Content:        postReq.Content,
		Summary:        postReq.Summary,
		ImgThumbnailID: 2,
		ImgContentID:   3,
		CreatedByID:    postReq.CreatedByID,
		UpdatedById:    postReq.UpdatedByID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
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

	accessErr := h.postService.CanUserAccessThisPost(user, fetchedPost)
	if accessErr != nil {
		_ = c.Error(httperror.UnauthorizedError())
		return
	}

	_, actErr := h.postService.AddActivity(user, fetchedPost)
	if actErr != nil {
		_ = c.Error(actErr)
		return
	}

	postRes := new(dto.GetPostRes).FromPost(fetchedPost)
	helpers.StandardResponse(c, http.StatusOK, postRes)
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
		_ = c.Error(httperror.UnauthorizedError())
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
